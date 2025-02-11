// Code generated by ent, DO NOT EDIT.

package loginlog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the loginlog type in the database.
	Label = "login_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTenantCode holds the string denoting the tenant_code field in the database.
	FieldTenantCode = "tenant_code"
	// FieldLogID holds the string denoting the log_id field in the database.
	FieldLogID = "log_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldUserName holds the string denoting the user_name field in the database.
	FieldUserName = "user_name"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldUserAgent holds the string denoting the user_agent field in the database.
	FieldUserAgent = "user_agent"
	// FieldBrowser holds the string denoting the browser field in the database.
	FieldBrowser = "browser"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldDevice holds the string denoting the device field in the database.
	FieldDevice = "device"
	// FieldLoginTime holds the string denoting the login_time field in the database.
	FieldLoginTime = "login_time"
	// Table holds the table name of the loginlog in the database.
	Table = "login_log"
)

// Columns holds all SQL columns for loginlog fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTenantCode,
	FieldLogID,
	FieldUserID,
	FieldUserName,
	FieldIP,
	FieldMessage,
	FieldUserAgent,
	FieldBrowser,
	FieldOs,
	FieldDevice,
	FieldLoginTime,
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
	// UserNameValidator is a validator for the "user_name" field. It is called by the builders before save.
	UserNameValidator func(string) error
	// IPValidator is a validator for the "ip" field. It is called by the builders before save.
	IPValidator func(string) error
)

// OrderOption defines the ordering options for the LoginLog queries.
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

// ByLogID orders the results by the log_id field.
func ByLogID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUserName orders the results by the user_name field.
func ByUserName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserName, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByUserAgent orders the results by the user_agent field.
func ByUserAgent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserAgent, opts...).ToFunc()
}

// ByBrowser orders the results by the browser field.
func ByBrowser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrowser, opts...).ToFunc()
}

// ByOs orders the results by the os field.
func ByOs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOs, opts...).ToFunc()
}

// ByDevice orders the results by the device field.
func ByDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevice, opts...).ToFunc()
}

// ByLoginTime orders the results by the login_time field.
func ByLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLoginTime, opts...).ToFunc()
}
