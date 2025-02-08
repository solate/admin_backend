// Code generated by ent, DO NOT EDIT.

package generated

import (
	"github.com/solate/admin_backend/pkg/ent/generated/loginlog"
	"github.com/solate/admin_backend/pkg/ent/generated/permission"
	"github.com/solate/admin_backend/pkg/ent/generated/role"
	"github.com/solate/admin_backend/pkg/ent/generated/systemlog"
	"github.com/solate/admin_backend/pkg/ent/generated/tenant"
	"github.com/solate/admin_backend/pkg/ent/generated/user"
	"github.com/solate/admin_backend/pkg/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	loginlogFields := schema.LoginLog{}.Fields()
	_ = loginlogFields
	// loginlogDescCreatedAt is the schema descriptor for created_at field.
	loginlogDescCreatedAt := loginlogFields[0].Descriptor()
	// loginlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	loginlog.DefaultCreatedAt = loginlogDescCreatedAt.Default.(int)
	// loginlogDescUpdatedAt is the schema descriptor for updated_at field.
	loginlogDescUpdatedAt := loginlogFields[1].Descriptor()
	// loginlog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	loginlog.DefaultUpdatedAt = loginlogDescUpdatedAt.Default.(int)
	// loginlogDescUserName is the schema descriptor for user_name field.
	loginlogDescUserName := loginlogFields[4].Descriptor()
	// loginlog.UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	loginlog.UserNameValidator = loginlogDescUserName.Validators[0].(func(string) error)
	// loginlogDescIP is the schema descriptor for ip field.
	loginlogDescIP := loginlogFields[5].Descriptor()
	// loginlog.IPValidator is a validator for the "ip" field. It is called by the builders before save.
	loginlog.IPValidator = loginlogDescIP.Validators[0].(func(string) error)
	// loginlogDescStatus is the schema descriptor for status field.
	loginlogDescStatus := loginlogFields[6].Descriptor()
	// loginlog.DefaultStatus holds the default value on creation for the status field.
	loginlog.DefaultStatus = loginlogDescStatus.Default.(int)
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionFields[0].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(int)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionFields[1].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(int)
	// permissionDescTenantCode is the schema descriptor for tenant_code field.
	permissionDescTenantCode := permissionFields[3].Descriptor()
	// permission.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	permission.TenantCodeValidator = permissionDescTenantCode.Validators[0].(func(string) error)
	// permissionDescName is the schema descriptor for name field.
	permissionDescName := permissionFields[4].Descriptor()
	// permission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	permission.NameValidator = permissionDescName.Validators[0].(func(string) error)
	// permissionDescCode is the schema descriptor for code field.
	permissionDescCode := permissionFields[5].Descriptor()
	// permission.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	permission.CodeValidator = permissionDescCode.Validators[0].(func(string) error)
	// permissionDescAction is the schema descriptor for action field.
	permissionDescAction := permissionFields[8].Descriptor()
	// permission.DefaultAction holds the default value on creation for the action field.
	permission.DefaultAction = permissionDescAction.Default.(int)
	// permissionDescStatus is the schema descriptor for status field.
	permissionDescStatus := permissionFields[11].Descriptor()
	// permission.DefaultStatus holds the default value on creation for the status field.
	permission.DefaultStatus = permissionDescStatus.Default.(int)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(int)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(int)
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
	// systemlogDescUpdatedAt is the schema descriptor for updated_at field.
	systemlogDescUpdatedAt := systemlogFields[1].Descriptor()
	// systemlog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	systemlog.DefaultUpdatedAt = systemlogDescUpdatedAt.Default.(int64)
	// systemlogDescTenantCode is the schema descriptor for tenant_code field.
	systemlogDescTenantCode := systemlogFields[2].Descriptor()
	// systemlog.TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	systemlog.TenantCodeValidator = systemlogDescTenantCode.Validators[0].(func(string) error)
	// systemlogDescModule is the schema descriptor for module field.
	systemlogDescModule := systemlogFields[3].Descriptor()
	// systemlog.DefaultModule holds the default value on creation for the module field.
	systemlog.DefaultModule = systemlogDescModule.Default.(string)
	// systemlogDescAction is the schema descriptor for action field.
	systemlogDescAction := systemlogFields[4].Descriptor()
	// systemlog.DefaultAction holds the default value on creation for the action field.
	systemlog.DefaultAction = systemlogDescAction.Default.(string)
	// systemlogDescContent is the schema descriptor for content field.
	systemlogDescContent := systemlogFields[5].Descriptor()
	// systemlog.DefaultContent holds the default value on creation for the content field.
	systemlog.DefaultContent = systemlogDescContent.Default.(string)
	// systemlogDescOperator is the schema descriptor for operator field.
	systemlogDescOperator := systemlogFields[6].Descriptor()
	// systemlog.DefaultOperator holds the default value on creation for the operator field.
	systemlog.DefaultOperator = systemlogDescOperator.Default.(string)
	// systemlogDescUserID is the schema descriptor for user_id field.
	systemlogDescUserID := systemlogFields[7].Descriptor()
	// systemlog.DefaultUserID holds the default value on creation for the user_id field.
	systemlog.DefaultUserID = systemlogDescUserID.Default.(uint64)
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
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
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
	// userDescRoleID is the schema descriptor for role_id field.
	userDescRoleID := userFields[15].Descriptor()
	// user.DefaultRoleID holds the default value on creation for the role_id field.
	user.DefaultRoleID = userDescRoleID.Default.(uint64)
	// userDescDeptID is the schema descriptor for dept_id field.
	userDescDeptID := userFields[16].Descriptor()
	// user.DefaultDeptID holds the default value on creation for the dept_id field.
	user.DefaultDeptID = userDescDeptID.Default.(uint64)
	// userDescPostID is the schema descriptor for post_id field.
	userDescPostID := userFields[17].Descriptor()
	// user.DefaultPostID holds the default value on creation for the post_id field.
	user.DefaultPostID = userDescPostID.Default.(uint64)
}
