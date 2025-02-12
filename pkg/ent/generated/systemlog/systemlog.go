// Code generated by ent, DO NOT EDIT.

package systemlog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the systemlog type in the database.
	Label = "system_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTenantCode holds the string denoting the tenant_code field in the database.
	FieldTenantCode = "tenant_code"
	// FieldModule holds the string denoting the module field in the database.
	FieldModule = "module"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldOperator holds the string denoting the operator field in the database.
	FieldOperator = "operator"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// Table holds the table name of the systemlog in the database.
	Table = "system_logs"
)

// Columns holds all SQL columns for systemlog fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTenantCode,
	FieldModule,
	FieldAction,
	FieldContent,
	FieldOperator,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt int64
	// TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	TenantCodeValidator func(string) error
	// DefaultModule holds the default value on creation for the "module" field.
	DefaultModule string
	// DefaultAction holds the default value on creation for the "action" field.
	DefaultAction string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultOperator holds the default value on creation for the "operator" field.
	DefaultOperator string
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID string
)

// OrderOption defines the ordering options for the SystemLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTenantCode orders the results by the tenant_code field.
func ByTenantCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantCode, opts...).ToFunc()
}

// ByModule orders the results by the module field.
func ByModule(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModule, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByOperator orders the results by the operator field.
func ByOperator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperator, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}
