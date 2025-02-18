// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/menu"
	"admin_backend/pkg/ent/generated/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MenuUpdate) SetUpdatedAt(i int64) *MenuUpdate {
	mu.mutation.ResetUpdatedAt()
	mu.mutation.SetUpdatedAt(i)
	return mu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableUpdatedAt(i *int64) *MenuUpdate {
	if i != nil {
		mu.SetUpdatedAt(*i)
	}
	return mu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (mu *MenuUpdate) AddUpdatedAt(i int64) *MenuUpdate {
	mu.mutation.AddUpdatedAt(i)
	return mu
}

// SetDeletedAt sets the "deleted_at" field.
func (mu *MenuUpdate) SetDeletedAt(i int64) *MenuUpdate {
	mu.mutation.ResetDeletedAt()
	mu.mutation.SetDeletedAt(i)
	return mu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableDeletedAt(i *int64) *MenuUpdate {
	if i != nil {
		mu.SetDeletedAt(*i)
	}
	return mu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (mu *MenuUpdate) AddDeletedAt(i int64) *MenuUpdate {
	mu.mutation.AddDeletedAt(i)
	return mu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mu *MenuUpdate) ClearDeletedAt() *MenuUpdate {
	mu.mutation.ClearDeletedAt()
	return mu
}

// SetTenantCode sets the "tenant_code" field.
func (mu *MenuUpdate) SetTenantCode(s string) *MenuUpdate {
	mu.mutation.SetTenantCode(s)
	return mu
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableTenantCode(s *string) *MenuUpdate {
	if s != nil {
		mu.SetTenantCode(*s)
	}
	return mu
}

// SetMenuID sets the "menu_id" field.
func (mu *MenuUpdate) SetMenuID(s string) *MenuUpdate {
	mu.mutation.SetMenuID(s)
	return mu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableMenuID(s *string) *MenuUpdate {
	if s != nil {
		mu.SetMenuID(*s)
	}
	return mu
}

// SetCode sets the "code" field.
func (mu *MenuUpdate) SetCode(s string) *MenuUpdate {
	mu.mutation.SetCode(s)
	return mu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableCode(s *string) *MenuUpdate {
	if s != nil {
		mu.SetCode(*s)
	}
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(s string) *MenuUpdate {
	mu.mutation.SetParentID(s)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(s *string) *MenuUpdate {
	if s != nil {
		mu.SetParentID(*s)
	}
	return mu
}

// SetName sets the "name" field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetPath sets the "path" field.
func (mu *MenuUpdate) SetPath(s string) *MenuUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mu *MenuUpdate) SetNillablePath(s *string) *MenuUpdate {
	if s != nil {
		mu.SetPath(*s)
	}
	return mu
}

// SetComponent sets the "component" field.
func (mu *MenuUpdate) SetComponent(s string) *MenuUpdate {
	mu.mutation.SetComponent(s)
	return mu
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableComponent(s *string) *MenuUpdate {
	if s != nil {
		mu.SetComponent(*s)
	}
	return mu
}

// SetRedirect sets the "redirect" field.
func (mu *MenuUpdate) SetRedirect(s string) *MenuUpdate {
	mu.mutation.SetRedirect(s)
	return mu
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableRedirect(s *string) *MenuUpdate {
	if s != nil {
		mu.SetRedirect(*s)
	}
	return mu
}

// SetIcon sets the "icon" field.
func (mu *MenuUpdate) SetIcon(s string) *MenuUpdate {
	mu.mutation.SetIcon(s)
	return mu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIcon(s *string) *MenuUpdate {
	if s != nil {
		mu.SetIcon(*s)
	}
	return mu
}

// SetSort sets the "sort" field.
func (mu *MenuUpdate) SetSort(i int) *MenuUpdate {
	mu.mutation.ResetSort()
	mu.mutation.SetSort(i)
	return mu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableSort(i *int) *MenuUpdate {
	if i != nil {
		mu.SetSort(*i)
	}
	return mu
}

// AddSort adds i to the "sort" field.
func (mu *MenuUpdate) AddSort(i int) *MenuUpdate {
	mu.mutation.AddSort(i)
	return mu
}

// SetType sets the "type" field.
func (mu *MenuUpdate) SetType(s string) *MenuUpdate {
	mu.mutation.SetType(s)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableType(s *string) *MenuUpdate {
	if s != nil {
		mu.SetType(*s)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MenuUpdate) SetStatus(i int) *MenuUpdate {
	mu.mutation.ResetStatus()
	mu.mutation.SetStatus(i)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableStatus(i *int) *MenuUpdate {
	if i != nil {
		mu.SetStatus(*i)
	}
	return mu
}

// AddStatus adds i to the "status" field.
func (mu *MenuUpdate) AddStatus(i int) *MenuUpdate {
	mu.mutation.AddStatus(i)
	return mu
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MenuUpdate) check() error {
	if v, ok := mu.mutation.Code(); ok {
		if err := menu.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Menu.code": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Menu.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(menu.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.DeletedAt(); ok {
		_spec.SetField(menu.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := mu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(menu.FieldDeletedAt, field.TypeInt64, value)
	}
	if mu.mutation.DeletedAtCleared() {
		_spec.ClearField(menu.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := mu.mutation.TenantCode(); ok {
		_spec.SetField(menu.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := mu.mutation.MenuID(); ok {
		_spec.SetField(menu.FieldMenuID, field.TypeString, value)
	}
	if value, ok := mu.mutation.Code(); ok {
		_spec.SetField(menu.FieldCode, field.TypeString, value)
	}
	if value, ok := mu.mutation.ParentID(); ok {
		_spec.SetField(menu.FieldParentID, field.TypeString, value)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if value, ok := mu.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
	}
	if value, ok := mu.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
	}
	if value, ok := mu.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
	}
	if value, ok := mu.mutation.Sort(); ok {
		_spec.SetField(menu.FieldSort, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedSort(); ok {
		_spec.AddField(menu.FieldSort, field.TypeInt, value)
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeString, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedStatus(); ok {
		_spec.AddField(menu.FieldStatus, field.TypeInt, value)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MenuUpdateOne) SetUpdatedAt(i int64) *MenuUpdateOne {
	muo.mutation.ResetUpdatedAt()
	muo.mutation.SetUpdatedAt(i)
	return muo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableUpdatedAt(i *int64) *MenuUpdateOne {
	if i != nil {
		muo.SetUpdatedAt(*i)
	}
	return muo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (muo *MenuUpdateOne) AddUpdatedAt(i int64) *MenuUpdateOne {
	muo.mutation.AddUpdatedAt(i)
	return muo
}

// SetDeletedAt sets the "deleted_at" field.
func (muo *MenuUpdateOne) SetDeletedAt(i int64) *MenuUpdateOne {
	muo.mutation.ResetDeletedAt()
	muo.mutation.SetDeletedAt(i)
	return muo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableDeletedAt(i *int64) *MenuUpdateOne {
	if i != nil {
		muo.SetDeletedAt(*i)
	}
	return muo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (muo *MenuUpdateOne) AddDeletedAt(i int64) *MenuUpdateOne {
	muo.mutation.AddDeletedAt(i)
	return muo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (muo *MenuUpdateOne) ClearDeletedAt() *MenuUpdateOne {
	muo.mutation.ClearDeletedAt()
	return muo
}

// SetTenantCode sets the "tenant_code" field.
func (muo *MenuUpdateOne) SetTenantCode(s string) *MenuUpdateOne {
	muo.mutation.SetTenantCode(s)
	return muo
}

// SetNillableTenantCode sets the "tenant_code" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableTenantCode(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetTenantCode(*s)
	}
	return muo
}

// SetMenuID sets the "menu_id" field.
func (muo *MenuUpdateOne) SetMenuID(s string) *MenuUpdateOne {
	muo.mutation.SetMenuID(s)
	return muo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableMenuID(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetMenuID(*s)
	}
	return muo
}

// SetCode sets the "code" field.
func (muo *MenuUpdateOne) SetCode(s string) *MenuUpdateOne {
	muo.mutation.SetCode(s)
	return muo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableCode(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetCode(*s)
	}
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(s string) *MenuUpdateOne {
	muo.mutation.SetParentID(s)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetParentID(*s)
	}
	return muo
}

// SetName sets the "name" field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetPath sets the "path" field.
func (muo *MenuUpdateOne) SetPath(s string) *MenuUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillablePath(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetPath(*s)
	}
	return muo
}

// SetComponent sets the "component" field.
func (muo *MenuUpdateOne) SetComponent(s string) *MenuUpdateOne {
	muo.mutation.SetComponent(s)
	return muo
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableComponent(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetComponent(*s)
	}
	return muo
}

// SetRedirect sets the "redirect" field.
func (muo *MenuUpdateOne) SetRedirect(s string) *MenuUpdateOne {
	muo.mutation.SetRedirect(s)
	return muo
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableRedirect(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetRedirect(*s)
	}
	return muo
}

// SetIcon sets the "icon" field.
func (muo *MenuUpdateOne) SetIcon(s string) *MenuUpdateOne {
	muo.mutation.SetIcon(s)
	return muo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIcon(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetIcon(*s)
	}
	return muo
}

// SetSort sets the "sort" field.
func (muo *MenuUpdateOne) SetSort(i int) *MenuUpdateOne {
	muo.mutation.ResetSort()
	muo.mutation.SetSort(i)
	return muo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableSort(i *int) *MenuUpdateOne {
	if i != nil {
		muo.SetSort(*i)
	}
	return muo
}

// AddSort adds i to the "sort" field.
func (muo *MenuUpdateOne) AddSort(i int) *MenuUpdateOne {
	muo.mutation.AddSort(i)
	return muo
}

// SetType sets the "type" field.
func (muo *MenuUpdateOne) SetType(s string) *MenuUpdateOne {
	muo.mutation.SetType(s)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableType(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetType(*s)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MenuUpdateOne) SetStatus(i int) *MenuUpdateOne {
	muo.mutation.ResetStatus()
	muo.mutation.SetStatus(i)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableStatus(i *int) *MenuUpdateOne {
	if i != nil {
		muo.SetStatus(*i)
	}
	return muo
}

// AddStatus adds i to the "status" field.
func (muo *MenuUpdateOne) AddStatus(i int) *MenuUpdateOne {
	muo.mutation.AddStatus(i)
	return muo
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// Where appends a list predicates to the MenuUpdate builder.
func (muo *MenuUpdateOne) Where(ps ...predicate.Menu) *MenuUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MenuUpdateOne) check() error {
	if v, ok := muo.mutation.Code(); ok {
		if err := menu.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`generated: validator failed for field "Menu.code": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Menu.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(menu.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.DeletedAt(); ok {
		_spec.SetField(menu.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := muo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(menu.FieldDeletedAt, field.TypeInt64, value)
	}
	if muo.mutation.DeletedAtCleared() {
		_spec.ClearField(menu.FieldDeletedAt, field.TypeInt64)
	}
	if value, ok := muo.mutation.TenantCode(); ok {
		_spec.SetField(menu.FieldTenantCode, field.TypeString, value)
	}
	if value, ok := muo.mutation.MenuID(); ok {
		_spec.SetField(menu.FieldMenuID, field.TypeString, value)
	}
	if value, ok := muo.mutation.Code(); ok {
		_spec.SetField(menu.FieldCode, field.TypeString, value)
	}
	if value, ok := muo.mutation.ParentID(); ok {
		_spec.SetField(menu.FieldParentID, field.TypeString, value)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if value, ok := muo.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
	}
	if value, ok := muo.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
	}
	if value, ok := muo.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
	}
	if value, ok := muo.mutation.Sort(); ok {
		_spec.SetField(menu.FieldSort, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedSort(); ok {
		_spec.AddField(menu.FieldSort, field.TypeInt, value)
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeString, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedStatus(); ok {
		_spec.AddField(menu.FieldStatus, field.TypeInt, value)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
