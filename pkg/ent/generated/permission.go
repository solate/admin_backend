// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/permission"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 权限
type Permission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 修改时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt *int64 `json:"deleted_at,omitempty"`
	// 租户编码
	TenantCode string `json:"tenant_code,omitempty"`
	// 权限ID
	PermissionID string `json:"permission_id,omitempty"`
	// 权限名称
	Name string `json:"name,omitempty"`
	// 权限编码
	Code string `json:"code,omitempty"`
	// 类型类型: dir/menu/button/api/data
	Type string `json:"type,omitempty"`
	// 资源
	Resource string `json:"resource,omitempty"`
	// 操作类型
	Action string `json:"action,omitempty"`
	// 父级ID
	ParentID string `json:"parent_id,omitempty"`
	// 描述
	Description string `json:"description,omitempty"`
	// 状态 1:启用 2:禁用
	Status int `json:"status,omitempty"`
	// 菜单ID
	MenuID       string `json:"menu_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID, permission.FieldCreatedAt, permission.FieldUpdatedAt, permission.FieldDeletedAt, permission.FieldStatus:
			values[i] = new(sql.NullInt64)
		case permission.FieldTenantCode, permission.FieldPermissionID, permission.FieldName, permission.FieldCode, permission.FieldType, permission.FieldResource, permission.FieldAction, permission.FieldParentID, permission.FieldDescription, permission.FieldMenuID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Int64
			}
		case permission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Int64
			}
		case permission.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pe.DeletedAt = new(int64)
				*pe.DeletedAt = value.Int64
			}
		case permission.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				pe.TenantCode = value.String
			}
		case permission.FieldPermissionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permission_id", values[i])
			} else if value.Valid {
				pe.PermissionID = value.String
			}
		case permission.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case permission.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				pe.Code = value.String
			}
		case permission.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pe.Type = value.String
			}
		case permission.FieldResource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource", values[i])
			} else if value.Valid {
				pe.Resource = value.String
			}
		case permission.FieldAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field action", values[i])
			} else if value.Valid {
				pe.Action = value.String
			}
		case permission.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				pe.ParentID = value.String
			}
		case permission.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pe.Description = value.String
			}
		case permission.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pe.Status = int(value.Int64)
			}
		case permission.FieldMenuID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_id", values[i])
			} else if value.Valid {
				pe.MenuID = value.String
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Permission.
// This includes values selected through modifiers, order, etc.
func (pe *Permission) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("generated: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pe.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pe.UpdatedAt))
	builder.WriteString(", ")
	if v := pe.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("tenant_code=")
	builder.WriteString(pe.TenantCode)
	builder.WriteString(", ")
	builder.WriteString("permission_id=")
	builder.WriteString(pe.PermissionID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(pe.Code)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pe.Type)
	builder.WriteString(", ")
	builder.WriteString("resource=")
	builder.WriteString(pe.Resource)
	builder.WriteString(", ")
	builder.WriteString("action=")
	builder.WriteString(pe.Action)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(pe.ParentID)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pe.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pe.Status))
	builder.WriteString(", ")
	builder.WriteString("menu_id=")
	builder.WriteString(pe.MenuID)
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
