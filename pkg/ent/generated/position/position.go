// Code generated by ent, DO NOT EDIT.

package position

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the position type in the database.
	Label = "position"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldTenantCode holds the string denoting the tenant_code field in the database.
	FieldTenantCode = "tenant_code"
	// FieldPositionID holds the string denoting the position_id field in the database.
	FieldPositionID = "position_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDepartmentID holds the string denoting the department_id field in the database.
	FieldDepartmentID = "department_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// Table holds the table name of the position in the database.
	Table = "sys_positions"
)

// Columns holds all SQL columns for position fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldTenantCode,
	FieldPositionID,
	FieldName,
	FieldDepartmentID,
	FieldDescription,
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
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// TenantCodeValidator is a validator for the "tenant_code" field. It is called by the builders before save.
	TenantCodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	DepartmentIDValidator func(string) error
)

// OrderOption defines the ordering options for the Position queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByTenantCode orders the results by the tenant_code field.
func ByTenantCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantCode, opts...).ToFunc()
}

// ByPositionID orders the results by the position_id field.
func ByPositionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPositionID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDepartmentID orders the results by the department_id field.
func ByDepartmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartmentID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}
