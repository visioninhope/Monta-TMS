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
	"github.com/emoss08/trenova/internal/ent/tablechangealert"
	"github.com/google/uuid"
)

// TableChangeAlertQuery is the builder for querying TableChangeAlert entities.
type TableChangeAlertQuery struct {
	config
	ctx              *QueryContext
	order            []tablechangealert.OrderOption
	inters           []Interceptor
	predicates       []predicate.TableChangeAlert
	withBusinessUnit *BusinessUnitQuery
	withOrganization *OrganizationQuery
	modifiers        []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TableChangeAlertQuery builder.
func (tcaq *TableChangeAlertQuery) Where(ps ...predicate.TableChangeAlert) *TableChangeAlertQuery {
	tcaq.predicates = append(tcaq.predicates, ps...)
	return tcaq
}

// Limit the number of records to be returned by this query.
func (tcaq *TableChangeAlertQuery) Limit(limit int) *TableChangeAlertQuery {
	tcaq.ctx.Limit = &limit
	return tcaq
}

// Offset to start from.
func (tcaq *TableChangeAlertQuery) Offset(offset int) *TableChangeAlertQuery {
	tcaq.ctx.Offset = &offset
	return tcaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcaq *TableChangeAlertQuery) Unique(unique bool) *TableChangeAlertQuery {
	tcaq.ctx.Unique = &unique
	return tcaq
}

// Order specifies how the records should be ordered.
func (tcaq *TableChangeAlertQuery) Order(o ...tablechangealert.OrderOption) *TableChangeAlertQuery {
	tcaq.order = append(tcaq.order, o...)
	return tcaq
}

// QueryBusinessUnit chains the current query on the "business_unit" edge.
func (tcaq *TableChangeAlertQuery) QueryBusinessUnit() *BusinessUnitQuery {
	query := (&BusinessUnitClient{config: tcaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tablechangealert.Table, tablechangealert.FieldID, selector),
			sqlgraph.To(businessunit.Table, businessunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tablechangealert.BusinessUnitTable, tablechangealert.BusinessUnitColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrganization chains the current query on the "organization" edge.
func (tcaq *TableChangeAlertQuery) QueryOrganization() *OrganizationQuery {
	query := (&OrganizationClient{config: tcaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tablechangealert.Table, tablechangealert.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, tablechangealert.OrganizationTable, tablechangealert.OrganizationColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TableChangeAlert entity from the query.
// Returns a *NotFoundError when no TableChangeAlert was found.
func (tcaq *TableChangeAlertQuery) First(ctx context.Context) (*TableChangeAlert, error) {
	nodes, err := tcaq.Limit(1).All(setContextOp(ctx, tcaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tablechangealert.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) FirstX(ctx context.Context) *TableChangeAlert {
	node, err := tcaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TableChangeAlert ID from the query.
// Returns a *NotFoundError when no TableChangeAlert ID was found.
func (tcaq *TableChangeAlertQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tcaq.Limit(1).IDs(setContextOp(ctx, tcaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tablechangealert.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tcaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TableChangeAlert entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TableChangeAlert entity is found.
// Returns a *NotFoundError when no TableChangeAlert entities are found.
func (tcaq *TableChangeAlertQuery) Only(ctx context.Context) (*TableChangeAlert, error) {
	nodes, err := tcaq.Limit(2).All(setContextOp(ctx, tcaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tablechangealert.Label}
	default:
		return nil, &NotSingularError{tablechangealert.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) OnlyX(ctx context.Context) *TableChangeAlert {
	node, err := tcaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TableChangeAlert ID in the query.
// Returns a *NotSingularError when more than one TableChangeAlert ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcaq *TableChangeAlertQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tcaq.Limit(2).IDs(setContextOp(ctx, tcaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tablechangealert.Label}
	default:
		err = &NotSingularError{tablechangealert.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tcaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TableChangeAlerts.
func (tcaq *TableChangeAlertQuery) All(ctx context.Context) ([]*TableChangeAlert, error) {
	ctx = setContextOp(ctx, tcaq.ctx, "All")
	if err := tcaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TableChangeAlert, *TableChangeAlertQuery]()
	return withInterceptors[[]*TableChangeAlert](ctx, tcaq, qr, tcaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) AllX(ctx context.Context) []*TableChangeAlert {
	nodes, err := tcaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TableChangeAlert IDs.
func (tcaq *TableChangeAlertQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tcaq.ctx.Unique == nil && tcaq.path != nil {
		tcaq.Unique(true)
	}
	ctx = setContextOp(ctx, tcaq.ctx, "IDs")
	if err = tcaq.Select(tablechangealert.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tcaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcaq *TableChangeAlertQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcaq.ctx, "Count")
	if err := tcaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcaq, querierCount[*TableChangeAlertQuery](), tcaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) CountX(ctx context.Context) int {
	count, err := tcaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcaq *TableChangeAlertQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tcaq.ctx, "Exist")
	switch _, err := tcaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tcaq *TableChangeAlertQuery) ExistX(ctx context.Context) bool {
	exist, err := tcaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TableChangeAlertQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcaq *TableChangeAlertQuery) Clone() *TableChangeAlertQuery {
	if tcaq == nil {
		return nil
	}
	return &TableChangeAlertQuery{
		config:           tcaq.config,
		ctx:              tcaq.ctx.Clone(),
		order:            append([]tablechangealert.OrderOption{}, tcaq.order...),
		inters:           append([]Interceptor{}, tcaq.inters...),
		predicates:       append([]predicate.TableChangeAlert{}, tcaq.predicates...),
		withBusinessUnit: tcaq.withBusinessUnit.Clone(),
		withOrganization: tcaq.withOrganization.Clone(),
		// clone intermediate query.
		sql:  tcaq.sql.Clone(),
		path: tcaq.path,
	}
}

// WithBusinessUnit tells the query-builder to eager-load the nodes that are connected to
// the "business_unit" edge. The optional arguments are used to configure the query builder of the edge.
func (tcaq *TableChangeAlertQuery) WithBusinessUnit(opts ...func(*BusinessUnitQuery)) *TableChangeAlertQuery {
	query := (&BusinessUnitClient{config: tcaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcaq.withBusinessUnit = query
	return tcaq
}

// WithOrganization tells the query-builder to eager-load the nodes that are connected to
// the "organization" edge. The optional arguments are used to configure the query builder of the edge.
func (tcaq *TableChangeAlertQuery) WithOrganization(opts ...func(*OrganizationQuery)) *TableChangeAlertQuery {
	query := (&OrganizationClient{config: tcaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcaq.withOrganization = query
	return tcaq
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
//	client.TableChangeAlert.Query().
//		GroupBy(tablechangealert.FieldBusinessUnitID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcaq *TableChangeAlertQuery) GroupBy(field string, fields ...string) *TableChangeAlertGroupBy {
	tcaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TableChangeAlertGroupBy{build: tcaq}
	grbuild.flds = &tcaq.ctx.Fields
	grbuild.label = tablechangealert.Label
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
//	client.TableChangeAlert.Query().
//		Select(tablechangealert.FieldBusinessUnitID).
//		Scan(ctx, &v)
func (tcaq *TableChangeAlertQuery) Select(fields ...string) *TableChangeAlertSelect {
	tcaq.ctx.Fields = append(tcaq.ctx.Fields, fields...)
	sbuild := &TableChangeAlertSelect{TableChangeAlertQuery: tcaq}
	sbuild.label = tablechangealert.Label
	sbuild.flds, sbuild.scan = &tcaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TableChangeAlertSelect configured with the given aggregations.
func (tcaq *TableChangeAlertQuery) Aggregate(fns ...AggregateFunc) *TableChangeAlertSelect {
	return tcaq.Select().Aggregate(fns...)
}

func (tcaq *TableChangeAlertQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tcaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tcaq); err != nil {
				return err
			}
		}
	}
	for _, f := range tcaq.ctx.Fields {
		if !tablechangealert.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcaq.path != nil {
		prev, err := tcaq.path(ctx)
		if err != nil {
			return err
		}
		tcaq.sql = prev
	}
	return nil
}

func (tcaq *TableChangeAlertQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TableChangeAlert, error) {
	var (
		nodes       = []*TableChangeAlert{}
		_spec       = tcaq.querySpec()
		loadedTypes = [2]bool{
			tcaq.withBusinessUnit != nil,
			tcaq.withOrganization != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TableChangeAlert).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TableChangeAlert{config: tcaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tcaq.modifiers) > 0 {
		_spec.Modifiers = tcaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tcaq.withBusinessUnit; query != nil {
		if err := tcaq.loadBusinessUnit(ctx, query, nodes, nil,
			func(n *TableChangeAlert, e *BusinessUnit) { n.Edges.BusinessUnit = e }); err != nil {
			return nil, err
		}
	}
	if query := tcaq.withOrganization; query != nil {
		if err := tcaq.loadOrganization(ctx, query, nodes, nil,
			func(n *TableChangeAlert, e *Organization) { n.Edges.Organization = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tcaq *TableChangeAlertQuery) loadBusinessUnit(ctx context.Context, query *BusinessUnitQuery, nodes []*TableChangeAlert, init func(*TableChangeAlert), assign func(*TableChangeAlert, *BusinessUnit)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TableChangeAlert)
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
func (tcaq *TableChangeAlertQuery) loadOrganization(ctx context.Context, query *OrganizationQuery, nodes []*TableChangeAlert, init func(*TableChangeAlert), assign func(*TableChangeAlert, *Organization)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*TableChangeAlert)
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

func (tcaq *TableChangeAlertQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcaq.querySpec()
	if len(tcaq.modifiers) > 0 {
		_spec.Modifiers = tcaq.modifiers
	}
	_spec.Node.Columns = tcaq.ctx.Fields
	if len(tcaq.ctx.Fields) > 0 {
		_spec.Unique = tcaq.ctx.Unique != nil && *tcaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tcaq.driver, _spec)
}

func (tcaq *TableChangeAlertQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tablechangealert.Table, tablechangealert.Columns, sqlgraph.NewFieldSpec(tablechangealert.FieldID, field.TypeUUID))
	_spec.From = tcaq.sql
	if unique := tcaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcaq.path != nil {
		_spec.Unique = true
	}
	if fields := tcaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tablechangealert.FieldID)
		for i := range fields {
			if fields[i] != tablechangealert.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tcaq.withBusinessUnit != nil {
			_spec.Node.AddColumnOnce(tablechangealert.FieldBusinessUnitID)
		}
		if tcaq.withOrganization != nil {
			_spec.Node.AddColumnOnce(tablechangealert.FieldOrganizationID)
		}
	}
	if ps := tcaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcaq *TableChangeAlertQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcaq.driver.Dialect())
	t1 := builder.Table(tablechangealert.Table)
	columns := tcaq.ctx.Fields
	if len(columns) == 0 {
		columns = tablechangealert.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcaq.sql != nil {
		selector = tcaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcaq.ctx.Unique != nil && *tcaq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range tcaq.modifiers {
		m(selector)
	}
	for _, p := range tcaq.predicates {
		p(selector)
	}
	for _, p := range tcaq.order {
		p(selector)
	}
	if offset := tcaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tcaq *TableChangeAlertQuery) Modify(modifiers ...func(s *sql.Selector)) *TableChangeAlertSelect {
	tcaq.modifiers = append(tcaq.modifiers, modifiers...)
	return tcaq.Select()
}

// TableChangeAlertGroupBy is the group-by builder for TableChangeAlert entities.
type TableChangeAlertGroupBy struct {
	selector
	build *TableChangeAlertQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcagb *TableChangeAlertGroupBy) Aggregate(fns ...AggregateFunc) *TableChangeAlertGroupBy {
	tcagb.fns = append(tcagb.fns, fns...)
	return tcagb
}

// Scan applies the selector query and scans the result into the given value.
func (tcagb *TableChangeAlertGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcagb.build.ctx, "GroupBy")
	if err := tcagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TableChangeAlertQuery, *TableChangeAlertGroupBy](ctx, tcagb.build, tcagb, tcagb.build.inters, v)
}

func (tcagb *TableChangeAlertGroupBy) sqlScan(ctx context.Context, root *TableChangeAlertQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tcagb.fns))
	for _, fn := range tcagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tcagb.flds)+len(tcagb.fns))
		for _, f := range *tcagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tcagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TableChangeAlertSelect is the builder for selecting fields of TableChangeAlert entities.
type TableChangeAlertSelect struct {
	*TableChangeAlertQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcas *TableChangeAlertSelect) Aggregate(fns ...AggregateFunc) *TableChangeAlertSelect {
	tcas.fns = append(tcas.fns, fns...)
	return tcas
}

// Scan applies the selector query and scans the result into the given value.
func (tcas *TableChangeAlertSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcas.ctx, "Select")
	if err := tcas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TableChangeAlertQuery, *TableChangeAlertSelect](ctx, tcas.TableChangeAlertQuery, tcas, tcas.inters, v)
}

func (tcas *TableChangeAlertSelect) sqlScan(ctx context.Context, root *TableChangeAlertQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tcas.fns))
	for _, fn := range tcas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tcas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tcas *TableChangeAlertSelect) Modify(modifiers ...func(s *sql.Selector)) *TableChangeAlertSelect {
	tcas.modifiers = append(tcas.modifiers, modifiers...)
	return tcas
}
