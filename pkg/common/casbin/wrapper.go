package casbin

import (
	"admin_backend/pkg/ent/generated"

	"github.com/casbin/casbin/v2"
)

// PermissionManager 权限管理器
type PermissionManager struct {
	enforcer *casbin.Enforcer
}

// NewPermissionManager 创建权限管理器实例
func NewPermissionManager(db *generated.Client) (*PermissionManager, error) {
	e, err := NewCasbin(db)
	if err != nil {
		return nil, err
	}
	return &PermissionManager{enforcer: e}, nil
}

// AddRoleForUser 为用户分配角色
func (pm *PermissionManager) AddRoleForUser(user, role, domain string) error {
	_, err := pm.enforcer.AddGroupingPolicy(user, role, domain)
	return err
}

// RemoveRoleForUser 移除用户的角色
func (pm *PermissionManager) RemoveRoleForUser(user, role, domain string) error {
	_, err := pm.enforcer.RemoveGroupingPolicy(user, role, domain)
	return err
}

// AddPermissionForRole 为角色添加权限
func (pm *PermissionManager) AddPermissionForRole(role, domain, resource, action, tType string) error {
	_, err := pm.enforcer.AddPolicy(role, domain, resource, action, tType)
	return err
}

// RemovePermissionForRole 移除角色的权限
func (pm *PermissionManager) RemovePermissionForRole(role, domain, resource, action, tType string) error {
	_, err := pm.enforcer.RemovePolicy(role, domain, resource, action, tType)
	return err
}

// GetUserRoles 获取用户的所有角色
func (pm *PermissionManager) GetUserRoles(user, domain string) ([]string, error) {
	return pm.enforcer.GetRolesForUser(user, domain)
}

// GetRolePermissions 获取角色的所有权限
func (pm *PermissionManager) GetRolePermissions(role, domain string) ([][]string, error) {
	// 使用 GetFilteredPolicy 获取指定角色和域的所有权限
	return pm.enforcer.GetFilteredPolicy(0, role, domain)
}

// BatchAddPermissions 批量添加权限
// BatchAddPermissions 批量添加权限
func (pm *PermissionManager) BatchAddPermissions(policies [][]string) error {
    // 批量添加新策略
    _, err := pm.enforcer.AddPolicies(policies)
    if err != nil {
        return err
    }

    // 保存策略到存储
    return pm.enforcer.SavePolicy()
}

// AddMenuPermission 添加菜单权限
func (pm *PermissionManager) AddMenuPermission(role, tenantCode, menuCode string) error {
	return pm.AddPermissionForRole(role, tenantCode, menuCode, ActionView, PermTypeMenu)
}

// AddPagePermission 添加页面权限
func (pm *PermissionManager) AddPagePermission(role, tenantCode, pageCode string) error {
	return pm.AddPermissionForRole(role, tenantCode, pageCode, ActionView, PermTypePage)
}

// AddButtonPermission 添加按钮权限
func (pm *PermissionManager) AddButtonPermission(role, tenantCode, resourceCode, action string) error {
	return pm.AddPermissionForRole(role, tenantCode, resourceCode, action, PermTypeButton)
}

// AddAPIPermission 添加API权限
func (pm *PermissionManager) AddAPIPermission(role, tenantCode, apiPath, method string) error {
	return pm.AddPermissionForRole(role, tenantCode, apiPath, method, PermTypeAPI)
}

// AddDataPermission 添加数据权限
func (pm *PermissionManager) AddDataPermission(role, tenantCode, resourceCode, rule string) error {
	return pm.AddPermissionForRole(role, tenantCode, resourceCode, rule, PermTypeData)
}

// GetUserAllPermissions 获取用户所有权限
func (pm *PermissionManager) GetUserAllPermissions(user, tenant string) ([][]string, error) {
	// 直接使用 enforcer 的 GetImplicitPermissionsForUser 方法获取所有权限
	return pm.enforcer.GetImplicitPermissionsForUser(user, tenant)
}

// CheckPermission 检查用户是否有指定权限
func (pm *PermissionManager) CheckPermission(user, domain, resource, action, tType string) (bool, error) {
	return pm.enforcer.Enforce(user, domain, resource, action, tType)
}

// BatchRemovePermissions 批量移除权限
func (pm *PermissionManager) BatchRemovePermissions(policies [][]string) error {
	_, err := pm.enforcer.RemovePolicies(policies)
	return err
}

// ClearUserPermissions 清除用户所有权限
func (pm *PermissionManager) ClearUserPermissions(user, domain string) error {
	roles, err := pm.GetUserRoles(user, domain)
	if err != nil {
		return err
	}

	for _, role := range roles {
		_, err = pm.enforcer.RemoveGroupingPolicy(user, role, domain)
		if err != nil {
			return err
		}
	}
	return nil
}

// BatchAddDataPermissions 批量添加数据权限
func (pm *PermissionManager) BatchAddDataPermissions(role, tenantCode string, rules []DataPermissionRule) error {
	var policies [][]string
	for _, rule := range rules {
		// 简化策略结构：role, domain, resource, rule, type
		policy := []string{
			role,
			tenantCode,
			rule.Resource,
			rule.Rule,
			PermTypeData,
		}
		policies = append(policies, policy)
	}
	return pm.BatchAddPermissions(policies)
}

// BatchCheckDataPermissions 批量检查数据权限
func (pm *PermissionManager) BatchCheckDataPermissions(user, tenantCode string, rules []DataPermissionRule) (map[string]bool, error) {
	results := make(map[string]bool)
	for _, rule := range rules {
		// 检查权限时使用统一的参数结构
		hasPermission, err := pm.enforcer.Enforce(
			user,
			tenantCode,
			rule.Resource,
			rule.Rule,
			PermTypeData,
		)
		if err != nil {
			return nil, err
		}
		results[rule.Resource] = hasPermission
	}
	return results, nil
}

// UpdateRolePermissions 更新角色权限
func (pm *PermissionManager) UpdateRolePermissions(role, domain string, permissions []Permission) error {
	// 先清除旧权限
	oldPermissions, err := pm.GetRolePermissions(role, domain)
	if err != nil {
		return err
	}

	if len(oldPermissions) > 0 {
		err = pm.BatchRemovePermissions(oldPermissions)
		if err != nil {
			return err
		}
	}

	// 添加新权限
	var newPolicies [][]string
	for _, p := range permissions {
		newPolicies = append(newPolicies, []string{
			role,
			domain,
			p.Resource,
			p.Action,
			p.Type, // 添加权限类型
		})
	}

	if len(newPolicies) > 0 {
		return pm.BatchAddPermissions(newPolicies)
	}
	return nil
}
