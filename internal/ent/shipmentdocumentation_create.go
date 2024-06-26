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
	"github.com/emoss08/trenova/internal/ent/documentclassification"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/shipment"
	"github.com/emoss08/trenova/internal/ent/shipmentdocumentation"
	"github.com/google/uuid"
)

// ShipmentDocumentationCreate is the builder for creating a ShipmentDocumentation entity.
type ShipmentDocumentationCreate struct {
	config
	mutation *ShipmentDocumentationMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (sdc *ShipmentDocumentationCreate) SetBusinessUnitID(u uuid.UUID) *ShipmentDocumentationCreate {
	sdc.mutation.SetBusinessUnitID(u)
	return sdc
}

// SetOrganizationID sets the "organization_id" field.
func (sdc *ShipmentDocumentationCreate) SetOrganizationID(u uuid.UUID) *ShipmentDocumentationCreate {
	sdc.mutation.SetOrganizationID(u)
	return sdc
}

// SetCreatedAt sets the "created_at" field.
func (sdc *ShipmentDocumentationCreate) SetCreatedAt(t time.Time) *ShipmentDocumentationCreate {
	sdc.mutation.SetCreatedAt(t)
	return sdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sdc *ShipmentDocumentationCreate) SetNillableCreatedAt(t *time.Time) *ShipmentDocumentationCreate {
	if t != nil {
		sdc.SetCreatedAt(*t)
	}
	return sdc
}

// SetUpdatedAt sets the "updated_at" field.
func (sdc *ShipmentDocumentationCreate) SetUpdatedAt(t time.Time) *ShipmentDocumentationCreate {
	sdc.mutation.SetUpdatedAt(t)
	return sdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sdc *ShipmentDocumentationCreate) SetNillableUpdatedAt(t *time.Time) *ShipmentDocumentationCreate {
	if t != nil {
		sdc.SetUpdatedAt(*t)
	}
	return sdc
}

// SetVersion sets the "version" field.
func (sdc *ShipmentDocumentationCreate) SetVersion(i int) *ShipmentDocumentationCreate {
	sdc.mutation.SetVersion(i)
	return sdc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (sdc *ShipmentDocumentationCreate) SetNillableVersion(i *int) *ShipmentDocumentationCreate {
	if i != nil {
		sdc.SetVersion(*i)
	}
	return sdc
}

// SetShipmentID sets the "shipment_id" field.
func (sdc *ShipmentDocumentationCreate) SetShipmentID(u uuid.UUID) *ShipmentDocumentationCreate {
	sdc.mutation.SetShipmentID(u)
	return sdc
}

// SetDocumentURL sets the "document_url" field.
func (sdc *ShipmentDocumentationCreate) SetDocumentURL(s string) *ShipmentDocumentationCreate {
	sdc.mutation.SetDocumentURL(s)
	return sdc
}

// SetDocumentClassificationID sets the "document_classification_id" field.
func (sdc *ShipmentDocumentationCreate) SetDocumentClassificationID(u uuid.UUID) *ShipmentDocumentationCreate {
	sdc.mutation.SetDocumentClassificationID(u)
	return sdc
}

// SetID sets the "id" field.
func (sdc *ShipmentDocumentationCreate) SetID(u uuid.UUID) *ShipmentDocumentationCreate {
	sdc.mutation.SetID(u)
	return sdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdc *ShipmentDocumentationCreate) SetNillableID(u *uuid.UUID) *ShipmentDocumentationCreate {
	if u != nil {
		sdc.SetID(*u)
	}
	return sdc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (sdc *ShipmentDocumentationCreate) SetBusinessUnit(b *BusinessUnit) *ShipmentDocumentationCreate {
	return sdc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (sdc *ShipmentDocumentationCreate) SetOrganization(o *Organization) *ShipmentDocumentationCreate {
	return sdc.SetOrganizationID(o.ID)
}

// SetShipment sets the "shipment" edge to the Shipment entity.
func (sdc *ShipmentDocumentationCreate) SetShipment(s *Shipment) *ShipmentDocumentationCreate {
	return sdc.SetShipmentID(s.ID)
}

// SetDocumentClassification sets the "document_classification" edge to the DocumentClassification entity.
func (sdc *ShipmentDocumentationCreate) SetDocumentClassification(d *DocumentClassification) *ShipmentDocumentationCreate {
	return sdc.SetDocumentClassificationID(d.ID)
}

// Mutation returns the ShipmentDocumentationMutation object of the builder.
func (sdc *ShipmentDocumentationCreate) Mutation() *ShipmentDocumentationMutation {
	return sdc.mutation
}

// Save creates the ShipmentDocumentation in the database.
func (sdc *ShipmentDocumentationCreate) Save(ctx context.Context) (*ShipmentDocumentation, error) {
	sdc.defaults()
	return withHooks(ctx, sdc.sqlSave, sdc.mutation, sdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sdc *ShipmentDocumentationCreate) SaveX(ctx context.Context) *ShipmentDocumentation {
	v, err := sdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdc *ShipmentDocumentationCreate) Exec(ctx context.Context) error {
	_, err := sdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdc *ShipmentDocumentationCreate) ExecX(ctx context.Context) {
	if err := sdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sdc *ShipmentDocumentationCreate) defaults() {
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		v := shipmentdocumentation.DefaultCreatedAt()
		sdc.mutation.SetCreatedAt(v)
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		v := shipmentdocumentation.DefaultUpdatedAt()
		sdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sdc.mutation.Version(); !ok {
		v := shipmentdocumentation.DefaultVersion
		sdc.mutation.SetVersion(v)
	}
	if _, ok := sdc.mutation.ID(); !ok {
		v := shipmentdocumentation.DefaultID()
		sdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdc *ShipmentDocumentationCreate) check() error {
	if _, ok := sdc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "ShipmentDocumentation.business_unit_id"`)}
	}
	if _, ok := sdc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "ShipmentDocumentation.organization_id"`)}
	}
	if _, ok := sdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ShipmentDocumentation.created_at"`)}
	}
	if _, ok := sdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ShipmentDocumentation.updated_at"`)}
	}
	if _, ok := sdc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ShipmentDocumentation.version"`)}
	}
	if _, ok := sdc.mutation.ShipmentID(); !ok {
		return &ValidationError{Name: "shipment_id", err: errors.New(`ent: missing required field "ShipmentDocumentation.shipment_id"`)}
	}
	if _, ok := sdc.mutation.DocumentURL(); !ok {
		return &ValidationError{Name: "document_url", err: errors.New(`ent: missing required field "ShipmentDocumentation.document_url"`)}
	}
	if v, ok := sdc.mutation.DocumentURL(); ok {
		if err := shipmentdocumentation.DocumentURLValidator(v); err != nil {
			return &ValidationError{Name: "document_url", err: fmt.Errorf(`ent: validator failed for field "ShipmentDocumentation.document_url": %w`, err)}
		}
	}
	if _, ok := sdc.mutation.DocumentClassificationID(); !ok {
		return &ValidationError{Name: "document_classification_id", err: errors.New(`ent: missing required field "ShipmentDocumentation.document_classification_id"`)}
	}
	if _, ok := sdc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "ShipmentDocumentation.business_unit"`)}
	}
	if _, ok := sdc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "ShipmentDocumentation.organization"`)}
	}
	if _, ok := sdc.mutation.ShipmentID(); !ok {
		return &ValidationError{Name: "shipment", err: errors.New(`ent: missing required edge "ShipmentDocumentation.shipment"`)}
	}
	if _, ok := sdc.mutation.DocumentClassificationID(); !ok {
		return &ValidationError{Name: "document_classification", err: errors.New(`ent: missing required edge "ShipmentDocumentation.document_classification"`)}
	}
	return nil
}

func (sdc *ShipmentDocumentationCreate) sqlSave(ctx context.Context) (*ShipmentDocumentation, error) {
	if err := sdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdc.driver, _spec); err != nil {
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
	sdc.mutation.id = &_node.ID
	sdc.mutation.done = true
	return _node, nil
}

func (sdc *ShipmentDocumentationCreate) createSpec() (*ShipmentDocumentation, *sqlgraph.CreateSpec) {
	var (
		_node = &ShipmentDocumentation{config: sdc.config}
		_spec = sqlgraph.NewCreateSpec(shipmentdocumentation.Table, sqlgraph.NewFieldSpec(shipmentdocumentation.FieldID, field.TypeUUID))
	)
	if id, ok := sdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sdc.mutation.CreatedAt(); ok {
		_spec.SetField(shipmentdocumentation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sdc.mutation.UpdatedAt(); ok {
		_spec.SetField(shipmentdocumentation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sdc.mutation.Version(); ok {
		_spec.SetField(shipmentdocumentation.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := sdc.mutation.DocumentURL(); ok {
		_spec.SetField(shipmentdocumentation.FieldDocumentURL, field.TypeString, value)
		_node.DocumentURL = value
	}
	if nodes := sdc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmentdocumentation.BusinessUnitTable,
			Columns: []string{shipmentdocumentation.BusinessUnitColumn},
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
	if nodes := sdc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shipmentdocumentation.OrganizationTable,
			Columns: []string{shipmentdocumentation.OrganizationColumn},
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
	if nodes := sdc.mutation.ShipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shipmentdocumentation.ShipmentTable,
			Columns: []string{shipmentdocumentation.ShipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ShipmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sdc.mutation.DocumentClassificationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   shipmentdocumentation.DocumentClassificationTable,
			Columns: []string{shipmentdocumentation.DocumentClassificationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentclassification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DocumentClassificationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShipmentDocumentationCreateBulk is the builder for creating many ShipmentDocumentation entities in bulk.
type ShipmentDocumentationCreateBulk struct {
	config
	err      error
	builders []*ShipmentDocumentationCreate
}

// Save creates the ShipmentDocumentation entities in the database.
func (sdcb *ShipmentDocumentationCreateBulk) Save(ctx context.Context) ([]*ShipmentDocumentation, error) {
	if sdcb.err != nil {
		return nil, sdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sdcb.builders))
	nodes := make([]*ShipmentDocumentation, len(sdcb.builders))
	mutators := make([]Mutator, len(sdcb.builders))
	for i := range sdcb.builders {
		func(i int, root context.Context) {
			builder := sdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShipmentDocumentationMutation)
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
					_, err = mutators[i+1].Mutate(root, sdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdcb *ShipmentDocumentationCreateBulk) SaveX(ctx context.Context) []*ShipmentDocumentation {
	v, err := sdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdcb *ShipmentDocumentationCreateBulk) Exec(ctx context.Context) error {
	_, err := sdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdcb *ShipmentDocumentationCreateBulk) ExecX(ctx context.Context) {
	if err := sdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
