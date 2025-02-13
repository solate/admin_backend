package casbin

import (
	"admin_backend/pkg/ent/generated"
	"fmt"

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
func (pm *PermissionManager) AddPermissionForRole(role, domain, resource, action string) error {
	_, err := pm.enforcer.AddPolicy(role, domain, resource, action)
	return err
}

// RemovePermissionForRole 移除角色的权限
func (pm *PermissionManager) RemovePermissionForRole(role, domain, resource, action string) error {
	_, err := pm.enforcer.RemovePolicy(role, domain, resource, action)
	return err
}

// GetUserRoles 获取用户的所有角色
func (pm *PermissionManager) GetUserRoles(user, domain string) ([]string, error) {
	return pm.enforcer.GetRolesForUser(user, domain)
}

// GetRolePermissions 获取角色的所有权限
func (pm *PermissionManager) GetRolePermissions(role, domain string) ([][]string, error) {
	return pm.enforcer.GetPermissionsForUser(role, domain)
}

// CheckPermission 检查用户是否有指定权限
func (pm *PermissionManager) CheckPermission(user, domain, resource, action string) (bool, error) {
	return pm.enforcer.Enforce(user, domain, resource, action)
}

// BatchAddPermissions 批量添加权限
func (pm *PermissionManager) BatchAddPermissions(policies [][]string) error {
	_, err := pm.enforcer.AddPolicies(policies)
	return err
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

func (pm *PermissionManager) AddABACPolicy(role, domain string, attributes map[string]interface{}) error {
	policy := []string{role, domain}
	for _, v := range attributes {
		policy = append(policy, fmt.Sprintf("%v", v))
	}
	_, err := pm.enforcer.AddPolicy(policy)
	return err
}

// CheckABACPermission 检查基于属性的权限
func (pm *PermissionManager) CheckABACPermission(user string, domain string, attributes map[string]interface{}) (bool, error) {
	params := []interface{}{user, domain}
	for _, v := range attributes {
		params = append(params, v)
	}
	return pm.enforcer.Enforce(params...)
}

// 更新角色权限
func UpdateRolePermissions(pm *PermissionManager, role, domain string, permissions []Permission) error {
	// 先清除旧权限
	oldPermissions, _ := pm.GetRolePermissions(role, domain)
	pm.BatchRemovePermissions(oldPermissions)

	// 添加新权限
	var newPolicies [][]string
	for _, p := range permissions {
		newPolicies = append(newPolicies, []string{role, domain, p.Resource, p.Action})
	}
	return pm.BatchAddPermissions(newPolicies)
}
