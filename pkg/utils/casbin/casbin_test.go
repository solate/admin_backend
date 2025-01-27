package casbin_sdk

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewCasbin(t *testing.T) {

	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	e, err := NewCasbin(dataSource)
	if err != nil {
		t.Error(err)
		return
	}

	sub := "18811111111" // the user that wants to access a resource.
	tenant := "demo"
	obj := "Dashboard" // the resource that is going to be accessed.
	act := "read"      // the operation that the user performs on the resource.

	// e.AddPolicy(sub, tenant, obj, act)

	ok, err := e.Enforce(sub, tenant, obj, act)
	if err != nil {
		t.Error(err)
		return
	}

	if ok == true {
		t.Log("ok") // true
	} else {
		t.Log("not ok")
	}

	// 根据用户获取权限
	xx, err := e.GetPermissionsForUser(sub, tenant)
	fmt.Println(xx, err)

	fmt.Println("-------")

	// admin
	admin := "admin"
	// 根据角色获取权限
	xx, err = e.GetPermissionsForUser(admin, tenant)
	fmt.Println(xx, err)

}

// 测试初始化权限
func TestInitCasbin(t *testing.T) {

	str := `
# Policy for root (超级管理员)
p, root, demo, Dashboard, read

# Policy for admin (管理员)
p, admin, demo, Dashboard, read
p, admin, demo, Workspace, read

	`
	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	e, err := NewCasbin(dataSource)
	if err != nil {
		t.Error(err)
		return
	}

	// 按行读取string
	for _, line := range strings.Split(str, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		// Grouping policy
		if strings.HasPrefix(line, "g") {
			s := strings.Split(line, ",")
			if len(s) != 4 {
				continue
			}
			e.AddGroupingPolicy(strings.TrimSpace(s[1]), strings.TrimSpace(s[2]), strings.TrimSpace(s[3]))
			continue
		}

		if strings.HasPrefix(line, "p") {
			s := strings.Split(line, ",")
			if len(s) != 5 {
				continue
			}
			e.AddPolicy(strings.TrimSpace(s[1]), strings.TrimSpace(s[2]), strings.TrimSpace(s[3]), strings.TrimSpace(s[4]))
			continue
		}

	}

	e.SavePolicy()

}

// 测试用户权限
func TestEnforce(t *testing.T) {

	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	e, err := NewCasbin(dataSource)
	if err != nil {
		t.Error(err)
		return
	}

	// 批量测试用例
	testCases := []struct {
		sub    string
		tenant string
		obj    string
		act    string
		expect bool
		reason string
	}{
		{"root", "bbb", "xxx", "read", true, "admin"},

		// 添加更多测试用例
	}
	// roles, _ := e.GetRolesForUser(sub, tenant)
	// var role string
	// if len(roles) > 0 {
	// 	role = roles[0]
	// }

	for _, tc := range testCases {
		ok, reson, err := e.EnforceEx(tc.sub, tc.tenant, tc.obj, tc.act)
		if err != nil {
			t.Errorf("Error for sub: %s, tenant: %s, obj: %s, act: %s - %v", tc.sub, tc.tenant, tc.obj, tc.act, err)
			continue
		}

		if ok != tc.expect {
			t.Errorf("Expected %v for sub: %s, tenant: %s, obj: %s, act: %s, but got %v", tc.expect, tc.sub, tc.tenant, tc.obj, tc.act, ok)
		}

		if tc.expect && len(reson) > 0 && reson[0] != tc.reason {
			t.Errorf("Expected reason %s for sub: %s, tenant: %s, obj: %s, act: %s, but got %s", tc.reason, tc.sub, tc.tenant, tc.obj, tc.act, reson[0])
		}
	}

}

// 测试角色权限
func TestEnforceRole(t *testing.T) {

	dataSource := "user=root password=root host=127.0.0.1 port=5432 dbname=testdb sslmode=disable"
	e, err := NewCasbin(dataSource)
	if err != nil {
		t.Error(err)
		return
	}

	// 批量测试用例
	testCases := []struct {
		sub    string
		tenant string
		obj    string
		act    string
		expect bool
		reason string
	}{
		// 角色测试用例
		{"root", "demo", "Workspace", "read", true, "admin"},
	}

	for _, tc := range testCases {
		ok, reson, err := e.EnforceEx(tc.sub, tc.tenant, tc.obj, tc.act)
		if err != nil {
			t.Errorf("Error for sub: %s, tenant: %s, obj: %s, act: %s - %v", tc.sub, tc.tenant, tc.obj, tc.act, err)
			continue
		}

		if ok != tc.expect {
			t.Errorf("Expected %v for sub: %s, tenant: %s, obj: %s, act: %s, but got %v", tc.expect, tc.sub, tc.tenant, tc.obj, tc.act, ok)
		}

		if tc.expect && len(reson) > 0 && reson[0] != tc.reason {
			t.Errorf("Expected reason %s for sub: %s, tenant: %s, obj: %s, act: %s, but got %s", tc.reason, tc.sub, tc.tenant, tc.obj, tc.act, reson[0])
		}
	}

}
