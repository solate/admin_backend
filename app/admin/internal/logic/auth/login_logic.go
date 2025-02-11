package auth

import (
	"context"
	"net/http"
	"time"

	"admin_backend/app/admin/internal/repository/loginLogRepo"
	"admin_backend/app/admin/internal/repository/userRepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/context_util"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"
	"admin_backend/pkg/utils/jwt"
	"admin_backend/pkg/utils/passwordgen"
	"admin_backend/pkg/utils/userAgent"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userRepo     *userRepo.UserRepo
	loginLogRepo *loginLogRepo.LoginLogRepo
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		userRepo:     userRepo.NewUserRepo(svcCtx.DB),
		loginLogRepo: loginLogRepo.NewLoginLogRepo(svcCtx.DB),
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	// 1. 查找用户
	user, err := l.userRepo.GetByUserName(l.ctx, req.UserName)
	if err != nil {
		l.Error("Login userRepo.GetByUserName err:", err.Error())
		if generated.IsNotFound(err) {
			return nil, xerr.NewErrMsg("用户不存在")
		}
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 2. 验证密码
	if passwordgen.VerifyPassword(req.Password, user.PwdSalt) {
		l.Error("Login passwordgen.VerifyPassword err:", err.Error())
		return nil, xerr.NewErrMsg("密码错误")
	}

	// 3. 检查用户状态
	if user.Status != 1 {
		l.Error("Login user.Status != 1 err:", err.Error())
		return nil, xerr.NewErrMsg("用户已被禁用")
	}

	// 4. 生成JWT Token
	token, err := jwt.GenerateToken(user.UserID, user.TenantCode, jwt.JWTConfig{
		AccessSecret: []byte(l.svcCtx.Config.JwtAuth.AccessSecret),
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
	})
	if err != nil {
		l.Error("Login jwt.GenerateToken err:", err.Error())
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	// 5. 更新用户Token
	_, err = l.svcCtx.DB.User.UpdateOne(user).
		SetToken(token).
		SetUpdatedAt(time.Now().UnixMilli()).
		Save(l.ctx)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 添加登录日志
	err = l.addLoginLog(user)
	if err != nil {
		l.Error("Login addLoginLog err:", err.Error())
	}

	// 7. 返回结果
	return &types.LoginResp{
		UserID:   user.UserID,
		Token:    token,
		UserName: user.UserName,
		Phone:    user.Phone,
		Email:    user.Email,
	}, nil
}

// 添加登录日志
func (l *LoginLogic) addLoginLog(user *generated.User) error {
	tenantCode, err := context_util.GetTenantCodeFromCtx(l.ctx)
	if err != nil {
		l.Error("addLoginLog context_util.GetTenantIDFromCtx err: ", err.Error())
		return xerr.NewErrCodeMsg(xerr.ServerError, "获取租户码失败")
	}

	id, err := idgen.GenerateUUID()
	if err != nil {
		l.Error("addLoginLog GenerateUUID err:", err.Error())
		return xerr.NewErrCodeMsg(xerr.ServerError, "生成ID失败")
	}

	r := l.ctx.Value("request").(*http.Request)
	// 获取客户端信息
	clientInfo := userAgent.GetClientInfo(r)

	log := &generated.LoginLog{
		TenantCode: tenantCode,
		LogID:      id,
		UserID:     user.UserID,
		UserName:   user.UserName,
		IP:         clientInfo.IP,
		Message:    "登录成功",
		UserAgent:  clientInfo.UserAgent,
		Browser:    clientInfo.Browser + " " + clientInfo.BrowserVer,
		Os:         clientInfo.OS,
		Device:     clientInfo.Device,
		LoginTime:  time.Now().UnixMilli(),
	}

	// 创建登录日志记录
	_, err = l.loginLogRepo.Create(l.ctx, log)
	return err
}
