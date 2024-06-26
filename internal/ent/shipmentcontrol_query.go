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
	"github.com/emoss08/trenova/internal/ent/shipmentcontrol"
	"github.com/google/uuid"
)

// ShipmentControlQuery is the builder for querying ShipmentControl entities.
type ShipmentControlQuery struct {
	config
	ctx              *QueryContext
	order            []shipmentcontrol.OrderOption
	inters           []Interceptor
	predicates       []predicate.ShipmentControl
	withOrganization *OrganizationQuery
	withBusinessUnit *BusinessUnitQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShipmentControlQuery builder.
func (scq *ShipmentControlQuery) Where(ps ...predicate.ShipmentControl) *ShipmentControlQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ShipmentControlQuery) Limit(limit int) *ShipmentControlQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ShipmentControlQuery) Offset(offset int) *ShipmentControlQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ShipmentControlQuery) Unique(unique bool) *ShipmentControlQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ShipmentControlQuery) Order(o ...shipmentcontrol.OrderOption) *ShipmentControlQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryOrganization chains the current query on the "organization" edge.
func (scq *ShipmentControlQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcontrol.Table, shipmentcontrol.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, shipmentcontrol.OrganizationTable, shipmentcontrol.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (scq *ShipmentControlQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shipmentcontrol.Table, shipmentcontrol.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, shipmentcontrol.BusinessUnitTable, shipmentcontrol.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShipmentControl entity from the query.
// Returns a *NotFoundError when no ShipmentControl was found.
func (scq *ShipmentControlQuery) First(ctx context.Context) (*ShipmentControl, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shipmentcontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ShipmentControlQuery) FirstX(ctx context.Context) *ShipmentControl {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShipmentControl ID from the query.
// Returns a *NotFoundError when no ShipmentControl ID was found.
func (scq *ShipmentControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shipmentcontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ShipmentControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShipmentControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShipmentControl entity is found.
// Returns a *NotFoundError when no ShipmentControl entities are found.
func (scq *ShipmentControlQuery) Only(ctx context.Context) (*ShipmentControl, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shipmentcontrol.Label}
	default:
		return nil, &NotSingularError{shipmentcontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ShipmentControlQuery) OnlyX(ctx context.Context) *ShipmentControl {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShipmentControl ID in the query.
// Returns a *NotSingularError when more than one ShipmentControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ShipmentControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shipmentcontrol.Label}
	default:
		err = &NotSingularError{shipmentcontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ShipmentControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShipmentControls.
func (scq *ShipmentControlQuery) All(ctx context.Context) ([]*ShipmentControl, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShipmentControl, *ShipmentControlQuery]()
	return withInterceptors[[]*ShipmentControl](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ShipmentControlQuery) AllX(ctx context.Context) []*ShipmentControl {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShipmentControl IDs.
func (scq *ShipmentControlQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(shipmentcontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ShipmentControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ShipmentControlQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ShipmentControlQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ShipmentControlQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ShipmentControlQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ShipmentControlQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShipmentControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ShipmentControlQuery) Clone() *ShipmentControlQuery {
	if scq == nil {
		return nil
	}
	return &ShipmentControlQuery{
		config:           scq.config,
		ctx:              scq.ctx.Clone(),
		order:            append([]shipmentcontrol.OrderOption{}, scq.order...),
		inters:           append([]Interceptor{}, scq.inters...),
		predicates:       append([]predicate.ShipmentControl{}, scq.predicates...),
		withOrganization: scq.withOrganization.Clone(),
		withBusinessUnit: scq.withBusinessUnit.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentControlQuery) WithOrganization(opts ...func(*OrganizationQuery)) *ShipmentControlQuery {
	query := (&OrganizationClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withOrganization = query
	return scq
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ShipmentControlQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *ShipmentControlQuery {
	query := (&BusinessUnitClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withBusinessUnit = query
	return scq
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
//	client.ShipmentControl.Query().
//		GroupBy(shipmentcontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ShipmentControlQuery) GroupBy(field string, fields ...string) *ShipmentControlGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShipmentControlGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = shipmentcontrol.Label
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
//	client.ShipmentControl.Query().
//		Select(shipmentcontrol.FieldCreatedAt).
//		Scan(ctx, &v)
func (scq *ShipmentControlQuery) Select(fields ...string) *ShipmentControlSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ShipmentControlSelect{ShipmentControlQuery: scq}
	sbuild.label = shipmentcontrol.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShipmentControlSelect configured with the given aggregations.
func (scq *ShipmentControlQuery) Aggregate(fns ...AggregateFunc) *ShipmentControlSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ShipmentControlQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !shipmentcontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ShipmentControlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShipmentControl, error) {
	var (
		nodes       = []*ShipmentControl{}
		withFKs     = scq.withFKs
		_spec       = scq.querySpec()
		loadedTypes = [2]bool{
			scq.withOrganization != nil,
			scq.withBusinessUnit != nil,
		}
	)
	if scq.withOrganization != nil || scq.withBusinessUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentcontrol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShipmentControl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShipmentControl{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withOrganization; query != nil {
		if err := scq.loadOrganization(ctx, query, nodes, nil,
			func(n *ShipmentControl, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := scq.withBusinessUnit; query != nil {
		if err := scq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *ShipmentControl, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ShipmentControlQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*ShipmentControl, init func(*ShipmentControl), assign func(*ShipmentControl, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentControl)
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
func (scq *ShipmentControlQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*ShipmentControl, init func(*ShipmentControl), assign func(*ShipmentControl, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShipmentControl)
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

func (scq *ShipmentControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ShipmentControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shipmentcontrol.Table, shipmentcontrol.Columns, sqlgraph.NewFieldSpec(shipmentcontrol.FieldID, field.TypeUUID))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shipmentcontrol.FieldID)
		for i := range fields {
			if fields[i] != shipmentcontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ShipmentControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(shipmentcontrol.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = shipmentcontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range scq.modifiers {
		m(selector)
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scq *ShipmentControlQuery) Modify(modifiers ...func(s *sql.Selector)) *ShipmentControlSelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ShipmentControlGroupBy is the group-by builder for ShipmentControl entities.
type ShipmentControlGroupBy struct {
	selector
	build *ShipmentControlQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ShipmentControlGroupBy) Aggregate(fns ...AggregateFunc) *ShipmentControlGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ShipmentControlGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentControlQuery, *ShipmentControlGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ShipmentControlGroupBy) sqlScan(ctx context.Context, root *ShipmentControlQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShipmentControlSelect is the builder for selecting fields of ShipmentControl entities.
type ShipmentControlSelect struct {
	*ShipmentControlQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ShipmentControlSelect) Aggregate(fns ...AggregateFunc) *ShipmentControlSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ShipmentControlSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShipmentControlQuery, *ShipmentControlSelect](ctx, scs.ShipmentControlQuery, scs, scs.inters, v)
}

func (scs *ShipmentControlSelect) sqlScan(ctx context.Context, root *ShipmentControlQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scs *ShipmentControlSelect) Modify(modifiers ...func(s *sql.Selector)) *ShipmentControlSelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
