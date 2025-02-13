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

type UpdateUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 更新用户
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp bool, err error) {
	// 1. 检查用户是否存在
	user, err := l.userRepo.GetByUserID(l.ctx, req.UserID)
	if err != nil {
		l.Error("UpdateUser userRepo.GetByUserID err:", err.Error())
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCode(xerr.DbError)
	}

	if user.UserName == "admin" && req.Status != 1 {
		return false, xerr.NewErrMsg("管理员账号不能禁用")
	}

	// 2. 更新用户信息
	user.NickName = req.Name
	user.Email = req.Email
	user.Status = req.Status
	user.Sex = req.Sex
	user.Avatar = req.Avatar

	_, err = l.userRepo.Update(l.ctx, user)
	if err != nil {
		l.Error("UpdateUser Update err:", err.Error())
		return false, xerr.NewErrCodeMsg(xerr.DbError, "更新用户失败")
	}

	return true, nil
}
