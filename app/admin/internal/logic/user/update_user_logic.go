package user

import (
	"context"
	"time"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/ent/generated/user"

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
	user, err := l.svcCtx.DB.User.Query().Where(user.UserID(uint64(req.UserID))).Only(l.ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return false, xerr.NewErrMsg("用户不存在")
		}
		return false, xerr.NewErrCode(xerr.DbError)
	}

	// 2. 更新用户信息
	updateBuilder := l.svcCtx.DB.User.UpdateOne(user).
		SetUpdatedAt(time.Now().UnixMilli())

	if req.Name != "" {
		updateBuilder.SetUserName(req.Name)
		updateBuilder.SetNickName(req.Name)
	}

	// 3. 执行更新
	_, err = updateBuilder.Save(l.ctx)
	if err != nil {
		return false, xerr.NewErrCode(xerr.DbError)
	}

	return true, nil
}
