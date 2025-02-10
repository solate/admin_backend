// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/role"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks     []Hook
	mutation  *RoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(i int) *RoleUpdate {
	ru.mutation.ResetUpdatedAt()
	ru.mutation.SetUpdatedAt(i)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdatedAt(i *int) *RoleUpdate {
	if i != nil {
		ru.SetUpdatedAt(*i)
	}
	return ru
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ru *RoleUpdate) AddUpdatedAt(i int) *RoleUpdate {
	ru.mutation.AddUpdatedAt(i)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RoleUpdate) SetDeletedAt(i int) *RoleUpdate {
	ru.mutation.ResetDeletedAt()
	ru.mutation.SetDeletedAt(i)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDeletedAt(i *int) *RoleUpdate {
	if i != nil {
		ru.SetDeletedAt(*i)
	}
	return ru
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ru *RoleUpdate) AddDeletedAt(i int) *RoleUpdate {
	ru.mutation.AddDeletedAt(i)
	return ru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ru *RoleUpdate) ClearDeletedAt() *RoleUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// SetTenantCode sets the "tenant_code" field.
func (ru *RoleUpdate) SetTenantCode(s string) *RoleUpdate {
	ru.mutation.SetTenantCode(s)
	return ru
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableTenantCode(s *string) *RoleUpdate {
	if s != nil {
		ru.SetTenantCode(*s)
	}
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetCode sets the "code" field.
func (ru *RoleUpdate) SetCode(s string) *RoleUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCode(s *string) *RoleUpdate {
	if s != nil {
		ru.SetCode(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RoleUpdate) SetDescription(s string) *RoleUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableDescription(s *string) *RoleUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// ClearDescription clears the value of the "description" field.
func (ru *RoleUpdate) ClearDescription() *RoleUpdate {
	ru.mutation.ClearDescription()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(i int) *RoleUpdate {
	ru.mutation.ResetStatus()
	ru.mutation.SetStatus(i)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(i *int) *RoleUpdate {
	if i != nil {
		ru.SetStatus(*i)
	}
	return ru
}

// AddStatus adds i to the "status" field.
func (ru *RoleUpdate) AddStatus(i int) *RoleUpdate {
	ru.mutation.AddStatus(i)
	return ru
}

// SetSort sets the "sort" field.
func (ru *RoleUpdate) SetSort(i int) *RoleUpdate {
	ru.mutation.ResetSort()
	ru.mutation.SetSort(i)
	return ru
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableSort(i *int) *RoleUpdate {
	if i != nil {
		ru.SetSort(*i)
	}
	return ru
}

// AddSort adds i to the "sort" field.
func (ru *RoleUpdate) AddSort(i int) *RoleUpdate {
	ru.mutation.AddSort(i)
	return ru
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.TenantCode(); ok {
		if err := role.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Role.tenant_code": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedDeletedAt(); ok {
		_spec.AddField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.ClearField(role.FieldDeletedAt, field.TypeInt)
	}
	if value, ok := ru.mutation.TenantCode(); ok {
		_spec.SetField(role.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
	}
	if ru.mutation.DescriptionCleared() {
		_spec.ClearField(role.FieldDescription, field.TypeString)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeInt, value)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(i int) *RoleUpdateOne {
	ruo.mutation.ResetUpdatedAt()
	ruo.mutation.SetUpdatedAt(i)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdatedAt(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetUpdatedAt(*i)
	}
	return ruo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ruo *RoleUpdateOne) AddUpdatedAt(i int) *RoleUpdateOne {
	ruo.mutation.AddUpdatedAt(i)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RoleUpdateOne) SetDeletedAt(i int) *RoleUpdateOne {
	ruo.mutation.ResetDeletedAt()
	ruo.mutation.SetDeletedAt(i)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDeletedAt(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetDeletedAt(*i)
	}
	return ruo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ruo *RoleUpdateOne) AddDeletedAt(i int) *RoleUpdateOne {
	ruo.mutation.AddDeletedAt(i)
	return ruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruo *RoleUpdateOne) ClearDeletedAt() *RoleUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// SetTenantCode sets the "tenant_code" field.
func (ruo *RoleUpdateOne) SetTenantCode(s string) *RoleUpdateOne {
	ruo.mutation.SetTenantCode(s)
	return ruo
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableTenantCode(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetTenantCode(*s)
	}
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RoleUpdateOne) SetCode(s string) *RoleUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCode(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetCode(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RoleUpdateOne) SetDescription(s string) *RoleUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableDescription(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// ClearDescription clears the value of the "description" field.
func (ruo *RoleUpdateOne) ClearDescription() *RoleUpdateOne {
	ruo.mutation.ClearDescription()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(i int) *RoleUpdateOne {
	ruo.mutation.ResetStatus()
	ruo.mutation.SetStatus(i)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetStatus(*i)
	}
	return ruo
}

// AddStatus adds i to the "status" field.
func (ruo *RoleUpdateOne) AddStatus(i int) *RoleUpdateOne {
	ruo.mutation.AddStatus(i)
	return ruo
}

// SetSort sets the "sort" field.
func (ruo *RoleUpdateOne) SetSort(i int) *RoleUpdateOne {
	ruo.mutation.ResetSort()
	ruo.mutation.SetSort(i)
	return ruo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableSort(i *int) *RoleUpdateOne {
	if i != nil {
		ruo.SetSort(*i)
	}
	return ruo
}

// AddSort adds i to the "sort" field.
func (ruo *RoleUpdateOne) AddSort(i int) *RoleUpdateOne {
	ruo.mutation.AddSort(i)
	return ruo
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.TenantCode(); ok {
		if err := role.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Role.tenant_code": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role.FieldUpdatedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(role.FieldDeletedAt, field.TypeInt, value)
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.ClearField(role.FieldDeletedAt, field.TypeInt)
	}
	if value, ok := ruo.mutation.TenantCode(); ok {
		_spec.SetField(role.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
	}
	if ruo.mutation.DescriptionCleared() {
		_spec.ClearField(role.FieldDescription, field.TypeString)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedStatus(); ok {
		_spec.AddField(role.FieldStatus, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeInt, value)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
