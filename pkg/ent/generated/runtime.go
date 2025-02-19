// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/casbinrule"
	"admin_backend/pkg/ent/generated/department"
	"admin_backend/pkg/ent/generated/dictitem"
	"admin_backend/pkg/ent/generated/dicttype"
	"admin_backend/pkg/ent/generated/loginlog"
	"admin_backend/pkg/ent/generated/menu"
	"admin_backend/pkg/ent/generated/permission"
	"admin_backend/pkg/ent/generated/position"
	"admin_backend/pkg/ent/generated/role"
	"admin_backend/pkg/ent/generated/systemlog"
	"admin_backend/pkg/ent/generated/tenant"
	"admin_backend/pkg/ent/generated/user"
	"admin_backend/pkg/ent/generated/userposition"
	"admin_backend/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbinruleFields := schema.CasbinRule{}.Fields()
	_ = casbinruleFields
	// casbinruleDescPtype is the schema descriptor for Ptype field.
	casbinruleDescPtype := casbinruleFields[0].Descriptor()
	// casbinrule.DefaultPtype holds the default value on creation for the Ptype field.
	casbinrule.DefaultPtype = casbinruleDescPtype.Default.(string)
	// casbinruleDescV0 is the schema descriptor for V0 field.
	casbinruleDescV0 := casbinruleFields[1].Descriptor()
	// casbinrule.DefaultV0 holds the default value on creation for the V0 field.
	casbinrule.DefaultV0 = casbinruleDescV0.Default.(string)
	// casbinruleDescV1 is the schema descriptor for V1 field.
	casbinruleDescV1 := casbinruleFields[2].Descriptor()
	// casbinrule.DefaultV1 holds the default value on creation for the V1 field.
	casbinrule.DefaultV1 = casbinruleDescV1.Default.(string)
	// casbinruleDescV2 is the schema descriptor for V2 field.
	casbinruleDescV2 := casbinruleFields[3].Descriptor()
	// casbinrule.DefaultV2 holds the default value on creation for the V2 field.
	casbinrule.DefaultV2 = casbinruleDescV2.Default.(string)
	// casbinruleDescV3 is the schema descriptor for V3 field.
	casbinruleDescV3 := casbinruleFields[4].Descriptor()
	// casbinrule.DefaultV3 holds the default value on creation for the V3 field.
	casbinrule.DefaultV3 = casbinruleDescV3.Default.(string)
	// casbinruleDescV4 is the schema descriptor for V4 field.
	casbinruleDescV4 := casbinruleFields[5].Descriptor()
	// casbinrule.DefaultV4 holds the default value on creation for the V4 field.
	casbinrule.DefaultV4 = casbinruleDescV4.Default.(string)
	// casbinruleDescV5 is the schema descriptor for V5 field.
	casbinruleDescV5 := casbinruleFields[6].Descriptor()
	// casbinrule.DefaultV5 holds the default value on creation for the V5 field.
	casbinrule.DefaultV5 = casbinruleDescV5.Default.(string)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescCreatedAt is the schema descriptor for created_at field.
	departmentDescCreatedAt := departmentFields[0].Descriptor()
	// department.DefaultCreatedAt holds the default value on creation for the created_at field.
	department.DefaultCreatedAt = departmentDescCreatedAt.Default.(int64)
	// departmentDescUpdatedAt is the schema descriptor for updated_at field.
	departmentDescUpdatedAt := departmentFields[1].Descriptor()
	// department.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	department.DefaultUpdatedAt = departmentDescUpdatedAt.Default.(int64)
	// departmentDescTenantCode is the schema descriptor for tenant_code field.
	departmentDescTenantCode := departmentFields[3].Descriptor()
	// department.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	department.TenantCodeValidator = departmentDescTenantCode.Validators[0].(func(string) error)
	// departmentDescName is the schema descriptor for name field.
	departmentDescName := departmentFields[5].Descriptor()
	// department.NameValidator is a validator for the "name" field. It is called by the builders before save.
	department.NameValidator = departmentDescName.Validators[0].(func(string) error)
	// departmentDescParentID is the schema descriptor for parent_id field.
	departmentDescParentID := departmentFields[6].Descriptor()
	// department.DefaultParentID holds the default value on creation for the parent_id field.
	department.DefaultParentID = departmentDescParentID.Default.(string)
	// departmentDescSort is the schema descriptor for sort field.
	departmentDescSort := departmentFields[7].Descriptor()
	// department.DefaultSort holds the default value on creation for the sort field.
	department.DefaultSort = departmentDescSort.Default.(int)
	dictitemFields := schema.DictItem{}.Fields()
	_ = dictitemFields
	// dictitemDescCreatedAt is the schema descriptor for created_at field.
	dictitemDescCreatedAt := dictitemFields[0].Descriptor()
	// dictitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	dictitem.DefaultCreatedAt = dictitemDescCreatedAt.Default.(int64)
	// dictitemDescUpdatedAt is the schema descriptor for updated_at field.
	dictitemDescUpdatedAt := dictitemFields[1].Descriptor()
	// dictitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dictitem.DefaultUpdatedAt = dictitemDescUpdatedAt.Default.(int64)
	// dictitemDescTenantCode is the schema descriptor for tenant_code field.
	dictitemDescTenantCode := dictitemFields[3].Descriptor()
	// dictitem.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	dictitem.TenantCodeValidator = dictitemDescTenantCode.Validators[0].(func(string) error)
	// dictitemDescTypeCode is the schema descriptor for type_code field.
	dictitemDescTypeCode := dictitemFields[5].Descriptor()
	// dictitem.TypeCodeValidator is a validator for the "type_code" field. It is called by the builders before save.
	dictitem.TypeCodeValidator = dictitemDescTypeCode.Validators[0].(func(string) error)
	// dictitemDescLabel is the schema descriptor for label field.
	dictitemDescLabel := dictitemFields[6].Descriptor()
	// dictitem.LabelValidator is a validator for the "label" field. It is called by the builders before save.
	dictitem.LabelValidator = dictitemDescLabel.Validators[0].(func(string) error)
	// dictitemDescValue is the schema descriptor for value field.
	dictitemDescValue := dictitemFields[7].Descriptor()
	// dictitem.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	dictitem.ValueValidator = dictitemDescValue.Validators[0].(func(string) error)
	// dictitemDescSort is the schema descriptor for sort field.
	dictitemDescSort := dictitemFields[9].Descriptor()
	// dictitem.DefaultSort holds the default value on creation for the sort field.
	dictitem.DefaultSort = dictitemDescSort.Default.(int)
	// dictitemDescStatus is the schema descriptor for status field.
	dictitemDescStatus := dictitemFields[10].Descriptor()
	// dictitem.DefaultStatus holds the default value on creation for the status field.
	dictitem.DefaultStatus = dictitemDescStatus.Default.(int)
	dicttypeFields := schema.DictType{}.Fields()
	_ = dicttypeFields
	// dicttypeDescCreatedAt is the schema descriptor for created_at field.
	dicttypeDescCreatedAt := dicttypeFields[0].Descriptor()
	// dicttype.DefaultCreatedAt holds the default value on creation for the created_at field.
	dicttype.DefaultCreatedAt = dicttypeDescCreatedAt.Default.(int64)
	// dicttypeDescUpdatedAt is the schema descriptor for updated_at field.
	dicttypeDescUpdatedAt := dicttypeFields[1].Descriptor()
	// dicttype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	dicttype.DefaultUpdatedAt = dicttypeDescUpdatedAt.Default.(int64)
	// dicttypeDescTenantCode is the schema descriptor for tenant_code field.
	dicttypeDescTenantCode := dicttypeFields[3].Descriptor()
	// dicttype.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	dicttype.TenantCodeValidator = dicttypeDescTenantCode.Validators[0].(func(string) error)
	// dicttypeDescName is the schema descriptor for name field.
	dicttypeDescName := dicttypeFields[5].Descriptor()
	// dicttype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	dicttype.NameValidator = dicttypeDescName.Validators[0].(func(string) error)
	// dicttypeDescCode is the schema descriptor for code field.
	dicttypeDescCode := dicttypeFields[6].Descriptor()
	// dicttype.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	dicttype.CodeValidator = dicttypeDescCode.Validators[0].(func(string) error)
	// dicttypeDescStatus is the schema descriptor for status field.
	dicttypeDescStatus := dicttypeFields[8].Descriptor()
	// dicttype.DefaultStatus holds the default value on creation for the status field.
	dicttype.DefaultStatus = dicttypeDescStatus.Default.(int)
	loginlogFields := schema.LoginLog{}.Fields()
	_ = loginlogFields
	// loginlogDescCreatedAt is the schema descriptor for created_at field.
	loginlogDescCreatedAt := loginlogFields[0].Descriptor()
	// loginlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	loginlog.DefaultCreatedAt = loginlogDescCreatedAt.Default.(int64)
	// loginlogDescTenantCode is the schema descriptor for tenant_code field.
	loginlogDescTenantCode := loginlogFields[1].Descriptor()
	// loginlog.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	loginlog.TenantCodeValidator = loginlogDescTenantCode.Validators[0].(func(string) error)
	// loginlogDescUserName is the schema descriptor for user_name field.
	loginlogDescUserName := loginlogFields[4].Descriptor()
	// loginlog.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	loginlog.UserNameValidator = loginlogDescUserName.Validators[0].(func(string) error)
	// loginlogDescIP is the schema descriptor for ip field.
	loginlogDescIP := loginlogFields[5].Descriptor()
	// loginlog.IPValidator is a validator for the "ip" field. It is called by the builders before save.
	loginlog.IPValidator = loginlogDescIP.Validators[0].(func(string) error)
	menuFields := schema.Menu{}.Fields()
	_ = menuFields
	// menuDescCreatedAt is the schema descriptor for created_at field.
	menuDescCreatedAt := menuFields[0].Descriptor()
	// menu.DefaultCreatedAt holds the default value on creation for the created_at field.
	menu.DefaultCreatedAt = menuDescCreatedAt.Default.(int64)
	// menuDescUpdatedAt is the schema descriptor for updated_at field.
	menuDescUpdatedAt := menuFields[1].Descriptor()
	// menu.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menu.DefaultUpdatedAt = menuDescUpdatedAt.Default.(int64)
	// menuDescCode is the schema descriptor for code field.
	menuDescCode := menuFields[5].Descriptor()
	// menu.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	menu.CodeValidator = menuDescCode.Validators[0].(func(string) error)
	// menuDescParentID is the schema descriptor for parent_id field.
	menuDescParentID := menuFields[6].Descriptor()
	// menu.DefaultParentID holds the default value on creation for the parent_id field.
	menu.DefaultParentID = menuDescParentID.Default.(string)
	// menuDescName is the schema descriptor for name field.
	menuDescName := menuFields[7].Descriptor()
	// menu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menu.NameValidator = menuDescName.Validators[0].(func(string) error)
	// menuDescPath is the schema descriptor for path field.
	menuDescPath := menuFields[8].Descriptor()
	// menu.DefaultPath holds the default value on creation for the path field.
	menu.DefaultPath = menuDescPath.Default.(string)
	// menuDescComponent is the schema descriptor for component field.
	menuDescComponent := menuFields[9].Descriptor()
	// menu.DefaultComponent holds the default value on creation for the component field.
	menu.DefaultComponent = menuDescComponent.Default.(string)
	// menuDescRedirect is the schema descriptor for redirect field.
	menuDescRedirect := menuFields[10].Descriptor()
	// menu.DefaultRedirect holds the default value on creation for the redirect field.
	menu.DefaultRedirect = menuDescRedirect.Default.(string)
	// menuDescIcon is the schema descriptor for icon field.
	menuDescIcon := menuFields[11].Descriptor()
	// menu.DefaultIcon holds the default value on creation for the icon field.
	menu.DefaultIcon = menuDescIcon.Default.(string)
	// menuDescSort is the schema descriptor for sort field.
	menuDescSort := menuFields[12].Descriptor()
	// menu.DefaultSort holds the default value on creation for the sort field.
	menu.DefaultSort = menuDescSort.Default.(int)
	// menuDescType is the schema descriptor for type field.
	menuDescType := menuFields[13].Descriptor()
	// menu.DefaultType holds the default value on creation for the type field.
	menu.DefaultType = menuDescType.Default.(string)
	// menuDescStatus is the schema descriptor for status field.
	menuDescStatus := menuFields[14].Descriptor()
	// menu.DefaultStatus holds the default value on creation for the status field.
	menu.DefaultStatus = menuDescStatus.Default.(int)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionFields[0].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(int64)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionFields[1].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(int64)
	// permissionDescTenantCode is the schema descriptor for tenant_code field.
	permissionDescTenantCode := permissionFields[3].Descriptor()
	// permission.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	permission.TenantCodeValidator = permissionDescTenantCode.Validators[0].(func(string) error)
	// permissionDescName is the schema descriptor for name field.
	permissionDescName := permissionFields[5].Descriptor()
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescCode is the schema descriptor for code field.
	permissionDescCode := permissionFields[6].Descriptor()
	// permission.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	permission.CodeValidator = permissionDescCode.Validators[0].(func(string) error)
	// permissionDescResource is the schema descriptor for resource field.
	permissionDescResource := permissionFields[8].Descriptor()
	// permission.ResourceValidator is a validator for the "resource" field. It is called by the builders before save.
	permission.ResourceValidator = permissionDescResource.Validators[0].(func(string) error)
	// permissionDescAction is the schema descriptor for action field.
	permissionDescAction := permissionFields[9].Descriptor()
	// permission.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	permission.ActionValidator = permissionDescAction.Validators[0].(func(string) error)
	// permissionDescStatus is the schema descriptor for status field.
	permissionDescStatus := permissionFields[12].Descriptor()
	// permission.DefaultStatus holds the default value on creation for the status field.
	permission.DefaultStatus = permissionDescStatus.Default.(int)
	positionFields := schema.Position{}.Fields()
	_ = positionFields
	// positionDescCreatedAt is the schema descriptor for created_at field.
	positionDescCreatedAt := positionFields[0].Descriptor()
	// position.DefaultCreatedAt holds the default value on creation for the created_at field.
	position.DefaultCreatedAt = positionDescCreatedAt.Default.(int64)
	// positionDescUpdatedAt is the schema descriptor for updated_at field.
	positionDescUpdatedAt := positionFields[1].Descriptor()
	// position.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	position.DefaultUpdatedAt = positionDescUpdatedAt.Default.(int64)
	// positionDescTenantCode is the schema descriptor for tenant_code field.
	positionDescTenantCode := positionFields[3].Descriptor()
	// position.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	position.TenantCodeValidator = positionDescTenantCode.Validators[0].(func(string) error)
	// positionDescName is the schema descriptor for name field.
	positionDescName := positionFields[5].Descriptor()
	// position.NameValidator is a validator for the "name" field. It is called by the builders before save.
	position.NameValidator = positionDescName.Validators[0].(func(string) error)
	// positionDescDepartmentID is the schema descriptor for department_id field.
	positionDescDepartmentID := positionFields[6].Descriptor()
	// position.DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	position.DepartmentIDValidator = positionDescDepartmentID.Validators[0].(func(string) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(int64)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(int64)
	// roleDescTenantCode is the schema descriptor for tenant_code field.
	roleDescTenantCode := roleFields[3].Descriptor()
	// role.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	role.TenantCodeValidator = roleDescTenantCode.Validators[0].(func(string) error)
	// roleDescName is the schema descriptor for name field.
	roleDescName := roleFields[5].Descriptor()
	// role.NameValidator is a validator for the "name" field. It is called by the builders before save.
	role.NameValidator = roleDescName.Validators[0].(func(string) error)
	// roleDescCode is the schema descriptor for code field.
	roleDescCode := roleFields[6].Descriptor()
	// role.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	role.CodeValidator = roleDescCode.Validators[0].(func(string) error)
	// roleDescStatus is the schema descriptor for status field.
	roleDescStatus := roleFields[8].Descriptor()
	// role.DefaultStatus holds the default value on creation for the status field.
	role.DefaultStatus = roleDescStatus.Default.(int)
	// roleDescSort is the schema descriptor for sort field.
	roleDescSort := roleFields[9].Descriptor()
	// role.DefaultSort holds the default value on creation for the sort field.
	role.DefaultSort = roleDescSort.Default.(int)
	systemlogFields := schema.SystemLog{}.Fields()
	_ = systemlogFields
	// systemlogDescCreatedAt is the schema descriptor for created_at field.
	systemlogDescCreatedAt := systemlogFields[0].Descriptor()
	// systemlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	systemlog.DefaultCreatedAt = systemlogDescCreatedAt.Default.(int64)
	// systemlogDescTenantCode is the schema descriptor for tenant_code field.
	systemlogDescTenantCode := systemlogFields[1].Descriptor()
	// systemlog.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	systemlog.TenantCodeValidator = systemlogDescTenantCode.Validators[0].(func(string) error)
	// systemlogDescModule is the schema descriptor for module field.
	systemlogDescModule := systemlogFields[2].Descriptor()
	// systemlog.DefaultModule holds the default value on creation for the module field.
	systemlog.DefaultModule = systemlogDescModule.Default.(string)
	// systemlogDescAction is the schema descriptor for action field.
	systemlogDescAction := systemlogFields[3].Descriptor()
	// systemlog.DefaultAction holds the default value on creation for the action field.
	systemlog.DefaultAction = systemlogDescAction.Default.(string)
	// systemlogDescContent is the schema descriptor for content field.
	systemlogDescContent := systemlogFields[4].Descriptor()
	// systemlog.DefaultContent holds the default value on creation for the content field.
	systemlog.DefaultContent = systemlogDescContent.Default.(string)
	// systemlogDescOperator is the schema descriptor for operator field.
	systemlogDescOperator := systemlogFields[5].Descriptor()
	// systemlog.DefaultOperator holds the default value on creation for the operator field.
	systemlog.DefaultOperator = systemlogDescOperator.Default.(string)
	// systemlogDescUserID is the schema descriptor for user_id field.
	systemlogDescUserID := systemlogFields[6].Descriptor()
	// systemlog.DefaultUserID holds the default value on creation for the user_id field.
	systemlog.DefaultUserID = systemlogDescUserID.Default.(string)
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescCreatedAt is the schema descriptor for created_at field.
	tenantDescCreatedAt := tenantFields[0].Descriptor()
	// tenant.DefaultCreatedAt holds the default value on creation for the created_at field.
	tenant.DefaultCreatedAt = tenantDescCreatedAt.Default.(int64)
	// tenantDescUpdatedAt is the schema descriptor for updated_at field.
	tenantDescUpdatedAt := tenantFields[1].Descriptor()
	// tenant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tenant.DefaultUpdatedAt = tenantDescUpdatedAt.Default.(int64)
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[4].Descriptor()
	// tenant.DefaultName holds the default value on creation for the name field.
	tenant.DefaultName = tenantDescName.Default.(string)
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
	// tenantDescCode is the schema descriptor for code field.
	tenantDescCode := tenantFields[5].Descriptor()
	// tenant.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	tenant.CodeValidator = tenantDescCode.Validators[0].(func(string) error)
	// tenantDescDescription is the schema descriptor for description field.
	tenantDescDescription := tenantFields[6].Descriptor()
	// tenant.DefaultDescription holds the default value on creation for the description field.
	tenant.DefaultDescription = tenantDescDescription.Default.(string)
	// tenantDescStatus is the schema descriptor for status field.
	tenantDescStatus := tenantFields[7].Descriptor()
	// tenant.DefaultStatus holds the default value on creation for the status field.
	tenant.DefaultStatus = tenantDescStatus.Default.(int)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(int64)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(int64)
	// userDescUserName is the schema descriptor for user_name field.
	userDescUserName := userFields[5].Descriptor()
	// user.DefaultUserName holds the default value on creation for the user_name field.
	user.DefaultUserName = userDescUserName.Default.(string)
	// user.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	user.UserNameValidator = userDescUserName.Validators[0].(func(string) error)
	// userDescPwdHashed is the schema descriptor for pwd_hashed field.
	userDescPwdHashed := userFields[6].Descriptor()
	// user.DefaultPwdHashed holds the default value on creation for the pwd_hashed field.
	user.DefaultPwdHashed = userDescPwdHashed.Default.(string)
	// user.PwdHashedValidator is a validator for the "pwd_hashed" field. It is called by the builders before save.
	user.PwdHashedValidator = userDescPwdHashed.Validators[0].(func(string) error)
	// userDescPwdSalt is the schema descriptor for pwd_salt field.
	userDescPwdSalt := userFields[7].Descriptor()
	// user.DefaultPwdSalt holds the default value on creation for the pwd_salt field.
	user.DefaultPwdSalt = userDescPwdSalt.Default.(string)
	// user.PwdSaltValidator is a validator for the "pwd_salt" field. It is called by the builders before save.
	user.PwdSaltValidator = userDescPwdSalt.Validators[0].(func(string) error)
	// userDescToken is the schema descriptor for token field.
	userDescToken := userFields[8].Descriptor()
	// user.DefaultToken holds the default value on creation for the token field.
	user.DefaultToken = userDescToken.Default.(string)
	// userDescNickName is the schema descriptor for nick_name field.
	userDescNickName := userFields[9].Descriptor()
	// user.DefaultNickName holds the default value on creation for the nick_name field.
	user.DefaultNickName = userDescNickName.Default.(string)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[10].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[11].Descriptor()
	// user.DefaultPhone holds the default value on creation for the phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[12].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
	// userDescSex is the schema descriptor for sex field.
	userDescSex := userFields[13].Descriptor()
	// user.DefaultSex holds the default value on creation for the sex field.
	user.DefaultSex = userDescSex.Default.(int)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[14].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int)
	userpositionFields := schema.UserPosition{}.Fields()
	_ = userpositionFields
	// userpositionDescUserID is the schema descriptor for user_id field.
	userpositionDescUserID := userpositionFields[0].Descriptor()
	// userposition.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	userposition.UserIDValidator = userpositionDescUserID.Validators[0].(func(string) error)
	// userpositionDescPositionID is the schema descriptor for position_id field.
	userpositionDescPositionID := userpositionFields[1].Descriptor()
	// userposition.PositionIDValidator is a validator for the "position_id" field. It is called by the builders before save.
	userposition.PositionIDValidator = userpositionDescPositionID.Validators[0].(func(string) error)
}
