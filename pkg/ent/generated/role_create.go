// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"admin_backend/pkg/ent/generated/role"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoleCreate) SetCreatedAt(i int) *RoleCreate {
	rc.mutation.SetCreatedAt(i)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreatedAt(i *int) *RoleCreate {
	if i != nil {
		rc.SetCreatedAt(*i)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoleCreate) SetUpdatedAt(i int) *RoleCreate {
	rc.mutation.SetUpdatedAt(i)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdatedAt(i *int) *RoleCreate {
	if i != nil {
		rc.SetUpdatedAt(*i)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RoleCreate) SetDeletedAt(i int) *RoleCreate {
	rc.mutation.SetDeletedAt(i)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDeletedAt(i *int) *RoleCreate {
	if i != nil {
		rc.SetDeletedAt(*i)
	}
	return rc
}

// SetTenantCode sets the "tenant_code" field.
func (rc *RoleCreate) SetTenantCode(s string) *RoleCreate {
	rc.mutation.SetTenantCode(s)
	return rc
}

// SetRoleID sets the "role_id" field.
func (rc *RoleCreate) SetRoleID(u uint64) *RoleCreate {
	rc.mutation.SetRoleID(u)
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetCode sets the "code" field.
func (rc *RoleCreate) SetCode(s string) *RoleCreate {
	rc.mutation.SetCode(s)
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoleCreate) SetDescription(s string) *RoleCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDescription(s *string) *RoleCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoleCreate) SetStatus(i int) *RoleCreate {
	rc.mutation.SetStatus(i)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoleCreate) SetNillableStatus(i *int) *RoleCreate {
	if i != nil {
		rc.SetStatus(*i)
	}
	return rc
}

// SetSort sets the "sort" field.
func (rc *RoleCreate) SetSort(i int) *RoleCreate {
	rc.mutation.SetSort(i)
	return rc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (rc *RoleCreate) SetNillableSort(i *int) *RoleCreate {
	if i != nil {
		rc.SetSort(*i)
	}
	return rc
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := role.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := role.DefaultUpdatedAt
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := role.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.Sort(); !ok {
		v := role.DefaultSort
		rc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Role.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Role.updated_at"`)}
	}
	if _, ok := rc.mutation.TenantCode(); !ok {
		return &ValidationError{Name: "tenant_code", err: errors.New(`generated: missing required field "Role.tenant_code"`)}
	}
	if v, ok := rc.mutation.TenantCode(); ok {
		if err := role.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Role.tenant_code": %w`, err)}
		}
	}
	if _, ok := rc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`generated: missing required field "Role.role_id"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Role.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Role.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`generated: missing required field "Role.code"`)}
	}
	if v, ok := rc.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Role.code": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "Role.status"`)}
	}
	if _, ok := rc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`generated: missing required field "Role.sort"`)}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt))
	)
	_spec.OnConflict = rc.conflict
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(role.FieldCreatedAt, field.TypeInt, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeInt, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(role.FieldDeletedAt, field.TypeInt, value)
		_node.DeletedAt = &value
	}
	if value, ok := rc.mutation.TenantCode(); ok {
		_spec.SetField(role.FieldTenantCode, field.TypeString, value)
		_node.TenantCode = value
	}
	if value, ok := rc.mutation.RoleID(); ok {
		_spec.SetField(role.FieldRoleID, field.TypeUint64, value)
		_node.RoleID = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt, value)
		_node.Sort = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RoleCreate) OnConflict(opts ...sql.ConflictOption) *RoleUpsertOne {
	rc.conflict = opts
	return &RoleUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RoleCreate) OnConflictColumns(columns ...string) *RoleUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertOne{
		create: rc,
	}
}

type (
	// RoleUpsertOne is the builder for "upsert"-ing
	//  one Role node.
	RoleUpsertOne struct {
		create *RoleCreate
	}

	// RoleUpsert is the "OnConflict" setter.
	RoleUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsert) SetUpdatedAt(v int) *RoleUpsert {
	u.Set(role.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateUpdatedAt() *RoleUpsert {
	u.SetExcluded(role.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RoleUpsert) AddUpdatedAt(v int) *RoleUpsert {
	u.Add(role.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsert) SetDeletedAt(v int) *RoleUpsert {
	u.Set(role.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDeletedAt() *RoleUpsert {
	u.SetExcluded(role.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsert) AddDeletedAt(v int) *RoleUpsert {
	u.Add(role.FieldDeletedAt, v)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsert) ClearDeletedAt() *RoleUpsert {
	u.SetNull(role.FieldDeletedAt)
	return u
}

// SetTenantCode sets the "tenant_code" field.
func (u *RoleUpsert) SetTenantCode(v string) *RoleUpsert {
	u.Set(role.FieldTenantCode, v)
	return u
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *RoleUpsert) UpdateTenantCode() *RoleUpsert {
	u.SetExcluded(role.FieldTenantCode)
	return u
}

// SetName sets the "name" field.
func (u *RoleUpsert) SetName(v string) *RoleUpsert {
	u.Set(role.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsert) UpdateName() *RoleUpsert {
	u.SetExcluded(role.FieldName)
	return u
}

// SetCode sets the "code" field.
func (u *RoleUpsert) SetCode(v string) *RoleUpsert {
	u.Set(role.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RoleUpsert) UpdateCode() *RoleUpsert {
	u.SetExcluded(role.FieldCode)
	return u
}

// SetDescription sets the "description" field.
func (u *RoleUpsert) SetDescription(v string) *RoleUpsert {
	u.Set(role.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsert) UpdateDescription() *RoleUpsert {
	u.SetExcluded(role.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsert) ClearDescription() *RoleUpsert {
	u.SetNull(role.FieldDescription)
	return u
}

// SetStatus sets the "status" field.
func (u *RoleUpsert) SetStatus(v int) *RoleUpsert {
	u.Set(role.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RoleUpsert) UpdateStatus() *RoleUpsert {
	u.SetExcluded(role.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *RoleUpsert) AddStatus(v int) *RoleUpsert {
	u.Add(role.FieldStatus, v)
	return u
}

// SetSort sets the "sort" field.
func (u *RoleUpsert) SetSort(v int) *RoleUpsert {
	u.Set(role.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsert) UpdateSort() *RoleUpsert {
	u.SetExcluded(role.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsert) AddSort(v int) *RoleUpsert {
	u.Add(role.FieldSort, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RoleUpsertOne) UpdateNewValues() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(role.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.RoleID(); exists {
			s.SetIgnore(role.FieldRoleID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RoleUpsertOne) Ignore() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertOne) DoNothing() *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreate.OnConflict
// documentation for more info.
func (u *RoleUpsertOne) Update(set func(*RoleUpsert)) *RoleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertOne) SetUpdatedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RoleUpsertOne) AddUpdatedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateUpdatedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertOne) SetDeletedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsertOne) AddDeletedAt(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDeletedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsertOne) ClearDeletedAt() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *RoleUpsertOne) SetTenantCode(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateTenantCode() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateTenantCode()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertOne) SetName(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateName() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetCode sets the "code" field.
func (u *RoleUpsertOne) SetCode(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateCode() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateCode()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertOne) SetDescription(v string) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertOne) ClearDescription() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetStatus sets the "status" field.
func (u *RoleUpsertOne) SetStatus(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *RoleUpsertOne) AddStatus(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateStatus() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateStatus()
	})
}

// SetSort sets the "sort" field.
func (u *RoleUpsertOne) SetSort(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsertOne) AddSort(v int) *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsertOne) UpdateSort() *RoleUpsertOne {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *RoleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for RoleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RoleUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RoleUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	err      error
	builders []*RoleCreate
	conflict []sql.ConflictOption
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Role.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RoleUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflict(opts ...sql.ConflictOption) *RoleUpsertBulk {
	rcb.conflict = opts
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RoleCreateBulk) OnConflictColumns(columns ...string) *RoleUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RoleUpsertBulk{
		create: rcb,
	}
}

// RoleUpsertBulk is the builder for "upsert"-ing
// a bulk of Role nodes.
type RoleUpsertBulk struct {
	create *RoleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RoleUpsertBulk) UpdateNewValues() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(role.FieldCreatedAt)
			}
			if _, exists := b.mutation.RoleID(); exists {
				s.SetIgnore(role.FieldRoleID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Role.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RoleUpsertBulk) Ignore() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RoleUpsertBulk) DoNothing() *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RoleCreateBulk.OnConflict
// documentation for more info.
func (u *RoleUpsertBulk) Update(set func(*RoleUpsert)) *RoleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RoleUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RoleUpsertBulk) SetUpdatedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RoleUpsertBulk) AddUpdatedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateUpdatedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RoleUpsertBulk) SetDeletedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RoleUpsertBulk) AddDeletedAt(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDeletedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *RoleUpsertBulk) ClearDeletedAt() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *RoleUpsertBulk) SetTenantCode(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateTenantCode() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateTenantCode()
	})
}

// SetName sets the "name" field.
func (u *RoleUpsertBulk) SetName(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateName() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateName()
	})
}

// SetCode sets the "code" field.
func (u *RoleUpsertBulk) SetCode(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateCode() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateCode()
	})
}

// SetDescription sets the "description" field.
func (u *RoleUpsertBulk) SetDescription(v string) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *RoleUpsertBulk) ClearDescription() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.ClearDescription()
	})
}

// SetStatus sets the "status" field.
func (u *RoleUpsertBulk) SetStatus(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *RoleUpsertBulk) AddStatus(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateStatus() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateStatus()
	})
}

// SetSort sets the "sort" field.
func (u *RoleUpsertBulk) SetSort(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *RoleUpsertBulk) AddSort(v int) *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *RoleUpsertBulk) UpdateSort() *RoleUpsertBulk {
	return u.Update(func(s *RoleUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *RoleUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the RoleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for RoleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RoleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
