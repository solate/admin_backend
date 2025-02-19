// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/department"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dc *DepartmentCreate) SetCreatedAt(i int64) *DepartmentCreate {
	dc.mutation.SetCreatedAt(i)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreatedAt(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetCreatedAt(*i)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DepartmentCreate) SetUpdatedAt(i int64) *DepartmentCreate {
	dc.mutation.SetUpdatedAt(i)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdatedAt(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetUpdatedAt(*i)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DepartmentCreate) SetDeletedAt(i int64) *DepartmentCreate {
	dc.mutation.SetDeletedAt(i)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableDeletedAt(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetDeletedAt(*i)
	}
	return dc
}

// SetTenantCode sets the "tenant_code" field.
func (dc *DepartmentCreate) SetTenantCode(s string) *DepartmentCreate {
	dc.mutation.SetTenantCode(s)
	return dc
}

// SetDepartmentID sets the "department_id" field.
func (dc *DepartmentCreate) SetDepartmentID(s string) *DepartmentCreate {
	dc.mutation.SetDepartmentID(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetParentID sets the "parent_id" field.
func (dc *DepartmentCreate) SetParentID(s string) *DepartmentCreate {
	dc.mutation.SetParentID(s)
	return dc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableParentID(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetParentID(*s)
	}
	return dc
}

// SetSort sets the "sort" field.
func (dc *DepartmentCreate) SetSort(i int) *DepartmentCreate {
	dc.mutation.SetSort(i)
	return dc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableSort(i *int) *DepartmentCreate {
	if i != nil {
		dc.SetSort(*i)
	}
	return dc
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := department.DefaultCreatedAt
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := department.DefaultUpdatedAt
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.ParentID(); !ok {
		v := department.DefaultParentID
		dc.mutation.SetParentID(v)
	}
	if _, ok := dc.mutation.Sort(); !ok {
		v := department.DefaultSort
		dc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Department.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Department.updated_at"`)}
	}
	if _, ok := dc.mutation.TenantCode(); !ok {
		return &ValidationError{Name: "tenant_code", err: errors.New(`generated: missing required field "Department.tenant_code"`)}
	}
	if v, ok := dc.mutation.TenantCode(); ok {
		if err := department.TenantCodeValidator(v); err != nil {
			return &ValidationError{Name: "tenant_code", err: fmt.Errorf(`generated: validator failed for field "Department.tenant_code": %w`, err)}
		}
	}
	if _, ok := dc.mutation.DepartmentID(); !ok {
		return &ValidationError{Name: "department_id", err: errors.New(`generated: missing required field "Department.department_id"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Department.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Department.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`generated: missing required field "Department.parent_id"`)}
	}
	if _, ok := dc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`generated: missing required field "Department.sort"`)}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dc.conflict
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(department.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(department.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.SetField(department.FieldDeletedAt, field.TypeInt64, value)
		_node.DeletedAt = &value
	}
	if value, ok := dc.mutation.TenantCode(); ok {
		_spec.SetField(department.FieldTenantCode, field.TypeString, value)
		_node.TenantCode = value
	}
	if value, ok := dc.mutation.DepartmentID(); ok {
		_spec.SetField(department.FieldDepartmentID, field.TypeString, value)
		_node.DepartmentID = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.ParentID(); ok {
		_spec.SetField(department.FieldParentID, field.TypeString, value)
		_node.ParentID = value
	}
	if value, ok := dc.mutation.Sort(); ok {
		_spec.SetField(department.FieldSort, field.TypeInt, value)
		_node.Sort = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Department.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DepartmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dc *DepartmentCreate) OnConflict(opts ...sql.ConflictOption) *DepartmentUpsertOne {
	dc.conflict = opts
	return &DepartmentUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DepartmentCreate) OnConflictColumns(columns ...string) *DepartmentUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DepartmentUpsertOne{
		create: dc,
	}
}

type (
	// DepartmentUpsertOne is the builder for "upsert"-ing
	//  one Department node.
	DepartmentUpsertOne struct {
		create *DepartmentCreate
	}

	// DepartmentUpsert is the "OnConflict" setter.
	DepartmentUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsert) SetUpdatedAt(v int64) *DepartmentUpsert {
	u.Set(department.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateUpdatedAt() *DepartmentUpsert {
	u.SetExcluded(department.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DepartmentUpsert) AddUpdatedAt(v int64) *DepartmentUpsert {
	u.Add(department.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsert) SetDeletedAt(v int64) *DepartmentUpsert {
	u.Set(department.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateDeletedAt() *DepartmentUpsert {
	u.SetExcluded(department.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsert) AddDeletedAt(v int64) *DepartmentUpsert {
	u.Add(department.FieldDeletedAt, v)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DepartmentUpsert) ClearDeletedAt() *DepartmentUpsert {
	u.SetNull(department.FieldDeletedAt)
	return u
}

// SetTenantCode sets the "tenant_code" field.
func (u *DepartmentUpsert) SetTenantCode(v string) *DepartmentUpsert {
	u.Set(department.FieldTenantCode, v)
	return u
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateTenantCode() *DepartmentUpsert {
	u.SetExcluded(department.FieldTenantCode)
	return u
}

// SetDepartmentID sets the "department_id" field.
func (u *DepartmentUpsert) SetDepartmentID(v string) *DepartmentUpsert {
	u.Set(department.FieldDepartmentID, v)
	return u
}

// UpdateDepartmentID sets the "department_id" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateDepartmentID() *DepartmentUpsert {
	u.SetExcluded(department.FieldDepartmentID)
	return u
}

// SetName sets the "name" field.
func (u *DepartmentUpsert) SetName(v string) *DepartmentUpsert {
	u.Set(department.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateName() *DepartmentUpsert {
	u.SetExcluded(department.FieldName)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *DepartmentUpsert) SetParentID(v string) *DepartmentUpsert {
	u.Set(department.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateParentID() *DepartmentUpsert {
	u.SetExcluded(department.FieldParentID)
	return u
}

// SetSort sets the "sort" field.
func (u *DepartmentUpsert) SetSort(v int) *DepartmentUpsert {
	u.Set(department.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DepartmentUpsert) UpdateSort() *DepartmentUpsert {
	u.SetExcluded(department.FieldSort)
	return u
}

// AddSort adds v to the "sort" field.
func (u *DepartmentUpsert) AddSort(v int) *DepartmentUpsert {
	u.Add(department.FieldSort, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DepartmentUpsertOne) UpdateNewValues() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(department.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DepartmentUpsertOne) Ignore() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DepartmentUpsertOne) DoNothing() *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DepartmentCreate.OnConflict
// documentation for more info.
func (u *DepartmentUpsertOne) Update(set func(*DepartmentUpsert)) *DepartmentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DepartmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsertOne) SetUpdatedAt(v int64) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DepartmentUpsertOne) AddUpdatedAt(v int64) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateUpdatedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsertOne) SetDeletedAt(v int64) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsertOne) AddDeletedAt(v int64) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateDeletedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DepartmentUpsertOne) ClearDeletedAt() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *DepartmentUpsertOne) SetTenantCode(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateTenantCode() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateTenantCode()
	})
}

// SetDepartmentID sets the "department_id" field.
func (u *DepartmentUpsertOne) SetDepartmentID(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDepartmentID(v)
	})
}

// UpdateDepartmentID sets the "department_id" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateDepartmentID() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDepartmentID()
	})
}

// SetName sets the "name" field.
func (u *DepartmentUpsertOne) SetName(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateName() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *DepartmentUpsertOne) SetParentID(v string) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateParentID() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateParentID()
	})
}

// SetSort sets the "sort" field.
func (u *DepartmentUpsertOne) SetSort(v int) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *DepartmentUpsertOne) AddSort(v int) *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DepartmentUpsertOne) UpdateSort() *DepartmentUpsertOne {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *DepartmentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for DepartmentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DepartmentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DepartmentUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DepartmentUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	err      error
	builders []*DepartmentCreate
	conflict []sql.ConflictOption
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Department.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DepartmentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dcb *DepartmentCreateBulk) OnConflict(opts ...sql.ConflictOption) *DepartmentUpsertBulk {
	dcb.conflict = opts
	return &DepartmentUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DepartmentCreateBulk) OnConflictColumns(columns ...string) *DepartmentUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DepartmentUpsertBulk{
		create: dcb,
	}
}

// DepartmentUpsertBulk is the builder for "upsert"-ing
// a bulk of Department nodes.
type DepartmentUpsertBulk struct {
	create *DepartmentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DepartmentUpsertBulk) UpdateNewValues() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(department.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Department.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DepartmentUpsertBulk) Ignore() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DepartmentUpsertBulk) DoNothing() *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DepartmentCreateBulk.OnConflict
// documentation for more info.
func (u *DepartmentUpsertBulk) Update(set func(*DepartmentUpsert)) *DepartmentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DepartmentUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DepartmentUpsertBulk) SetUpdatedAt(v int64) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DepartmentUpsertBulk) AddUpdatedAt(v int64) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateUpdatedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DepartmentUpsertBulk) SetDeletedAt(v int64) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DepartmentUpsertBulk) AddDeletedAt(v int64) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateDeletedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DepartmentUpsertBulk) ClearDeletedAt() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTenantCode sets the "tenant_code" field.
func (u *DepartmentUpsertBulk) SetTenantCode(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetTenantCode(v)
	})
}

// UpdateTenantCode sets the "tenant_code" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateTenantCode() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateTenantCode()
	})
}

// SetDepartmentID sets the "department_id" field.
func (u *DepartmentUpsertBulk) SetDepartmentID(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetDepartmentID(v)
	})
}

// UpdateDepartmentID sets the "department_id" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateDepartmentID() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateDepartmentID()
	})
}

// SetName sets the "name" field.
func (u *DepartmentUpsertBulk) SetName(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateName() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateName()
	})
}

// SetParentID sets the "parent_id" field.
func (u *DepartmentUpsertBulk) SetParentID(v string) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateParentID() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateParentID()
	})
}

// SetSort sets the "sort" field.
func (u *DepartmentUpsertBulk) SetSort(v int) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.SetSort(v)
	})
}

// AddSort adds v to the "sort" field.
func (u *DepartmentUpsertBulk) AddSort(v int) *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.AddSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *DepartmentUpsertBulk) UpdateSort() *DepartmentUpsertBulk {
	return u.Update(func(s *DepartmentUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *DepartmentUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("generated: OnConflict was set for builder %d. Set it on the DepartmentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("generated: missing options for DepartmentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DepartmentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
