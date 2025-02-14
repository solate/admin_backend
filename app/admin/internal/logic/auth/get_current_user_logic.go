package auth

import (
	"context"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/contextutil"
	"admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCurrentUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 获取当前用户信息
func NewGetCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCurrentUserLogic {
	return &GetCurrentUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *GetCurrentUserLogic) GetCurrentUser() (resp *types.GetUserResp, err error) {
	// 从上下文中获取当前用户ID
	userID := contextutil.GetUserIDFromCtx(l.ctx)

	// 查询用户信息
	user, err := l.userRepo.GetByUserID(l.ctx, userID)
	if err != nil {
		l.Error("GetCurrentUser l.userRepo.GetByUserID err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "获取用户信息失败")
	}

	resp = &types.GetUserResp{
		UserInfo: types.UserInfo{
			UserID:    user.UserID,
			UserName:  user.UserName,
			Name:      user.NickName,
			Phone:     user.Phone,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		},
	}

	return
}
