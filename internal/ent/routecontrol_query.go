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
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/routecontrol"
	"github.com/google/uuid"
)

// RouteControlQuery is the builder for querying RouteControl entities.
type RouteControlQuery struct {
	config
	ctx              *QueryContext
	order            []routecontrol.OrderOption
	inters           []Interceptor
	predicates       []predicate.RouteControl
	withOrganization *OrganizationQuery
	withBusinessUnit *BusinessUnitQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RouteControlQuery builder.
func (rcq *RouteControlQuery) Where(ps ...predicate.RouteControl) *RouteControlQuery {
	rcq.predicates = append(rcq.predicates, ps...)
	return rcq
}

// Limit the number of records to be returned by this query.
func (rcq *RouteControlQuery) Limit(limit int) *RouteControlQuery {
	rcq.ctx.Limit = &limit
	return rcq
}

// Offset to start from.
func (rcq *RouteControlQuery) Offset(offset int) *RouteControlQuery {
	rcq.ctx.Offset = &offset
	return rcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rcq *RouteControlQuery) Unique(unique bool) *RouteControlQuery {
	rcq.ctx.Unique = &unique
	return rcq
}

// Order specifies how the records should be ordered.
func (rcq *RouteControlQuery) Order(o ...routecontrol.OrderOption) *RouteControlQuery {
	rcq.order = append(rcq.order, o...)
	return rcq
}

// QueryOrganization chains the current query on the "organization" edge.
func (rcq *RouteControlQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: rcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routecontrol.Table, routecontrol.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, routecontrol.OrganizationTable, routecontrol.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(rcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (rcq *RouteControlQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: rcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(routecontrol.Table, routecontrol.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, routecontrol.BusinessUnitTable, routecontrol.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(rcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RouteControl entity from the query.
// Returns a *NotFoundError when no RouteControl was found.
func (rcq *RouteControlQuery) First(ctx context.Context) (*RouteControl, error) {
	nodes, err := rcq.Limit(1).All(setContextOp(ctx, rcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{routecontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rcq *RouteControlQuery) FirstX(ctx context.Context) *RouteControl {
	node, err := rcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RouteControl ID from the query.
// Returns a *NotFoundError when no RouteControl ID was found.
func (rcq *RouteControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rcq.Limit(1).IDs(setContextOp(ctx, rcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{routecontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rcq *RouteControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RouteControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RouteControl entity is found.
// Returns a *NotFoundError when no RouteControl entities are found.
func (rcq *RouteControlQuery) Only(ctx context.Context) (*RouteControl, error) {
	nodes, err := rcq.Limit(2).All(setContextOp(ctx, rcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{routecontrol.Label}
	default:
		return nil, &NotSingularError{routecontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rcq *RouteControlQuery) OnlyX(ctx context.Context) *RouteControl {
	node, err := rcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RouteControl ID in the query.
// Returns a *NotSingularError when more than one RouteControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (rcq *RouteControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rcq.Limit(2).IDs(setContextOp(ctx, rcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{routecontrol.Label}
	default:
		err = &NotSingularError{routecontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rcq *RouteControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RouteControls.
func (rcq *RouteControlQuery) All(ctx context.Context) ([]*RouteControl, error) {
	ctx = setContextOp(ctx, rcq.ctx, "All")
	if err := rcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RouteControl, *RouteControlQuery]()
	return withInterceptors[[]*RouteControl](ctx, rcq, qr, rcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rcq *RouteControlQuery) AllX(ctx context.Context) []*RouteControl {
	nodes, err := rcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RouteControl IDs.
func (rcq *RouteControlQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rcq.ctx.Unique == nil && rcq.path != nil {
		rcq.Unique(true)
	}
	ctx = setContextOp(ctx, rcq.ctx, "IDs")
	if err = rcq.Select(routecontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rcq *RouteControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rcq *RouteControlQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rcq.ctx, "Count")
	if err := rcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rcq, querierCount[*RouteControlQuery](), rcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rcq *RouteControlQuery) CountX(ctx context.Context) int {
	count, err := rcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rcq *RouteControlQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rcq.ctx, "Exist")
	switch _, err := rcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rcq *RouteControlQuery) ExistX(ctx context.Context) bool {
	exist, err := rcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RouteControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rcq *RouteControlQuery) Clone() *RouteControlQuery {
	if rcq == nil {
		return nil
	}
	return &RouteControlQuery{
		config:           rcq.config,
		ctx:              rcq.ctx.Clone(),
		order:            append([]routecontrol.OrderOption{}, rcq.order...),
		inters:           append([]Interceptor{}, rcq.inters...),
		predicates:       append([]predicate.RouteControl{}, rcq.predicates...),
		withOrganization: rcq.withOrganization.Clone(),
		withBusinessUnit: rcq.withBusinessUnit.Clone(),
		// clone intermediate query.
		sql:  rcq.sql.Clone(),
		path: rcq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (rcq *RouteControlQuery) WithOrganization(opts ...func(*OrganizationQuery)) *RouteControlQuery {
	query := (&OrganizationClient{config: rcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rcq.withOrganization = query
	return rcq
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (rcq *RouteControlQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *RouteControlQuery {
	query := (&BusinessUnitClient{config: rcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rcq.withBusinessUnit = query
	return rcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RouteControl.Query().
//		GroupBy(routecontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rcq *RouteControlQuery) GroupBy(field string, fields ...string) *RouteControlGroupBy {
	rcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RouteControlGroupBy{build: rcq}
	grbuild.flds = &rcq.ctx.Fields
	grbuild.label = routecontrol.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
//	}
//
//	client.RouteControl.Query().
//		Select(routecontrol.FieldCreatedAt).
//		Scan(ctx, &v)
func (rcq *RouteControlQuery) Select(fields ...string) *RouteControlSelect {
	rcq.ctx.Fields = append(rcq.ctx.Fields, fields...)
	sbuild := &RouteControlSelect{RouteControlQuery: rcq}
	sbuild.label = routecontrol.Label
	sbuild.flds, sbuild.scan = &rcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RouteControlSelect configured with the given aggregations.
func (rcq *RouteControlQuery) Aggregate(fns ...AggregateFunc) *RouteControlSelect {
	return rcq.Select().Aggregate(fns...)
}

func (rcq *RouteControlQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rcq); err != nil {
				return err
			}
		}
	}
	for _, f := range rcq.ctx.Fields {
		if !routecontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rcq.path != nil {
		prev, err := rcq.path(ctx)
		if err != nil {
			return err
		}
		rcq.sql = prev
	}
	return nil
}

func (rcq *RouteControlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RouteControl, error) {
	var (
		nodes       = []*RouteControl{}
		withFKs     = rcq.withFKs
		_spec       = rcq.querySpec()
		loadedTypes = [2]bool{
			rcq.withOrganization != nil,
			rcq.withBusinessUnit != nil,
		}
	)
	if rcq.withOrganization != nil || rcq.withBusinessUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, routecontrol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RouteControl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RouteControl{config: rcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rcq.modifiers) > 0 {
		_spec.Modifiers = rcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rcq.withOrganization; query != nil {
		if err := rcq.loadOrganization(ctx, query, nodes, nil,
			func(n *RouteControl, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := rcq.withBusinessUnit; query != nil {
		if err := rcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *RouteControl, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rcq *RouteControlQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*RouteControl, init func(*RouteControl), assign func(*RouteControl, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RouteControl)
	for i := range nodes {
		if nodes[i].organization_id == nil {
			continue
		}
		fk := *nodes[i].organization_id
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
func (rcq *RouteControlQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*RouteControl, init func(*RouteControl), assign func(*RouteControl, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*RouteControl)
	for i := range nodes {
		if nodes[i].business_unit_id == nil {
			continue
		}
		fk := *nodes[i].business_unit_id
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

func (rcq *RouteControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rcq.querySpec()
	if len(rcq.modifiers) > 0 {
		_spec.Modifiers = rcq.modifiers
	}
	_spec.Node.Columns = rcq.ctx.Fields
	if len(rcq.ctx.Fields) > 0 {
		_spec.Unique = rcq.ctx.Unique != nil && *rcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rcq.driver, _spec)
}

func (rcq *RouteControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(routecontrol.Table, routecontrol.Columns, sqlgraph.NewFieldSpec(routecontrol.FieldID, field.TypeUUID))
	_spec.From = rcq.sql
	if unique := rcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rcq.path != nil {
		_spec.Unique = true
	}
	if fields := rcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routecontrol.FieldID)
		for i := range fields {
			if fields[i] != routecontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rcq *RouteControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rcq.driver.Dialect())
	t1 := builder.Table(routecontrol.Table)
	columns := rcq.ctx.Fields
	if len(columns) == 0 {
		columns = routecontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rcq.sql != nil {
		selector = rcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rcq.ctx.Unique != nil && *rcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rcq.modifiers {
		m(selector)
	}
	for _, p := range rcq.predicates {
		p(selector)
	}
	for _, p := range rcq.order {
		p(selector)
	}
	if offset := rcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rcq *RouteControlQuery) Modify(modifiers ...func(s *sql.Selector)) *RouteControlSelect {
	rcq.modifiers = append(rcq.modifiers, modifiers...)
	return rcq.Select()
}

// RouteControlGroupBy is the group-by builder for RouteControl entities.
type RouteControlGroupBy struct {
	selector
	build *RouteControlQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rcgb *RouteControlGroupBy) Aggregate(fns ...AggregateFunc) *RouteControlGroupBy {
	rcgb.fns = append(rcgb.fns, fns...)
	return rcgb
}

// Scan applies the selector query and scans the result into the given value.
func (rcgb *RouteControlGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rcgb.build.ctx, "GroupBy")
	if err := rcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteControlQuery, *RouteControlGroupBy](ctx, rcgb.build, rcgb, rcgb.build.inters, v)
}

func (rcgb *RouteControlGroupBy) sqlScan(ctx context.Context, root *RouteControlQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rcgb.fns))
	for _, fn := range rcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rcgb.flds)+len(rcgb.fns))
		for _, f := range *rcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RouteControlSelect is the builder for selecting fields of RouteControl entities.
type RouteControlSelect struct {
	*RouteControlQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rcs *RouteControlSelect) Aggregate(fns ...AggregateFunc) *RouteControlSelect {
	rcs.fns = append(rcs.fns, fns...)
	return rcs
}

// Scan applies the selector query and scans the result into the given value.
func (rcs *RouteControlSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rcs.ctx, "Select")
	if err := rcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RouteControlQuery, *RouteControlSelect](ctx, rcs.RouteControlQuery, rcs, rcs.inters, v)
}

func (rcs *RouteControlSelect) sqlScan(ctx context.Context, root *RouteControlQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rcs.fns))
	for _, fn := range rcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rcs *RouteControlSelect) Modify(modifiers ...func(s *sql.Selector)) *RouteControlSelect {
	rcs.modifiers = append(rcs.modifiers, modifiers...)
	return rcs
}
