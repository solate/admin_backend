// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package handler

import (
	"net/http"

	auth "admin_backend/app/admin/internal/handler/auth"
	menu "admin_backend/app/admin/internal/handler/menu"
	permission "admin_backend/app/admin/internal/handler/permission"
	role "admin_backend/app/admin/internal/handler/role"
	tenant "admin_backend/app/admin/internal/handler/tenant"
	user "admin_backend/app/admin/internal/handler/user"
	"admin_backend/app/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 获取验证码
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: auth.GetCaptchaHandler(serverCtx),
			},
			{
				// 用户登录
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
			{
				// 用户注册
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: auth.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/admin/api/v1/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 修改密码
					Method:  http.MethodPost,
					Path:    "/change-password",
					Handler: auth.ChangePasswordHandler(serverCtx),
				},
				{
					// 用户登出
					Method:  http.MethodPost,
					Path:    "/logout",
					Handler: auth.LogoutHandler(serverCtx),
				},
				{
					// 重置密码
					Method:  http.MethodPost,
					Path:    "/reset-password",
					Handler: auth.ResetPasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1/auth"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 创建菜单
					Method:  http.MethodPost,
					Path:    "/menus",
					Handler: menu.CreateMenuHandler(serverCtx),
				},
				{
					// 获取菜单列表
					Method:  http.MethodGet,
					Path:    "/menus",
					Handler: menu.ListMenuHandler(serverCtx),
				},
				{
					// 更新菜单
					Method:  http.MethodPut,
					Path:    "/menus/:menu_id",
					Handler: menu.UpdateMenuHandler(serverCtx),
				},
				{
					// 删除菜单
					Method:  http.MethodDelete,
					Path:    "/menus/:menu_id",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					// 获取菜单详情
					Method:  http.MethodGet,
					Path:    "/menus/:menu_id",
					Handler: menu.GetMenuHandler(serverCtx),
				},
				{
					// 获取菜单树
					Method:  http.MethodGet,
					Path:    "/menus/tree",
					Handler: menu.GetMenuTreeHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 设置角色权限
					Method:  http.MethodPost,
					Path:    "/roles/:role_id/permissions",
					Handler: permission.SetRolePermissionsHandler(serverCtx),
				},
				{
					// 获取角色权限列表
					Method:  http.MethodGet,
					Path:    "/roles/:role_id/permissions",
					Handler: permission.GetRolePermissionsHandler(serverCtx),
				},
				{
					// 获取资源类型列表
					Method:  http.MethodGet,
					Path:    "/rules/resource-types",
					Handler: permission.GetResourceTypesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 创建角色
					Method:  http.MethodPost,
					Path:    "/roles",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					// 获取角色列表
					Method:  http.MethodGet,
					Path:    "/roles",
					Handler: role.ListRoleHandler(serverCtx),
				},
				{
					// 更新角色
					Method:  http.MethodPut,
					Path:    "/roles/:role_id",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					// 删除角色
					Method:  http.MethodDelete,
					Path:    "/roles/:role_id",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					// 获取角色详情
					Method:  http.MethodGet,
					Path:    "/roles/:role_id",
					Handler: role.GetRoleHandler(serverCtx),
				},
				{
					// 设置用户角色
					Method:  http.MethodPost,
					Path:    "/users/:user_id/roles",
					Handler: role.SetUserRolesHandler(serverCtx),
				},
				{
					// 获取用户角色列表
					Method:  http.MethodGet,
					Path:    "/users/:user_id/roles",
					Handler: role.GetUserRolesHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 创建租户
					Method:  http.MethodPost,
					Path:    "/tenants",
					Handler: tenant.CreateTenantHandler(serverCtx),
				},
				{
					// 获取租户列表
					Method:  http.MethodGet,
					Path:    "/tenants",
					Handler: tenant.ListTenantHandler(serverCtx),
				},
				{
					// 更新租户
					Method:  http.MethodPut,
					Path:    "/tenants/:tenant_id",
					Handler: tenant.UpdateTenantHandler(serverCtx),
				},
				{
					// 获取租户详情
					Method:  http.MethodGet,
					Path:    "/tenants/:tenant_id",
					Handler: tenant.GetTenantHandler(serverCtx),
				},
				{
					// 删除租户
					Method:  http.MethodDelete,
					Path:    "/tenants/:tenant_id",
					Handler: tenant.DeleteTenantHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthMiddleware},
			[]rest.Route{
				{
					// 查询登录记录
					Method:  http.MethodGet,
					Path:    "/login-logs",
					Handler: user.ListLoginLogHandler(serverCtx),
				},
				{
					// 创建用户
					Method:  http.MethodPost,
					Path:    "/users",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					// 获取用户列表
					Method:  http.MethodGet,
					Path:    "/users",
					Handler: user.ListUserHandler(serverCtx),
				},
				{
					// 更新用户
					Method:  http.MethodPut,
					Path:    "/users/:user_id",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					// 删除用户
					Method:  http.MethodDelete,
					Path:    "/users/:user_id",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					// 获取用户详情
					Method:  http.MethodGet,
					Path:    "/users/:user_id",
					Handler: user.GetUserHandler(serverCtx),
				},
				{
					// 获取当前用户信息
					Method:  http.MethodGet,
					Path:    "/users/me",
					Handler: user.GetCurrentUserHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/admin/api/v1"),
	)
}
