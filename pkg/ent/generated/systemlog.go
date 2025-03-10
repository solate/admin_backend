// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/systemlog"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 系统日志
type SystemLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 租户编码
	TenantCode string `json:"tenant_code,omitempty"`
	// 所属模块
	Module string `json:"module,omitempty"`
	// 操作类型
	Action string `json:"action,omitempty"`
	// 操作内容
	Content string `json:"content,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty"`
	// 用户ID
	UserID       string `json:"user_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SystemLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case systemlog.FieldID, systemlog.FieldCreatedAt:
			values[i] = new(sql.NullInt64)
		case systemlog.FieldTenantCode, systemlog.FieldModule, systemlog.FieldAction, systemlog.FieldContent, systemlog.FieldOperator, systemlog.FieldUserID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SystemLog fields.
func (sl *SystemLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case systemlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sl.ID = int(value.Int64)
		case systemlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sl.CreatedAt = value.Int64
			}
		case systemlog.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				sl.TenantCode = value.String
			}
		case systemlog.FieldModule:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field module", values[i])
			} else if value.Valid {
				sl.Module = value.String
			}
		case systemlog.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				sl.Action = value.String
			}
		case systemlog.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				sl.Content = value.String
			}
		case systemlog.FieldOperator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operator", values[i])
			} else if value.Valid {
				sl.Operator = value.String
			}
		case systemlog.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				sl.UserID = value.String
			}
		default:
			sl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SystemLog.
// This includes values selected through modifiers, order, etc.
func (sl *SystemLog) Value(name string) (ent.Value, error) {
	return sl.selectValues.Get(name)
}

// Update returns a builder for updating this SystemLog.
// Note that you need to call SystemLog.Unwrap() before calling this method if this SystemLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (sl *SystemLog) Update() *SystemLogUpdateOne {
	return NewSystemLogClient(sl.config).UpdateOne(sl)
}

// Unwrap unwraps the SystemLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sl *SystemLog) Unwrap() *SystemLog {
	_tx, ok := sl.config.driver.(*txDriver)
	if !ok {
		panic("generated: SystemLog is not a transactional entity")
	}
	sl.config.driver = _tx.drv
	return sl
}

// String implements the fmt.Stringer.
func (sl *SystemLog) String() string {
	var builder strings.Builder
	builder.WriteString("SystemLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sl.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", sl.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("tenant_code=")
	builder.WriteString(sl.TenantCode)
	builder.WriteString(", ")
	builder.WriteString("module=")
	builder.WriteString(sl.Module)
	builder.WriteString(", ")
	builder.WriteString("action=")
	builder.WriteString(sl.Action)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(sl.Content)
	builder.WriteString(", ")
	builder.WriteString("operator=")
	builder.WriteString(sl.Operator)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(sl.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// SystemLogs is a parsable slice of SystemLog.
type SystemLogs []*SystemLog
