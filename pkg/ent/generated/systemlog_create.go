// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/solate/admin_backend/pkg/ent/generated/systemlog"
)

// SystemLogCreate is the builder for creating a SystemLog entity.
type SystemLogCreate struct {
	config
	mutation *SystemLogMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (slc *SystemLogCreate) SetCreatedAt(i int64) *SystemLogCreate {
	slc.mutation.SetCreatedAt(i)
	return slc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableCreatedAt(i *int64) *SystemLogCreate {
	if i != nil {
		slc.SetCreatedAt(*i)
	}
	return slc
}

// SetUpdatedAt sets the "updated_at" field.
func (slc *SystemLogCreate) SetUpdatedAt(i int64) *SystemLogCreate {
	slc.mutation.SetUpdatedAt(i)
	return slc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableUpdatedAt(i *int64) *SystemLogCreate {
	if i != nil {
		slc.SetUpdatedAt(*i)
	}
	return slc
}

// SetTenantCode sets the "tenant_code" field.
func (slc *SystemLogCreate) SetTenantCode(s string) *SystemLogCreate {
	slc.mutation.SetTenantCode(s)
	return slc
}

// SetModule sets the "module" field.
func (slc *SystemLogCreate) SetModule(s string) *SystemLogCreate {
	slc.mutation.SetModule(s)
	return slc
}

// SetNillableModule sets the "module" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableModule(s *string) *SystemLogCreate {
	if s != nil {
		slc.SetModule(*s)
	}
	return slc
}

// SetAction sets the "action" field.
func (slc *SystemLogCreate) SetAction(s string) *SystemLogCreate {
	slc.mutation.SetAction(s)
	return slc
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableAction(s *string) *SystemLogCreate {
	if s != nil {
		slc.SetAction(*s)
	}
	return slc
}

// SetContent sets the "content" field.
func (slc *SystemLogCreate) SetContent(s string) *SystemLogCreate {
	slc.mutation.SetContent(s)
	return slc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableContent(s *string) *SystemLogCreate {
	if s != nil {
		slc.SetContent(*s)
	}
	return slc
}

// SetOperator sets the "operator" field.
func (slc *SystemLogCreate) SetOperator(s string) *SystemLogCreate {
	slc.mutation.SetOperator(s)
	return slc
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableOperator(s *string) *SystemLogCreate {
	if s != nil {
		slc.SetOperator(*s)
	}
	return slc
}

// SetUserID sets the "user_id" field.
func (slc *SystemLogCreate) SetUserID(u uint64) *SystemLogCreate {
	slc.mutation.SetUserID(u)
	return slc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (slc *SystemLogCreate) SetNillableUserID(u *uint64) *SystemLogCreate {
	if u != nil {
		slc.SetUserID(*u)
	}
	return slc
}

// Mutation returns the SystemLogMutation object of the builder.
func (slc *SystemLogCreate) Mutation() *SystemLogMutation {
	return slc.mutation
}

// Save creates the SystemLog in the database.
func (slc *SystemLogCreate) Save(ctx context.Context) (*SystemLog, error) {
	slc.defaults()
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *SystemLogCreate) SaveX(ctx context.Context) *SystemLog {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *SystemLogCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *SystemLogCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *SystemLogCreate) defaults() {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		v := systemlog.DefaultCreatedAt
		slc.mutation.SetCreatedAt(v)
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		v := systemlog.DefaultUpdatedAt
		slc.mutation.SetUpdatedAt(v)
	}
	if _, ok := slc.mutation.Module(); !ok {
		v := systemlog.DefaultModule
		slc.mutation.SetModule(v)
	}
	if _, ok := slc.mutation.Action(); !ok {
		v := systemlog.DefaultAction
		slc.mutation.SetAction(v)
	}
	if _, ok := slc.mutation.Content(); !ok {
		v := systemlog.DefaultContent
		slc.mutation.SetContent(v)
	}
	if _, ok := slc.mutation.Operator(); !ok {
		v := systemlog.DefaultOperator
		slc.mutation.SetOperator(v)
	}
	if _, ok := slc.mutation.UserID(); !ok {
		v := systemlog.DefaultUserID
		slc.mutation.SetUserID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *SystemLogCreate) check() error {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "SystemLog.created_at"`)}
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "SystemLog.updated_at"`)}
	}
	if _, ok := slc.mutation.TenantCode(); !ok {
		return &ValidationError{Name: "tenant_code", err: errors.New(`generated: missing required field "SystemLog.tenant_code"`)}
	}
	if v, ok := slc.mutation.TenantCode(); ok {
		if err := systemlog.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "SystemLog.tenant_code": %w`, err)}
		}
	}
	if _, ok := slc.mutation.Module(); !ok {
		return &ValidationError{Name: "module", err: errors.New(`generated: missing required field "SystemLog.module"`)}
	}
	if _, ok := slc.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`generated: missing required field "SystemLog.action"`)}
	}
	if _, ok := slc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`generated: missing required field "SystemLog.content"`)}
	}
	if _, ok := slc.mutation.Operator(); !ok {
		return &ValidationError{Name: "operator", err: errors.New(`generated: missing required field "SystemLog.operator"`)}
	}
	if _, ok := slc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "SystemLog.user_id"`)}
	}
	return nil
}

func (slc *SystemLogCreate) sqlSave(ctx context.Context) (*SystemLog, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *SystemLogCreate) createSpec() (*SystemLog, *sqlgraph.CreateSpec) {
	var (
		_node = &SystemLog{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(systemlog.Table, sqlgraph.NewFieldSpec(systemlog.FieldID, field.TypeInt))
	)
	_spec.OnConflict = slc.conflict
	if value, ok := slc.mutation.CreatedAt(); ok {
		_spec.SetField(systemlog.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := slc.mutation.UpdatedAt(); ok {
		_spec.SetField(systemlog.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := slc.mutation.TenantCode(); ok {
		_spec.SetField(systemlog.FieldTenantCode, field.TypeString, value)
		_node.TenantCode = value
	}
	if value, ok := slc.mutation.Module(); ok {
		_spec.SetField(systemlog.FieldModule, field.TypeString, value)
		_node.Module = value
	}
	if value, ok := slc.mutation.Action(); ok {
		_spec.SetField(systemlog.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	if value, ok := slc.mutation.Content(); ok {
		_spec.SetField(systemlog.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := slc.mutation.Operator(); ok {
		_spec.SetField(systemlog.FieldOperator, field.TypeString, value)
		_node.Operator = value
	}
	if value, ok := slc.mutation.UserID(); ok {
		_spec.SetField(systemlog.FieldUserID, field.TypeUint64, value)
		_node.UserID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SystemLog.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SystemLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (slc *SystemLogCreate) OnConflict(opts ...sql.ConflictOption) *SystemLogUpsertOne {
	slc.conflict = opts
	return &SystemLogUpsertOne{
		create: slc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (slc *SystemLogCreate) OnConflictColumns(columns ...string) *SystemLogUpsertOne {
	slc.conflict = append(slc.conflict, sql.ConflictColumns(columns...))
	return &SystemLogUpsertOne{
		create: slc,
	}
}

type (
	// SystemLogUpsertOne is the builder for "upsert"-ing
	//  one SystemLog node.
	SystemLogUpsertOne struct {
		create *SystemLogCreate
	}

	// SystemLogUpsert is the "OnConflict" setter.
	SystemLogUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SystemLogUpsert) SetUpdatedAt(v int64) *SystemLogUpsert {
	u.Set(systemlog.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateUpdatedAt() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SystemLogUpsert) AddUpdatedAt(v int64) *SystemLogUpsert {
	u.Add(systemlog.FieldUpdatedAt, v)
	return u
}

// SetTenantCode sets the "tenant_code" field.
func (u *SystemLogUpsert) SetTenantCode(v string) *SystemLogUpsert {
	u.Set(systemlog.FieldTenantCode, v)
	return u
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateTenantCode() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldTenantCode)
	return u
}

// SetModule sets the "module" field.
func (u *SystemLogUpsert) SetModule(v string) *SystemLogUpsert {
	u.Set(systemlog.FieldModule, v)
	return u
}

// UpdateModule sets the "module" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateModule() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldModule)
	return u
}

// SetAction sets the "action" field.
func (u *SystemLogUpsert) SetAction(v string) *SystemLogUpsert {
	u.Set(systemlog.FieldAction, v)
	return u
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateAction() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldAction)
	return u
}

// SetContent sets the "content" field.
func (u *SystemLogUpsert) SetContent(v string) *SystemLogUpsert {
	u.Set(systemlog.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateContent() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldContent)
	return u
}

// SetOperator sets the "operator" field.
func (u *SystemLogUpsert) SetOperator(v string) *SystemLogUpsert {
	u.Set(systemlog.FieldOperator, v)
	return u
}

// UpdateOperator sets the "operator" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateOperator() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldOperator)
	return u
}

// SetUserID sets the "user_id" field.
func (u *SystemLogUpsert) SetUserID(v uint64) *SystemLogUpsert {
	u.Set(systemlog.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SystemLogUpsert) UpdateUserID() *SystemLogUpsert {
	u.SetExcluded(systemlog.FieldUserID)
	return u
}

// AddUserID adds v to the "user_id" field.
func (u *SystemLogUpsert) AddUserID(v uint64) *SystemLogUpsert {
	u.Add(systemlog.FieldUserID, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SystemLogUpsertOne) UpdateNewValues() *SystemLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(systemlog.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SystemLogUpsertOne) Ignore() *SystemLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SystemLogUpsertOne) DoNothing() *SystemLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SystemLogCreate.OnConflict
// documentation for more info.
func (u *SystemLogUpsertOne) Update(set func(*SystemLogUpsert)) *SystemLogUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SystemLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SystemLogUpsertOne) SetUpdatedAt(v int64) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SystemLogUpsertOne) AddUpdatedAt(v int64) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateUpdatedAt() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *SystemLogUpsertOne) SetTenantCode(v string) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateTenantCode() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateTenantCode()
	})
}

// SetModule sets the "module" field.
func (u *SystemLogUpsertOne) SetModule(v string) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetModule(v)
	})
}

// UpdateModule sets the "module" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateModule() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateModule()
	})
}

// SetAction sets the "action" field.
func (u *SystemLogUpsertOne) SetAction(v string) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateAction() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateAction()
	})
}

// SetContent sets the "content" field.
func (u *SystemLogUpsertOne) SetContent(v string) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateContent() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateContent()
	})
}

// SetOperator sets the "operator" field.
func (u *SystemLogUpsertOne) SetOperator(v string) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetOperator(v)
	})
}

// UpdateOperator sets the "operator" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateOperator() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateOperator()
	})
}

// SetUserID sets the "user_id" field.
func (u *SystemLogUpsertOne) SetUserID(v uint64) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *SystemLogUpsertOne) AddUserID(v uint64) *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SystemLogUpsertOne) UpdateUserID() *SystemLogUpsertOne {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *SystemLogUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for SystemLogCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SystemLogUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SystemLogUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SystemLogUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SystemLogCreateBulk is the builder for creating many SystemLog entities in bulk.
type SystemLogCreateBulk struct {
	config
	err      error
	builders []*SystemLogCreate
	conflict []sql.ConflictOption
}

// Save creates the SystemLog entities in the database.
func (slcb *SystemLogCreateBulk) Save(ctx context.Context) ([]*SystemLog, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*SystemLog, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SystemLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = slcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *SystemLogCreateBulk) SaveX(ctx context.Context) []*SystemLog {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *SystemLogCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *SystemLogCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.SystemLog.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SystemLogUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (slcb *SystemLogCreateBulk) OnConflict(opts ...sql.ConflictOption) *SystemLogUpsertBulk {
	slcb.conflict = opts
	return &SystemLogUpsertBulk{
		create: slcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (slcb *SystemLogCreateBulk) OnConflictColumns(columns ...string) *SystemLogUpsertBulk {
	slcb.conflict = append(slcb.conflict, sql.ConflictColumns(columns...))
	return &SystemLogUpsertBulk{
		create: slcb,
	}
}

// SystemLogUpsertBulk is the builder for "upsert"-ing
// a bulk of SystemLog nodes.
type SystemLogUpsertBulk struct {
	create *SystemLogCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SystemLogUpsertBulk) UpdateNewValues() *SystemLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(systemlog.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.SystemLog.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SystemLogUpsertBulk) Ignore() *SystemLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SystemLogUpsertBulk) DoNothing() *SystemLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SystemLogCreateBulk.OnConflict
// documentation for more info.
func (u *SystemLogUpsertBulk) Update(set func(*SystemLogUpsert)) *SystemLogUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SystemLogUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SystemLogUpsertBulk) SetUpdatedAt(v int64) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SystemLogUpsertBulk) AddUpdatedAt(v int64) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateUpdatedAt() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *SystemLogUpsertBulk) SetTenantCode(v string) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateTenantCode() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateTenantCode()
	})
}

// SetModule sets the "module" field.
func (u *SystemLogUpsertBulk) SetModule(v string) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetModule(v)
	})
}

// UpdateModule sets the "module" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateModule() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateModule()
	})
}

// SetAction sets the "action" field.
func (u *SystemLogUpsertBulk) SetAction(v string) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetAction(v)
	})
}

// UpdateAction sets the "action" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateAction() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateAction()
	})
}

// SetContent sets the "content" field.
func (u *SystemLogUpsertBulk) SetContent(v string) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateContent() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateContent()
	})
}

// SetOperator sets the "operator" field.
func (u *SystemLogUpsertBulk) SetOperator(v string) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetOperator(v)
	})
}

// UpdateOperator sets the "operator" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateOperator() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateOperator()
	})
}

// SetUserID sets the "user_id" field.
func (u *SystemLogUpsertBulk) SetUserID(v uint64) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.SetUserID(v)
	})
}

// AddUserID adds v to the "user_id" field.
func (u *SystemLogUpsertBulk) AddUserID(v uint64) *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.AddUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *SystemLogUpsertBulk) UpdateUserID() *SystemLogUpsertBulk {
	return u.Update(func(s *SystemLogUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *SystemLogUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the SystemLogCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for SystemLogCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SystemLogUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
