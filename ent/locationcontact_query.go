// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/location"
	"github.com/emoss08/trenova/ent/locationcontact"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// LocationContactQuery is the builder for querying LocationContact entities.
type LocationContactQuery struct {
	config
	ctx              *QueryContext
	order            []locationcontact.OrderOption
	inters           []Interceptor
	predicates       []predicate.LocationContact
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	withLocation     *LocationQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LocationContactQuery builder.
func (lcq *LocationContactQuery) Where(ps ...predicate.LocationContact) *LocationContactQuery {
	lcq.predicates = append(lcq.predicates, ps...)
	return lcq
}

// Limit the number of records to be returned by this query.
func (lcq *LocationContactQuery) Limit(limit int) *LocationContactQuery {
	lcq.ctx.Limit = &limit
	return lcq
}

// Offset to start from.
func (lcq *LocationContactQuery) Offset(offset int) *LocationContactQuery {
	lcq.ctx.Offset = &offset
	return lcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lcq *LocationContactQuery) Unique(unique bool) *LocationContactQuery {
	lcq.ctx.Unique = &unique
	return lcq
}

// Order specifies how the records should be ordered.
func (lcq *LocationContactQuery) Order(o ...locationcontact.OrderOption) *LocationContactQuery {
	lcq.order = append(lcq.order, o...)
	return lcq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (lcq *LocationContactQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(locationcontact.Table, locationcontact.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, locationcontact.BusinessUnitTable, locationcontact.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (lcq *LocationContactQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(locationcontact.Table, locationcontact.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, locationcontact.OrganizationTable, locationcontact.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLocation chains the current query on the "location" edge.
func (lcq *LocationContactQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: lcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(locationcontact.Table, locationcontact.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, locationcontact.LocationTable, locationcontact.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(lcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LocationContact entity from the query.
// Returns a *NotFoundError when no LocationContact was found.
func (lcq *LocationContactQuery) First(ctx context.Context) (*LocationContact, error) {
	nodes, err := lcq.Limit(1).All(setContextOp(ctx, lcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{locationcontact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lcq *LocationContactQuery) FirstX(ctx context.Context) *LocationContact {
	node, err := lcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LocationContact ID from the query.
// Returns a *NotFoundError when no LocationContact ID was found.
func (lcq *LocationContactQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(1).IDs(setContextOp(ctx, lcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{locationcontact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lcq *LocationContactQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LocationContact entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LocationContact entity is found.
// Returns a *NotFoundError when no LocationContact entities are found.
func (lcq *LocationContactQuery) Only(ctx context.Context) (*LocationContact, error) {
	nodes, err := lcq.Limit(2).All(setContextOp(ctx, lcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{locationcontact.Label}
	default:
		return nil, &NotSingularError{locationcontact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lcq *LocationContactQuery) OnlyX(ctx context.Context) *LocationContact {
	node, err := lcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LocationContact ID in the query.
// Returns a *NotSingularError when more than one LocationContact ID is found.
// Returns a *NotFoundError when no entities are found.
func (lcq *LocationContactQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lcq.Limit(2).IDs(setContextOp(ctx, lcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{locationcontact.Label}
	default:
		err = &NotSingularError{locationcontact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lcq *LocationContactQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LocationContacts.
func (lcq *LocationContactQuery) All(ctx context.Context) ([]*LocationContact, error) {
	ctx = setContextOp(ctx, lcq.ctx, "All")
	if err := lcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*LocationContact, *LocationContactQuery]()
	return withInterceptors[[]*LocationContact](ctx, lcq, qr, lcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lcq *LocationContactQuery) AllX(ctx context.Context) []*LocationContact {
	nodes, err := lcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LocationContact IDs.
func (lcq *LocationContactQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if lcq.ctx.Unique == nil && lcq.path != nil {
		lcq.Unique(true)
	}
	ctx = setContextOp(ctx, lcq.ctx, "IDs")
	if err = lcq.Select(locationcontact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lcq *LocationContactQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lcq *LocationContactQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lcq.ctx, "Count")
	if err := lcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lcq, querierCount[*LocationContactQuery](), lcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lcq *LocationContactQuery) CountX(ctx context.Context) int {
	count, err := lcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lcq *LocationContactQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lcq.ctx, "Exist")
	switch _, err := lcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lcq *LocationContactQuery) ExistX(ctx context.Context) bool {
	exist, err := lcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LocationContactQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lcq *LocationContactQuery) Clone() *LocationContactQuery {
	if lcq == nil {
		return nil
	}
	return &LocationContactQuery{
		config:           lcq.config,
		ctx:              lcq.ctx.Clone(),
		order:            append([]locationcontact.OrderOption{}, lcq.order...),
		inters:           append([]Interceptor{}, lcq.inters...),
		predicates:       append([]predicate.LocationContact{}, lcq.predicates...),
		withBusinessUnit: lcq.withBusinessUnit.Clone(),
		withOrganization: lcq.withOrganization.Clone(),
		withLocation:     lcq.withLocation.Clone(),
		// clone intermediate query.
		sql:  lcq.sql.Clone(),
		path: lcq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LocationContactQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *LocationContactQuery {
	query := (&BusinessUnitClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withBusinessUnit = query
	return lcq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LocationContactQuery) WithOrganization(opts ...func(*OrganizationQuery)) *LocationContactQuery {
	query := (&OrganizationClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withOrganization = query
	return lcq
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (lcq *LocationContactQuery) WithLocation(opts ...func(*LocationQuery)) *LocationContactQuery {
	query := (&LocationClient{config: lcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lcq.withLocation = query
	return lcq
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
//	client.LocationContact.Query().
//		GroupBy(locationcontact.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lcq *LocationContactQuery) GroupBy(field string, fields ...string) *LocationContactGroupBy {
	lcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LocationContactGroupBy{build: lcq}
	grbuild.flds = &lcq.ctx.Fields
	grbuild.label = locationcontact.Label
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
//	client.LocationContact.Query().
//		Select(locationcontact.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (lcq *LocationContactQuery) Select(fields ...string) *LocationContactSelect {
	lcq.ctx.Fields = append(lcq.ctx.Fields, fields...)
	sbuild := &LocationContactSelect{LocationContactQuery: lcq}
	sbuild.label = locationcontact.Label
	sbuild.flds, sbuild.scan = &lcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LocationContactSelect configured with the given aggregations.
func (lcq *LocationContactQuery) Aggregate(fns ...AggregateFunc) *LocationContactSelect {
	return lcq.Select().Aggregate(fns...)
}

func (lcq *LocationContactQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lcq); err != nil {
				return err
			}
		}
	}
	for _, f := range lcq.ctx.Fields {
		if !locationcontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lcq.path != nil {
		prev, err := lcq.path(ctx)
		if err != nil {
			return err
		}
		lcq.sql = prev
	}
	return nil
}

func (lcq *LocationContactQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LocationContact, error) {
	var (
		nodes       = []*LocationContact{}
		_spec       = lcq.querySpec()
		loadedTypes = [3]bool{
			lcq.withBusinessUnit != nil,
			lcq.withOrganization != nil,
			lcq.withLocation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LocationContact).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LocationContact{config: lcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lcq.withBusinessUnit; query != nil {
		if err := lcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *LocationContact, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := lcq.withOrganization; query != nil {
		if err := lcq.loadOrganization(ctx, query, nodes, nil,
			func(n *LocationContact, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := lcq.withLocation; query != nil {
		if err := lcq.loadLocation(ctx, query, nodes, nil,
			func(n *LocationContact, e *Location) { n.Edges.Location = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lcq *LocationContactQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*LocationContact, init func(*LocationContact), assign func(*LocationContact, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LocationContact)
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
func (lcq *LocationContactQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*LocationContact, init func(*LocationContact), assign func(*LocationContact, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LocationContact)
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
func (lcq *LocationContactQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*LocationContact, init func(*LocationContact), assign func(*LocationContact, *Location)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*LocationContact)
	for i := range nodes {
		fk := nodes[i].LocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lcq *LocationContactQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lcq.querySpec()
	if len(lcq.modifiers) > 0 {
		_spec.Modifiers = lcq.modifiers
	}
	_spec.Node.Columns = lcq.ctx.Fields
	if len(lcq.ctx.Fields) > 0 {
		_spec.Unique = lcq.ctx.Unique != nil && *lcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lcq.driver, _spec)
}

func (lcq *LocationContactQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(locationcontact.Table, locationcontact.Columns, sqlgraph.NewFieldSpec(locationcontact.FieldID, field.TypeUUID))
	_spec.From = lcq.sql
	if unique := lcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lcq.path != nil {
		_spec.Unique = true
	}
	if fields := lcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, locationcontact.FieldID)
		for i := range fields {
			if fields[i] != locationcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if lcq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(locationcontact.FieldBusinessUnitID)
		}
		if lcq.withOrganization != nil {
			_spec.Node.AddColumnOnce(locationcontact.FieldOrganizationID)
		}
		if lcq.withLocation != nil {
			_spec.Node.AddColumnOnce(locationcontact.FieldLocationID)
		}
	}
	if ps := lcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lcq *LocationContactQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lcq.driver.Dialect())
	t1 := builder.Table(locationcontact.Table)
	columns := lcq.ctx.Fields
	if len(columns) == 0 {
		columns = locationcontact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lcq.sql != nil {
		selector = lcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lcq.ctx.Unique != nil && *lcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range lcq.modifiers {
		m(selector)
	}
	for _, p := range lcq.predicates {
		p(selector)
	}
	for _, p := range lcq.order {
		p(selector)
	}
	if offset := lcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcq *LocationContactQuery) Modify(modifiers ...func(s *sql.Selector)) *LocationContactSelect {
	lcq.modifiers = append(lcq.modifiers, modifiers...)
	return lcq.Select()
}

// LocationContactGroupBy is the group-by builder for LocationContact entities.
type LocationContactGroupBy struct {
	selector
	build *LocationContactQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lcgb *LocationContactGroupBy) Aggregate(fns ...AggregateFunc) *LocationContactGroupBy {
	lcgb.fns = append(lcgb.fns, fns...)
	return lcgb
}

// Scan applies the selector query and scans the result into the given value.
func (lcgb *LocationContactGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcgb.build.ctx, "GroupBy")
	if err := lcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationContactQuery, *LocationContactGroupBy](ctx, lcgb.build, lcgb, lcgb.build.inters, v)
}

func (lcgb *LocationContactGroupBy) sqlScan(ctx context.Context, root *LocationContactQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lcgb.fns))
	for _, fn := range lcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lcgb.flds)+len(lcgb.fns))
		for _, f := range *lcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LocationContactSelect is the builder for selecting fields of LocationContact entities.
type LocationContactSelect struct {
	*LocationContactQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (lcs *LocationContactSelect) Aggregate(fns ...AggregateFunc) *LocationContactSelect {
	lcs.fns = append(lcs.fns, fns...)
	return lcs
}

// Scan applies the selector query and scans the result into the given value.
func (lcs *LocationContactSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lcs.ctx, "Select")
	if err := lcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LocationContactQuery, *LocationContactSelect](ctx, lcs.LocationContactQuery, lcs, lcs.inters, v)
}

func (lcs *LocationContactSelect) sqlScan(ctx context.Context, root *LocationContactQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(lcs.fns))
	for _, fn := range lcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*lcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lcs *LocationContactSelect) Modify(modifiers ...func(s *sql.Selector)) *LocationContactSelect {
	lcs.modifiers = append(lcs.modifiers, modifiers...)
	return lcs
}