package casbin

// 权限类型常量
const (
	PermTypeMenu   = "menu"   // 菜单权限
	PermTypePage   = "page"   // 页面权限
	PermTypeButton = "button" // 按钮权限
	PermTypeAPI    = "api"    // 接口权限
	PermTypeData   = "data"   // 数据权限
)

// 操作类型常量
const (
	ActionView   = "view"   // 查看
	ActionAdd    = "add"    // 增加
	ActionEdit   = "edit"   // 编辑
	ActionDelete = "delete" // 删除
	ActionExport = "export" // 导出
	ActionImport = "import" // 导入
)

// DataRule 数据权限规则
const (
	DataRuleSelf          = "self"           // 仅本人数据
	DataRuleDepartment    = "department"     // 本部门数据
	DataRuleDepartmentSub = "department_sub" // 本部门及下级部门
	DataRuleAll           = "all"            // 所有数据
)

// ResourceType 资源类型
type ResourceType struct {
	Type      string
	Actions   []string
	DataRules []string
}

// 预定义资源类型
var ResourceTypes = map[string]ResourceType{
	PermTypeMenu: {
		Type:    PermTypeMenu,
		Actions: []string{ActionView},
	},
	PermTypePage: {
		Type:    PermTypePage,
		Actions: []string{ActionView},
	},
	PermTypeButton: {
		Type:    PermTypeButton,
		Actions: []string{ActionView, ActionAdd, ActionEdit, ActionDelete, ActionExport, ActionImport},
	},
	PermTypeAPI: {
		Type:    PermTypeAPI,
		Actions: []string{"GET", "POST", "PUT", "DELETE"},
	},
	PermTypeData: {
		Type:    PermTypeData,
		Actions: []string{ActionView, ActionAdd, ActionEdit, ActionDelete},
		DataRules: []string{
			DataRuleSelf,          // 仅本人数据
			DataRuleDepartment,    // 本部门数据
			DataRuleDepartmentSub, // 本部门及下级部门
			DataRuleAll,           // 所有数据
		},
	},
}

// Permission 权限定义
type Permission struct {
	Resource string
	Action   string
	Type     string // 添加权限类型字段
}

/**

// 使用示例
rule := DataPermissionRule{
    Resource: "employee",
    Attributes: []interface{}{
        "dept_id = 1",
        "position in ('manager', 'director')",
    },
    Conditions: []string{
        "working_hours between 9 and 18",
    },
    Operations: []string{"view", "edit"},
}
*/
// DataPermissionRule 数据权限规则
type DataPermissionRule struct {
	Resource   string   // 资源名称
	Rule       string   // 权限规则（如：dept_id in (1,2,3)）
	Conditions []string // 条件
	Operations []string // 操作
}

// Menu 菜单定义
type Menu struct {
	ID       string `json:"id"`
	Path     string `json:"path"`     // 路由路径
	Name     string `json:"name"`     // 菜单名称
	ParentID string `json:"parentId"` // 父菜单ID
}

// Button 按钮定义
type Button struct {
	ID     string `json:"id"`
	MenuID string `json:"menuId"` // 所属菜单ID
	Code   string `json:"code"`   // 按钮标识
	Name   string `json:"name"`   // 按钮名称
}

// DataAttribute 数据属性
type DataAttribute struct {
	Field    string `json:"field"`    // 字段名
	Operator string `json:"operator"` // 操作符
	Value    string `json:"value"`    // 值
}

// Role 角色定义
type Role struct {
	ID          string
	Name        string
	Description string
	Domain      string
	Permissions []Permission
}

// User 用户权限相关信息
type User struct {
	ID     string
	Roles  []string
	Domain string
}
