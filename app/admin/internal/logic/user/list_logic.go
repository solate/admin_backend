package user

import (
	"context"

	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
	"github.com/solate/admin_backend/pkg/ent/generated/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.UserListReq) (resp *types.UserListResp, err error) {

	tenantID := l.ctx.Value("tenantID").(string)
	// 1. 构建查询条件
	where := []predicate.User{
		user.DeletedAt(0), // 未删除的用户
		user.TenantCode(tenantID),
	}

	list, total, err := l.PageList(req.Current, req.PageSize, where)
	if err != nil {
		l.Error("list agency page err: ", err)
		return nil, xerr.NewErrMsg("list agency page err.")
	}

	// 4. 构建返回结果
	userList := make([]*types.UserInfoResponse, 0)
	for _, user := range list {
		userList = append(userList, &types.UserInfoResponse{
			UserID: user.ID,
			Name:   user.UserName,
			Phone:  user.Phone,
			Email:  user.Email,
			Status: user.Status,
		})
	}

	return &types.UserListResp{
		List: userList,
		Page: &types.PageResponse{
			Total:           int32(total),
			PageSize:        int32(len(list)),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}

func (s *ListLogic) PageList(current, pageSize int32, where []predicate.User) (list []*generated.User, total int, err error) {
	offset := common.Offset(current, pageSize)

	query := s.svcCtx.Orm.User.Query().Where(where...).Order(generated.Desc(user.FieldCreatedAt))

	// 2. 查询总数
	total, err = query.Count(s.ctx)
	if err != nil || total == 0 {
		return
	}

	// 3. 分页查询
	list, err = query.Offset(int(offset)).Limit(int(pageSize)).All(s.ctx)

	return
}
