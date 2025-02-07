package tenant_repo

import (
	"context"
	"time"

	"github.com/solate/admin_backend/pkg/ent/generated"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
	"github.com/solate/admin_backend/pkg/ent/generated/tenant"
)

type TenantRepo struct {
	db *generated.Client
}

// NewTenantRepo 创建租户仓储实例
func NewTenantRepo(db *generated.Client) *TenantRepo {
	return &TenantRepo{db: db}
}

func (r *TenantRepo) Create(ctx context.Context, tenant *generated.Tenant) (*generated.Tenant, error) {
	now := time.Now().UnixMilli()
	return r.db.Tenant.Create().
		SetCreatedAt(now).
		SetUpdatedAt(now).
		SetTenantID(tenant.TenantID).
		SetName(tenant.Name).
		SetDescription(tenant.Description).
		SetStatus(tenant.Status).
		Save(ctx)
}

func (r *TenantRepo) Update(ctx context.Context, tenant *generated.Tenant) (*generated.Tenant, error) {
	now := time.Now().UnixMilli()
	tenant.UpdatedAt = now
	return r.db.Tenant.UpdateOne(tenant).Save(ctx)
}

// func (r *TenantRepo) GetByID(ctx context.Context, id uint64) (*generated.Tenant, error) {
// 	return r.db.Tenant.Query().Where(tenant.ID(id)).Only(ctx)
// }

func (r *TenantRepo) GetByTenantID(ctx context.Context, tenantID uint64) (*generated.Tenant, error) {
	return r.db.Tenant.Query().Where(tenant.TenantID(tenantID)).Only(ctx)
}

func (r *TenantRepo) PageList(ctx context.Context, offset, limit int, where []predicate.Tenant) ([]*generated.Tenant, int, error) {
	query := r.db.Tenant.Query().Where(where...).Order(generated.Desc(tenant.FieldCreatedAt))

	// 查询总数
	total, err := query.Count(ctx)
	if err != nil || total == 0 {
		return nil, 0, err
	}

	// 分页查询
	list, err := query.Offset(offset).Limit(limit).All(ctx)
	return list, total, err
}

func (r *TenantRepo) DeleteByTenantID(ctx context.Context, tenant *generated.Tenant) error {
	now := time.Now().UnixMilli()
	tenant.DeletedAt = &now
	_, err := r.Update(ctx, tenant)
	return err
}
