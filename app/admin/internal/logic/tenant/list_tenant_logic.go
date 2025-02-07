package tenant

import (
	"context"

	"github.com/solate/admin_backend/app/admin/internal/repository/tenant_repo"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
	"github.com/solate/admin_backend/pkg/ent/generated/tenant"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTenantLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	tenantRepo *tenant_repo.TenantRepo
}

// 获取租户列表
func NewListTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTenantLogic {
	return &ListTenantLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		tenantRepo: tenant_repo.NewTenantRepo(svcCtx.DB),
	}
}

func (l *ListTenantLogic) ListTenant(req *types.ListTenantReq) (*types.ListTenantResp, error) {
	// 1. 构建查询条件
	where := []predicate.Tenant{
		tenant.DeletedAtIsNil(), // 未删除的租户
	}
	// 3. 状态筛选
	if req.Status > 0 {
		where = append(where, tenant.Status(req.Status))
	}

	// 4. 分页查询
	offset := common.Offset(req.Current, req.PageSize)
	list, total, err := l.tenantRepo.PageList(l.ctx, offset, req.PageSize, where)
	if err != nil {
		l.Error("ListTenant PageList err:", err.Error())
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// 5. 构建返回结果
	tenantList := make([]*types.TenantInfo, 0)
	for _, t := range list {
		tenantList = append(tenantList, &types.TenantInfo{
			TenantID:    t.TenantID,
			Name:        t.Name,
			Description: t.Description,
			Status:      t.Status,
		})
	}

	return &types.ListTenantResp{
		List: tenantList,
		Page: &types.PageResponse{
			Total:           total,
			PageSize:        len(list),
			Current:         req.Current,
			RequestPageSize: req.PageSize,
		},
	}, nil
}
