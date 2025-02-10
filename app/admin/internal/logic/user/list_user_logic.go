package user

import (
	"context"

	"admin_backend/app/admin/internal/repository/user_repo"
	"admin_backend/app/admin/internal/svc"
	"admin_backend/app/admin/internal/types"
	"admin_backend/pkg/common"
	"admin_backend/pkg/common/context_util"
	"admin_backend/pkg/common/xerr"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	userRepo *user_repo.UserRepo
}

// 获取用户列表
func NewListUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserLogic {
	return &ListUserLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		userRepo: user_repo.NewUserRepo(svcCtx.DB),
	}
}

func (l *ListUserLogic) ListUser(req *types.UserListReq) (resp *types.UserListResp, err error) {

	tenantCode, err := context_util.GetTenantCodeFromCtx(l.ctx)
	if err != nil {
		l.Error("ListUser context_util.GetTenantIDFromCtx err: ", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "get tenant id from ctx err.")
	}

	// 1. 构建查询条件
	where := []predicate.User{
		user.DeletedAtIsNil(), // 未删除的用户
		user.TenantCode(tenantCode),
	}

	offset := common.Offset(req.Current, req.PageSize)
	list, total, err := l.userRepo.PageList(l.ctx, offset, req.PageSize, where)
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
