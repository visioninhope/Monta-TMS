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
	"github.com/emoss08/trenova/internal/ent/hazardousmaterial"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/google/uuid"
)

// HazardousMaterialCreate is the builder for creating a HazardousMaterial entity.
type HazardousMaterialCreate struct {
	config
	mutation *HazardousMaterialMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (hmc *HazardousMaterialCreate) SetBusinessUnitID(u uuid.UUID) *HazardousMaterialCreate {
	hmc.mutation.SetBusinessUnitID(u)
	return hmc
}

// SetOrganizationID sets the "organization_id" field.
func (hmc *HazardousMaterialCreate) SetOrganizationID(u uuid.UUID) *HazardousMaterialCreate {
	hmc.mutation.SetOrganizationID(u)
	return hmc
}

// SetCreatedAt sets the "created_at" field.
func (hmc *HazardousMaterialCreate) SetCreatedAt(t time.Time) *HazardousMaterialCreate {
	hmc.mutation.SetCreatedAt(t)
	return hmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableCreatedAt(t *time.Time) *HazardousMaterialCreate {
	if t != nil {
		hmc.SetCreatedAt(*t)
	}
	return hmc
}

// SetUpdatedAt sets the "updated_at" field.
func (hmc *HazardousMaterialCreate) SetUpdatedAt(t time.Time) *HazardousMaterialCreate {
	hmc.mutation.SetUpdatedAt(t)
	return hmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableUpdatedAt(t *time.Time) *HazardousMaterialCreate {
	if t != nil {
		hmc.SetUpdatedAt(*t)
	}
	return hmc
}

// SetVersion sets the "version" field.
func (hmc *HazardousMaterialCreate) SetVersion(i int) *HazardousMaterialCreate {
	hmc.mutation.SetVersion(i)
	return hmc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableVersion(i *int) *HazardousMaterialCreate {
	if i != nil {
		hmc.SetVersion(*i)
	}
	return hmc
}

// SetStatus sets the "status" field.
func (hmc *HazardousMaterialCreate) SetStatus(h hazardousmaterial.Status) *HazardousMaterialCreate {
	hmc.mutation.SetStatus(h)
	return hmc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableStatus(h *hazardousmaterial.Status) *HazardousMaterialCreate {
	if h != nil {
		hmc.SetStatus(*h)
	}
	return hmc
}

// SetName sets the "name" field.
func (hmc *HazardousMaterialCreate) SetName(s string) *HazardousMaterialCreate {
	hmc.mutation.SetName(s)
	return hmc
}

// SetHazardClass sets the "hazard_class" field.
func (hmc *HazardousMaterialCreate) SetHazardClass(hc hazardousmaterial.HazardClass) *HazardousMaterialCreate {
	hmc.mutation.SetHazardClass(hc)
	return hmc
}

// SetNillableHazardClass sets the "hazard_class" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableHazardClass(hc *hazardousmaterial.HazardClass) *HazardousMaterialCreate {
	if hc != nil {
		hmc.SetHazardClass(*hc)
	}
	return hmc
}

// SetErgNumber sets the "erg_number" field.
func (hmc *HazardousMaterialCreate) SetErgNumber(s string) *HazardousMaterialCreate {
	hmc.mutation.SetErgNumber(s)
	return hmc
}

// SetNillableErgNumber sets the "erg_number" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableErgNumber(s *string) *HazardousMaterialCreate {
	if s != nil {
		hmc.SetErgNumber(*s)
	}
	return hmc
}

// SetDescription sets the "description" field.
func (hmc *HazardousMaterialCreate) SetDescription(s string) *HazardousMaterialCreate {
	hmc.mutation.SetDescription(s)
	return hmc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableDescription(s *string) *HazardousMaterialCreate {
	if s != nil {
		hmc.SetDescription(*s)
	}
	return hmc
}

// SetPackingGroup sets the "packing_group" field.
func (hmc *HazardousMaterialCreate) SetPackingGroup(s string) *HazardousMaterialCreate {
	hmc.mutation.SetPackingGroup(s)
	return hmc
}

// SetNillablePackingGroup sets the "packing_group" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillablePackingGroup(s *string) *HazardousMaterialCreate {
	if s != nil {
		hmc.SetPackingGroup(*s)
	}
	return hmc
}

// SetProperShippingName sets the "proper_shipping_name" field.
func (hmc *HazardousMaterialCreate) SetProperShippingName(s string) *HazardousMaterialCreate {
	hmc.mutation.SetProperShippingName(s)
	return hmc
}

// SetNillableProperShippingName sets the "proper_shipping_name" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableProperShippingName(s *string) *HazardousMaterialCreate {
	if s != nil {
		hmc.SetProperShippingName(*s)
	}
	return hmc
}

// SetID sets the "id" field.
func (hmc *HazardousMaterialCreate) SetID(u uuid.UUID) *HazardousMaterialCreate {
	hmc.mutation.SetID(u)
	return hmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hmc *HazardousMaterialCreate) SetNillableID(u *uuid.UUID) *HazardousMaterialCreate {
	if u != nil {
		hmc.SetID(*u)
	}
	return hmc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (hmc *HazardousMaterialCreate) SetBusinessUnit(b *BusinessUnit) *HazardousMaterialCreate {
	return hmc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (hmc *HazardousMaterialCreate) SetOrganization(o *Organization) *HazardousMaterialCreate {
	return hmc.SetOrganizationID(o.ID)
}

// Mutation returns the HazardousMaterialMutation object of the builder.
func (hmc *HazardousMaterialCreate) Mutation() *HazardousMaterialMutation {
	return hmc.mutation
}

// Save creates the HazardousMaterial in the database.
func (hmc *HazardousMaterialCreate) Save(ctx context.Context) (*HazardousMaterial, error) {
	hmc.defaults()
	return withHooks(ctx, hmc.sqlSave, hmc.mutation, hmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hmc *HazardousMaterialCreate) SaveX(ctx context.Context) *HazardousMaterial {
	v, err := hmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hmc *HazardousMaterialCreate) Exec(ctx context.Context) error {
	_, err := hmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hmc *HazardousMaterialCreate) ExecX(ctx context.Context) {
	if err := hmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hmc *HazardousMaterialCreate) defaults() {
	if _, ok := hmc.mutation.CreatedAt(); !ok {
		v := hazardousmaterial.DefaultCreatedAt()
		hmc.mutation.SetCreatedAt(v)
	}
	if _, ok := hmc.mutation.UpdatedAt(); !ok {
		v := hazardousmaterial.DefaultUpdatedAt()
		hmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := hmc.mutation.Version(); !ok {
		v := hazardousmaterial.DefaultVersion
		hmc.mutation.SetVersion(v)
	}
	if _, ok := hmc.mutation.Status(); !ok {
		v := hazardousmaterial.DefaultStatus
		hmc.mutation.SetStatus(v)
	}
	if _, ok := hmc.mutation.HazardClass(); !ok {
		v := hazardousmaterial.DefaultHazardClass
		hmc.mutation.SetHazardClass(v)
	}
	if _, ok := hmc.mutation.ID(); !ok {
		v := hazardousmaterial.DefaultID()
		hmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hmc *HazardousMaterialCreate) check() error {
	if _, ok := hmc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "HazardousMaterial.business_unit_id"`)}
	}
	if _, ok := hmc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "HazardousMaterial.organization_id"`)}
	}
	if _, ok := hmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "HazardousMaterial.created_at"`)}
	}
	if _, ok := hmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "HazardousMaterial.updated_at"`)}
	}
	if _, ok := hmc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "HazardousMaterial.version"`)}
	}
	if _, ok := hmc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "HazardousMaterial.status"`)}
	}
	if v, ok := hmc.mutation.Status(); ok {
		if err := hazardousmaterial.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "HazardousMaterial.status": %w`, err)}
		}
	}
	if _, ok := hmc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "HazardousMaterial.name"`)}
	}
	if v, ok := hmc.mutation.Name(); ok {
		if err := hazardousmaterial.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "HazardousMaterial.name": %w`, err)}
		}
	}
	if _, ok := hmc.mutation.HazardClass(); !ok {
		return &ValidationError{Name: "hazard_class", err: errors.New(`ent: missing required field "HazardousMaterial.hazard_class"`)}
	}
	if v, ok := hmc.mutation.HazardClass(); ok {
		if err := hazardousmaterial.HazardClassValidator(v); err != nil {
			return &ValidationError{Name: "hazard_class", err: fmt.Errorf(`ent: validator failed for field "HazardousMaterial.hazard_class": %w`, err)}
		}
	}
	if _, ok := hmc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "HazardousMaterial.business_unit"`)}
	}
	if _, ok := hmc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "HazardousMaterial.organization"`)}
	}
	return nil
}

func (hmc *HazardousMaterialCreate) sqlSave(ctx context.Context) (*HazardousMaterial, error) {
	if err := hmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hmc.driver, _spec); err != nil {
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
	hmc.mutation.id = &_node.ID
	hmc.mutation.done = true
	return _node, nil
}

func (hmc *HazardousMaterialCreate) createSpec() (*HazardousMaterial, *sqlgraph.CreateSpec) {
	var (
		_node = &HazardousMaterial{config: hmc.config}
		_spec = sqlgraph.NewCreateSpec(hazardousmaterial.Table, sqlgraph.NewFieldSpec(hazardousmaterial.FieldID, field.TypeUUID))
	)
	if id, ok := hmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hmc.mutation.CreatedAt(); ok {
		_spec.SetField(hazardousmaterial.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hmc.mutation.UpdatedAt(); ok {
		_spec.SetField(hazardousmaterial.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := hmc.mutation.Version(); ok {
		_spec.SetField(hazardousmaterial.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := hmc.mutation.Status(); ok {
		_spec.SetField(hazardousmaterial.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := hmc.mutation.Name(); ok {
		_spec.SetField(hazardousmaterial.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := hmc.mutation.HazardClass(); ok {
		_spec.SetField(hazardousmaterial.FieldHazardClass, field.TypeEnum, value)
		_node.HazardClass = value
	}
	if value, ok := hmc.mutation.ErgNumber(); ok {
		_spec.SetField(hazardousmaterial.FieldErgNumber, field.TypeString, value)
		_node.ErgNumber = value
	}
	if value, ok := hmc.mutation.Description(); ok {
		_spec.SetField(hazardousmaterial.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := hmc.mutation.PackingGroup(); ok {
		_spec.SetField(hazardousmaterial.FieldPackingGroup, field.TypeString, value)
		_node.PackingGroup = value
	}
	if value, ok := hmc.mutation.ProperShippingName(); ok {
		_spec.SetField(hazardousmaterial.FieldProperShippingName, field.TypeString, value)
		_node.ProperShippingName = value
	}
	if nodes := hmc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hazardousmaterial.BusinessUnitTable,
			Columns: []string{hazardousmaterial.BusinessUnitColumn},
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
	if nodes := hmc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hazardousmaterial.OrganizationTable,
			Columns: []string{hazardousmaterial.OrganizationColumn},
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
	return _node, _spec
}

// HazardousMaterialCreateBulk is the builder for creating many HazardousMaterial entities in bulk.
type HazardousMaterialCreateBulk struct {
	config
	err      error
	builders []*HazardousMaterialCreate
}

// Save creates the HazardousMaterial entities in the database.
func (hmcb *HazardousMaterialCreateBulk) Save(ctx context.Context) ([]*HazardousMaterial, error) {
	if hmcb.err != nil {
		return nil, hmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hmcb.builders))
	nodes := make([]*HazardousMaterial, len(hmcb.builders))
	mutators := make([]Mutator, len(hmcb.builders))
	for i := range hmcb.builders {
		func(i int, root context.Context) {
			builder := hmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HazardousMaterialMutation)
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
					_, err = mutators[i+1].Mutate(root, hmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hmcb *HazardousMaterialCreateBulk) SaveX(ctx context.Context) []*HazardousMaterial {
	v, err := hmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hmcb *HazardousMaterialCreateBulk) Exec(ctx context.Context) error {
	_, err := hmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hmcb *HazardousMaterialCreateBulk) ExecX(ctx context.Context) {
	if err := hmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
