// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/accessorialcharge"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/emoss08/trenova/internal/ent/shipmentcharges"
	"github.com/google/uuid"
)

// AccessorialChargeQuery is the builder for querying AccessorialCharge entities.
type AccessorialChargeQuery struct {
	config
	ctx                      *QueryContext
	order                    []accessorialcharge.OrderOption
	inters                   []Interceptor
	predicates               []predicate.AccessorialCharge
	withBusinessUnit         *BusinessUnitQuery
	withOrganization         *OrganizationQuery
	withShipmentCharges      *ShipmentChargesQuery
	modifiers                []func(*sql.Selector)
	withNamedShipmentCharges map[string]*ShipmentChargesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccessorialChargeQuery builder.
func (acq *AccessorialChargeQuery) Where(ps ...predicate.AccessorialCharge) *AccessorialChargeQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AccessorialChargeQuery) Limit(limit int) *AccessorialChargeQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AccessorialChargeQuery) Offset(offset int) *AccessorialChargeQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AccessorialChargeQuery) Unique(unique bool) *AccessorialChargeQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AccessorialChargeQuery) Order(o ...accessorialcharge.OrderOption) *AccessorialChargeQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (acq *AccessorialChargeQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accessorialcharge.Table, accessorialcharge.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, accessorialcharge.BusinessUnitTable, accessorialcharge.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (acq *AccessorialChargeQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accessorialcharge.Table, accessorialcharge.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, accessorialcharge.OrganizationTable, accessorialcharge.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShipmentCharges chains the current query on the "shipment_charges" edge.
func (acq *AccessorialChargeQuery) QueryShipmentCharges() *ShipmentChargesQuery {
	query := (&ShipmentChargesClient{config: acq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(accessorialcharge.Table, accessorialcharge.FieldID, selector),
			sqlgraph.To(shipmentcharges.Table, shipmentcharges.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, accessorialcharge.ShipmentChargesTable, accessorialcharge.ShipmentChargesColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AccessorialCharge entity from the query.
// Returns a *NotFoundError when no AccessorialCharge was found.
func (acq *AccessorialChargeQuery) First(ctx context.Context) (*AccessorialCharge, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accessorialcharge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AccessorialChargeQuery) FirstX(ctx context.Context) *AccessorialCharge {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccessorialCharge ID from the query.
// Returns a *NotFoundError when no AccessorialCharge ID was found.
func (acq *AccessorialChargeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accessorialcharge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AccessorialChargeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccessorialCharge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccessorialCharge entity is found.
// Returns a *NotFoundError when no AccessorialCharge entities are found.
func (acq *AccessorialChargeQuery) Only(ctx context.Context) (*AccessorialCharge, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accessorialcharge.Label}
	default:
		return nil, &NotSingularError{accessorialcharge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AccessorialChargeQuery) OnlyX(ctx context.Context) *AccessorialCharge {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccessorialCharge ID in the query.
// Returns a *NotSingularError when more than one AccessorialCharge ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AccessorialChargeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accessorialcharge.Label}
	default:
		err = &NotSingularError{accessorialcharge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AccessorialChargeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccessorialCharges.
func (acq *AccessorialChargeQuery) All(ctx context.Context) ([]*AccessorialCharge, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccessorialCharge, *AccessorialChargeQuery]()
	return withInterceptors[[]*AccessorialCharge](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AccessorialChargeQuery) AllX(ctx context.Context) []*AccessorialCharge {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccessorialCharge IDs.
func (acq *AccessorialChargeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(accessorialcharge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AccessorialChargeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AccessorialChargeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AccessorialChargeQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AccessorialChargeQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AccessorialChargeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AccessorialChargeQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccessorialChargeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AccessorialChargeQuery) Clone() *AccessorialChargeQuery {
	if acq == nil {
		return nil
	}
	return &AccessorialChargeQuery{
		config:              acq.config,
		ctx:                 acq.ctx.Clone(),
		order:               append([]accessorialcharge.OrderOption{}, acq.order...),
		inters:              append([]Interceptor{}, acq.inters...),
		predicates:          append([]predicate.AccessorialCharge{}, acq.predicates...),
		withBusinessUnit:    acq.withBusinessUnit.Clone(),
		withOrganization:    acq.withOrganization.Clone(),
		withShipmentCharges: acq.withShipmentCharges.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccessorialChargeQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *AccessorialChargeQuery {
	query := (&BusinessUnitClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withBusinessUnit = query
	return acq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccessorialChargeQuery) WithOrganization(opts ...func(*OrganizationQuery)) *AccessorialChargeQuery {
	query := (&OrganizationClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withOrganization = query
	return acq
}

// WithShipmentCharges tells the query-builder to eager-load the nodes that are connected to
// the "shipment_charges" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AccessorialChargeQuery) WithShipmentCharges(opts ...func(*ShipmentChargesQuery)) *AccessorialChargeQuery {
	query := (&ShipmentChargesClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	acq.withShipmentCharges = query
	return acq
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
//	client.AccessorialCharge.Query().
//		GroupBy(accessorialcharge.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *AccessorialChargeQuery) GroupBy(field string, fields ...string) *AccessorialChargeGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccessorialChargeGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = accessorialcharge.Label
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
//	client.AccessorialCharge.Query().
//		Select(accessorialcharge.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (acq *AccessorialChargeQuery) Select(fields ...string) *AccessorialChargeSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AccessorialChargeSelect{AccessorialChargeQuery: acq}
	sbuild.label = accessorialcharge.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccessorialChargeSelect configured with the given aggregations.
func (acq *AccessorialChargeQuery) Aggregate(fns ...AggregateFunc) *AccessorialChargeSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AccessorialChargeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !accessorialcharge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AccessorialChargeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccessorialCharge, error) {
	var (
		nodes       = []*AccessorialCharge{}
		_spec       = acq.querySpec()
		loadedTypes = [3]bool{
			acq.withBusinessUnit != nil,
			acq.withOrganization != nil,
			acq.withShipmentCharges != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccessorialCharge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccessorialCharge{config: acq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := acq.withBusinessUnit; query != nil {
		if err := acq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *AccessorialCharge, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := acq.withOrganization; query != nil {
		if err := acq.loadOrganization(ctx, query, nodes, nil,
			func(n *AccessorialCharge, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := acq.withShipmentCharges; query != nil {
		if err := acq.loadShipmentCharges(ctx, query, nodes,
			func(n *AccessorialCharge) { n.Edges.ShipmentCharges = []*ShipmentCharges{} },
			func(n *AccessorialCharge, e *ShipmentCharges) {
				n.Edges.ShipmentCharges = append(n.Edges.ShipmentCharges, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range acq.withNamedShipmentCharges {
		if err := acq.loadShipmentCharges(ctx, query, nodes,
			func(n *AccessorialCharge) { n.appendNamedShipmentCharges(name) },
			func(n *AccessorialCharge, e *ShipmentCharges) { n.appendNamedShipmentCharges(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *AccessorialChargeQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*AccessorialCharge, init func(*AccessorialCharge), assign func(*AccessorialCharge, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccessorialCharge)
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
func (acq *AccessorialChargeQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*AccessorialCharge, init func(*AccessorialCharge), assign func(*AccessorialCharge, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AccessorialCharge)
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
func (acq *AccessorialChargeQuery) loadShipmentCharges(ctx context.Context, query *ShipmentChargesQuery, nodes []*AccessorialCharge, init func(*AccessorialCharge), assign func(*AccessorialCharge, *ShipmentCharges)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*AccessorialCharge)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(shipmentcharges.FieldAccessorialChargeID)
	}
	query.Where(predicate.ShipmentCharges(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(accessorialcharge.ShipmentChargesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AccessorialChargeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "accessorial_charge_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (acq *AccessorialChargeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AccessorialChargeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accessorialcharge.Table, accessorialcharge.Columns, sqlgraph.NewFieldSpec(accessorialcharge.FieldID, field.TypeUUID))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accessorialcharge.FieldID)
		for i := range fields {
			if fields[i] != accessorialcharge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if acq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(accessorialcharge.FieldBusinessUnitID)
		}
		if acq.withOrganization != nil {
			_spec.Node.AddColumnOnce(accessorialcharge.FieldOrganizationID)
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AccessorialChargeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(accessorialcharge.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = accessorialcharge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range acq.modifiers {
		m(selector)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acq *AccessorialChargeQuery) Modify(modifiers ...func(s *sql.Selector)) *AccessorialChargeSelect {
	acq.modifiers = append(acq.modifiers, modifiers...)
	return acq.Select()
}

// WithNamedShipmentCharges tells the query-builder to eager-load the nodes that are connected to the "shipment_charges"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (acq *AccessorialChargeQuery) WithNamedShipmentCharges(name string, opts ...func(*ShipmentChargesQuery)) *AccessorialChargeQuery {
	query := (&ShipmentChargesClient{config: acq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if acq.withNamedShipmentCharges == nil {
		acq.withNamedShipmentCharges = make(map[string]*ShipmentChargesQuery)
	}
	acq.withNamedShipmentCharges[name] = query
	return acq
}

// AccessorialChargeGroupBy is the group-by builder for AccessorialCharge entities.
type AccessorialChargeGroupBy struct {
	selector
	build *AccessorialChargeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AccessorialChargeGroupBy) Aggregate(fns ...AggregateFunc) *AccessorialChargeGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AccessorialChargeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessorialChargeQuery, *AccessorialChargeGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AccessorialChargeGroupBy) sqlScan(ctx context.Context, root *AccessorialChargeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccessorialChargeSelect is the builder for selecting fields of AccessorialCharge entities.
type AccessorialChargeSelect struct {
	*AccessorialChargeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AccessorialChargeSelect) Aggregate(fns ...AggregateFunc) *AccessorialChargeSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AccessorialChargeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccessorialChargeQuery, *AccessorialChargeSelect](ctx, acs.AccessorialChargeQuery, acs, acs.inters, v)
}

func (acs *AccessorialChargeSelect) sqlScan(ctx context.Context, root *AccessorialChargeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acs *AccessorialChargeSelect) Modify(modifiers ...func(s *sql.Selector)) *AccessorialChargeSelect {
	acs.modifiers = append(acs.modifiers, modifiers...)
	return acs
}
