// Code generated by ent, DO NOT EDIT.

package department

import (
	"admin_backend/pkg/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDeletedAt, v))
}

// TenantCode applies equality check predicate on the "tenant_code" field. It's identical to TenantCodeEQ.
func TenantCode(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldTenantCode, v))
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDepartmentID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int64) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int64) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int64) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int64) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int64) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Department {
	return predicate.Department(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Department {
	return predicate.Department(sql.FieldNotNull(FieldDeletedAt))
}

// TenantCodeEQ applies the EQ predicate on the "tenant_code" field.
func TenantCodeEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldTenantCode, v))
}

// TenantCodeNEQ applies the NEQ predicate on the "tenant_code" field.
func TenantCodeNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldTenantCode, v))
}

// TenantCodeIn applies the In predicate on the "tenant_code" field.
func TenantCodeIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldTenantCode, vs...))
}

// TenantCodeNotIn applies the NotIn predicate on the "tenant_code" field.
func TenantCodeNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldTenantCode, vs...))
}

// TenantCodeGT applies the GT predicate on the "tenant_code" field.
func TenantCodeGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldTenantCode, v))
}

// TenantCodeGTE applies the GTE predicate on the "tenant_code" field.
func TenantCodeGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldTenantCode, v))
}

// TenantCodeLT applies the LT predicate on the "tenant_code" field.
func TenantCodeLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldTenantCode, v))
}

// TenantCodeLTE applies the LTE predicate on the "tenant_code" field.
func TenantCodeLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldTenantCode, v))
}

// TenantCodeContains applies the Contains predicate on the "tenant_code" field.
func TenantCodeContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldTenantCode, v))
}

// TenantCodeHasPrefix applies the HasPrefix predicate on the "tenant_code" field.
func TenantCodeHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldTenantCode, v))
}

// TenantCodeHasSuffix applies the HasSuffix predicate on the "tenant_code" field.
func TenantCodeHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldTenantCode, v))
}

// TenantCodeEqualFold applies the EqualFold predicate on the "tenant_code" field.
func TenantCodeEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldTenantCode, v))
}

// TenantCodeContainsFold applies the ContainsFold predicate on the "tenant_code" field.
func TenantCodeContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldTenantCode, v))
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldDepartmentID, v))
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldDepartmentID, v))
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldDepartmentID, vs...))
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldDepartmentID, vs...))
}

// DepartmentIDGT applies the GT predicate on the "department_id" field.
func DepartmentIDGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldDepartmentID, v))
}

// DepartmentIDGTE applies the GTE predicate on the "department_id" field.
func DepartmentIDGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldDepartmentID, v))
}

// DepartmentIDLT applies the LT predicate on the "department_id" field.
func DepartmentIDLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldDepartmentID, v))
}

// DepartmentIDLTE applies the LTE predicate on the "department_id" field.
func DepartmentIDLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldDepartmentID, v))
}

// DepartmentIDContains applies the Contains predicate on the "department_id" field.
func DepartmentIDContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldDepartmentID, v))
}

// DepartmentIDHasPrefix applies the HasPrefix predicate on the "department_id" field.
func DepartmentIDHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldDepartmentID, v))
}

// DepartmentIDHasSuffix applies the HasSuffix predicate on the "department_id" field.
func DepartmentIDHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldDepartmentID, v))
}

// DepartmentIDEqualFold applies the EqualFold predicate on the "department_id" field.
func DepartmentIDEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldDepartmentID, v))
}

// DepartmentIDContainsFold applies the ContainsFold predicate on the "department_id" field.
func DepartmentIDContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldDepartmentID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldName, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v string) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...string) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v string) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v string) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v string) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v string) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldParentID, v))
}

// ParentIDContains applies the Contains predicate on the "parent_id" field.
func ParentIDContains(v string) predicate.Department {
	return predicate.Department(sql.FieldContains(FieldParentID, v))
}

// ParentIDHasPrefix applies the HasPrefix predicate on the "parent_id" field.
func ParentIDHasPrefix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasPrefix(FieldParentID, v))
}

// ParentIDHasSuffix applies the HasSuffix predicate on the "parent_id" field.
func ParentIDHasSuffix(v string) predicate.Department {
	return predicate.Department(sql.FieldHasSuffix(FieldParentID, v))
}

// ParentIDEqualFold applies the EqualFold predicate on the "parent_id" field.
func ParentIDEqualFold(v string) predicate.Department {
	return predicate.Department(sql.FieldEqualFold(FieldParentID, v))
}

// ParentIDContainsFold applies the ContainsFold predicate on the "parent_id" field.
func ParentIDContainsFold(v string) predicate.Department {
	return predicate.Department(sql.FieldContainsFold(FieldParentID, v))
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldEQ(FieldSort, v))
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.Department {
	return predicate.Department(sql.FieldNEQ(FieldSort, v))
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldIn(FieldSort, vs...))
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.Department {
	return predicate.Department(sql.FieldNotIn(FieldSort, vs...))
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.Department {
	return predicate.Department(sql.FieldGT(FieldSort, v))
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.Department {
	return predicate.Department(sql.FieldGTE(FieldSort, v))
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.Department {
	return predicate.Department(sql.FieldLT(FieldSort, v))
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.Department {
	return predicate.Department(sql.FieldLTE(FieldSort, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Department) predicate.Department {
	return predicate.Department(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Department) predicate.Department {
	return predicate.Department(sql.NotPredicates(p))
}
