package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserReq) (resp bool, err error) {
	// 1. 检查用户是否存在
	user, err := l.userRepo.GetByUserID(l.ctx, req.UserID)
	if err != nil {
		l.Error("DeleteUser userRepo.GetByUserID err:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCodeMsg(xerr.DbError, "查询用户失败")
	}

	// 2. 软删除用户
	_, err = l.userRepo.Delete(l.ctx, user.UserID)
	if err != nil {
		l.Error("DeleteUser userRepo.Update err:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, "删除用户失败")
	}

	return true, nil
}
