package casbin_sdk

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewCasbin(t *testing.T) {

	dataSource := "root:123456@tcp(127.0.0.1:13306)/testdb?parseTime=true&loc=Asia%2FShanghai"
	e, err := NewCasbin(dataSource)
	if err != nil {
		t.Error(err)
		return
	}

	sub := "18811111111" // the user that wants to access a resource.
	tenant := "ldx"
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
# p, root, ldx, Dashboard, read

# Policy for admin (管理员)
p, admin, ldx, Dashboard, read
p, admin, ldx, Workspace, read
p, admin, ldx, merchantManage, read
p, admin, ldx, merchantManageUser, read
p, admin, ldx, takeawayExpert, read
p, admin, ldx, takeawayExpertAgent, read
p, admin, ldx, agentManage, read
p, admin, ldx, agentManageOne, read
p, admin, ldx, agentManageTwo, read
p, admin, ldx, fansManage, read
p, admin, ldx, fansKanban, read
p, admin, ldx, fansGroupPosting, read
p, admin, ldx, fansGroupPostingCreate, read
p, admin, ldx, fansGroupPostingUpdate, read
p, admin, ldx, enterpriseMicroEntities, read
p, admin, ldx, enterpriseMicroEntitiesFansList, read
p, admin, ldx, dataManage, read
p, admin, ldx, commissionSettlement, read
p, admin, ldx, systemManage, read
p, admin, ldx, systemManagePay, read
p, admin, ldx, logManage, read
p, admin, ldx, logManageList, read

# Policy for service (服务商)
p, service, ldx, Dashboard, read
p, service, ldx, Workspace, read
p, service, ldx, merchantManage, read
p, service, ldx, merchantManageUser, read
p, service, ldx, takeawayExpertAgentList, read
p, service, ldx, agentManage, read
p, service, ldx, agentManageTwo, read
p, service, ldx, fansManage, read
p, service, ldx, enterpriseMicroEntities, read
p, service, ldx, enterpriseMicroEntitiesFansList, read
p, service, ldx, dataManage, read
p, service, ldx, commissionSettlement, read
p, service, ldx, systemManage, read
p, service, ldx, systemManageUser, read
p, service, ldx, logManage, read
p, service, ldx, logManageList, read

# Policy for agency (代理商)
p, agency, ldx, Dashboard, read
p, agency, ldx, Workspace, read
p, agency, ldx, merchantManage, read
p, agency, ldx, merchantManageUser, read
p, agency, ldx, takeawayExpertAgentList, read
p, agency, ldx, fansManage, read
p, agency, ldx, enterpriseMicroEntities, read
p, agency, ldx, enterpriseMicroEntitiesFansList, read
p, agency, ldx, dataManage, read
p, agency, ldx, commissionSettlement, read
p, agency, ldx, systemManage, read
p, agency, ldx, systemManageUser, read
p, agency, ldx, logManage, read
p, agency, ldx, logManageList, read

# Policy for store (商家)
p, store, ldx, Dashboard, read
p, store, ldx, Workspace, read
p, store, ldx, fansManage, read
p, store, ldx, fansList, read
p, store, ldx, dataManage, read
p, store, ldx, commissionSettlement, read
p, store, ldx, systemManage, read
p, store, ldx, systemManageUser, read
p, store, ldx, logManage, read
p, store, ldx, logManageList, read

# Policy for operation (运营)
p, operation, ldx, Dashboard, read
p, operation, ldx, Workspace, read


# Grouping policy
g, 18888888888, root, ldx
g, 18811111111, admin, ldx
g, 18382131111, service, ldx
g, 18382132222, agency, ldx
g, 18382133333, store, ldx
g, 18833333333, operation, ldx
	

	`
	dataSource := "root:ldx2024..@tcp(43.136.182.87:3306)/hdz_cloud_service?parseTime=true&loc=Asia%2FShanghai"
	// dataSource := "root:123456@tcp(127.0.0.1:13306)/testdb?parseTime=true&loc=Asia%2FShanghai"
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
			// g, 18888888888, root, ldx
			s := strings.Split(line, ",")
			if len(s) != 4 {
				continue
			}
			e.AddGroupingPolicy(strings.TrimSpace(s[1]), strings.TrimSpace(s[2]), strings.TrimSpace(s[3]))
			continue
		}

		if strings.HasPrefix(line, "p") {
			// p, root, ldx, Dashboard, read
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

	dataSource := "root:123456@tcp(127.0.0.1:13306)/testdb?parseTime=true&loc=Asia%2FShanghai"
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
		{"18888888888", "ldx", "Dashboard", "read", false, "root"},
		{"18811111111", "ldx", "Dashboard", "read", true, "admin"},
		{"18382131111", "ldx", "Dashboard", "read", true, "service"},
		{"18382132222", "ldx", "Dashboard", "read", true, "agency"},
		{"18382133333", "ldx", "Dashboard", "read", true, "store"},
		{"18833333333", "ldx", "Dashboard", "read", true, "operation"},
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

	dataSource := "root:123456@tcp(127.0.0.1:13306)/testdb?parseTime=true&loc=Asia%2FShanghai"
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
		{"root", "ldx", "Workspace", "read", true, "admin"},
		{"admin", "ldx", "Workspace", "read", true, "admin"},
		{"service", "ldx", "Workspace", "read", true, "service"},
		{"agency", "ldx", "Workspace", "read", true, "agency"},
		{"store", "ldx", "Workspace", "read", true, "store"},
		{"operation", "ldx", "Workspace", "read", true, "operation"},
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
