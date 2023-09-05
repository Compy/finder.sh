// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/compy/finder.sh/ent/datasource"
	"github.com/compy/finder.sh/ent/predicate"
)

// DataSourceQuery is the builder for querying DataSource entities.
type DataSourceQuery struct {
	config
	ctx        *QueryContext
	order      []datasource.OrderOption
	inters     []Interceptor
	predicates []predicate.DataSource
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DataSourceQuery builder.
func (dsq *DataSourceQuery) Where(ps ...predicate.DataSource) *DataSourceQuery {
	dsq.predicates = append(dsq.predicates, ps...)
	return dsq
}

// Limit the number of records to be returned by this query.
func (dsq *DataSourceQuery) Limit(limit int) *DataSourceQuery {
	dsq.ctx.Limit = &limit
	return dsq
}

// Offset to start from.
func (dsq *DataSourceQuery) Offset(offset int) *DataSourceQuery {
	dsq.ctx.Offset = &offset
	return dsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dsq *DataSourceQuery) Unique(unique bool) *DataSourceQuery {
	dsq.ctx.Unique = &unique
	return dsq
}

// Order specifies how the records should be ordered.
func (dsq *DataSourceQuery) Order(o ...datasource.OrderOption) *DataSourceQuery {
	dsq.order = append(dsq.order, o...)
	return dsq
}

// First returns the first DataSource entity from the query.
// Returns a *NotFoundError when no DataSource was found.
func (dsq *DataSourceQuery) First(ctx context.Context) (*DataSource, error) {
	nodes, err := dsq.Limit(1).All(setContextOp(ctx, dsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{datasource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dsq *DataSourceQuery) FirstX(ctx context.Context) *DataSource {
	node, err := dsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DataSource ID from the query.
// Returns a *NotFoundError when no DataSource ID was found.
func (dsq *DataSourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(1).IDs(setContextOp(ctx, dsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{datasource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dsq *DataSourceQuery) FirstIDX(ctx context.Context) int {
	id, err := dsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DataSource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DataSource entity is found.
// Returns a *NotFoundError when no DataSource entities are found.
func (dsq *DataSourceQuery) Only(ctx context.Context) (*DataSource, error) {
	nodes, err := dsq.Limit(2).All(setContextOp(ctx, dsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{datasource.Label}
	default:
		return nil, &NotSingularError{datasource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dsq *DataSourceQuery) OnlyX(ctx context.Context) *DataSource {
	node, err := dsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DataSource ID in the query.
// Returns a *NotSingularError when more than one DataSource ID is found.
// Returns a *NotFoundError when no entities are found.
func (dsq *DataSourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dsq.Limit(2).IDs(setContextOp(ctx, dsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = &NotSingularError{datasource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dsq *DataSourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := dsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DataSources.
func (dsq *DataSourceQuery) All(ctx context.Context) ([]*DataSource, error) {
	ctx = setContextOp(ctx, dsq.ctx, "All")
	if err := dsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DataSource, *DataSourceQuery]()
	return withInterceptors[[]*DataSource](ctx, dsq, qr, dsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dsq *DataSourceQuery) AllX(ctx context.Context) []*DataSource {
	nodes, err := dsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DataSource IDs.
func (dsq *DataSourceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dsq.ctx.Unique == nil && dsq.path != nil {
		dsq.Unique(true)
	}
	ctx = setContextOp(ctx, dsq.ctx, "IDs")
	if err = dsq.Select(datasource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dsq *DataSourceQuery) IDsX(ctx context.Context) []int {
	ids, err := dsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dsq *DataSourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dsq.ctx, "Count")
	if err := dsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dsq, querierCount[*DataSourceQuery](), dsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dsq *DataSourceQuery) CountX(ctx context.Context) int {
	count, err := dsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dsq *DataSourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dsq.ctx, "Exist")
	switch _, err := dsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dsq *DataSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := dsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DataSourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dsq *DataSourceQuery) Clone() *DataSourceQuery {
	if dsq == nil {
		return nil
	}
	return &DataSourceQuery{
		config:     dsq.config,
		ctx:        dsq.ctx.Clone(),
		order:      append([]datasource.OrderOption{}, dsq.order...),
		inters:     append([]Interceptor{}, dsq.inters...),
		predicates: append([]predicate.DataSource{}, dsq.predicates...),
		// clone intermediate query.
		sql:  dsq.sql.Clone(),
		path: dsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DataSource.Query().
//		GroupBy(datasource.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dsq *DataSourceQuery) GroupBy(field string, fields ...string) *DataSourceGroupBy {
	dsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DataSourceGroupBy{build: dsq}
	grbuild.flds = &dsq.ctx.Fields
	grbuild.label = datasource.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DataSource.Query().
//		Select(datasource.FieldName).
//		Scan(ctx, &v)
func (dsq *DataSourceQuery) Select(fields ...string) *DataSourceSelect {
	dsq.ctx.Fields = append(dsq.ctx.Fields, fields...)
	sbuild := &DataSourceSelect{DataSourceQuery: dsq}
	sbuild.label = datasource.Label
	sbuild.flds, sbuild.scan = &dsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DataSourceSelect configured with the given aggregations.
func (dsq *DataSourceQuery) Aggregate(fns ...AggregateFunc) *DataSourceSelect {
	return dsq.Select().Aggregate(fns...)
}

func (dsq *DataSourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dsq); err != nil {
				return err
			}
		}
	}
	for _, f := range dsq.ctx.Fields {
		if !datasource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dsq.path != nil {
		prev, err := dsq.path(ctx)
		if err != nil {
			return err
		}
		dsq.sql = prev
	}
	return nil
}

func (dsq *DataSourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DataSource, error) {
	var (
		nodes = []*DataSource{}
		_spec = dsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DataSource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DataSource{config: dsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dsq *DataSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dsq.querySpec()
	_spec.Node.Columns = dsq.ctx.Fields
	if len(dsq.ctx.Fields) > 0 {
		_spec.Unique = dsq.ctx.Unique != nil && *dsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dsq.driver, _spec)
}

func (dsq *DataSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(datasource.Table, datasource.Columns, sqlgraph.NewFieldSpec(datasource.FieldID, field.TypeInt))
	_spec.From = dsq.sql
	if unique := dsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dsq.path != nil {
		_spec.Unique = true
	}
	if fields := dsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, datasource.FieldID)
		for i := range fields {
			if fields[i] != datasource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dsq *DataSourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dsq.driver.Dialect())
	t1 := builder.Table(datasource.Table)
	columns := dsq.ctx.Fields
	if len(columns) == 0 {
		columns = datasource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dsq.sql != nil {
		selector = dsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dsq.ctx.Unique != nil && *dsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dsq.predicates {
		p(selector)
	}
	for _, p := range dsq.order {
		p(selector)
	}
	if offset := dsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DataSourceGroupBy is the group-by builder for DataSource entities.
type DataSourceGroupBy struct {
	selector
	build *DataSourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dsgb *DataSourceGroupBy) Aggregate(fns ...AggregateFunc) *DataSourceGroupBy {
	dsgb.fns = append(dsgb.fns, fns...)
	return dsgb
}

// Scan applies the selector query and scans the result into the given value.
func (dsgb *DataSourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dsgb.build.ctx, "GroupBy")
	if err := dsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DataSourceQuery, *DataSourceGroupBy](ctx, dsgb.build, dsgb, dsgb.build.inters, v)
}

func (dsgb *DataSourceGroupBy) sqlScan(ctx context.Context, root *DataSourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dsgb.fns))
	for _, fn := range dsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dsgb.flds)+len(dsgb.fns))
		for _, f := range *dsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DataSourceSelect is the builder for selecting fields of DataSource entities.
type DataSourceSelect struct {
	*DataSourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dss *DataSourceSelect) Aggregate(fns ...AggregateFunc) *DataSourceSelect {
	dss.fns = append(dss.fns, fns...)
	return dss
}

// Scan applies the selector query and scans the result into the given value.
func (dss *DataSourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dss.ctx, "Select")
	if err := dss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DataSourceQuery, *DataSourceSelect](ctx, dss.DataSourceQuery, dss, dss.inters, v)
}

func (dss *DataSourceSelect) sqlScan(ctx context.Context, root *DataSourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dss.fns))
	for _, fn := range dss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
