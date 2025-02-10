// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"admin_backend/pkg/ent/generated/predicate"
	"admin_backend/pkg/ent/generated/systemlog"
)

// SystemLogDelete is the builder for deleting a SystemLog entity.
type SystemLogDelete struct {
	config
	hooks    []Hook
	mutation *SystemLogMutation
}

// Where appends a list predicates to the SystemLogDelete builder.
func (sld *SystemLogDelete) Where(ps ...predicate.SystemLog) *SystemLogDelete {
	sld.mutation.Where(ps...)
	return sld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sld *SystemLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sld.sqlExec, sld.mutation, sld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sld *SystemLogDelete) ExecX(ctx context.Context) int {
	n, err := sld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sld *SystemLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(systemlog.Table, sqlgraph.NewFieldSpec(systemlog.FieldID, field.TypeInt))
	if ps := sld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sld.mutation.done = true
	return affected, err
}

// SystemLogDeleteOne is the builder for deleting a single SystemLog entity.
type SystemLogDeleteOne struct {
	sld *SystemLogDelete
}

// Where appends a list predicates to the SystemLogDelete builder.
func (sldo *SystemLogDeleteOne) Where(ps ...predicate.SystemLog) *SystemLogDeleteOne {
	sldo.sld.mutation.Where(ps...)
	return sldo
}

// Exec executes the deletion query.
func (sldo *SystemLogDeleteOne) Exec(ctx context.Context) error {
	n, err := sldo.sld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{systemlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sldo *SystemLogDeleteOne) ExecX(ctx context.Context) {
	if err := sldo.Exec(ctx); err != nil {
		panic(err)
	}
}
