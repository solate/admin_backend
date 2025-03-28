// Code generated by ent, DO NOT EDIT.

package userposition

import (
	"admin_backend/pkg/ent/generated/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldUserID, v))
}

// PositionID applies equality check predicate on the "position_id" field. It's identical to PositionIDEQ.
func PositionID(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldPositionID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldContainsFold(FieldUserID, v))
}

// PositionIDEQ applies the EQ predicate on the "position_id" field.
func PositionIDEQ(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEQ(FieldPositionID, v))
}

// PositionIDNEQ applies the NEQ predicate on the "position_id" field.
func PositionIDNEQ(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNEQ(FieldPositionID, v))
}

// PositionIDIn applies the In predicate on the "position_id" field.
func PositionIDIn(vs ...string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldIn(FieldPositionID, vs...))
}

// PositionIDNotIn applies the NotIn predicate on the "position_id" field.
func PositionIDNotIn(vs ...string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldNotIn(FieldPositionID, vs...))
}

// PositionIDGT applies the GT predicate on the "position_id" field.
func PositionIDGT(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGT(FieldPositionID, v))
}

// PositionIDGTE applies the GTE predicate on the "position_id" field.
func PositionIDGTE(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldGTE(FieldPositionID, v))
}

// PositionIDLT applies the LT predicate on the "position_id" field.
func PositionIDLT(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLT(FieldPositionID, v))
}

// PositionIDLTE applies the LTE predicate on the "position_id" field.
func PositionIDLTE(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldLTE(FieldPositionID, v))
}

// PositionIDContains applies the Contains predicate on the "position_id" field.
func PositionIDContains(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldContains(FieldPositionID, v))
}

// PositionIDHasPrefix applies the HasPrefix predicate on the "position_id" field.
func PositionIDHasPrefix(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldHasPrefix(FieldPositionID, v))
}

// PositionIDHasSuffix applies the HasSuffix predicate on the "position_id" field.
func PositionIDHasSuffix(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldHasSuffix(FieldPositionID, v))
}

// PositionIDEqualFold applies the EqualFold predicate on the "position_id" field.
func PositionIDEqualFold(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldEqualFold(FieldPositionID, v))
}

// PositionIDContainsFold applies the ContainsFold predicate on the "position_id" field.
func PositionIDContainsFold(v string) predicate.UserPosition {
	return predicate.UserPosition(sql.FieldContainsFold(FieldPositionID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPosition) predicate.UserPosition {
	return predicate.UserPosition(sql.NotPredicates(p))
}
