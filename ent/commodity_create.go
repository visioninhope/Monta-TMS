// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/commodity"
	"github.com/emoss08/trenova/ent/hazardousmaterial"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// CommodityCreate is the builder for creating a Commodity entity.
type CommodityCreate struct {
	config
	mutation *CommodityMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (cc *CommodityCreate) SetBusinessUnitID(u uuid.UUID) *CommodityCreate {
	cc.mutation.SetBusinessUnitID(u)
	return cc
}

// SetOrganizationID sets the "organization_id" field.
func (cc *CommodityCreate) SetOrganizationID(u uuid.UUID) *CommodityCreate {
	cc.mutation.SetOrganizationID(u)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommodityCreate) SetCreatedAt(t time.Time) *CommodityCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableCreatedAt(t *time.Time) *CommodityCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommodityCreate) SetUpdatedAt(t time.Time) *CommodityCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableUpdatedAt(t *time.Time) *CommodityCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CommodityCreate) SetStatus(c commodity.Status) *CommodityCreate {
	cc.mutation.SetStatus(c)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableStatus(c *commodity.Status) *CommodityCreate {
	if c != nil {
		cc.SetStatus(*c)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CommodityCreate) SetName(s string) *CommodityCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIsHazmat sets the "is_hazmat" field.
func (cc *CommodityCreate) SetIsHazmat(b bool) *CommodityCreate {
	cc.mutation.SetIsHazmat(b)
	return cc
}

// SetNillableIsHazmat sets the "is_hazmat" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableIsHazmat(b *bool) *CommodityCreate {
	if b != nil {
		cc.SetIsHazmat(*b)
	}
	return cc
}

// SetUnitOfMeasure sets the "unit_of_measure" field.
func (cc *CommodityCreate) SetUnitOfMeasure(s string) *CommodityCreate {
	cc.mutation.SetUnitOfMeasure(s)
	return cc
}

// SetNillableUnitOfMeasure sets the "unit_of_measure" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableUnitOfMeasure(s *string) *CommodityCreate {
	if s != nil {
		cc.SetUnitOfMeasure(*s)
	}
	return cc
}

// SetMinTemp sets the "min_temp" field.
func (cc *CommodityCreate) SetMinTemp(i int8) *CommodityCreate {
	cc.mutation.SetMinTemp(i)
	return cc
}

// SetNillableMinTemp sets the "min_temp" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableMinTemp(i *int8) *CommodityCreate {
	if i != nil {
		cc.SetMinTemp(*i)
	}
	return cc
}

// SetMaxTemp sets the "max_temp" field.
func (cc *CommodityCreate) SetMaxTemp(i int8) *CommodityCreate {
	cc.mutation.SetMaxTemp(i)
	return cc
}

// SetNillableMaxTemp sets the "max_temp" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableMaxTemp(i *int8) *CommodityCreate {
	if i != nil {
		cc.SetMaxTemp(*i)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CommodityCreate) SetDescription(s string) *CommodityCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableDescription(s *string) *CommodityCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetHazardousMaterialID sets the "hazardous_material_id" field.
func (cc *CommodityCreate) SetHazardousMaterialID(u uuid.UUID) *CommodityCreate {
	cc.mutation.SetHazardousMaterialID(u)
	return cc
}

// SetNillableHazardousMaterialID sets the "hazardous_material_id" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableHazardousMaterialID(u *uuid.UUID) *CommodityCreate {
	if u != nil {
		cc.SetHazardousMaterialID(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommodityCreate) SetID(u uuid.UUID) *CommodityCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CommodityCreate) SetNillableID(u *uuid.UUID) *CommodityCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (cc *CommodityCreate) SetBusinessUnit(b *BusinessUnit) *CommodityCreate {
	return cc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (cc *CommodityCreate) SetOrganization(o *Organization) *CommodityCreate {
	return cc.SetOrganizationID(o.ID)
}

// SetHazardousMaterial sets the "hazardous_material" edge to the HazardousMaterial entity.
func (cc *CommodityCreate) SetHazardousMaterial(h *HazardousMaterial) *CommodityCreate {
	return cc.SetHazardousMaterialID(h.ID)
}

// Mutation returns the CommodityMutation object of the builder.
func (cc *CommodityCreate) Mutation() *CommodityMutation {
	return cc.mutation
}

// Save creates the Commodity in the database.
func (cc *CommodityCreate) Save(ctx context.Context) (*Commodity, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommodityCreate) SaveX(ctx context.Context) *Commodity {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommodityCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommodityCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommodityCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := commodity.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := commodity.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := commodity.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.IsHazmat(); !ok {
		v := commodity.DefaultIsHazmat
		cc.mutation.SetIsHazmat(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := commodity.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommodityCreate) check() error {
	if _, ok := cc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "Commodity.business_unit_id"`)}
	}
	if _, ok := cc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "Commodity.organization_id"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Commodity.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Commodity.updated_at"`)}
	}
	if _, ok := cc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Commodity.status"`)}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := commodity.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Commodity.status": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Commodity.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := commodity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Commodity.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsHazmat(); !ok {
		return &ValidationError{Name: "is_hazmat", err: errors.New(`ent: missing required field "Commodity.is_hazmat"`)}
	}
	if _, ok := cc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "Commodity.business_unit"`)}
	}
	if _, ok := cc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Commodity.organization"`)}
	}
	return nil
}

func (cc *CommodityCreate) sqlSave(ctx context.Context) (*Commodity, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommodityCreate) createSpec() (*Commodity, *sqlgraph.CreateSpec) {
	var (
		_node = &Commodity{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(commodity.Table, sqlgraph.NewFieldSpec(commodity.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(commodity.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(commodity.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(commodity.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(commodity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.IsHazmat(); ok {
		_spec.SetField(commodity.FieldIsHazmat, field.TypeBool, value)
		_node.IsHazmat = value
	}
	if value, ok := cc.mutation.UnitOfMeasure(); ok {
		_spec.SetField(commodity.FieldUnitOfMeasure, field.TypeString, value)
		_node.UnitOfMeasure = value
	}
	if value, ok := cc.mutation.MinTemp(); ok {
		_spec.SetField(commodity.FieldMinTemp, field.TypeInt8, value)
		_node.MinTemp = value
	}
	if value, ok := cc.mutation.MaxTemp(); ok {
		_spec.SetField(commodity.FieldMaxTemp, field.TypeInt8, value)
		_node.MaxTemp = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(commodity.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := cc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.BusinessUnitTable,
			Columns: []string{commodity.BusinessUnitColumn},
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
	if nodes := cc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.OrganizationTable,
			Columns: []string{commodity.OrganizationColumn},
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
	if nodes := cc.mutation.HazardousMaterialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   commodity.HazardousMaterialTable,
			Columns: []string{commodity.HazardousMaterialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.HazardousMaterialID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommodityCreateBulk is the builder for creating many Commodity entities in bulk.
type CommodityCreateBulk struct {
	config
	err      error
	builders []*CommodityCreate
}

// Save creates the Commodity entities in the database.
func (ccb *CommodityCreateBulk) Save(ctx context.Context) ([]*Commodity, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Commodity, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommodityMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommodityCreateBulk) SaveX(ctx context.Context) []*Commodity {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommodityCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommodityCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
