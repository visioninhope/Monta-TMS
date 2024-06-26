// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/generalledgeraccount"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/revenuecode"
	"github.com/google/uuid"
)

// RevenueCodeCreate is the builder for creating a RevenueCode entity.
type RevenueCodeCreate struct {
	config
	mutation *RevenueCodeMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (rcc *RevenueCodeCreate) SetBusinessUnitID(u uuid.UUID) *RevenueCodeCreate {
	rcc.mutation.SetBusinessUnitID(u)
	return rcc
}

// SetOrganizationID sets the "organization_id" field.
func (rcc *RevenueCodeCreate) SetOrganizationID(u uuid.UUID) *RevenueCodeCreate {
	rcc.mutation.SetOrganizationID(u)
	return rcc
}

// SetCreatedAt sets the "created_at" field.
func (rcc *RevenueCodeCreate) SetCreatedAt(t time.Time) *RevenueCodeCreate {
	rcc.mutation.SetCreatedAt(t)
	return rcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableCreatedAt(t *time.Time) *RevenueCodeCreate {
	if t != nil {
		rcc.SetCreatedAt(*t)
	}
	return rcc
}

// SetUpdatedAt sets the "updated_at" field.
func (rcc *RevenueCodeCreate) SetUpdatedAt(t time.Time) *RevenueCodeCreate {
	rcc.mutation.SetUpdatedAt(t)
	return rcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableUpdatedAt(t *time.Time) *RevenueCodeCreate {
	if t != nil {
		rcc.SetUpdatedAt(*t)
	}
	return rcc
}

// SetVersion sets the "version" field.
func (rcc *RevenueCodeCreate) SetVersion(i int) *RevenueCodeCreate {
	rcc.mutation.SetVersion(i)
	return rcc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableVersion(i *int) *RevenueCodeCreate {
	if i != nil {
		rcc.SetVersion(*i)
	}
	return rcc
}

// SetStatus sets the "status" field.
func (rcc *RevenueCodeCreate) SetStatus(r revenuecode.Status) *RevenueCodeCreate {
	rcc.mutation.SetStatus(r)
	return rcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableStatus(r *revenuecode.Status) *RevenueCodeCreate {
	if r != nil {
		rcc.SetStatus(*r)
	}
	return rcc
}

// SetCode sets the "code" field.
func (rcc *RevenueCodeCreate) SetCode(s string) *RevenueCodeCreate {
	rcc.mutation.SetCode(s)
	return rcc
}

// SetDescription sets the "description" field.
func (rcc *RevenueCodeCreate) SetDescription(s string) *RevenueCodeCreate {
	rcc.mutation.SetDescription(s)
	return rcc
}

// SetExpenseAccountID sets the "expense_account_id" field.
func (rcc *RevenueCodeCreate) SetExpenseAccountID(u uuid.UUID) *RevenueCodeCreate {
	rcc.mutation.SetExpenseAccountID(u)
	return rcc
}

// SetNillableExpenseAccountID sets the "expense_account_id" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableExpenseAccountID(u *uuid.UUID) *RevenueCodeCreate {
	if u != nil {
		rcc.SetExpenseAccountID(*u)
	}
	return rcc
}

// SetRevenueAccountID sets the "revenue_account_id" field.
func (rcc *RevenueCodeCreate) SetRevenueAccountID(u uuid.UUID) *RevenueCodeCreate {
	rcc.mutation.SetRevenueAccountID(u)
	return rcc
}

// SetNillableRevenueAccountID sets the "revenue_account_id" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableRevenueAccountID(u *uuid.UUID) *RevenueCodeCreate {
	if u != nil {
		rcc.SetRevenueAccountID(*u)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *RevenueCodeCreate) SetID(u uuid.UUID) *RevenueCodeCreate {
	rcc.mutation.SetID(u)
	return rcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rcc *RevenueCodeCreate) SetNillableID(u *uuid.UUID) *RevenueCodeCreate {
	if u != nil {
		rcc.SetID(*u)
	}
	return rcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (rcc *RevenueCodeCreate) SetBusinessUnit(b *BusinessUnit) *RevenueCodeCreate {
	return rcc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (rcc *RevenueCodeCreate) SetOrganization(o *Organization) *RevenueCodeCreate {
	return rcc.SetOrganizationID(o.ID)
}

// SetExpenseAccount sets the "expense_account" edge to the GeneralLedgerAccount entity.
func (rcc *RevenueCodeCreate) SetExpenseAccount(g *GeneralLedgerAccount) *RevenueCodeCreate {
	return rcc.SetExpenseAccountID(g.ID)
}

// SetRevenueAccount sets the "revenue_account" edge to the GeneralLedgerAccount entity.
func (rcc *RevenueCodeCreate) SetRevenueAccount(g *GeneralLedgerAccount) *RevenueCodeCreate {
	return rcc.SetRevenueAccountID(g.ID)
}

// Mutation returns the RevenueCodeMutation object of the builder.
func (rcc *RevenueCodeCreate) Mutation() *RevenueCodeMutation {
	return rcc.mutation
}

// Save creates the RevenueCode in the database.
func (rcc *RevenueCodeCreate) Save(ctx context.Context) (*RevenueCode, error) {
	if err := rcc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rcc.sqlSave, rcc.mutation, rcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *RevenueCodeCreate) SaveX(ctx context.Context) *RevenueCode {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *RevenueCodeCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *RevenueCodeCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcc *RevenueCodeCreate) defaults() error {
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		if revenuecode.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized revenuecode.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := revenuecode.DefaultCreatedAt()
		rcc.mutation.SetCreatedAt(v)
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		if revenuecode.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized revenuecode.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := revenuecode.DefaultUpdatedAt()
		rcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rcc.mutation.Version(); !ok {
		v := revenuecode.DefaultVersion
		rcc.mutation.SetVersion(v)
	}
	if _, ok := rcc.mutation.Status(); !ok {
		v := revenuecode.DefaultStatus
		rcc.mutation.SetStatus(v)
	}
	if _, ok := rcc.mutation.ID(); !ok {
		if revenuecode.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized revenuecode.DefaultID (forgotten import ent/runtime?)")
		}
		v := revenuecode.DefaultID()
		rcc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rcc *RevenueCodeCreate) check() error {
	if _, ok := rcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "RevenueCode.business_unit_id"`)}
	}
	if _, ok := rcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "RevenueCode.organization_id"`)}
	}
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RevenueCode.created_at"`)}
	}
	if _, ok := rcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RevenueCode.updated_at"`)}
	}
	if _, ok := rcc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "RevenueCode.version"`)}
	}
	if _, ok := rcc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "RevenueCode.status"`)}
	}
	if v, ok := rcc.mutation.Status(); ok {
		if err := revenuecode.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "RevenueCode.status": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "RevenueCode.code"`)}
	}
	if v, ok := rcc.mutation.Code(); ok {
		if err := revenuecode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "RevenueCode.code": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "RevenueCode.description"`)}
	}
	if v, ok := rcc.mutation.Description(); ok {
		if err := revenuecode.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "RevenueCode.description": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "RevenueCode.business_unit"`)}
	}
	if _, ok := rcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "RevenueCode.organization"`)}
	}
	return nil
}

func (rcc *RevenueCodeCreate) sqlSave(ctx context.Context) (*RevenueCode, error) {
	if err := rcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rcc.mutation.id = &_node.ID
	rcc.mutation.done = true
	return _node, nil
}

func (rcc *RevenueCodeCreate) createSpec() (*RevenueCode, *sqlgraph.CreateSpec) {
	var (
		_node = &RevenueCode{config: rcc.config}
		_spec = sqlgraph.NewCreateSpec(revenuecode.Table, sqlgraph.NewFieldSpec(revenuecode.FieldID, field.TypeUUID))
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rcc.mutation.CreatedAt(); ok {
		_spec.SetField(revenuecode.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rcc.mutation.UpdatedAt(); ok {
		_spec.SetField(revenuecode.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rcc.mutation.Version(); ok {
		_spec.SetField(revenuecode.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := rcc.mutation.Status(); ok {
		_spec.SetField(revenuecode.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rcc.mutation.Code(); ok {
		_spec.SetField(revenuecode.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := rcc.mutation.Description(); ok {
		_spec.SetField(revenuecode.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := rcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   revenuecode.BusinessUnitTable,
			Columns: []string{revenuecode.BusinessUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(businessunit.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BusinessUnitID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   revenuecode.OrganizationTable,
			Columns: []string{revenuecode.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.ExpenseAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   revenuecode.ExpenseAccountTable,
			Columns: []string{revenuecode.ExpenseAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExpenseAccountID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.RevenueAccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   revenuecode.RevenueAccountTable,
			Columns: []string{revenuecode.RevenueAccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(generalledgeraccount.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RevenueAccountID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RevenueCodeCreateBulk is the builder for creating many RevenueCode entities in bulk.
type RevenueCodeCreateBulk struct {
	config
	err      error
	builders []*RevenueCodeCreate
}

// Save creates the RevenueCode entities in the database.
func (rccb *RevenueCodeCreateBulk) Save(ctx context.Context) ([]*RevenueCode, error) {
	if rccb.err != nil {
		return nil, rccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*RevenueCode, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RevenueCodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *RevenueCodeCreateBulk) SaveX(ctx context.Context) []*RevenueCode {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *RevenueCodeCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *RevenueCodeCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}
