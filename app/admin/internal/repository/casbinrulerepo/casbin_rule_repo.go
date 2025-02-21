package casbinrulerepo

import (
	"context"

	"admin_backend/pkg/ent"
	"admin_backend/pkg/ent/generated"
)

type CasbinRuleRepo struct {
	db *ent.Client
}

// NewCasbinRuleRepo 创建 casbin 规则仓储实例
func NewCasbinRuleRepo(db *ent.Client) *CasbinRuleRepo {
	return &CasbinRuleRepo{db: db}
}

// QueryBySQL 直接执行 SQL 查询
func (r *CasbinRuleRepo) QueryBySQL(ctx context.Context, query string, args ...interface{}) ([]*generated.CasbinRule, error) {
	var rules []*generated.CasbinRule
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 扫描结果到结构体切片
	for rows.Next() {
		rule := &generated.CasbinRule{}
		err := rows.Scan(
			&rule.ID,
			&rule.Ptype,
			&rule.V0,
			&rule.V1,
			&rule.V2,
			&rule.V3,
			&rule.V4,
			&rule.V5,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	return rules, rows.Err()
}
