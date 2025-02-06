package auth

import (
	"context"
	"time"

	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
	"github.com/solate/admin_backend/pkg/utils/jwt"
	"github.com/solate/admin_backend/pkg/utils/passwordgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 1. 查找用户
	user, err := l.svcCtx.Orm.User.Query().
		Where(user.UserName(req.UserName)).
		Where(user.DeletedAt(0)).
		Only(l.ctx)

	if err != nil {
		if generated.IsNotFound(err) {
			return nil, xerr.NewErrMsg("用户不存在")
		}
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 2. 验证密码
	hashedPassword, err := passwordgen.Argon2Hash(req.Password, []byte(user.Salt))
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	if hashedPassword != user.Password {
		return nil, xerr.NewErrMsg("密码错误")
	}

	// 3. 检查用户状态
	if user.Status != 1 {
		return nil, xerr.NewErrMsg("用户已被禁用")
	}

	// 4. 生成JWT Token
	token, err := jwt.GenerateToken(int64(user.UserID), user.TenantCode, jwt.JWTConfig{
		AccessSecret: []byte(l.svcCtx.Config.JwtAuth.AccessSecret),
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ServerError)
	}

	// 5. 更新用户Token
	_, err = l.svcCtx.Orm.User.UpdateOne(user).
		SetToken(token).
		SetUpdatedAt(time.Now().UnixMilli()).
		Save(l.ctx)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 6. 返回结果
	return &types.LoginResp{
		UserID:   user.UserID,
		Token:    token,
		UserName: user.UserName,
		Phone:    user.Phone,
		Email:    user.Email,
	}, nil
}
