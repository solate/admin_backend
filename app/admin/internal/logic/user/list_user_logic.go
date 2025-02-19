package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/userrepo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *userrepo.UserRepo
}

// 获取用户列表
func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: userrepo.NewUserRepo(svcCtx.DB),
	}
}

func (l *ListUserLogic) ListUser(req *types.UserListReq) (resp *types.UserListResp, err error) {

	// 1. 构建查询条件
	where := []predicate.User{}

	if req.Name != "" {
		where = append(where, user.NickNameContains(req.Name))
	}

	if req.Phone != "" {
		where = append(where, user.Phone(req.Phone))
	}

	if req.Status != 0 {
		where = append(where, user.StatusEQ(req.Status))
	}

	list, total, err := l.userRepo.PageList(l.ctx, req.Current, req.PageSize, where)
	if err != nil {
		l.Error("ListUser Logic PageList err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "list user page err.")
	}

	// 4. 构建返回结果
	userList := make([]*types.UserInfo, 0)
	for _, user := range list {
		userList = append(userList, &types.UserInfo{
			UserID:   user.UserID,
			UserName: user.UserName,
			Phone:    user.Phone,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return &types.UserListResp{
		List: userList,
		Page: &types.PageResponse{
			Total:           total,
			PageSize:        len(list),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}
