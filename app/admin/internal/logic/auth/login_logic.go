package auth

import (
	"context"
	"net/http"

	"admin_backend/app/admin/internal/repository/loginlogrepo"
	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/jwt"
	"admin_backend/pkg/utils/passwordgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userRepo     *userrepo.UserRepo
	loginLogRepo *loginlogrepo.LoginLogRepo
	r            *http.Request
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginLogic {
	return &LoginLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		userRepo:     userrepo.NewUserRepo(svcCtx.DB),
		loginLogRepo: loginlogrepo.NewLoginLogRepo(svcCtx.DB),
		r:            r,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {

	// // 1. 验证验证码
	// if !l.svcCtx.CaptchaManager.Verify(req.CaptchaId, req.Captcha) {
	// 	return nil, xerr.NewErrMsg("验证码错误或已过期")
	// }

	// 2. 查找用户
	user, err := l.userRepo.GetByUserName(l.ctx, req.UserName)
	if err != nil {
		l.Error("Login userRepo.GetByUserName err:", err.Error())
		if generated.IsNotFound(err) {
			return nil, xerr.NewErrMsg("用户不存在")

		}
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 2. 验证密码
	if !passwordgen.VerifyPassword(req.Password, user.PwdHashed) {
		l.Error("Login passwordgen.VerifyPassword err")
		return nil, xerr.NewErrMsg("密码错误")
	}

	// 3. 检查用户状态
	if user.Status != 1 {
		l.Error("Login user.Status != 1 err")
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
	_, err = l.userRepo.UpdateToken(l.ctx, user.UserID, token)
	if err != nil {
		l.Error("Login userRepo.UpdateToken err:", err.Error())
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 添加登录日志
	err = l.loginLogRepo.AddLoginLog(l.ctx, l.r, user, "登录成功")
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
