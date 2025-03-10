// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/permission"
	"admin_backend/pkg/ent/generated/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(i int64) *PermissionUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(i)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableUpdatedAt(i *int64) *PermissionUpdate {
	if i != nil {
		pu.SetUpdatedAt(*i)
	}
	return pu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (pu *PermissionUpdate) AddUpdatedAt(i int64) *PermissionUpdate {
	pu.mutation.AddUpdatedAt(i)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PermissionUpdate) SetDeletedAt(i int64) *PermissionUpdate {
	pu.mutation.ResetDeletedAt()
	pu.mutation.SetDeletedAt(i)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDeletedAt(i *int64) *PermissionUpdate {
	if i != nil {
		pu.SetDeletedAt(*i)
	}
	return pu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (pu *PermissionUpdate) AddDeletedAt(i int64) *PermissionUpdate {
	pu.mutation.AddDeletedAt(i)
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *PermissionUpdate) ClearDeletedAt() *PermissionUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetTenantCode sets the "tenant_code" field.
func (pu *PermissionUpdate) SetTenantCode(s string) *PermissionUpdate {
	pu.mutation.SetTenantCode(s)
	return pu
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableTenantCode(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetTenantCode(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCode sets the "code" field.
func (pu *PermissionUpdate) SetCode(s string) *PermissionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableCode(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetCode(*s)
	}
	return pu
}

// SetType sets the "type" field.
func (pu *PermissionUpdate) SetType(s string) *PermissionUpdate {
	pu.mutation.SetType(s)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableType(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetType(*s)
	}
	return pu
}

// SetResource sets the "resource" field.
func (pu *PermissionUpdate) SetResource(s string) *PermissionUpdate {
	pu.mutation.SetResource(s)
	return pu
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableResource(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetResource(*s)
	}
	return pu
}

// SetAction sets the "action" field.
func (pu *PermissionUpdate) SetAction(s string) *PermissionUpdate {
	pu.mutation.SetAction(s)
	return pu
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableAction(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetAction(*s)
	}
	return pu
}

// SetParentID sets the "parent_id" field.
func (pu *PermissionUpdate) SetParentID(s string) *PermissionUpdate {
	pu.mutation.SetParentID(s)
	return pu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableParentID(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetParentID(*s)
	}
	return pu
}

// ClearParentID clears the value of the "parent_id" field.
func (pu *PermissionUpdate) ClearParentID() *PermissionUpdate {
	pu.mutation.ClearParentID()
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PermissionUpdate) ClearDescription() *PermissionUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetStatus sets the "status" field.
func (pu *PermissionUpdate) SetStatus(i int) *PermissionUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableStatus(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to the "status" field.
func (pu *PermissionUpdate) AddStatus(i int) *PermissionUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// SetMenuID sets the "menu_id" field.
func (pu *PermissionUpdate) SetMenuID(s string) *PermissionUpdate {
	pu.mutation.SetMenuID(s)
	return pu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableMenuID(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetMenuID(*s)
	}
	return pu
}

// ClearMenuID clears the value of the "menu_id" field.
func (pu *PermissionUpdate) ClearMenuID() *PermissionUpdate {
	pu.mutation.ClearMenuID()
	return pu
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.TenantCode(); ok {
		if err := permission.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Permission.tenant_code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Code(); ok {
		if err := permission.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Permission.code": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Resource(); ok {
		if err := permission.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf(`generated: validator failed for field "Permission.resource": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`generated: validator failed for field "Permission.action": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PermissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(permission.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(permission.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(permission.FieldDeletedAt, field.TypeInt64, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(permission.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := pu.mutation.TenantCode(); ok {
		_spec.SetField(permission.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(permission.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(permission.FieldType, field.TypeString, value)
	}
	if value, ok := pu.mutation.Resource(); ok {
		_spec.SetField(permission.FieldResource, field.TypeString, value)
	}
	if value, ok := pu.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if value, ok := pu.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeString, value)
	}
	if pu.mutation.ParentIDCleared() {
		_spec.ClearField(permission.FieldParentID, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(permission.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(permission.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pu.mutation.MenuID(); ok {
		_spec.SetField(permission.FieldMenuID, field.TypeString, value)
	}
	if pu.mutation.MenuIDCleared() {
		_spec.ClearField(permission.FieldMenuID, field.TypeString)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(i int64) *PermissionUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(i)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableUpdatedAt(i *int64) *PermissionUpdateOne {
	if i != nil {
		puo.SetUpdatedAt(*i)
	}
	return puo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (puo *PermissionUpdateOne) AddUpdatedAt(i int64) *PermissionUpdateOne {
	puo.mutation.AddUpdatedAt(i)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PermissionUpdateOne) SetDeletedAt(i int64) *PermissionUpdateOne {
	puo.mutation.ResetDeletedAt()
	puo.mutation.SetDeletedAt(i)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDeletedAt(i *int64) *PermissionUpdateOne {
	if i != nil {
		puo.SetDeletedAt(*i)
	}
	return puo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (puo *PermissionUpdateOne) AddDeletedAt(i int64) *PermissionUpdateOne {
	puo.mutation.AddDeletedAt(i)
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *PermissionUpdateOne) ClearDeletedAt() *PermissionUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetTenantCode sets the "tenant_code" field.
func (puo *PermissionUpdateOne) SetTenantCode(s string) *PermissionUpdateOne {
	puo.mutation.SetTenantCode(s)
	return puo
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableTenantCode(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetTenantCode(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCode sets the "code" field.
func (puo *PermissionUpdateOne) SetCode(s string) *PermissionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableCode(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetCode(*s)
	}
	return puo
}

// SetType sets the "type" field.
func (puo *PermissionUpdateOne) SetType(s string) *PermissionUpdateOne {
	puo.mutation.SetType(s)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableType(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetType(*s)
	}
	return puo
}

// SetResource sets the "resource" field.
func (puo *PermissionUpdateOne) SetResource(s string) *PermissionUpdateOne {
	puo.mutation.SetResource(s)
	return puo
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableResource(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetResource(*s)
	}
	return puo
}

// SetAction sets the "action" field.
func (puo *PermissionUpdateOne) SetAction(s string) *PermissionUpdateOne {
	puo.mutation.SetAction(s)
	return puo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableAction(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetAction(*s)
	}
	return puo
}

// SetParentID sets the "parent_id" field.
func (puo *PermissionUpdateOne) SetParentID(s string) *PermissionUpdateOne {
	puo.mutation.SetParentID(s)
	return puo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableParentID(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetParentID(*s)
	}
	return puo
}

// ClearParentID clears the value of the "parent_id" field.
func (puo *PermissionUpdateOne) ClearParentID() *PermissionUpdateOne {
	puo.mutation.ClearParentID()
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PermissionUpdateOne) ClearDescription() *PermissionUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetStatus sets the "status" field.
func (puo *PermissionUpdateOne) SetStatus(i int) *PermissionUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableStatus(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to the "status" field.
func (puo *PermissionUpdateOne) AddStatus(i int) *PermissionUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// SetMenuID sets the "menu_id" field.
func (puo *PermissionUpdateOne) SetMenuID(s string) *PermissionUpdateOne {
	puo.mutation.SetMenuID(s)
	return puo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableMenuID(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetMenuID(*s)
	}
	return puo
}

// ClearMenuID clears the value of the "menu_id" field.
func (puo *PermissionUpdateOne) ClearMenuID() *PermissionUpdateOne {
	puo.mutation.ClearMenuID()
	return puo
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.TenantCode(); ok {
		if err := permission.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Permission.tenant_code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Code(); ok {
		if err := permission.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Permission.code": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Resource(); ok {
		if err := permission.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf(`generated: validator failed for field "Permission.resource": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Action(); ok {
		if err := permission.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`generated: validator failed for field "Permission.action": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PermissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(permission.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(permission.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(permission.FieldDeletedAt, field.TypeInt64, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(permission.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := puo.mutation.TenantCode(); ok {
		_spec.SetField(permission.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(permission.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(permission.FieldType, field.TypeString, value)
	}
	if value, ok := puo.mutation.Resource(); ok {
		_spec.SetField(permission.FieldResource, field.TypeString, value)
	}
	if value, ok := puo.mutation.Action(); ok {
		_spec.SetField(permission.FieldAction, field.TypeString, value)
	}
	if value, ok := puo.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeString, value)
	}
	if puo.mutation.ParentIDCleared() {
		_spec.ClearField(permission.FieldParentID, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(permission.FieldStatus, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(permission.FieldStatus, field.TypeInt, value)
	}
	if value, ok := puo.mutation.MenuID(); ok {
		_spec.SetField(permission.FieldMenuID, field.TypeString, value)
	}
	if puo.mutation.MenuIDCleared() {
		_spec.ClearField(permission.FieldMenuID, field.TypeString)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
