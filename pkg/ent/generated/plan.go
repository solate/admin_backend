// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/plan"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 计划
type Plan struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 修改时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *int64 `json:"deleted_at,omitempty"`
	// 租户code
	TenantCode string `json:"tenant_code,omitempty"`
	// 计划ID
	PlanID string `json:"plan_id,omitempty"`
	// 计划名称
	Name string `json:"name,omitempty"`
	// 计划描述
	Description string `json:"description,omitempty"`
	// 任务分组
	Group string `json:"group,omitempty"`
	// cron表达式
	CronSpec string `json:"cron_spec,omitempty"`
	// 状态: 1:启用, 2:禁用
	Status int `json:"status,omitempty"`
	// 计划类型: 例行 routine/特殊special
	PlanType string `json:"plan_type,omitempty"`
	// 任务优先级
	Priority int `json:"priority,omitempty"`
	// 任务超时时间(秒)
	Timeout int `json:"timeout,omitempty"`
	// 重试次数
	RetryTimes int `json:"retry_times,omitempty"`
	// 重试间隔(秒)
	RetryInterval int `json:"retry_interval,omitempty"`
	// 生效开始时间
	StartTime int64 `json:"start_time,omitempty"`
	// 生效结束时间
	EndTime int64 `json:"end_time,omitempty"`
	// 要执行的命令或方法
	Command string `json:"command,omitempty"`
	// 执行参数，支持JSON格式
	Params       string `json:"params,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Plan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plan.FieldID, plan.FieldCreatedAt, plan.FieldUpdatedAt, plan.FieldDeletedAt, plan.FieldStatus, plan.FieldPriority, plan.FieldTimeout, plan.FieldRetryTimes, plan.FieldRetryInterval, plan.FieldStartTime, plan.FieldEndTime:
			values[i] = new(sql.NullInt64)
		case plan.FieldTenantCode, plan.FieldPlanID, plan.FieldName, plan.FieldDescription, plan.FieldGroup, plan.FieldCronSpec, plan.FieldPlanType, plan.FieldCommand, plan.FieldParams:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Plan fields.
func (pl *Plan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plan.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case plan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pl.CreatedAt = value.Int64
			}
		case plan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pl.UpdatedAt = value.Int64
			}
		case plan.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pl.DeletedAt = new(int64)
				*pl.DeletedAt = value.Int64
			}
		case plan.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				pl.TenantCode = value.String
			}
		case plan.FieldPlanID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plan_id", values[i])
			} else if value.Valid {
				pl.PlanID = value.String
			}
		case plan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case plan.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pl.Description = value.String
			}
		case plan.FieldGroup:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value.Valid {
				pl.Group = value.String
			}
		case plan.FieldCronSpec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cron_spec", values[i])
			} else if value.Valid {
				pl.CronSpec = value.String
			}
		case plan.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pl.Status = int(value.Int64)
			}
		case plan.FieldPlanType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field plan_type", values[i])
			} else if value.Valid {
				pl.PlanType = value.String
			}
		case plan.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				pl.Priority = int(value.Int64)
			}
		case plan.FieldTimeout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field timeout", values[i])
			} else if value.Valid {
				pl.Timeout = int(value.Int64)
			}
		case plan.FieldRetryTimes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field retry_times", values[i])
			} else if value.Valid {
				pl.RetryTimes = int(value.Int64)
			}
		case plan.FieldRetryInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field retry_interval", values[i])
			} else if value.Valid {
				pl.RetryInterval = int(value.Int64)
			}
		case plan.FieldStartTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				pl.StartTime = value.Int64
			}
		case plan.FieldEndTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				pl.EndTime = value.Int64
			}
		case plan.FieldCommand:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field command", values[i])
			} else if value.Valid {
				pl.Command = value.String
			}
		case plan.FieldParams:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field params", values[i])
			} else if value.Valid {
				pl.Params = value.String
			}
		default:
			pl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Plan.
// This includes values selected through modifiers, order, etc.
func (pl *Plan) Value(name string) (ent.Value, error) {
	return pl.selectValues.Get(name)
}

// Update returns a builder for updating this Plan.
// Note that you need to call Plan.Unwrap() before calling this method if this Plan
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Plan) Update() *PlanUpdateOne {
	return NewPlanClient(pl.config).UpdateOne(pl)
}

// Unwrap unwraps the Plan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Plan) Unwrap() *Plan {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("generated: Plan is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Plan) String() string {
	var builder strings.Builder
	builder.WriteString("Plan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pl.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pl.UpdatedAt))
	builder.WriteString(", ")
	if v := pl.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("tenant_code=")
	builder.WriteString(pl.TenantCode)
	builder.WriteString(", ")
	builder.WriteString("plan_id=")
	builder.WriteString(pl.PlanID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pl.Description)
	builder.WriteString(", ")
	builder.WriteString("group=")
	builder.WriteString(pl.Group)
	builder.WriteString(", ")
	builder.WriteString("cron_spec=")
	builder.WriteString(pl.CronSpec)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pl.Status))
	builder.WriteString(", ")
	builder.WriteString("plan_type=")
	builder.WriteString(pl.PlanType)
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", pl.Priority))
	builder.WriteString(", ")
	builder.WriteString("timeout=")
	builder.WriteString(fmt.Sprintf("%v", pl.Timeout))
	builder.WriteString(", ")
	builder.WriteString("retry_times=")
	builder.WriteString(fmt.Sprintf("%v", pl.RetryTimes))
	builder.WriteString(", ")
	builder.WriteString("retry_interval=")
	builder.WriteString(fmt.Sprintf("%v", pl.RetryInterval))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(fmt.Sprintf("%v", pl.StartTime))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(fmt.Sprintf("%v", pl.EndTime))
	builder.WriteString(", ")
	builder.WriteString("command=")
	builder.WriteString(pl.Command)
	builder.WriteString(", ")
	builder.WriteString("params=")
	builder.WriteString(pl.Params)
	builder.WriteByte(')')
	return builder.String()
}

// Plans is a parsable slice of Plan.
type Plans []*Plan
