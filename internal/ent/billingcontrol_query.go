// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/billingcontrol"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// BillingControlQuery is the builder for querying BillingControl entities.
type BillingControlQuery struct {
	config
	ctx              *QueryContext
	order            []billingcontrol.OrderOption
	inters           []Interceptor
	predicates       []predicate.BillingControl
	withOrganization *OrganizationQuery
	withBusinessUnit *BusinessUnitQuery
	withFKs          bool
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingControlQuery builder.
func (bcq *BillingControlQuery) Where(ps ...predicate.BillingControl) *BillingControlQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit the number of records to be returned by this query.
func (bcq *BillingControlQuery) Limit(limit int) *BillingControlQuery {
	bcq.ctx.Limit = &limit
	return bcq
}

// Offset to start from.
func (bcq *BillingControlQuery) Offset(offset int) *BillingControlQuery {
	bcq.ctx.Offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BillingControlQuery) Unique(unique bool) *BillingControlQuery {
	bcq.ctx.Unique = &unique
	return bcq
}

// Order specifies how the records should be ordered.
func (bcq *BillingControlQuery) Order(o ...billingcontrol.OrderOption) *BillingControlQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// QueryOrganization chains the current query on the "organization" edge.
func (bcq *BillingControlQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: bcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billingcontrol.Table, billingcontrol.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, billingcontrol.OrganizationTable, billingcontrol.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (bcq *BillingControlQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: bcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billingcontrol.Table, billingcontrol.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, billingcontrol.BusinessUnitTable, billingcontrol.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(bcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillingControl entity from the query.
// Returns a *NotFoundError when no BillingControl was found.
func (bcq *BillingControlQuery) First(ctx context.Context) (*BillingControl, error) {
	nodes, err := bcq.Limit(1).All(setContextOp(ctx, bcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billingcontrol.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BillingControlQuery) FirstX(ctx context.Context) *BillingControl {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingControl ID from the query.
// Returns a *NotFoundError when no BillingControl ID was found.
func (bcq *BillingControlQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bcq.Limit(1).IDs(setContextOp(ctx, bcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billingcontrol.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BillingControlQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingControl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingControl entity is found.
// Returns a *NotFoundError when no BillingControl entities are found.
func (bcq *BillingControlQuery) Only(ctx context.Context) (*BillingControl, error) {
	nodes, err := bcq.Limit(2).All(setContextOp(ctx, bcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billingcontrol.Label}
	default:
		return nil, &NotSingularError{billingcontrol.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BillingControlQuery) OnlyX(ctx context.Context) *BillingControl {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingControl ID in the query.
// Returns a *NotSingularError when more than one BillingControl ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BillingControlQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = bcq.Limit(2).IDs(setContextOp(ctx, bcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billingcontrol.Label}
	default:
		err = &NotSingularError{billingcontrol.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BillingControlQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingControls.
func (bcq *BillingControlQuery) All(ctx context.Context) ([]*BillingControl, error) {
	ctx = setContextOp(ctx, bcq.ctx, "All")
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingControl, *BillingControlQuery]()
	return withInterceptors[[]*BillingControl](ctx, bcq, qr, bcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BillingControlQuery) AllX(ctx context.Context) []*BillingControl {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingControl IDs.
func (bcq *BillingControlQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if bcq.ctx.Unique == nil && bcq.path != nil {
		bcq.Unique(true)
	}
	ctx = setContextOp(ctx, bcq.ctx, "IDs")
	if err = bcq.Select(billingcontrol.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BillingControlQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BillingControlQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bcq.ctx, "Count")
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bcq, querierCount[*BillingControlQuery](), bcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BillingControlQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BillingControlQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bcq.ctx, "Exist")
	switch _, err := bcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BillingControlQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingControlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BillingControlQuery) Clone() *BillingControlQuery {
	if bcq == nil {
		return nil
	}
	return &BillingControlQuery{
		config:           bcq.config,
		ctx:              bcq.ctx.Clone(),
		order:            append([]billingcontrol.OrderOption{}, bcq.order...),
		inters:           append([]Interceptor{}, bcq.inters...),
		predicates:       append([]predicate.BillingControl{}, bcq.predicates...),
		withOrganization: bcq.withOrganization.Clone(),
		withBusinessUnit: bcq.withBusinessUnit.Clone(),
		// clone intermediate query.
		sql:  bcq.sql.Clone(),
		path: bcq.path,
	}
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BillingControlQuery) WithOrganization(opts ...func(*OrganizationQuery)) *BillingControlQuery {
	query := (&OrganizationClient{config: bcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bcq.withOrganization = query
	return bcq
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (bcq *BillingControlQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *BillingControlQuery {
	query := (&BusinessUnitClient{config: bcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bcq.withBusinessUnit = query
	return bcq
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
//	client.BillingControl.Query().
//		GroupBy(billingcontrol.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bcq *BillingControlQuery) GroupBy(field string, fields ...string) *BillingControlGroupBy {
	bcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingControlGroupBy{build: bcq}
	grbuild.flds = &bcq.ctx.Fields
	grbuild.label = billingcontrol.Label
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
//	client.BillingControl.Query().
//		Select(billingcontrol.FieldCreatedAt).
//		Scan(ctx, &v)
func (bcq *BillingControlQuery) Select(fields ...string) *BillingControlSelect {
	bcq.ctx.Fields = append(bcq.ctx.Fields, fields...)
	sbuild := &BillingControlSelect{BillingControlQuery: bcq}
	sbuild.label = billingcontrol.Label
	sbuild.flds, sbuild.scan = &bcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingControlSelect configured with the given aggregations.
func (bcq *BillingControlQuery) Aggregate(fns ...AggregateFunc) *BillingControlSelect {
	return bcq.Select().Aggregate(fns...)
}

func (bcq *BillingControlQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bcq); err != nil {
				return err
			}
		}
	}
	for _, f := range bcq.ctx.Fields {
		if !billingcontrol.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BillingControlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingControl, error) {
	var (
		nodes       = []*BillingControl{}
		withFKs     = bcq.withFKs
		_spec       = bcq.querySpec()
		loadedTypes = [2]bool{
			bcq.withOrganization != nil,
			bcq.withBusinessUnit != nil,
		}
	)
	if bcq.withOrganization != nil || bcq.withBusinessUnit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, billingcontrol.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingControl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingControl{config: bcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bcq.modifiers) > 0 {
		_spec.Modifiers = bcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bcq.withOrganization; query != nil {
		if err := bcq.loadOrganization(ctx, query, nodes, nil,
			func(n *BillingControl, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	if query := bcq.withBusinessUnit; query != nil {
		if err := bcq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *BillingControl, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bcq *BillingControlQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*BillingControl, init func(*BillingControl), assign func(*BillingControl, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*BillingControl)
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
func (bcq *BillingControlQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*BillingControl, init func(*BillingControl), assign func(*BillingControl, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*BillingControl)
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

func (bcq *BillingControlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	if len(bcq.modifiers) > 0 {
		_spec.Modifiers = bcq.modifiers
	}
	_spec.Node.Columns = bcq.ctx.Fields
	if len(bcq.ctx.Fields) > 0 {
		_spec.Unique = bcq.ctx.Unique != nil && *bcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BillingControlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billingcontrol.Table, billingcontrol.Columns, sqlgraph.NewFieldSpec(billingcontrol.FieldID, field.TypeUUID))
	_spec.From = bcq.sql
	if unique := bcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bcq.path != nil {
		_spec.Unique = true
	}
	if fields := bcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billingcontrol.FieldID)
		for i := range fields {
			if fields[i] != billingcontrol.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BillingControlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(billingcontrol.Table)
	columns := bcq.ctx.Fields
	if len(columns) == 0 {
		columns = billingcontrol.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcq.ctx.Unique != nil && *bcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range bcq.modifiers {
		m(selector)
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (bcq *BillingControlQuery) Modify(modifiers ...func(s *sql.Selector)) *BillingControlSelect {
	bcq.modifiers = append(bcq.modifiers, modifiers...)
	return bcq.Select()
}

// BillingControlGroupBy is the group-by builder for BillingControl entities.
type BillingControlGroupBy struct {
	selector
	build *BillingControlQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BillingControlGroupBy) Aggregate(fns ...AggregateFunc) *BillingControlGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the selector query and scans the result into the given value.
func (bcgb *BillingControlGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bcgb.build.ctx, "GroupBy")
	if err := bcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingControlQuery, *BillingControlGroupBy](ctx, bcgb.build, bcgb, bcgb.build.inters, v)
}

func (bcgb *BillingControlGroupBy) sqlScan(ctx context.Context, root *BillingControlQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bcgb.flds)+len(bcgb.fns))
		for _, f := range *bcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingControlSelect is the builder for selecting fields of BillingControl entities.
type BillingControlSelect struct {
	*BillingControlQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bcs *BillingControlSelect) Aggregate(fns ...AggregateFunc) *BillingControlSelect {
	bcs.fns = append(bcs.fns, fns...)
	return bcs
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BillingControlSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bcs.ctx, "Select")
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingControlQuery, *BillingControlSelect](ctx, bcs.BillingControlQuery, bcs, bcs.inters, v)
}

func (bcs *BillingControlSelect) sqlScan(ctx context.Context, root *BillingControlQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bcs.fns))
	for _, fn := range bcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (bcs *BillingControlSelect) Modify(modifiers ...func(s *sql.Selector)) *BillingControlSelect {
	bcs.modifiers = append(bcs.modifiers, modifiers...)
	return bcs
}
