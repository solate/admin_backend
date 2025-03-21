// Code generated by ent, DO NOT EDIT.

package dicttype

import (
	"admin_backend/pkg/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldDeletedAt, v))
}

// TenantCode applies equality check predicate on the "tenant_code" field. It's identical to TenantCodeEQ.
func TenantCode(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldTenantCode, v))
}

// TypeID applies equality check predicate on the "type_id" field. It's identical to TypeIDEQ.
func TypeID(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldTypeID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldDescription, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldStatus, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int64) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int64) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.DictType {
	return predicate.DictType(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.DictType {
	return predicate.DictType(sql.FieldNotNull(FieldDeletedAt))
}

// TenantCodeEQ applies the EQ predicate on the "tenant_code" field.
func TenantCodeEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldTenantCode, v))
}

// TenantCodeNEQ applies the NEQ predicate on the "tenant_code" field.
func TenantCodeNEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldTenantCode, v))
}

// TenantCodeIn applies the In predicate on the "tenant_code" field.
func TenantCodeIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldTenantCode, vs...))
}

// TenantCodeNotIn applies the NotIn predicate on the "tenant_code" field.
func TenantCodeNotIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldTenantCode, vs...))
}

// TenantCodeGT applies the GT predicate on the "tenant_code" field.
func TenantCodeGT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldTenantCode, v))
}

// TenantCodeGTE applies the GTE predicate on the "tenant_code" field.
func TenantCodeGTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldTenantCode, v))
}

// TenantCodeLT applies the LT predicate on the "tenant_code" field.
func TenantCodeLT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldTenantCode, v))
}

// TenantCodeLTE applies the LTE predicate on the "tenant_code" field.
func TenantCodeLTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldTenantCode, v))
}

// TenantCodeContains applies the Contains predicate on the "tenant_code" field.
func TenantCodeContains(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContains(FieldTenantCode, v))
}

// TenantCodeHasPrefix applies the HasPrefix predicate on the "tenant_code" field.
func TenantCodeHasPrefix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasPrefix(FieldTenantCode, v))
}

// TenantCodeHasSuffix applies the HasSuffix predicate on the "tenant_code" field.
func TenantCodeHasSuffix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasSuffix(FieldTenantCode, v))
}

// TenantCodeEqualFold applies the EqualFold predicate on the "tenant_code" field.
func TenantCodeEqualFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEqualFold(FieldTenantCode, v))
}

// TenantCodeContainsFold applies the ContainsFold predicate on the "tenant_code" field.
func TenantCodeContainsFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContainsFold(FieldTenantCode, v))
}

// TypeIDEQ applies the EQ predicate on the "type_id" field.
func TypeIDEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldTypeID, v))
}

// TypeIDNEQ applies the NEQ predicate on the "type_id" field.
func TypeIDNEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldTypeID, v))
}

// TypeIDIn applies the In predicate on the "type_id" field.
func TypeIDIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldTypeID, vs...))
}

// TypeIDNotIn applies the NotIn predicate on the "type_id" field.
func TypeIDNotIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldTypeID, vs...))
}

// TypeIDGT applies the GT predicate on the "type_id" field.
func TypeIDGT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldTypeID, v))
}

// TypeIDGTE applies the GTE predicate on the "type_id" field.
func TypeIDGTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldTypeID, v))
}

// TypeIDLT applies the LT predicate on the "type_id" field.
func TypeIDLT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldTypeID, v))
}

// TypeIDLTE applies the LTE predicate on the "type_id" field.
func TypeIDLTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldTypeID, v))
}

// TypeIDContains applies the Contains predicate on the "type_id" field.
func TypeIDContains(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContains(FieldTypeID, v))
}

// TypeIDHasPrefix applies the HasPrefix predicate on the "type_id" field.
func TypeIDHasPrefix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasPrefix(FieldTypeID, v))
}

// TypeIDHasSuffix applies the HasSuffix predicate on the "type_id" field.
func TypeIDHasSuffix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasSuffix(FieldTypeID, v))
}

// TypeIDEqualFold applies the EqualFold predicate on the "type_id" field.
func TypeIDEqualFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEqualFold(FieldTypeID, v))
}

// TypeIDContainsFold applies the ContainsFold predicate on the "type_id" field.
func TypeIDContainsFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContainsFold(FieldTypeID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContainsFold(FieldCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.DictType {
	return predicate.DictType(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.DictType {
	return predicate.DictType(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.DictType {
	return predicate.DictType(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.DictType {
	return predicate.DictType(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.DictType {
	return predicate.DictType(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.DictType {
	return predicate.DictType(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.DictType {
	return predicate.DictType(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.DictType {
	return predicate.DictType(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.DictType {
	return predicate.DictType(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.DictType {
	return predicate.DictType(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.DictType {
	return predicate.DictType(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.DictType {
	return predicate.DictType(sql.FieldLTE(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DictType) predicate.DictType {
	return predicate.DictType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DictType) predicate.DictType {
	return predicate.DictType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DictType) predicate.DictType {
	return predicate.DictType(sql.NotPredicates(p))
}
