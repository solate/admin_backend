package auth

import (
	"context"

	"admin_backend/app/admin/internal/repository/login_log_repo"
	"admin_backend/app/admin/internal/repository/user_repo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx          context.Context
	svcCtx       *svc.ServiceContext
	userRepo     *user_repo.UserRepo
	loginLogRepo *login_log_repo.LoginLogRepo
}

// 用户登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger:       logx.WithContext(ctx),
		ctx:          ctx,
		svcCtx:       svcCtx,
		userRepo:     user_repo.NewUserRepo(svcCtx.DB),
		loginLogRepo: login_log_repo.NewLoginLogRepo(svcCtx.DB),
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp bool, err error) {

	// 1. 查找用户
	user, err := l.userRepo.GetByUserID(l.ctx, req.UserID)
	if err != nil {
		l.Error("Logout User.Query Error:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCodeMsg(xerr.DbError, err.Error())
	}

	// 更新user
	_, err = l.userRepo.Logout(l.ctx, req.UserID)
	if err != nil {
		l.Error("Logout User.Logout Error:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, err.Error())
	}

	// 添加登出日志
	err = l.loginLogRepo.AddLoginLog(l.ctx, user, "登出成功")
	if err != nil {
		l.Error("Logout addLoginLog err:", err.Error())
	}

	return true, nil
}
