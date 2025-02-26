package initialize

import (
	"context"
	"fmt"
	"time"

	"admin_backend/pkg/ent/generated"
	"admin_backend/pkg/utils/idgen"
)

// InitMenus 初始化菜单数据
func InitMenus(ctx context.Context, tx *generated.Tx, tenantCode string) error {
	// 一级菜单数据
	level1Menus := []struct {
		name      string
		code      string
		path      string
		component string
		icon      string
		type_     string
		sort      int
	}{
		{name: "工作台", code: "dashboard", path: "/dashboard", component: "Layout", icon: "dashboard", type_: "dir", sort: 1},
		{name: "系统管理", code: "system", path: "/system", component: "Layout", icon: "setting", type_: "dir", sort: 2},
		{name: "商品管理", code: "product", path: "/product", component: "Layout", icon: "shopping", type_: "dir", sort: 3},
	}

	// 二级菜单数据映射
	level2Menus := map[string][]struct {
		name      string
		code      string
		path      string
		component string
		type_     string
		sort      int
	}{
		"dashboard": {
			{name: "分析页", code: "dashboard_analysis", path: "analysis", component: "/dashboard/analysis/index", type_: "menu", sort: 1},
			{name: "工作台", code: "dashboard_workplace", path: "workplace", component: "/dashboard/workplace/index", type_: "menu", sort: 2},
		},
		"system": {
			{name: "用户管理", code: "system_user", path: "user", component: "/system/user/index", type_: "menu", sort: 1},
			{name: "角色管理", code: "system_role", path: "role", component: "/system/role/index", type_: "menu", sort: 2},
			{name: "菜单管理", code: "system_menu", path: "menu", component: "/system/menu/index", type_: "menu", sort: 3},
			{name: "租户管理", code: "system_tenant", path: "tenant", component: "/system/tenant/index", type_: "menu", sort: 4},
		},
		"product": {
			{name: "商品列表", code: "product_list", path: "list", component: "/product/list/index", type_: "menu", sort: 1},
			{name: "商品分类", code: "product_category", path: "category", component: "/product/category/index", type_: "menu", sort: 2},
		},
	}

	// 三级菜单（按钮）数据映射
	level3Menus := map[string][]struct {
		name  string
		code  string
		type_ string
		sort  int
	}{
		"system_user": {
			{name: "用户查询", code: "system_user_query", type_: "button", sort: 1},
			{name: "用户创建", code: "system_user_create", type_: "button", sort: 2},
			{name: "用户编辑", code: "system_user_update", type_: "button", sort: 3},
			{name: "用户删除", code: "system_user_delete", type_: "button", sort: 4},
		},
		"system_role": {
			{name: "角色查询", code: "system_role_query", type_: "button", sort: 1},
			{name: "角色创建", code: "system_role_create", type_: "button", sort: 2},
			{name: "角色编辑", code: "system_role_update", type_: "button", sort: 3},
			{name: "角色删除", code: "system_role_delete", type_: "button", sort: 4},
		},
	}

	// 创建一级菜单
	for _, menu1 := range level1Menus {
		menuID, err := idgen.GenerateUUID()
		if err != nil {
			return fmt.Errorf("生成一级菜单ID失败: %v", err)
		}

		now := time.Now().UnixMilli()
		level1, err := tx.Menu.Create().
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetTenantCode(tenantCode).
			SetMenuID(menuID).
			SetCode(menu1.code).
			SetName(menu1.name).
			SetPath(menu1.path).
			SetComponent(menu1.component).
			SetIcon(menu1.icon).
			SetType(menu1.type_).
			SetSort(menu1.sort).
			SetStatus(1).
			SetParentID("").
			Save(ctx)

		if err != nil {
			return fmt.Errorf("创建一级菜单失败: %v", err)
		}

		// 创建二级菜单
		if level2Items, ok := level2Menus[menu1.code]; ok {
			for _, menu2 := range level2Items {
				menuID, err := idgen.GenerateUUID()
				if err != nil {
					return fmt.Errorf("生成二级菜单ID失败: %v", err)
				}

				level2, err := tx.Menu.Create().
					SetCreatedAt(now).
					SetUpdatedAt(now).
					SetTenantCode(tenantCode).
					SetMenuID(menuID).
					SetCode(menu2.code).
					SetName(menu2.name).
					SetPath(menu2.path).
					SetComponent(menu2.component).
					SetType(menu2.type_).
					SetSort(menu2.sort).
					SetStatus(1).
					SetParentID(level1.MenuID).
					Save(ctx)

				if err != nil {
					return fmt.Errorf("创建二级菜单失败: %v", err)
				}

				// 创建三级菜单（按钮）
				if level3Items, ok := level3Menus[menu2.code]; ok {
					for _, menu3 := range level3Items {
						menuID, err := idgen.GenerateUUID()
						if err != nil {
							return fmt.Errorf("生成三级菜单ID失败: %v", err)
						}

						_, err = tx.Menu.Create().
							SetCreatedAt(now).
							SetUpdatedAt(now).
							SetTenantCode(tenantCode).
							SetMenuID(menuID).
							SetCode(menu3.code).
							SetName(menu3.name).
							SetType(menu3.type_).
							SetSort(menu3.sort).
							SetStatus(1).
							SetParentID(level2.MenuID).
							Save(ctx)

						if err != nil {
							return fmt.Errorf("创建三级菜单失败: %v", err)
						}
					}
				}
			}
		}
	}

	return nil
}
