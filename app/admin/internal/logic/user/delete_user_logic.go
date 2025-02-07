package user

import (
	"context"
	"time"

	"github.com/solate/admin_backend/app/admin/internal/repository/user_repo"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *user_repo.UserRepo
}

// 删除用户
func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: user_repo.NewUserRepo(svcCtx.DB),
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
	now := time.Now().UnixMilli()
	user.DeletedAt = &now
	_, err = l.userRepo.Update(l.ctx, user)
	if err != nil {
		l.Error("DeleteUser userRepo.Update err:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, "删除用户失败")
	}

	return true, nil
}
