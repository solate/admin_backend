// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/loginlog"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// LoginLog is the model entity for the LoginLog schema.
type LoginLog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 租户编码
	TenantCode string `json:"tenant_code,omitempty"`
	// 日志ID
	LogID string `json:"log_id,omitempty"`
	// 用户ID
	UserID string `json:"user_id,omitempty"`
	// 用户名
	UserName string `json:"user_name,omitempty"`
	// IP地址
	IP string `json:"ip,omitempty"`
	// 消息
	Message string `json:"message,omitempty"`
	// 用户代理
	UserAgent string `json:"user_agent,omitempty"`
	// 浏览器
	Browser string `json:"browser,omitempty"`
	// 操作系统
	Os string `json:"os,omitempty"`
	// 设备
	Device string `json:"device,omitempty"`
	// 登录时间
	LoginTime    int64 `json:"login_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LoginLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case loginlog.FieldID, loginlog.FieldCreatedAt, loginlog.FieldLoginTime:
			values[i] = new(sql.NullInt64)
		case loginlog.FieldTenantCode, loginlog.FieldLogID, loginlog.FieldUserID, loginlog.FieldUserName, loginlog.FieldIP, loginlog.FieldMessage, loginlog.FieldUserAgent, loginlog.FieldBrowser, loginlog.FieldOs, loginlog.FieldDevice:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoginLog fields.
func (ll *LoginLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loginlog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ll.ID = int(value.Int64)
		case loginlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ll.CreatedAt = value.Int64
			}
		case loginlog.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				ll.TenantCode = value.String
			}
		case loginlog.FieldLogID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field log_id", values[i])
			} else if value.Valid {
				ll.LogID = value.String
			}
		case loginlog.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ll.UserID = value.String
			}
		case loginlog.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				ll.UserName = value.String
			}
		case loginlog.FieldIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value.Valid {
				ll.IP = value.String
			}
		case loginlog.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				ll.Message = value.String
			}
		case loginlog.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_agent", values[i])
			} else if value.Valid {
				ll.UserAgent = value.String
			}
		case loginlog.FieldBrowser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field browser", values[i])
			} else if value.Valid {
				ll.Browser = value.String
			}
		case loginlog.FieldOs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field os", values[i])
			} else if value.Valid {
				ll.Os = value.String
			}
		case loginlog.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				ll.Device = value.String
			}
		case loginlog.FieldLoginTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field login_time", values[i])
			} else if value.Valid {
				ll.LoginTime = value.Int64
			}
		default:
			ll.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LoginLog.
// This includes values selected through modifiers, order, etc.
func (ll *LoginLog) Value(name string) (ent.Value, error) {
	return ll.selectValues.Get(name)
}

// Update returns a builder for updating this LoginLog.
// Note that you need to call LoginLog.Unwrap() before calling this method if this LoginLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (ll *LoginLog) Update() *LoginLogUpdateOne {
	return NewLoginLogClient(ll.config).UpdateOne(ll)
}

// Unwrap unwraps the LoginLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ll *LoginLog) Unwrap() *LoginLog {
	_tx, ok := ll.config.driver.(*txDriver)
	if !ok {
		panic("generated: LoginLog is not a transactional entity")
	}
	ll.config.driver = _tx.drv
	return ll
}

// String implements the fmt.Stringer.
func (ll *LoginLog) String() string {
	var builder strings.Builder
	builder.WriteString("LoginLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ll.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ll.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("tenant_code=")
	builder.WriteString(ll.TenantCode)
	builder.WriteString(", ")
	builder.WriteString("log_id=")
	builder.WriteString(ll.LogID)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(ll.UserID)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(ll.UserName)
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(ll.IP)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(ll.Message)
	builder.WriteString(", ")
	builder.WriteString("user_agent=")
	builder.WriteString(ll.UserAgent)
	builder.WriteString(", ")
	builder.WriteString("browser=")
	builder.WriteString(ll.Browser)
	builder.WriteString(", ")
	builder.WriteString("os=")
	builder.WriteString(ll.Os)
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(ll.Device)
	builder.WriteString(", ")
	builder.WriteString("login_time=")
	builder.WriteString(fmt.Sprintf("%v", ll.LoginTime))
	builder.WriteByte(')')
	return builder.String()
}

// LoginLogs is a parsable slice of LoginLog.
type LoginLogs []*LoginLog
