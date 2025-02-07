package tenant

import (
	"context"

	"github.com/solate/admin_backend/app/admin/internal/repository/tenant_repo"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTenantLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	tenantRepo *tenant_repo.TenantRepo
}

// 获取租户详情
func NewGetTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTenantLogic {
	return &GetTenantLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		tenantRepo: tenant_repo.NewTenantRepo(svcCtx.DB),
	}
}

func (l *GetTenantLogic) GetTenant(req *types.GetTenantReq) (*types.GetTenantResp, error) {

	// 获取租户信息
	tenant, err := l.tenantRepo.GetByTenantID(l.ctx, req.TenantID)
	if err != nil {
		l.Error("GetTenant l.tenantRepo.GetByTenantID err: ", err)
		return nil, xerr.NewErrCodeMsg(xerr.DbError, "获取租户信息失败")
	}

	return &types.GetTenantResp{
		TenantInfo: types.TenantInfo{
			TenantID:    tenant.TenantID,
			Name:        tenant.Name,
			Description: tenant.Description,
			Status:      tenant.Status,
		},
	}, nil
}
