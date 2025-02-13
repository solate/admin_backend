package casbin

import (
	"admin_backend/pkg/ent"
	"admin_backend/pkg/ent/generated"
	"context"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) (*generated.Client, error) {
	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	client, err := ent.NewClient(context.Background(), dataSource)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func cleanupTest(client *generated.Client) {
	if client != nil {
		client.Close()
	}
}

func TestNewCasbin(t *testing.T) {
	client, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}
	defer cleanupTest(client)

	e, err := NewCasbin(client)
	if err != nil {
		t.Error(err)
		return
	}

	// 测试添加策略
	sub := "18811111111"
	tenant := "demo"
	obj := "Dashboard"
	act := "read"

	// 添加策略
	ok, err := e.AddPolicy(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Failed to add policy")
		return
	}

	// 测试权限验证
	ok, err = e.Enforce(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Expected permission to be granted")
	}

	// 测试删除策略
	ok, err = e.RemovePolicy(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Failed to remove policy")
		return
	}

	// 验证策略已被删除
	ok, err = e.Enforce(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Error("Expected permission to be denied after removal")
	}
}

func TestRoleManagementWithTimeout(t *testing.T) {
	client, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}
	defer cleanupTest(client)

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e, err := NewCasbin(client)
	if err != nil {
		t.Error(err)
		return
	}

	// 测试角色管理
	user := "user123"
	role := "admin"
	tenant := "demo"

	// 添加角色
	ok, err := e.AddRoleForUser(user, role, tenant)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Failed to add role for user")
		return
	}

	// 验证角色是否添加成功
	roles, err := e.GetRolesForUser(user, tenant)
	if err != nil {
		t.Error(err)
		return
	}
	if len(roles) != 1 || roles[0] != role {
		t.Error("Expected role not found")
		return
	}

	// 删除角色
	ok, err = e.DeleteRoleForUser(user, role, tenant)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Failed to remove role for user")
		return
	}

	// 验证角色是否删除成功
	roles, err = e.GetRolesForUser(user, tenant)
	if err != nil {
		t.Error(err)
		return
	}
	if len(roles) != 0 {
		t.Error("Expected no roles after removal")
	}
}

func TestSuperAdminPermissionsAndEdgeCases(t *testing.T) {
	client, err := setupTestDB(t)
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}
	defer cleanupTest(client)

	e, err := NewCasbin(client)
	if err != nil {
		t.Error(err)
		return
	}

	// 测试root用户权限
	sub := "root"
	tenant := "demo"
	obj := "AnyResource"
	act := "any"

	// 测试root用户权限
	ok, err := e.Enforce(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Expected root user to have all permissions")
	}

	// 测试空用户名
	ok, err = e.Enforce("", tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Error("Empty username should not have permissions")
	}

	// 测试空租户 - root用户应该在任何租户下都有权限
	ok, err = e.Enforce(sub, "", obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Error("Root user should have permissions even with empty tenant")
	}

	// 测试普通用户在空租户下的权限
	ok, err = e.Enforce("normal_user", "", obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Error("Normal user should not have permissions with empty tenant")
	}

	// 测试特殊字符
	ok, err = e.Enforce("user!@#$%", tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}
	if ok {
		t.Error("User with special characters should not have permissions")
	}
}
