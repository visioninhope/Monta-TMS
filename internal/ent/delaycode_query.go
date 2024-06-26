// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/delaycode"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// DelayCodeQuery is the builder for querying DelayCode entities.
type DelayCodeQuery struct {
	config
	ctx              *QueryContext
	order            []delaycode.OrderOption
	inters           []Interceptor
	predicates       []predicate.DelayCode
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DelayCodeQuery builder.
func (dcq *DelayCodeQuery) Where(ps ...predicate.DelayCode) *DelayCodeQuery {
	dcq.predicates = append(dcq.predicates, ps...)
	return dcq
}

// Limit the number of records to be returned by this query.
func (dcq *DelayCodeQuery) Limit(limit int) *DelayCodeQuery {
	dcq.ctx.Limit = &limit
	return dcq
}

// Offset to start from.
func (dcq *DelayCodeQuery) Offset(offset int) *DelayCodeQuery {
	dcq.ctx.Offset = &offset
	return dcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcq *DelayCodeQuery) Unique(unique bool) *DelayCodeQuery {
	dcq.ctx.Unique = &unique
	return dcq
}

// Order specifies how the records should be ordered.
func (dcq *DelayCodeQuery) Order(o ...delaycode.OrderOption) *DelayCodeQuery {
	dcq.order = append(dcq.order, o...)
	return dcq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (dcq *DelayCodeQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(delaycode.Table, delaycode.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, delaycode.BusinessUnitTable, delaycode.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (dcq *DelayCodeQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(delaycode.Table, delaycode.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, delaycode.OrganizationTable, delaycode.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DelayCode entity from the query.
// Returns a *NotFoundError when no DelayCode was found.
func (dcq *DelayCodeQuery) First(ctx context.Context) (*DelayCode, error) {
	nodes, err := dcq.Limit(1).All(setContextOp(ctx, dcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{delaycode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcq *DelayCodeQuery) FirstX(ctx context.Context) *DelayCode {
	node, err := dcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DelayCode ID from the query.
// Returns a *NotFoundError when no DelayCode ID was found.
func (dcq *DelayCodeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dcq.Limit(1).IDs(setContextOp(ctx, dcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{delaycode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcq *DelayCodeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DelayCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DelayCode entity is found.
// Returns a *NotFoundError when no DelayCode entities are found.
func (dcq *DelayCodeQuery) Only(ctx context.Context) (*DelayCode, error) {
	nodes, err := dcq.Limit(2).All(setContextOp(ctx, dcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{delaycode.Label}
	default:
		return nil, &NotSingularError{delaycode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcq *DelayCodeQuery) OnlyX(ctx context.Context) *DelayCode {
	node, err := dcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DelayCode ID in the query.
// Returns a *NotSingularError when more than one DelayCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcq *DelayCodeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dcq.Limit(2).IDs(setContextOp(ctx, dcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{delaycode.Label}
	default:
		err = &NotSingularError{delaycode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcq *DelayCodeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DelayCodes.
func (dcq *DelayCodeQuery) All(ctx context.Context) ([]*DelayCode, error) {
	ctx = setContextOp(ctx, dcq.ctx, "All")
	if err := dcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DelayCode, *DelayCodeQuery]()
	return withInterceptors[[]*DelayCode](ctx, dcq, qr, dcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcq *DelayCodeQuery) AllX(ctx context.Context) []*DelayCode {
	nodes, err := dcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DelayCode IDs.
func (dcq *DelayCodeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dcq.ctx.Unique == nil && dcq.path != nil {
		dcq.Unique(true)
	}
	ctx = setContextOp(ctx, dcq.ctx, "IDs")
	if err = dcq.Select(delaycode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcq *DelayCodeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcq *DelayCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Count")
	if err := dcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcq, querierCount[*DelayCodeQuery](), dcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcq *DelayCodeQuery) CountX(ctx context.Context) int {
	count, err := dcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcq *DelayCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Exist")
	switch _, err := dcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcq *DelayCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := dcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DelayCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcq *DelayCodeQuery) Clone() *DelayCodeQuery {
	if dcq == nil {
		return nil
	}
	return &DelayCodeQuery{
		config:           dcq.config,
		ctx:              dcq.ctx.Clone(),
		order:            append([]delaycode.OrderOption{}, dcq.order...),
		inters:           append([]Interceptor{}, dcq.inters...),
		predicates:       append([]predicate.DelayCode{}, dcq.predicates...),
		withBusinessUnit: dcq.withBusinessUnit.Clone(),
		withOrganization: dcq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  dcq.sql.Clone(),
		path: dcq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DelayCodeQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *DelayCodeQuery {
	query := (&BusinessUnitClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withBusinessUnit = query
	return dcq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DelayCodeQuery) WithOrganization(opts ...func(*OrganizationQuery)) *DelayCodeQuery {
	query := (&OrganizationClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withOrganization = query
	return dcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DelayCode.Query().
//		GroupBy(delaycode.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcq *DelayCodeQuery) GroupBy(field string, fields ...string) *DelayCodeGroupBy {
	dcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DelayCodeGroupBy{build: dcq}
	grbuild.flds = &dcq.ctx.Fields
	grbuild.label = delaycode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BusinessUnitID uuid.UUID `json:"businessUnitId"`
//	}
//
//	client.DelayCode.Query().
//		Select(delaycode.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (dcq *DelayCodeQuery) Select(fields ...string) *DelayCodeSelect {
	dcq.ctx.Fields = append(dcq.ctx.Fields, fields...)
	sbuild := &DelayCodeSelect{DelayCodeQuery: dcq}
	sbuild.label = delaycode.Label
	sbuild.flds, sbuild.scan = &dcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DelayCodeSelect configured with the given aggregations.
func (dcq *DelayCodeQuery) Aggregate(fns ...AggregateFunc) *DelayCodeSelect {
	return dcq.Select().Aggregate(fns...)
}

func (dcq *DelayCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcq.ctx.Fields {
		if !delaycode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcq.path != nil {
		prev, err := dcq.path(ctx)
		if err != nil {
			return err
		}
		dcq.sql = prev
	}
	return nil
}

func (dcq *DelayCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DelayCode, error) {
	var (
		nodes       = []*DelayCode{}
		_spec       = dcq.querySpec()
		loadedTypes = [2]bool{
			dcq.withBusinessUnit != nil,
			dcq.withOrganization != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DelayCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DelayCode{config: dcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcq.withBusinessUnit; query != nil {
		if err := dcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *DelayCode, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := dcq.withOrganization; query != nil {
		if err := dcq.loadOrganization(ctx, query, nodes, nil,
			func(n *DelayCode, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcq *DelayCodeQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*DelayCode, init func(*DelayCode), assign func(*DelayCode, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DelayCode)
	for i := range nodes {
		fk := nodes[i].BusinessUnitID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(businessunit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "business_unit_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dcq *DelayCodeQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*DelayCode, init func(*DelayCode), assign func(*DelayCode, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DelayCode)
	for i := range nodes {
		fk := nodes[i].OrganizationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "organization_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dcq *DelayCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcq.querySpec()
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	_spec.Node.Columns = dcq.ctx.Fields
	if len(dcq.ctx.Fields) > 0 {
		_spec.Unique = dcq.ctx.Unique != nil && *dcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcq.driver, _spec)
}

func (dcq *DelayCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(delaycode.Table, delaycode.Columns, sqlgraph.NewFieldSpec(delaycode.FieldID, field.TypeUUID))
	_spec.From = dcq.sql
	if unique := dcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcq.path != nil {
		_spec.Unique = true
	}
	if fields := dcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, delaycode.FieldID)
		for i := range fields {
			if fields[i] != delaycode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(delaycode.FieldBusinessUnitID)
		}
		if dcq.withOrganization != nil {
			_spec.Node.AddColumnOnce(delaycode.FieldOrganizationID)
		}
	}
	if ps := dcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcq *DelayCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcq.driver.Dialect())
	t1 := builder.Table(delaycode.Table)
	columns := dcq.ctx.Fields
	if len(columns) == 0 {
		columns = delaycode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcq.sql != nil {
		selector = dcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcq.ctx.Unique != nil && *dcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dcq.modifiers {
		m(selector)
	}
	for _, p := range dcq.predicates {
		p(selector)
	}
	for _, p := range dcq.order {
		p(selector)
	}
	if offset := dcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcq *DelayCodeQuery) Modify(modifiers ...func(s *sql.Selector)) *DelayCodeSelect {
	dcq.modifiers = append(dcq.modifiers, modifiers...)
	return dcq.Select()
}

// DelayCodeGroupBy is the group-by builder for DelayCode entities.
type DelayCodeGroupBy struct {
	selector
	build *DelayCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcgb *DelayCodeGroupBy) Aggregate(fns ...AggregateFunc) *DelayCodeGroupBy {
	dcgb.fns = append(dcgb.fns, fns...)
	return dcgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcgb *DelayCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcgb.build.ctx, "GroupBy")
	if err := dcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DelayCodeQuery, *DelayCodeGroupBy](ctx, dcgb.build, dcgb, dcgb.build.inters, v)
}

func (dcgb *DelayCodeGroupBy) sqlScan(ctx context.Context, root *DelayCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcgb.fns))
	for _, fn := range dcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcgb.flds)+len(dcgb.fns))
		for _, f := range *dcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DelayCodeSelect is the builder for selecting fields of DelayCode entities.
type DelayCodeSelect struct {
	*DelayCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcs *DelayCodeSelect) Aggregate(fns ...AggregateFunc) *DelayCodeSelect {
	dcs.fns = append(dcs.fns, fns...)
	return dcs
}

// Scan applies the selector query and scans the result into the given value.
func (dcs *DelayCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcs.ctx, "Select")
	if err := dcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DelayCodeQuery, *DelayCodeSelect](ctx, dcs.DelayCodeQuery, dcs, dcs.inters, v)
}

func (dcs *DelayCodeSelect) sqlScan(ctx context.Context, root *DelayCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcs.fns))
	for _, fn := range dcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcs *DelayCodeSelect) Modify(modifiers ...func(s *sql.Selector)) *DelayCodeSelect {
	dcs.modifiers = append(dcs.modifiers, modifiers...)
	return dcs
}
