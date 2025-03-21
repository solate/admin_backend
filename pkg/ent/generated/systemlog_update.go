// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/systemlog"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SystemLogUpdate is the builder for updating SystemLog entities.
type SystemLogUpdate struct {
	config
	hooks     []Hook
	mutation  *SystemLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SystemLogUpdate builder.
func (slu *SystemLogUpdate) Where(ps ...predicate.SystemLog) *SystemLogUpdate {
	slu.mutation.Where(ps...)
	return slu
}

// SetTenantCode sets the "tenant_code" field.
func (slu *SystemLogUpdate) SetTenantCode(s string) *SystemLogUpdate {
	slu.mutation.SetTenantCode(s)
	return slu
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableTenantCode(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetTenantCode(*s)
	}
	return slu
}

// SetModule sets the "module" field.
func (slu *SystemLogUpdate) SetModule(s string) *SystemLogUpdate {
	slu.mutation.SetModule(s)
	return slu
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableModule(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetModule(*s)
	}
	return slu
}

// SetAction sets the "action" field.
func (slu *SystemLogUpdate) SetAction(s string) *SystemLogUpdate {
	slu.mutation.SetAction(s)
	return slu
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableAction(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetAction(*s)
	}
	return slu
}

// SetContent sets the "content" field.
func (slu *SystemLogUpdate) SetContent(s string) *SystemLogUpdate {
	slu.mutation.SetContent(s)
	return slu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableContent(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetContent(*s)
	}
	return slu
}

// SetOperator sets the "operator" field.
func (slu *SystemLogUpdate) SetOperator(s string) *SystemLogUpdate {
	slu.mutation.SetOperator(s)
	return slu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableOperator(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetOperator(*s)
	}
	return slu
}

// SetUserID sets the "user_id" field.
func (slu *SystemLogUpdate) SetUserID(s string) *SystemLogUpdate {
	slu.mutation.SetUserID(s)
	return slu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (slu *SystemLogUpdate) SetNillableUserID(s *string) *SystemLogUpdate {
	if s != nil {
		slu.SetUserID(*s)
	}
	return slu
}

// Mutation returns the SystemLogMutation object of the builder.
func (slu *SystemLogUpdate) Mutation() *SystemLogMutation {
	return slu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (slu *SystemLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, slu.sqlSave, slu.mutation, slu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (slu *SystemLogUpdate) SaveX(ctx context.Context) int {
	affected, err := slu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (slu *SystemLogUpdate) Exec(ctx context.Context) error {
	_, err := slu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slu *SystemLogUpdate) ExecX(ctx context.Context) {
	if err := slu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slu *SystemLogUpdate) check() error {
	if v, ok := slu.mutation.TenantCode(); ok {
		if err := systemlog.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "SystemLog.tenant_code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (slu *SystemLogUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemLogUpdate {
	slu.modifiers = append(slu.modifiers, modifiers...)
	return slu
}

func (slu *SystemLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := slu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(systemlog.Table, systemlog.Columns, sqlgraph.NewFieldSpec(systemlog.FieldID, field.TypeInt))
	if ps := slu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := slu.mutation.TenantCode(); ok {
		_spec.SetField(systemlog.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := slu.mutation.Module(); ok {
		_spec.SetField(systemlog.FieldModule, field.TypeString, value)
	}
	if value, ok := slu.mutation.Action(); ok {
		_spec.SetField(systemlog.FieldAction, field.TypeString, value)
	}
	if value, ok := slu.mutation.Content(); ok {
		_spec.SetField(systemlog.FieldContent, field.TypeString, value)
	}
	if value, ok := slu.mutation.Operator(); ok {
		_spec.SetField(systemlog.FieldOperator, field.TypeString, value)
	}
	if value, ok := slu.mutation.UserID(); ok {
		_spec.SetField(systemlog.FieldUserID, field.TypeString, value)
	}
	_spec.AddModifiers(slu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, slu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	slu.mutation.done = true
	return n, nil
}

// SystemLogUpdateOne is the builder for updating a single SystemLog entity.
type SystemLogUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SystemLogMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTenantCode sets the "tenant_code" field.
func (sluo *SystemLogUpdateOne) SetTenantCode(s string) *SystemLogUpdateOne {
	sluo.mutation.SetTenantCode(s)
	return sluo
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableTenantCode(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetTenantCode(*s)
	}
	return sluo
}

// SetModule sets the "module" field.
func (sluo *SystemLogUpdateOne) SetModule(s string) *SystemLogUpdateOne {
	sluo.mutation.SetModule(s)
	return sluo
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableModule(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetModule(*s)
	}
	return sluo
}

// SetAction sets the "action" field.
func (sluo *SystemLogUpdateOne) SetAction(s string) *SystemLogUpdateOne {
	sluo.mutation.SetAction(s)
	return sluo
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableAction(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetAction(*s)
	}
	return sluo
}

// SetContent sets the "content" field.
func (sluo *SystemLogUpdateOne) SetContent(s string) *SystemLogUpdateOne {
	sluo.mutation.SetContent(s)
	return sluo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableContent(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetContent(*s)
	}
	return sluo
}

// SetOperator sets the "operator" field.
func (sluo *SystemLogUpdateOne) SetOperator(s string) *SystemLogUpdateOne {
	sluo.mutation.SetOperator(s)
	return sluo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableOperator(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetOperator(*s)
	}
	return sluo
}

// SetUserID sets the "user_id" field.
func (sluo *SystemLogUpdateOne) SetUserID(s string) *SystemLogUpdateOne {
	sluo.mutation.SetUserID(s)
	return sluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (sluo *SystemLogUpdateOne) SetNillableUserID(s *string) *SystemLogUpdateOne {
	if s != nil {
		sluo.SetUserID(*s)
	}
	return sluo
}

// Mutation returns the SystemLogMutation object of the builder.
func (sluo *SystemLogUpdateOne) Mutation() *SystemLogMutation {
	return sluo.mutation
}

// Where appends a list predicates to the SystemLogUpdate builder.
func (sluo *SystemLogUpdateOne) Where(ps ...predicate.SystemLog) *SystemLogUpdateOne {
	sluo.mutation.Where(ps...)
	return sluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sluo *SystemLogUpdateOne) Select(field string, fields ...string) *SystemLogUpdateOne {
	sluo.fields = append([]string{field}, fields...)
	return sluo
}

// Save executes the query and returns the updated SystemLog entity.
func (sluo *SystemLogUpdateOne) Save(ctx context.Context) (*SystemLog, error) {
	return withHooks(ctx, sluo.sqlSave, sluo.mutation, sluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sluo *SystemLogUpdateOne) SaveX(ctx context.Context) *SystemLog {
	node, err := sluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sluo *SystemLogUpdateOne) Exec(ctx context.Context) error {
	_, err := sluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sluo *SystemLogUpdateOne) ExecX(ctx context.Context) {
	if err := sluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sluo *SystemLogUpdateOne) check() error {
	if v, ok := sluo.mutation.TenantCode(); ok {
		if err := systemlog.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "SystemLog.tenant_code": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sluo *SystemLogUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SystemLogUpdateOne {
	sluo.modifiers = append(sluo.modifiers, modifiers...)
	return sluo
}

func (sluo *SystemLogUpdateOne) sqlSave(ctx context.Context) (_node *SystemLog, err error) {
	if err := sluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(systemlog.Table, systemlog.Columns, sqlgraph.NewFieldSpec(systemlog.FieldID, field.TypeInt))
	id, ok := sluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "SystemLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemlog.FieldID)
		for _, f := range fields {
			if !systemlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != systemlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sluo.mutation.TenantCode(); ok {
		_spec.SetField(systemlog.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := sluo.mutation.Module(); ok {
		_spec.SetField(systemlog.FieldModule, field.TypeString, value)
	}
	if value, ok := sluo.mutation.Action(); ok {
		_spec.SetField(systemlog.FieldAction, field.TypeString, value)
	}
	if value, ok := sluo.mutation.Content(); ok {
		_spec.SetField(systemlog.FieldContent, field.TypeString, value)
	}
	if value, ok := sluo.mutation.Operator(); ok {
		_spec.SetField(systemlog.FieldOperator, field.TypeString, value)
	}
	if value, ok := sluo.mutation.UserID(); ok {
		_spec.SetField(systemlog.FieldUserID, field.TypeString, value)
	}
	_spec.AddModifiers(sluo.modifiers...)
	_node = &SystemLog{config: sluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sluo.mutation.done = true
	return _node, nil
}
