package casbin

import (
	"admin_backend/pkg/ent"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestEnv(t *testing.T) *CasbinManager {
	// 初始化测试数据库连接
	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	client, err := ent.NewClient(context.Background(), dataSource)
	if err != nil {
		t.Fatalf("failed creating permission manager: %v", err)
	}
	// 创建权限管理器
	pm := NewCasbinManager(client)

	return pm
}

func TestAddRoleForUser(t *testing.T) {
	pm := setupTestEnv(t)

	// 测试正常场景
	err := pm.AddRoleForUser("user1", "admin", "domain1")
	require.NoError(t, err)

	// 验证角色是否添加成功
	roles, err := pm.GetRolesForUser("user1", "domain1")
	t.Log(roles, "========")
	require.NoError(t, err)
	assert.Contains(t, roles, "admin")
}

func TestRemoveRoleForUser(t *testing.T) {
	pm := setupTestEnv(t)

	// 准备测试数据
	err := pm.AddRoleForUser("user1", "admin", "domain1")
	require.NoError(t, err)

	// 测试移除角色
	err = pm.RemoveRoleForUser("user1", "admin", "domain1")
	require.NoError(t, err)

	// 验证角色是否已移除
	roles, err := pm.GetRolesForUser("user1", "domain1")
	require.NoError(t, err)
	assert.NotContains(t, roles, "admin")
}

func TestAddPermissionForRole(t *testing.T) {
	pm := setupTestEnv(t)

	// 测试添加权限
	err := pm.AddPermissionForRole("admin", "default", "apiPath", "GET", "api")
	require.NoError(t, err)

	// 验证权限是否添加成功
	perms, err := pm.GetRolePermissions("admin", "default")
	require.NoError(t, err)
	assert.Contains(t, perms, []string{"admin", "default", "apiPath", "GET", "api"})
}

func TestCheckPermission(t *testing.T) {
	pm := setupTestEnv(t)

	// 准备测试数据
	_ = pm.AddRoleForUser("user1", "admin", "default")
	_ = pm.AddPermissionForRole("admin", "default", "apiPath", "GET", "api")

	// 测试有权限的场景
	hasPermission, err := pm.CheckPermission("user1", "default", "apiPath", "GET", "api")
	require.NoError(t, err)
	assert.True(t, hasPermission)

	// 测试无权限的场景
	hasPermission, err = pm.CheckPermission("user1", "default", "otherPath", "GET", "api")
	require.NoError(t, err)
	assert.False(t, hasPermission)
}

func TestBatchAddPermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// Clear any existing policies first
	pm.enforcer.ClearPolicy()

	// 准备测试数据
	policies := [][]string{
		{"admin", "default", "menuCode", "view", "menu"},
		{"admin", "default", "pageCode", "view", "page"},
	}

	// 测试批量添加权限
	err := pm.BatchAddPermissions(policies)
	require.NoError(t, err)

	// Force reload policies from storage
	err = pm.enforcer.LoadPolicy()
	require.NoError(t, err)

	// 验证权限是否都添加成功
	perms, err := pm.GetRolePermissions("admin", "default") // 修改这里
	require.NoError(t, err)

	// Simplified verification logic
	for _, policy := range policies {
		assert.Contains(t, perms, policy, "Policy should exist: %v", policy)
	}
}

func TestClearUserPermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// 准备测试数据
	_ = pm.AddRoleForUser("user1", "admin", "domain1")
	_ = pm.AddRoleForUser("user1", "editor", "domain1")

	// 测试清除权限
	err := pm.ClearUserPermissions("user1", "domain1")
	require.NoError(t, err)

	// 验证权限是否已清除
	roles, err := pm.GetRolesForUser("user1", "domain1")
	require.NoError(t, err)
	assert.Empty(t, roles)
}

// func TestAddMenuPermission(t *testing.T) {
// 	pm := setupTestEnv(t)

// 	tests := []struct {
// 		name       string
// 		role       string
// 		tenantCode string
// 		menuCode   string
// 		wantErr    bool
// 	}{
// 		{
// 			name:       "正常添加菜单权限",
// 			role:       "admin",
// 			tenantCode: "default",
// 			menuCode:   "menuCode",
// 			wantErr:    false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// 执行添加菜单权限
// 			err := pm.AddMenuPermission(tt.role, tt.tenantCode, tt.menuCode)

// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}
// 			assert.NoError(t, err)

// 			// 验证权限是否添加成功
// 			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
// 			require.NoError(t, err)

// 			expectedPolicy := []string{tt.role, tt.tenantCode, tt.menuCode, ActionView, PermTypeMenu}
// 			assert.Contains(t, perms, expectedPolicy, "Menu permission should exist")

// 			// 验证权限是否可以正确检查
// 			hasPermission, err := pm.CheckPermission(tt.role, tt.tenantCode, tt.menuCode, ActionView, PermTypeMenu)
// 			require.NoError(t, err)
// 			assert.True(t, hasPermission, "Should have menu permission")
// 		})
// 	}
// }

// func TestAddPagePermission(t *testing.T) {
// 	pm := setupTestEnv(t)

// 	tests := []struct {
// 		name       string
// 		role       string
// 		tenantCode string
// 		pageCode   string
// 		wantErr    bool
// 	}{
// 		{
// 			name:       "正常添加页面权限",
// 			role:       "admin",
// 			tenantCode: "tenant1",
// 			pageCode:   "page1",
// 			wantErr:    false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := pm.AddPagePermission(tt.role, tt.tenantCode, tt.pageCode)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}
// 			assert.NoError(t, err)

// 			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
// 			require.NoError(t, err)
// 			expectedPolicy := []string{tt.role, tt.tenantCode, tt.pageCode, ActionView, PermTypePage}
// 			assert.Contains(t, perms, expectedPolicy)
// 		})
// 	}
// }

// func TestAddButtonPermission(t *testing.T) {
// 	pm := setupTestEnv(t)

// 	tests := []struct {
// 		name         string
// 		role         string
// 		tenantCode   string
// 		resourceCode string
// 		action       string
// 		wantErr      bool
// 	}{
// 		{
// 			name:         "正常添加按钮权限",
// 			role:         "admin",
// 			tenantCode:   "tenant1",
// 			resourceCode: "addUser",
// 			action:       "add",
// 			wantErr:      false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := pm.AddButtonPermission(tt.role, tt.tenantCode, tt.resourceCode, tt.action)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}
// 			assert.NoError(t, err)

// 			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
// 			require.NoError(t, err)
// 			expectedPolicy := []string{tt.role, tt.tenantCode, tt.resourceCode, tt.action, PermTypeButton}
// 			assert.Contains(t, perms, expectedPolicy)
// 		})
// 	}
// }

// func TestAddAPIPermission(t *testing.T) {
// 	pm := setupTestEnv(t)

// 	tests := []struct {
// 		name       string
// 		role       string
// 		tenantCode string
// 		apiPath    string
// 		method     string
// 		wantErr    bool
// 	}{
// 		{
// 			name:       "正常添加API权限",
// 			role:       "admin",
// 			tenantCode: "tenant1",
// 			apiPath:    "/api/users",
// 			method:     "GET",
// 			wantErr:    false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := pm.AddAPIPermission(tt.role, tt.tenantCode, tt.apiPath, tt.method)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}
// 			assert.NoError(t, err)

// 			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
// 			require.NoError(t, err)
// 			expectedPolicy := []string{tt.role, tt.tenantCode, tt.apiPath, tt.method, PermTypeAPI}
// 			assert.Contains(t, perms, expectedPolicy)
// 		})
// 	}
// }

// func TestAddDataPermission(t *testing.T) {
// 	pm := setupTestEnv(t)

// 	tests := []struct {
// 		name         string
// 		role         string
// 		tenantCode   string
// 		resourceCode string
// 		rule         string
// 		wantErr      bool
// 	}{
// 		{
// 			name:         "正常添加数据权限",
// 			role:         "admin",
// 			tenantCode:   "tenant1",
// 			resourceCode: "users",
// 			rule:         "department",
// 			wantErr:      false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := pm.AddDataPermission(tt.role, tt.tenantCode, tt.resourceCode, tt.rule)
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				return
// 			}
// 			assert.NoError(t, err)

// 			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
// 			require.NoError(t, err)
// 			expectedPolicy := []string{tt.role, tt.tenantCode, tt.resourceCode, tt.rule, PermTypeData}
// 			assert.Contains(t, perms, expectedPolicy)
// 		})
// 	}
// }

func TestBatchAddDataPermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// 清除现有策略
	pm.enforcer.ClearPolicy()

	tests := []struct {
		name       string
		role       string
		tenantCode string
		rules      []DataPermissionRule
		wantErr    bool
	}{
		{
			name:       "正常批量添加数据权限",
			role:       "admin",
			tenantCode: "tenant1",
			rules: []DataPermissionRule{
				{
					Resource: "users",
					Rule:     "department", // 使用预定义的数据规则
				},
				{
					Resource: "orders",
					Rule:     "all", // 使用预定义的数据规则
				},
				{
					Resource: "products",
					Rule:     "self", // 使用预定义的数据规则
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清除之前的测试数据
			pm.enforcer.ClearPolicy()

			err := pm.BatchAddDataPermissions(tt.role, tt.tenantCode, tt.rules)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			// 强制重新加载策略
			err = pm.enforcer.LoadPolicy()
			require.NoError(t, err)

			perms, err := pm.GetRolePermissions(tt.role, tt.tenantCode)
			require.NoError(t, err)

			// 验证每条规则是否都正确添加
			for _, rule := range tt.rules {
				expectedPolicy := []string{
					tt.role,
					tt.tenantCode,
					rule.Resource,
					rule.Rule,
					PermTypeData,
				}
				assert.Contains(t, perms, expectedPolicy,
					"Policy should exist: %v", expectedPolicy)
			}
		})
	}
}

func TestBatchCheckDataPermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// 清除现有策略
	pm.enforcer.ClearPolicy()

	// 准备测试数据
	role := "admin"
	user := "user1"
	tenantCode := "tenant1"
	rules := []DataPermissionRule{
		{
			Resource: "users",
			Rule:     "department",
		},
		{
			Resource: "orders",
			Rule:     "all",
		},
		{
			Resource: "products",
			Rule:     "self",
		},
	}

	// 添加角色和权限
	err := pm.AddRoleForUser(user, role, tenantCode)
	require.NoError(t, err)
	err = pm.BatchAddDataPermissions(role, tenantCode, rules)
	require.NoError(t, err)

	// 强制重新加载策略
	err = pm.enforcer.LoadPolicy()
	require.NoError(t, err)

	tests := []struct {
		name       string
		user       string
		tenantCode string
		rules      []DataPermissionRule
		wantErr    bool
	}{
		{
			name:       "正常检查数据权限",
			user:       user,
			tenantCode: tenantCode,
			rules:      rules,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := pm.BatchCheckDataPermissions(tt.user, tt.tenantCode, tt.rules)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, results)

			// 验证每个资源的权限结果
			for _, rule := range tt.rules {
				hasPermission, exists := results[rule.Resource]
				assert.True(t, exists, "权限检查结果应该存在")
				assert.True(t, hasPermission, "应该具有权限")
			}
		})
	}
}

func TestGetUserAllPermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// 清除现有策略
	pm.enforcer.ClearPolicy()

	tests := []struct {
		name      string
		user      string
		tenant    string
		setupFunc func()
		wantPerms [][]string
		wantErr   bool
	}{
		{
			name:   "用户有多个角色和权限",
			user:   "user1",
			tenant: "tenant1",
			setupFunc: func() {
				// 清除所有策略（包括角色和权限）
				pm.enforcer.ClearPolicy()

				// 添加角色和权限
				pm.AddRoleForUser("user1", "admin", "tenant1")
				pm.AddRoleForUser("user1", "editor", "tenant1")

				// 只添加 API 权限
				policies := [][]string{
					{"admin", "tenant1", "/api/users", "GET", "api"},
					{"editor", "tenant1", "/api/posts", "POST", "api"},
				}
				pm.BatchAddPermissions(policies)

				// 重新加载策略确保生效
				pm.enforcer.LoadPolicy()
			},
			wantPerms: [][]string{
				{"admin", "tenant1", "/api/users", "GET", "api"},
				{"editor", "tenant1", "/api/posts", "POST", "api"},
			},
			wantErr: false,
		},
		{
			name:   "用户没有角色",
			user:   "user2",
			tenant: "tenant1",
			setupFunc: func() {
				pm.enforcer.ClearPolicy()
				pm.enforcer.LoadPolicy()
			},
			wantPerms: [][]string{},
			wantErr:   false,
		},
		{
			name:   "用户有角色但没有权限",
			user:   "user3",
			tenant: "tenant1",
			setupFunc: func() {
				pm.enforcer.ClearPolicy()
				pm.AddRoleForUser("user3", "viewer", "tenant1")
				pm.enforcer.LoadPolicy()
			},
			wantPerms: [][]string{},
			wantErr:   false,
		},
		{
			name:   "用户有多个角色相同权限",
			user:   "user4",
			tenant: "tenant1",
			setupFunc: func() {
				pm.enforcer.ClearPolicy()
				pm.AddRoleForUser("user4", "role1", "tenant1")
				pm.AddRoleForUser("user4", "role2", "tenant1")
				pm.AddPermissionForRole("role1", "tenant1", "/api/common", "GET", "api")
				pm.AddPermissionForRole("role2", "tenant1", "/api/common", "GET", "api")
				pm.enforcer.LoadPolicy()
			},
			wantPerms: [][]string{
				{"role1", "tenant1", "/api/common", "GET", "api"},
				{"role2", "tenant1", "/api/common", "GET", "api"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清除之前的所有策略
			pm.enforcer.ClearPolicy()

			// 设置测试数据
			if tt.setupFunc != nil {
				tt.setupFunc()
			}

			// 执行测试
			gotPerms, err := pm.GetUserAllPermissions(tt.user, tt.tenant)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)

			// 验证权限列表
			if len(tt.wantPerms) == 0 {
				assert.Empty(t, gotPerms)
			} else {
				// 验证权限数量和内容
				assert.ElementsMatch(t, tt.wantPerms, gotPerms,
					"权限列表应完全匹配")
			}
		})
	}
}

func TestUpdateRolePermissions(t *testing.T) {
	pm := setupTestEnv(t)

	// 清除现有策略
	pm.enforcer.ClearPolicy()

	tests := []struct {
		name        string
		role        string
		domain      string
		oldPerms    [][]string
		permissions []Permission
		wantErr     bool
	}{
		{
			name:   "正常更新角色权限",
			role:   "admin",
			domain: "tenant1",
			oldPerms: [][]string{
				{"admin", "tenant1", "/api/users", "GET", "api"},
				{"admin", "tenant1", "/api/roles", "POST", "api"},
			},
			permissions: []Permission{
				{Resource: "/api/users", Action: "POST", Type: "api"},
				{Resource: "/api/roles", Action: "GET", Type: "api"},
			},
			wantErr: false,
		},
		{
			name:   "清除所有权限",
			role:   "admin",
			domain: "tenant1",
			oldPerms: [][]string{
				{"admin", "tenant1", "/api/users", "GET", "api"},
			},
			permissions: []Permission{},
			wantErr:     false,
		},
		{
			name:     "添加新权限到空角色",
			role:     "editor",
			domain:   "tenant1",
			oldPerms: [][]string{},
			permissions: []Permission{
				{Resource: "/api/posts", Action: "POST", Type: "api"},
			},
			wantErr: false,
		},
		{
			name:   "更新多种类型的权限",
			role:   "admin",
			domain: "tenant1",
			oldPerms: [][]string{
				{"admin", "tenant1", "menu1", "view", "menu"},
				{"admin", "tenant1", "/api/users", "GET", "api"},
			},
			permissions: []Permission{
				{Resource: "menu2", Action: "view", Type: "menu"},
				{Resource: "/api/roles", Action: "POST", Type: "api"},
				{Resource: "button1", Action: "click", Type: "button"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清除之前的测试数据
			pm.enforcer.ClearPolicy()

			// 添加旧权限
			if len(tt.oldPerms) > 0 {
				err := pm.BatchAddPermissions(tt.oldPerms)
				require.NoError(t, err)
			}

			// 执行权限更新
			err := pm.UpdateRolePermissions(tt.role, tt.domain, tt.permissions)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			// 强制重新加载策略
			err = pm.enforcer.LoadPolicy()
			require.NoError(t, err)

			// 验证新权限
			perms, err := pm.GetRolePermissions(tt.role, tt.domain)
			require.NoError(t, err)

			// 验证旧权限已被移除
			for _, oldPerm := range tt.oldPerms {
				assert.NotContains(t, perms, oldPerm, "旧权限应该已被移除")
			}

			// 验证新权限已被添加
			expectedPolicies := make([][]string, 0)
			for _, p := range tt.permissions {
				expectedPolicy := []string{tt.role, tt.domain, p.Resource, p.Action, p.Type}
				expectedPolicies = append(expectedPolicies, expectedPolicy)
			}

			// 验证权限数量和内容
			assert.ElementsMatch(t, expectedPolicies, perms, "新权限列表应完全匹配")
		})
	}
}
