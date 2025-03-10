// Code generated by ent, DO NOT EDIT.

package generated

import (
	"admin_backend/pkg/ent/generated/dicttype"
	"admin_backend/pkg/ent/generated/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictTypeQuery is the builder for querying DictType entities.
type DictTypeQuery struct {
	config
	ctx        *QueryContext
	order      []dicttype.OrderOption
	inters     []Interceptor
	predicates []predicate.DictType
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DictTypeQuery builder.
func (dtq *DictTypeQuery) Where(ps ...predicate.DictType) *DictTypeQuery {
	dtq.predicates = append(dtq.predicates, ps...)
	return dtq
}

// Limit the number of records to be returned by this query.
func (dtq *DictTypeQuery) Limit(limit int) *DictTypeQuery {
	dtq.ctx.Limit = &limit
	return dtq
}

// Offset to start from.
func (dtq *DictTypeQuery) Offset(offset int) *DictTypeQuery {
	dtq.ctx.Offset = &offset
	return dtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dtq *DictTypeQuery) Unique(unique bool) *DictTypeQuery {
	dtq.ctx.Unique = &unique
	return dtq
}

// Order specifies how the records should be ordered.
func (dtq *DictTypeQuery) Order(o ...dicttype.OrderOption) *DictTypeQuery {
	dtq.order = append(dtq.order, o...)
	return dtq
}

// First returns the first DictType entity from the query.
// Returns a *NotFoundError when no DictType was found.
func (dtq *DictTypeQuery) First(ctx context.Context) (*DictType, error) {
	nodes, err := dtq.Limit(1).All(setContextOp(ctx, dtq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dicttype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dtq *DictTypeQuery) FirstX(ctx context.Context) *DictType {
	node, err := dtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DictType ID from the query.
// Returns a *NotFoundError when no DictType ID was found.
func (dtq *DictTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(1).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dicttype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dtq *DictTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := dtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DictType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DictType entity is found.
// Returns a *NotFoundError when no DictType entities are found.
func (dtq *DictTypeQuery) Only(ctx context.Context) (*DictType, error) {
	nodes, err := dtq.Limit(2).All(setContextOp(ctx, dtq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dicttype.Label}
	default:
		return nil, &NotSingularError{dicttype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dtq *DictTypeQuery) OnlyX(ctx context.Context) *DictType {
	node, err := dtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DictType ID in the query.
// Returns a *NotSingularError when more than one DictType ID is found.
// Returns a *NotFoundError when no entities are found.
func (dtq *DictTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dtq.Limit(2).IDs(setContextOp(ctx, dtq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dicttype.Label}
	default:
		err = &NotSingularError{dicttype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dtq *DictTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := dtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DictTypes.
func (dtq *DictTypeQuery) All(ctx context.Context) ([]*DictType, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryAll)
	if err := dtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DictType, *DictTypeQuery]()
	return withInterceptors[[]*DictType](ctx, dtq, qr, dtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dtq *DictTypeQuery) AllX(ctx context.Context) []*DictType {
	nodes, err := dtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DictType IDs.
func (dtq *DictTypeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dtq.ctx.Unique == nil && dtq.path != nil {
		dtq.Unique(true)
	}
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryIDs)
	if err = dtq.Select(dicttype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dtq *DictTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := dtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dtq *DictTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryCount)
	if err := dtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dtq, querierCount[*DictTypeQuery](), dtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dtq *DictTypeQuery) CountX(ctx context.Context) int {
	count, err := dtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dtq *DictTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dtq.ctx, ent.OpQueryExist)
	switch _, err := dtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dtq *DictTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := dtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DictTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dtq *DictTypeQuery) Clone() *DictTypeQuery {
	if dtq == nil {
		return nil
	}
	return &DictTypeQuery{
		config:     dtq.config,
		ctx:        dtq.ctx.Clone(),
		order:      append([]dicttype.OrderOption{}, dtq.order...),
		inters:     append([]Interceptor{}, dtq.inters...),
		predicates: append([]predicate.DictType{}, dtq.predicates...),
		// clone intermediate query.
		sql:       dtq.sql.Clone(),
		path:      dtq.path,
		modifiers: append([]func(*sql.Selector){}, dtq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DictType.Query().
//		GroupBy(dicttype.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (dtq *DictTypeQuery) GroupBy(field string, fields ...string) *DictTypeGroupBy {
	dtq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DictTypeGroupBy{build: dtq}
	grbuild.flds = &dtq.ctx.Fields
	grbuild.label = dicttype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//	}
//
//	client.DictType.Query().
//		Select(dicttype.FieldCreatedAt).
//		Scan(ctx, &v)
func (dtq *DictTypeQuery) Select(fields ...string) *DictTypeSelect {
	dtq.ctx.Fields = append(dtq.ctx.Fields, fields...)
	sbuild := &DictTypeSelect{DictTypeQuery: dtq}
	sbuild.label = dicttype.Label
	sbuild.flds, sbuild.scan = &dtq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DictTypeSelect configured with the given aggregations.
func (dtq *DictTypeQuery) Aggregate(fns ...AggregateFunc) *DictTypeSelect {
	return dtq.Select().Aggregate(fns...)
}

func (dtq *DictTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dtq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dtq); err != nil {
				return err
			}
		}
	}
	for _, f := range dtq.ctx.Fields {
		if !dicttype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if dtq.path != nil {
		prev, err := dtq.path(ctx)
		if err != nil {
			return err
		}
		dtq.sql = prev
	}
	return nil
}

func (dtq *DictTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DictType, error) {
	var (
		nodes = []*DictType{}
		_spec = dtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DictType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DictType{config: dtq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(dtq.modifiers) > 0 {
		_spec.Modifiers = dtq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dtq *DictTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dtq.querySpec()
	if len(dtq.modifiers) > 0 {
		_spec.Modifiers = dtq.modifiers
	}
	_spec.Node.Columns = dtq.ctx.Fields
	if len(dtq.ctx.Fields) > 0 {
		_spec.Unique = dtq.ctx.Unique != nil && *dtq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dtq.driver, _spec)
}

func (dtq *DictTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dicttype.Table, dicttype.Columns, sqlgraph.NewFieldSpec(dicttype.FieldID, field.TypeInt))
	_spec.From = dtq.sql
	if unique := dtq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dtq.path != nil {
		_spec.Unique = true
	}
	if fields := dtq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dicttype.FieldID)
		for i := range fields {
			if fields[i] != dicttype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dtq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dtq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dtq *DictTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dtq.driver.Dialect())
	t1 := builder.Table(dicttype.Table)
	columns := dtq.ctx.Fields
	if len(columns) == 0 {
		columns = dicttype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dtq.sql != nil {
		selector = dtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dtq.ctx.Unique != nil && *dtq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dtq.modifiers {
		m(selector)
	}
	for _, p := range dtq.predicates {
		p(selector)
	}
	for _, p := range dtq.order {
		p(selector)
	}
	if offset := dtq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dtq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dtq *DictTypeQuery) Modify(modifiers ...func(s *sql.Selector)) *DictTypeSelect {
	dtq.modifiers = append(dtq.modifiers, modifiers...)
	return dtq.Select()
}

// DictTypeGroupBy is the group-by builder for DictType entities.
type DictTypeGroupBy struct {
	selector
	build *DictTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dtgb *DictTypeGroupBy) Aggregate(fns ...AggregateFunc) *DictTypeGroupBy {
	dtgb.fns = append(dtgb.fns, fns...)
	return dtgb
}

// Scan applies the selector query and scans the result into the given value.
func (dtgb *DictTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dtgb.build.ctx, ent.OpQueryGroupBy)
	if err := dtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictTypeQuery, *DictTypeGroupBy](ctx, dtgb.build, dtgb, dtgb.build.inters, v)
}

func (dtgb *DictTypeGroupBy) sqlScan(ctx context.Context, root *DictTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dtgb.fns))
	for _, fn := range dtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dtgb.flds)+len(dtgb.fns))
		for _, f := range *dtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DictTypeSelect is the builder for selecting fields of DictType entities.
type DictTypeSelect struct {
	*DictTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dts *DictTypeSelect) Aggregate(fns ...AggregateFunc) *DictTypeSelect {
	dts.fns = append(dts.fns, fns...)
	return dts
}

// Scan applies the selector query and scans the result into the given value.
func (dts *DictTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dts.ctx, ent.OpQuerySelect)
	if err := dts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DictTypeQuery, *DictTypeSelect](ctx, dts.DictTypeQuery, dts, dts.inters, v)
}

func (dts *DictTypeSelect) sqlScan(ctx context.Context, root *DictTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dts.fns))
	for _, fn := range dts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dts *DictTypeSelect) Modify(modifiers ...func(s *sql.Selector)) *DictTypeSelect {
	dts.modifiers = append(dts.modifiers, modifiers...)
	return dts
}
