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
	"github.com/emoss08/trenova/internal/ent/fleetcode"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/user"
	"github.com/google/uuid"
)

// FleetCodeQuery is the builder for querying FleetCode entities.
type FleetCodeQuery struct {
	config
	ctx              *QueryContext
	order            []fleetcode.OrderOption
	inters           []Interceptor
	predicates       []predicate.FleetCode
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withManager      *UserQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FleetCodeQuery builder.
func (fcq *FleetCodeQuery) Where(ps ...predicate.FleetCode) *FleetCodeQuery {
	fcq.predicates = append(fcq.predicates, ps...)
	return fcq
}

// Limit the number of records to be returned by this query.
func (fcq *FleetCodeQuery) Limit(limit int) *FleetCodeQuery {
	fcq.ctx.Limit = &limit
	return fcq
}

// Offset to start from.
func (fcq *FleetCodeQuery) Offset(offset int) *FleetCodeQuery {
	fcq.ctx.Offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FleetCodeQuery) Unique(unique bool) *FleetCodeQuery {
	fcq.ctx.Unique = &unique
	return fcq
}

// Order specifies how the records should be ordered.
func (fcq *FleetCodeQuery) Order(o ...fleetcode.OrderOption) *FleetCodeQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (fcq *FleetCodeQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fleetcode.Table, fleetcode.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, fleetcode.BusinessUnitTable, fleetcode.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (fcq *FleetCodeQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fleetcode.Table, fleetcode.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, fleetcode.OrganizationTable, fleetcode.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryManager chains the current query on the "manager" edge.
func (fcq *FleetCodeQuery) QueryManager() *UserQuery {
	query := (&UserClient{config: fcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fleetcode.Table, fleetcode.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, fleetcode.ManagerTable, fleetcode.ManagerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FleetCode entity from the query.
// Returns a *NotFoundError when no FleetCode was found.
func (fcq *FleetCodeQuery) First(ctx context.Context) (*FleetCode, error) {
	nodes, err := fcq.Limit(1).All(setContextOp(ctx, fcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fleetcode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcq *FleetCodeQuery) FirstX(ctx context.Context) *FleetCode {
	node, err := fcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FleetCode ID from the query.
// Returns a *NotFoundError when no FleetCode ID was found.
func (fcq *FleetCodeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fcq.Limit(1).IDs(setContextOp(ctx, fcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fleetcode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcq *FleetCodeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FleetCode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FleetCode entity is found.
// Returns a *NotFoundError when no FleetCode entities are found.
func (fcq *FleetCodeQuery) Only(ctx context.Context) (*FleetCode, error) {
	nodes, err := fcq.Limit(2).All(setContextOp(ctx, fcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fleetcode.Label}
	default:
		return nil, &NotSingularError{fleetcode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcq *FleetCodeQuery) OnlyX(ctx context.Context) *FleetCode {
	node, err := fcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FleetCode ID in the query.
// Returns a *NotSingularError when more than one FleetCode ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcq *FleetCodeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fcq.Limit(2).IDs(setContextOp(ctx, fcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fleetcode.Label}
	default:
		err = &NotSingularError{fleetcode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcq *FleetCodeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FleetCodes.
func (fcq *FleetCodeQuery) All(ctx context.Context) ([]*FleetCode, error) {
	ctx = setContextOp(ctx, fcq.ctx, "All")
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FleetCode, *FleetCodeQuery]()
	return withInterceptors[[]*FleetCode](ctx, fcq, qr, fcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fcq *FleetCodeQuery) AllX(ctx context.Context) []*FleetCode {
	nodes, err := fcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FleetCode IDs.
func (fcq *FleetCodeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if fcq.ctx.Unique == nil && fcq.path != nil {
		fcq.Unique(true)
	}
	ctx = setContextOp(ctx, fcq.ctx, "IDs")
	if err = fcq.Select(fleetcode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcq *FleetCodeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcq *FleetCodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Count")
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fcq, querierCount[*FleetCodeQuery](), fcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fcq *FleetCodeQuery) CountX(ctx context.Context) int {
	count, err := fcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcq *FleetCodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fcq.ctx, "Exist")
	switch _, err := fcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fcq *FleetCodeQuery) ExistX(ctx context.Context) bool {
	exist, err := fcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FleetCodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcq *FleetCodeQuery) Clone() *FleetCodeQuery {
	if fcq == nil {
		return nil
	}
	return &FleetCodeQuery{
		config:           fcq.config,
		ctx:              fcq.ctx.Clone(),
		order:            append([]fleetcode.OrderOption{}, fcq.order...),
		inters:           append([]Interceptor{}, fcq.inters...),
		predicates:       append([]predicate.FleetCode{}, fcq.predicates...),
		withBusinessUnit: fcq.withBusinessUnit.Clone(),
		withOrganization: fcq.withOrganization.Clone(),
		withManager:      fcq.withManager.Clone(),
		// clone intermediate query.
		sql:  fcq.sql.Clone(),
		path: fcq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FleetCodeQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *FleetCodeQuery {
	query := (&BusinessUnitClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withBusinessUnit = query
	return fcq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FleetCodeQuery) WithOrganization(opts ...func(*OrganizationQuery)) *FleetCodeQuery {
	query := (&OrganizationClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withOrganization = query
	return fcq
}

// WithManager tells the query-builder to eager-load the nodes that are connected to
// the "manager" edge. The optional arguments are used to configure the query builder of the edge.
func (fcq *FleetCodeQuery) WithManager(opts ...func(*UserQuery)) *FleetCodeQuery {
	query := (&UserClient{config: fcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fcq.withManager = query
	return fcq
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
//	client.FleetCode.Query().
//		GroupBy(fleetcode.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fcq *FleetCodeQuery) GroupBy(field string, fields ...string) *FleetCodeGroupBy {
	fcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FleetCodeGroupBy{build: fcq}
	grbuild.flds = &fcq.ctx.Fields
	grbuild.label = fleetcode.Label
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
//	client.FleetCode.Query().
//		Select(fleetcode.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (fcq *FleetCodeQuery) Select(fields ...string) *FleetCodeSelect {
	fcq.ctx.Fields = append(fcq.ctx.Fields, fields...)
	sbuild := &FleetCodeSelect{FleetCodeQuery: fcq}
	sbuild.label = fleetcode.Label
	sbuild.flds, sbuild.scan = &fcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FleetCodeSelect configured with the given aggregations.
func (fcq *FleetCodeQuery) Aggregate(fns ...AggregateFunc) *FleetCodeSelect {
	return fcq.Select().Aggregate(fns...)
}

func (fcq *FleetCodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fcq); err != nil {
				return err
			}
		}
	}
	for _, f := range fcq.ctx.Fields {
		if !fleetcode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcq.path != nil {
		prev, err := fcq.path(ctx)
		if err != nil {
			return err
		}
		fcq.sql = prev
	}
	return nil
}

func (fcq *FleetCodeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FleetCode, error) {
	var (
		nodes       = []*FleetCode{}
		_spec       = fcq.querySpec()
		loadedTypes = [3]bool{
			fcq.withBusinessUnit != nil,
			fcq.withOrganization != nil,
			fcq.withManager != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FleetCode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FleetCode{config: fcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fcq.modifiers) > 0 {
		_spec.Modifiers = fcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fcq.withBusinessUnit; query != nil {
		if err := fcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *FleetCode, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withOrganization; query != nil {
		if err := fcq.loadOrganization(ctx, query, nodes, nil,
			func(n *FleetCode, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := fcq.withManager; query != nil {
		if err := fcq.loadManager(ctx, query, nodes, nil,
			func(n *FleetCode, e *User) { n.Edges.Manager = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fcq *FleetCodeQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*FleetCode, init func(*FleetCode), assign func(*FleetCode, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FleetCode)
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
func (fcq *FleetCodeQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*FleetCode, init func(*FleetCode), assign func(*FleetCode, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FleetCode)
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
func (fcq *FleetCodeQuery) loadManager(ctx context.Context, query *UserQuery, nodes []*FleetCode, init func(*FleetCode), assign func(*FleetCode, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*FleetCode)
	for i := range nodes {
		if nodes[i].ManagerID == nil {
			continue
		}
		fk := *nodes[i].ManagerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "manager_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fcq *FleetCodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	if len(fcq.modifiers) > 0 {
		_spec.Modifiers = fcq.modifiers
	}
	_spec.Node.Columns = fcq.ctx.Fields
	if len(fcq.ctx.Fields) > 0 {
		_spec.Unique = fcq.ctx.Unique != nil && *fcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FleetCodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fleetcode.Table, fleetcode.Columns, sqlgraph.NewFieldSpec(fleetcode.FieldID, field.TypeUUID))
	_spec.From = fcq.sql
	if unique := fcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fcq.path != nil {
		_spec.Unique = true
	}
	if fields := fcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fleetcode.FieldID)
		for i := range fields {
			if fields[i] != fleetcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fcq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(fleetcode.FieldBusinessUnitID)
		}
		if fcq.withOrganization != nil {
			_spec.Node.AddColumnOnce(fleetcode.FieldOrganizationID)
		}
		if fcq.withManager != nil {
			_spec.Node.AddColumnOnce(fleetcode.FieldManagerID)
		}
	}
	if ps := fcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcq *FleetCodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcq.driver.Dialect())
	t1 := builder.Table(fleetcode.Table)
	columns := fcq.ctx.Fields
	if len(columns) == 0 {
		columns = fleetcode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.ctx.Unique != nil && *fcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range fcq.modifiers {
		m(selector)
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcq *FleetCodeQuery) Modify(modifiers ...func(s *sql.Selector)) *FleetCodeSelect {
	fcq.modifiers = append(fcq.modifiers, modifiers...)
	return fcq.Select()
}

// FleetCodeGroupBy is the group-by builder for FleetCode entities.
type FleetCodeGroupBy struct {
	selector
	build *FleetCodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FleetCodeGroupBy) Aggregate(fns ...AggregateFunc) *FleetCodeGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the selector query and scans the result into the given value.
func (fcgb *FleetCodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcgb.build.ctx, "GroupBy")
	if err := fcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FleetCodeQuery, *FleetCodeGroupBy](ctx, fcgb.build, fcgb, fcgb.build.inters, v)
}

func (fcgb *FleetCodeGroupBy) sqlScan(ctx context.Context, root *FleetCodeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fcgb.flds)+len(fcgb.fns))
		for _, f := range *fcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FleetCodeSelect is the builder for selecting fields of FleetCode entities.
type FleetCodeSelect struct {
	*FleetCodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fcs *FleetCodeSelect) Aggregate(fns ...AggregateFunc) *FleetCodeSelect {
	fcs.fns = append(fcs.fns, fns...)
	return fcs
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FleetCodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fcs.ctx, "Select")
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FleetCodeQuery, *FleetCodeSelect](ctx, fcs.FleetCodeQuery, fcs, fcs.inters, v)
}

func (fcs *FleetCodeSelect) sqlScan(ctx context.Context, root *FleetCodeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fcs.fns))
	for _, fn := range fcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcs *FleetCodeSelect) Modify(modifiers ...func(s *sql.Selector)) *FleetCodeSelect {
	fcs.modifiers = append(fcs.modifiers, modifiers...)
	return fcs
}
