// Code generated by ent, DO NOT EDIT.

package systemlog

import (
	"entgo.io/ent/dialect/sql"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// TenantCode applies equality check predicate on the "tenant_code" field. It's identical to TenantCodeEQ.
func TenantCode(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldTenantCode, v))
}

// Module applies equality check predicate on the "module" field. It's identical to ModuleEQ.
func Module(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldModule, v))
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldAction, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldContent, v))
}

// Operator applies equality check predicate on the "operator" field. It's identical to OperatorEQ.
func Operator(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldOperator, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldUpdatedAt, v))
}

// TenantCodeEQ applies the EQ predicate on the "tenant_code" field.
func TenantCodeEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldTenantCode, v))
}

// TenantCodeNEQ applies the NEQ predicate on the "tenant_code" field.
func TenantCodeNEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldTenantCode, v))
}

// TenantCodeIn applies the In predicate on the "tenant_code" field.
func TenantCodeIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldTenantCode, vs...))
}

// TenantCodeNotIn applies the NotIn predicate on the "tenant_code" field.
func TenantCodeNotIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldTenantCode, vs...))
}

// TenantCodeGT applies the GT predicate on the "tenant_code" field.
func TenantCodeGT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldTenantCode, v))
}

// TenantCodeGTE applies the GTE predicate on the "tenant_code" field.
func TenantCodeGTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldTenantCode, v))
}

// TenantCodeLT applies the LT predicate on the "tenant_code" field.
func TenantCodeLT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldTenantCode, v))
}

// TenantCodeLTE applies the LTE predicate on the "tenant_code" field.
func TenantCodeLTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldTenantCode, v))
}

// TenantCodeContains applies the Contains predicate on the "tenant_code" field.
func TenantCodeContains(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContains(FieldTenantCode, v))
}

// TenantCodeHasPrefix applies the HasPrefix predicate on the "tenant_code" field.
func TenantCodeHasPrefix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasPrefix(FieldTenantCode, v))
}

// TenantCodeHasSuffix applies the HasSuffix predicate on the "tenant_code" field.
func TenantCodeHasSuffix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasSuffix(FieldTenantCode, v))
}

// TenantCodeEqualFold applies the EqualFold predicate on the "tenant_code" field.
func TenantCodeEqualFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEqualFold(FieldTenantCode, v))
}

// TenantCodeContainsFold applies the ContainsFold predicate on the "tenant_code" field.
func TenantCodeContainsFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContainsFold(FieldTenantCode, v))
}

// ModuleEQ applies the EQ predicate on the "module" field.
func ModuleEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldModule, v))
}

// ModuleNEQ applies the NEQ predicate on the "module" field.
func ModuleNEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldModule, v))
}

// ModuleIn applies the In predicate on the "module" field.
func ModuleIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldModule, vs...))
}

// ModuleNotIn applies the NotIn predicate on the "module" field.
func ModuleNotIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldModule, vs...))
}

// ModuleGT applies the GT predicate on the "module" field.
func ModuleGT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldModule, v))
}

// ModuleGTE applies the GTE predicate on the "module" field.
func ModuleGTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldModule, v))
}

// ModuleLT applies the LT predicate on the "module" field.
func ModuleLT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldModule, v))
}

// ModuleLTE applies the LTE predicate on the "module" field.
func ModuleLTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldModule, v))
}

// ModuleContains applies the Contains predicate on the "module" field.
func ModuleContains(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContains(FieldModule, v))
}

// ModuleHasPrefix applies the HasPrefix predicate on the "module" field.
func ModuleHasPrefix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasPrefix(FieldModule, v))
}

// ModuleHasSuffix applies the HasSuffix predicate on the "module" field.
func ModuleHasSuffix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasSuffix(FieldModule, v))
}

// ModuleEqualFold applies the EqualFold predicate on the "module" field.
func ModuleEqualFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEqualFold(FieldModule, v))
}

// ModuleContainsFold applies the ContainsFold predicate on the "module" field.
func ModuleContainsFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContainsFold(FieldModule, v))
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldAction, v))
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldAction, v))
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldAction, vs...))
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldAction, vs...))
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldAction, v))
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldAction, v))
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldAction, v))
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldAction, v))
}

// ActionContains applies the Contains predicate on the "action" field.
func ActionContains(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContains(FieldAction, v))
}

// ActionHasPrefix applies the HasPrefix predicate on the "action" field.
func ActionHasPrefix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasPrefix(FieldAction, v))
}

// ActionHasSuffix applies the HasSuffix predicate on the "action" field.
func ActionHasSuffix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasSuffix(FieldAction, v))
}

// ActionEqualFold applies the EqualFold predicate on the "action" field.
func ActionEqualFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEqualFold(FieldAction, v))
}

// ActionContainsFold applies the ContainsFold predicate on the "action" field.
func ActionContainsFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContainsFold(FieldAction, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContainsFold(FieldContent, v))
}

// OperatorEQ applies the EQ predicate on the "operator" field.
func OperatorEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldOperator, v))
}

// OperatorNEQ applies the NEQ predicate on the "operator" field.
func OperatorNEQ(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldOperator, v))
}

// OperatorIn applies the In predicate on the "operator" field.
func OperatorIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldOperator, vs...))
}

// OperatorNotIn applies the NotIn predicate on the "operator" field.
func OperatorNotIn(vs ...string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldOperator, vs...))
}

// OperatorGT applies the GT predicate on the "operator" field.
func OperatorGT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldOperator, v))
}

// OperatorGTE applies the GTE predicate on the "operator" field.
func OperatorGTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldOperator, v))
}

// OperatorLT applies the LT predicate on the "operator" field.
func OperatorLT(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldOperator, v))
}

// OperatorLTE applies the LTE predicate on the "operator" field.
func OperatorLTE(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldOperator, v))
}

// OperatorContains applies the Contains predicate on the "operator" field.
func OperatorContains(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContains(FieldOperator, v))
}

// OperatorHasPrefix applies the HasPrefix predicate on the "operator" field.
func OperatorHasPrefix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasPrefix(FieldOperator, v))
}

// OperatorHasSuffix applies the HasSuffix predicate on the "operator" field.
func OperatorHasSuffix(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldHasSuffix(FieldOperator, v))
}

// OperatorEqualFold applies the EqualFold predicate on the "operator" field.
func OperatorEqualFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEqualFold(FieldOperator, v))
}

// OperatorContainsFold applies the ContainsFold predicate on the "operator" field.
func OperatorContainsFold(v string) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldContainsFold(FieldOperator, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.SystemLog {
	return predicate.SystemLog(sql.FieldLTE(FieldUserID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SystemLog) predicate.SystemLog {
	return predicate.SystemLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SystemLog) predicate.SystemLog {
	return predicate.SystemLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SystemLog) predicate.SystemLog {
	return predicate.SystemLog(sql.NotPredicates(p))
}
