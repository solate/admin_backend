package tenant

import (
	"context"

	"github.com/solate/admin_backend/app/admin/internal/repository/tenant_repo"
	"github.com/solate/admin_backend/app/admin/internal/svc"
	"github.com/solate/admin_backend/app/admin/internal/types"
	"github.com/solate/admin_backend/pkg/common/xerr"
	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/utils/idgen"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTenantLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	tenantRepo *tenant_repo.TenantRepo
}

// 创建租户
func NewCreateTenantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTenantLogic {
	return &CreateTenantLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		tenantRepo: tenant_repo.NewTenantRepo(svcCtx.DB),
	}
}

func (l *CreateTenantLogic) CreateTenant(req *types.CreateTenantReq) (*types.CreateTenantResp, error) {

	tenantID, err := idgen.GenerateUUID()
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "生成租户ID失败")
	}
	// 创建租户
	newTenant := &generated.Tenant{
		TenantID:    tenantID,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}
	tenant, err := l.tenantRepo.Create(l.ctx, newTenant)

	if err != nil {
		l.Error("CreateTenant Create err:", err.Error())
		return nil, xerr.NewErrCodeMsg(xerr.ServerError, "创建租户失败")
	}

	return &types.CreateTenantResp{
		TenantID: tenant.TenantID,
	}, nil
}
