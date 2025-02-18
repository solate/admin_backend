// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type CaptchaResp struct {
	CaptchaId  string `json:"captcha_id"`  // 验证码ID
	CaptchaUrl string `json:"captcha_url"` // 验证码图片（base64）
}

type ChangePasswordReq struct {
	OldPassword string `json:"old_password" validate:"required"` // 原密码
	NewPassword string `json:"new_password" validate:"required"` // 新密码
}

type CreateMenuReq struct {
	Code      string `json:"code"`               // 菜单code
	Action    string `json:"action"`             // 权限标识
	Type      string `json:"type"`               // 菜单类型
	ParentID  string `json:"parent_id"`          // 父菜单ID
	Name      string `json:"name"`               // 菜单名称
	Path      string `json:"path,optional"`      // 路由路径
	Component string `json:"component,optional"` // 前端组件路径
	Redirect  string `json:"redirect,optional"`  // 重定向路径
	Icon      string `json:"icon,optional"`      // 菜单图标
	Sort      int    `json:"sort,optional"`      // 排序号
	Status    int    `json:"status,optional"`    // 状态 1:启用 2:禁用
}

type CreateMenuResp struct {
	MenuID string `json:"menu_id"` // 菜单ID
}

type CreatePermissionReq struct {
	Name        string `json:"name" validate:"required"`     // 权限名称
	Code        string `json:"code" validate:"required"`     // 权限编码
	Type        string `json:"type" validate:"required"`     // 权限类型(1:菜单menu 2:页面page 3:按钮button 4:接口api 5:数据data)
	Resource    string `json:"resource" validate:"required"` // 资源路径
	Action      string `json:"action" validate:"required"`   // 操作类型
	ParentID    string `json:"parent_id,optional"`           // 父级ID
	Description string `json:"description,optional"`         // 描述
	Status      int    `json:"status,optional"`              // 状态(1:启用 2:禁用)
	MenuID      string `json:"menu_id,optional"`             // 菜单ID
}

type CreatePermissionResp struct {
	PermissionID string `json:"permission_id"` // 权限规则ID
}

type CreateRoleReq struct {
	Name        string `json:"name"`                 // 角色名
	Code        string `json:"code"`                 // 角色编码
	Description string `json:"description,optional"` // 角色描述
	Status      int    `json:"status"`               // 状态: 1:启用, 2:禁用
	Sort        int    `json:"sort,optional"`        // 排序
}

type CreateRoleResp struct {
	RoleID string `json:"role_id"` // 角色ID
}

type CreateTenantReq struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description,optional" validate:"omitempty"`
	Status      int    `json:"status,optional" validate:"omitempty"`
}

type CreateTenantResp struct {
	TenantID string `json:"tenant_id"`
}

type CreateUserReq struct {
	UserName string   `json:"user_name"`         // 用户名
	Name     string   `json:"name"`              // 姓名
	Password string   `json:"password"`          // 密码
	Status   int      `json:"status"`            // 状态
	Phone    string   `json:"phone"`             // 手机号
	Email    string   `json:"email,optional"`    // 邮箱
	Sex      int      `json:"sex,optional"`      // 性别
	Avatar   string   `json:"avatar,optional"`   // 头像
	RoleIDs  []string `json:"role_ids,optional"` // 角色ID列表
}

type CreateUserResp struct {
	UserID string `json:"user_id"` // 用户ID
}

type DeleteMenuReq struct {
	MenuID string `path:"menu_id"`
}

type DeletePermissionReq struct {
	PermissionID string `path:"permission_id" validate:"required"` // 权限规则ID
}

type DeleteRoleReq struct {
	RoleID string `path:"role_id"`
}

type DeleteTenantReq struct {
	TenantID string `path:"tenant_id"`
}

type DeleteUserReq struct {
	UserID string `path:"user_id"`
}

type GetMenuReq struct {
	MenuID string `path:"menu_id"`
}

type GetPermissionReq struct {
	PermissionID string `path:"permission_id" validate:"required"` // 权限规则ID
}

type GetPermissionResp struct {
	PermissionInfo
}

type GetResourceTypesResp struct {
	Types map[string]ResourceTypeInfo `json:"types"` // 资源类型列表
}

type GetRolePermissionsReq struct {
	RoleCode string `path:"role_code"` // 角色ID
}

type GetRolePermissionsResp struct {
	List []*Permission `json:"list"` // 权限列表
}

type GetRoleReq struct {
	RoleID string `path:"role_id"`
}

type GetTenantReq struct {
	TenantID string `path:"tenant_id"`
}

type GetTenantResp struct {
	TenantInfo
}

type GetUserReq struct {
	UserID string `path:"user_id"`
}

type GetUserRolesReq struct {
	UserID string `path:"user_id"` // 用户ID
}

type GetUserRolesResp struct {
	List []*RoleInfo `json:"list"` // 角色列表
}

type IDRequest struct {
	ID int `json:"id"` // ID
}

type ListPermissionReq struct {
	PageRequest
	Name     string `form:"name,optional"`     // 权限名称
	Code     string `form:"code,optional"`     // 权限编码
	Type     string `form:"type,optional"`     // 权限类型
	Resource string `form:"resource,optional"` // 资源路径
	Action   int    `form:"action,optional"`   // 操作类型
	Status   int    `form:"status,optional"`   // 状态
}

type ListPermissionResp struct {
	Page *PageResponse     `json:"page"` // 分页信息
	List []*PermissionInfo `json:"list"` // 权限规则列表
}

type ListTenantReq struct {
	PageRequest
	Status int `form:"status" validate:"omitempty"` // 状态
}

type ListTenantResp struct {
	Page *PageResponse `json:"page"` // 分页
	List []*TenantInfo `json:"list"` // 租户列表
}

type LoginLogInfo struct {
	LogID     string `json:"log_id"`     // 记录ID
	UserID    string `json:"user_id"`    // 用户ID
	UserName  string `json:"user_name"`  // 用户名
	IP        string `json:"ip"`         // 登录IP
	UserAgent string `json:"user_agent"` // 用户代理
	Status    int    `json:"status"`     // 登录状态
	Message   string `json:"message"`    // 状态信息
	CreatedAt int64  `json:"created_at"` // 创建时间
}

type LoginLogListReq struct {
	PageRequest
	UserName  string `form:"user_name,optional"`  // 用户名
	IP        string `form:"ip,optional"`         // 登录IP
	Status    int    `form:"status,optional"`     // 登录状态
	StartTime string `form:"start_time,optional"` // 开始时间
	EndTime   string `form:"end_time,optional"`   // 结束时间
}

type LoginLogListResp struct {
	Page *PageResponse   `json:"page"` // 分页信息
	List []*LoginLogInfo `json:"list"` // 登录记录列表
}

type LoginReq struct {
	UserName  string `json:"username" validate:"required"`   // 用户名
	Password  string `json:"password" validate:"required"`   // 密码
	CaptchaId string `json:"captcha_id" validate:"required"` // 验证码ID
	Captcha   string `json:"captcha" validate:"required"`    // 验证码
}

type LoginResp struct {
	Token    string `json:"token"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type MenuInfo struct {
	MenuID    string `json:"menu_id"`    // 菜单ID
	Code      string `json:"code"`       // 菜单code
	ParentID  string `json:"parent_id"`  // 父菜单ID
	Name      string `json:"name"`       // 菜单名称
	Path      string `json:"path"`       // 路由路径
	Component string `json:"component"`  // 前端组件路径
	Redirect  string `json:"redirect"`   // 重定向路径
	Icon      string `json:"icon"`       // 菜单图标
	Sort      int    `json:"sort"`       // 排序号
	Type      string `json:"type"`       // 菜单类型
	Status    int    `json:"status"`     // 状态
	CreatedAt int64  `json:"created_at"` // 创建时间
}

type MenuListReq struct {
	PageRequest
	Name   string `form:"name,optional"`   // 菜单名称
	Type   string `form:"type,optional"`   // 菜单类型
	Status int    `form:"status,optional"` // 状态
}

type MenuListResp struct {
	Page *PageResponse `json:"page"` // 分页
	List []*MenuInfo   `json:"list"` // 菜单列表
}

type MenuTree struct {
	MenuID    string      `json:"menu_id"`   // 菜单ID
	Code      string      `json:"code"`      // 菜单code
	ParentID  string      `json:"parent_id"` // 父菜单ID
	Name      string      `json:"name"`      // 菜单名称
	Path      string      `json:"path"`      // 路由路径
	Component string      `json:"component"` // 前端组件路径
	Redirect  string      `json:"redirect"`  // 重定向路径
	Icon      string      `json:"icon"`      // 菜单图标
	Sort      int         `json:"sort"`      // 排序号
	Type      string      `json:"type"`      // 菜单类型
	Status    int         `json:"status"`    // 状态
	Children  []*MenuTree `json:"children"`  // 子菜单
}

type MenuTreeResp struct {
	List []*MenuTree `json:"list"` // 菜单树列表
}

type PageJsonRequest struct {
	Current  int `json:"current"`   // 当前页
	PageSize int `json:"page_size"` // 每页大小
}

type PageRequest struct {
	Current  int `form:"page,default=1" validate:"required,gte=1"`               // 当前页
	PageSize int `form:"page_size,default=10" validate:"required,gte=1,lte=100"` // 每页大小
}

type PageResponse struct {
	Total           int `json:"total"`             // 总数
	PageSize        int `json:"page_size"`         // 当前页大小
	RequestPageSize int `json:"request_page_size"` // 请求页大小
	Current         int `json:"current"`           // 当前页
}

type Permission struct {
	Resource string `json:"resource"` // 资源
	Action   string `json:"action"`   // 操作类型
	Type     string `json:"type"`     // 权限类型
}

type PermissionInfo struct {
	ID           string `json:"id"`            // 权限规则ID
	TenantCode   string `json:"tenant_code"`   // 租户编码
	PermissionID string `json:"permission_id"` // 权限规则ID
	Name         string `json:"name"`          // 权限名称
	Code         string `json:"code"`          // 权限编码
	Type         string `json:"type"`          // 权限类型
	Resource     string `json:"resource"`      // 资源路径
	Action       string `json:"action"`        // 操作类型
	ParentID     string `json:"parent_id"`     // 父级ID
	Description  string `json:"description"`   // 描述
	Status       int    `json:"status"`        // 状态
	MenuID       string `json:"menu_id"`       // 菜单ID
	CreatedAt    int64  `json:"created_at"`    // 创建时间
}

type RegisterReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar,optional"`
	Phone    string `json:"phone"`
	Email    string `json:"email,optional"`
	Sex      int    `json:"sex,optional"`
}

type RegisterResp struct {
	UserID string `json:"user_id"`
}

type ResetPasswordReq struct {
	UserID      string `json:"user_id" validate:"required"`      // 用户ID
	NewPassword string `json:"new_password" validate:"required"` // 新密码
}

type ResourceTypeInfo struct {
	Type      string   `json:"type"`       // 资源类型
	Actions   []string `json:"actions"`    // 可用操作
	DataRules []string `json:"data_rules"` // 数据规则（仅对数据权限类型有效）
}

type RoleInfo struct {
	RoleID      string `json:"role_id"`     // 角色ID
	Name        string `json:"name"`        // 角色名
	Code        string `json:"code"`        // 角色编码
	Description string `json:"description"` // 角色描述
	Status      int    `json:"status"`      // 状态
	Sort        int    `json:"sort"`        // 排序
	CreatedAt   int64  `json:"created_at"`  // 创建时间
}

type RoleListReq struct {
	PageRequest
	Name   string `form:"name,optional"`   // 角色名
	Code   string `form:"code,optional"`   // 角色编码
	Status int    `form:"status,optional"` // 状态
}

type RoleListResp struct {
	Page *PageResponse `json:"page"` // 分页
	List []*RoleInfo   `json:"list"` // 角色列表
}

type SetRolePermissionsReq struct {
	RoleCode       string        `path:"role_code"`                           // 角色ID
	PermissionList []*Permission `json:"permission_list" validate:"required"` // 权限列表
}

type SetUserRolesReq struct {
	UserID       string   `path:"user_id"`                            // 用户ID
	RoleCodeList []string `json:"role_code_list" validate:"required"` // 角色Code列表
}

type StatusRequest struct {
	ID     int `json:"id"`     // ID
	Status int `json:"status"` // 状态
}

type TenantInfo struct {
	TenantID    string `json:"tenant_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type TimeRange struct {
	StartTime string `form:"start_time,optional"` // 开始时间
	EndTime   string `form:"end_time,optional"`   // 结束时间
}

type UpdateMenuReq struct {
	MenuID    string `path:"menu_id"`
	Name      string `json:"name,optional"`      // 菜单名称
	Path      string `json:"path,optional"`      // 路由路径
	Component string `json:"component,optional"` // 前端组件路径
	Redirect  string `json:"redirect,optional"`  // 重定向路径
	Icon      string `json:"icon,optional"`      // 菜单图标
	Sort      int    `json:"sort,optional"`      // 排序号
	Status    int    `json:"status,optional"`    // 状态
}

type UpdatePermissionReq struct {
	PermissionID string `path:"permission_id"`        // 权限规则ID
	Name         string `json:"name,optional"`        // 权限名称
	Type         string `json:"type,optional"`        // 权限类型
	Resource     string `json:"resource,optional"`    // 资源路径
	Action       string `json:"action,optional"`      // 操作类型
	ParentID     string `json:"parent_id,optional"`   // 父级ID
	Description  string `json:"description,optional"` // 描述
	Status       int    `json:"status,optional"`      // 状态
	MenuID       string `json:"menu_id,optional"`     // 菜单ID
}

type UpdateRoleReq struct {
	RoleID      string `path:"role_id"`
	Name        string `json:"name,optional"`        // 角色名
	Description string `json:"description,optional"` // 角色描述
	Status      int    `json:"status,optional"`      // 状态
	Sort        int    `json:"sort,optional"`        // 排序
}

type UpdateTenantReq struct {
	TenantID    string `path:"tenant_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type UpdateUserReq struct {
	UserID  string   `path:"user_id"`
	Name    string   `json:"name,optional"`
	Email   string   `json:"email,optional"`
	Status  int      `json:"status,optional"`
	Sex     int      `json:"sex,optional"`
	Avatar  string   `json:"avatar,optional"`
	RoleIDs []string `json:"role_ids,optional"`
}

type UserInfo struct {
	UserID    string `json:"user_id"`    // 用户ID
	UserName  string `json:"user_name"`  // 用户名
	Name      string `json:"name"`       // 姓名
	Phone     string `json:"phone"`      // 手机号
	Email     string `json:"email"`      // 邮箱
	Avatar    string `json:"avatar"`     // 头像
	Status    int    `json:"status"`     // 状态
	CreatedAt int64  `json:"created_at"` // 创建时间
}

type UserListReq struct {
	PageRequest
	Name   string `form:"name,optional"`   // 姓名
	Phone  string `form:"phone,optional"`  // 手机号
	Status int    `form:"status,optional"` // 状态
}

type UserListResp struct {
	Page *PageResponse `json:"page"` // 分页
	List []*UserInfo   `json:"list"` // 用户列表
}
