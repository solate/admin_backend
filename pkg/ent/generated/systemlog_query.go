// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/solate/admin_backend/pkg/ent/generated/predicate"
	"github.com/solate/admin_backend/pkg/ent/generated/systemlog"
)

// SystemLogQuery is the builder for querying SystemLog entities.
type SystemLogQuery struct {
	config
	ctx        *QueryContext
	order      []systemlog.OrderOption
	inters     []Interceptor
	predicates []predicate.SystemLog
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SystemLogQuery builder.
func (slq *SystemLogQuery) Where(ps ...predicate.SystemLog) *SystemLogQuery {
	slq.predicates = append(slq.predicates, ps...)
	return slq
}

// Limit the number of records to be returned by this query.
func (slq *SystemLogQuery) Limit(limit int) *SystemLogQuery {
	slq.ctx.Limit = &limit
	return slq
}

// Offset to start from.
func (slq *SystemLogQuery) Offset(offset int) *SystemLogQuery {
	slq.ctx.Offset = &offset
	return slq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (slq *SystemLogQuery) Unique(unique bool) *SystemLogQuery {
	slq.ctx.Unique = &unique
	return slq
}

// Order specifies how the records should be ordered.
func (slq *SystemLogQuery) Order(o ...systemlog.OrderOption) *SystemLogQuery {
	slq.order = append(slq.order, o...)
	return slq
}

// First returns the first SystemLog entity from the query.
// Returns a *NotFoundError when no SystemLog was found.
func (slq *SystemLogQuery) First(ctx context.Context) (*SystemLog, error) {
	nodes, err := slq.Limit(1).All(setContextOp(ctx, slq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{systemlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (slq *SystemLogQuery) FirstX(ctx context.Context) *SystemLog {
	node, err := slq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SystemLog ID from the query.
// Returns a *NotFoundError when no SystemLog ID was found.
func (slq *SystemLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = slq.Limit(1).IDs(setContextOp(ctx, slq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{systemlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (slq *SystemLogQuery) FirstIDX(ctx context.Context) int {
	id, err := slq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SystemLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SystemLog entity is found.
// Returns a *NotFoundError when no SystemLog entities are found.
func (slq *SystemLogQuery) Only(ctx context.Context) (*SystemLog, error) {
	nodes, err := slq.Limit(2).All(setContextOp(ctx, slq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{systemlog.Label}
	default:
		return nil, &NotSingularError{systemlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (slq *SystemLogQuery) OnlyX(ctx context.Context) *SystemLog {
	node, err := slq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SystemLog ID in the query.
// Returns a *NotSingularError when more than one SystemLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (slq *SystemLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = slq.Limit(2).IDs(setContextOp(ctx, slq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{systemlog.Label}
	default:
		err = &NotSingularError{systemlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (slq *SystemLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := slq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SystemLogs.
func (slq *SystemLogQuery) All(ctx context.Context) ([]*SystemLog, error) {
	ctx = setContextOp(ctx, slq.ctx, ent.OpQueryAll)
	if err := slq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SystemLog, *SystemLogQuery]()
	return withInterceptors[[]*SystemLog](ctx, slq, qr, slq.inters)
}

// AllX is like All, but panics if an error occurs.
func (slq *SystemLogQuery) AllX(ctx context.Context) []*SystemLog {
	nodes, err := slq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SystemLog IDs.
func (slq *SystemLogQuery) IDs(ctx context.Context) (ids []int, err error) {
	if slq.ctx.Unique == nil && slq.path != nil {
		slq.Unique(true)
	}
	ctx = setContextOp(ctx, slq.ctx, ent.OpQueryIDs)
	if err = slq.Select(systemlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (slq *SystemLogQuery) IDsX(ctx context.Context) []int {
	ids, err := slq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (slq *SystemLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, slq.ctx, ent.OpQueryCount)
	if err := slq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, slq, querierCount[*SystemLogQuery](), slq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (slq *SystemLogQuery) CountX(ctx context.Context) int {
	count, err := slq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (slq *SystemLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, slq.ctx, ent.OpQueryExist)
	switch _, err := slq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (slq *SystemLogQuery) ExistX(ctx context.Context) bool {
	exist, err := slq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SystemLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (slq *SystemLogQuery) Clone() *SystemLogQuery {
	if slq == nil {
		return nil
	}
	return &SystemLogQuery{
		config:     slq.config,
		ctx:        slq.ctx.Clone(),
		order:      append([]systemlog.OrderOption{}, slq.order...),
		inters:     append([]Interceptor{}, slq.inters...),
		predicates: append([]predicate.SystemLog{}, slq.predicates...),
		// clone intermediate query.
		sql:       slq.sql.Clone(),
		path:      slq.path,
		modifiers: append([]func(*sql.Selector){}, slq.modifiers...),
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
//	client.SystemLog.Query().
//		GroupBy(systemlog.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (slq *SystemLogQuery) GroupBy(field string, fields ...string) *SystemLogGroupBy {
	slq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SystemLogGroupBy{build: slq}
	grbuild.flds = &slq.ctx.Fields
	grbuild.label = systemlog.Label
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
//	client.SystemLog.Query().
//		Select(systemlog.FieldCreatedAt).
//		Scan(ctx, &v)
func (slq *SystemLogQuery) Select(fields ...string) *SystemLogSelect {
	slq.ctx.Fields = append(slq.ctx.Fields, fields...)
	sbuild := &SystemLogSelect{SystemLogQuery: slq}
	sbuild.label = systemlog.Label
	sbuild.flds, sbuild.scan = &slq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SystemLogSelect configured with the given aggregations.
func (slq *SystemLogQuery) Aggregate(fns ...AggregateFunc) *SystemLogSelect {
	return slq.Select().Aggregate(fns...)
}

func (slq *SystemLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range slq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, slq); err != nil {
				return err
			}
		}
	}
	for _, f := range slq.ctx.Fields {
		if !systemlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if slq.path != nil {
		prev, err := slq.path(ctx)
		if err != nil {
			return err
		}
		slq.sql = prev
	}
	return nil
}

func (slq *SystemLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SystemLog, error) {
	var (
		nodes = []*SystemLog{}
		_spec = slq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SystemLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SystemLog{config: slq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(slq.modifiers) > 0 {
		_spec.Modifiers = slq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, slq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (slq *SystemLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := slq.querySpec()
	if len(slq.modifiers) > 0 {
		_spec.Modifiers = slq.modifiers
	}
	_spec.Node.Columns = slq.ctx.Fields
	if len(slq.ctx.Fields) > 0 {
		_spec.Unique = slq.ctx.Unique != nil && *slq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, slq.driver, _spec)
}

func (slq *SystemLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(systemlog.Table, systemlog.Columns, sqlgraph.NewFieldSpec(systemlog.FieldID, field.TypeInt))
	_spec.From = slq.sql
	if unique := slq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if slq.path != nil {
		_spec.Unique = true
	}
	if fields := slq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, systemlog.FieldID)
		for i := range fields {
			if fields[i] != systemlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := slq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := slq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := slq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := slq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (slq *SystemLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(slq.driver.Dialect())
	t1 := builder.Table(systemlog.Table)
	columns := slq.ctx.Fields
	if len(columns) == 0 {
		columns = systemlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if slq.sql != nil {
		selector = slq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if slq.ctx.Unique != nil && *slq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range slq.modifiers {
		m(selector)
	}
	for _, p := range slq.predicates {
		p(selector)
	}
	for _, p := range slq.order {
		p(selector)
	}
	if offset := slq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := slq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (slq *SystemLogQuery) Modify(modifiers ...func(s *sql.Selector)) *SystemLogSelect {
	slq.modifiers = append(slq.modifiers, modifiers...)
	return slq.Select()
}

// SystemLogGroupBy is the group-by builder for SystemLog entities.
type SystemLogGroupBy struct {
	selector
	build *SystemLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (slgb *SystemLogGroupBy) Aggregate(fns ...AggregateFunc) *SystemLogGroupBy {
	slgb.fns = append(slgb.fns, fns...)
	return slgb
}

// Scan applies the selector query and scans the result into the given value.
func (slgb *SystemLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, slgb.build.ctx, ent.OpQueryGroupBy)
	if err := slgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SystemLogQuery, *SystemLogGroupBy](ctx, slgb.build, slgb, slgb.build.inters, v)
}

func (slgb *SystemLogGroupBy) sqlScan(ctx context.Context, root *SystemLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(slgb.fns))
	for _, fn := range slgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*slgb.flds)+len(slgb.fns))
		for _, f := range *slgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*slgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := slgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SystemLogSelect is the builder for selecting fields of SystemLog entities.
type SystemLogSelect struct {
	*SystemLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sls *SystemLogSelect) Aggregate(fns ...AggregateFunc) *SystemLogSelect {
	sls.fns = append(sls.fns, fns...)
	return sls
}

// Scan applies the selector query and scans the result into the given value.
func (sls *SystemLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sls.ctx, ent.OpQuerySelect)
	if err := sls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SystemLogQuery, *SystemLogSelect](ctx, sls.SystemLogQuery, sls, sls.inters, v)
}

func (sls *SystemLogSelect) sqlScan(ctx context.Context, root *SystemLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sls.fns))
	for _, fn := range sls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (sls *SystemLogSelect) Modify(modifiers ...func(s *sql.Selector)) *SystemLogSelect {
	sls.modifiers = append(sls.modifiers, modifiers...)
	return sls
}
