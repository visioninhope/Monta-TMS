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
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/shipmentmove"
	"github.com/emoss08/trenova/internal/ent/stop"
	"github.com/google/uuid"
)

// StopCreate is the builder for creating a Stop entity.
type StopCreate struct {
	config
	mutation *StopMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (sc *StopCreate) SetBusinessUnitID(u uuid.UUID) *StopCreate {
	sc.mutation.SetBusinessUnitID(u)
	return sc
}

// SetOrganizationID sets the "organization_id" field.
func (sc *StopCreate) SetOrganizationID(u uuid.UUID) *StopCreate {
	sc.mutation.SetOrganizationID(u)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StopCreate) SetCreatedAt(t time.Time) *StopCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StopCreate) SetNillableCreatedAt(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StopCreate) SetUpdatedAt(t time.Time) *StopCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StopCreate) SetNillableUpdatedAt(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetVersion sets the "version" field.
func (sc *StopCreate) SetVersion(i int) *StopCreate {
	sc.mutation.SetVersion(i)
	return sc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (sc *StopCreate) SetNillableVersion(i *int) *StopCreate {
	if i != nil {
		sc.SetVersion(*i)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *StopCreate) SetStatus(s stop.Status) *StopCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StopCreate) SetNillableStatus(s *stop.Status) *StopCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetShipmentMoveID sets the "shipment_move_id" field.
func (sc *StopCreate) SetShipmentMoveID(u uuid.UUID) *StopCreate {
	sc.mutation.SetShipmentMoveID(u)
	return sc
}

// SetStopType sets the "stop_type" field.
func (sc *StopCreate) SetStopType(st stop.StopType) *StopCreate {
	sc.mutation.SetStopType(st)
	return sc
}

// SetSequence sets the "sequence" field.
func (sc *StopCreate) SetSequence(i int) *StopCreate {
	sc.mutation.SetSequence(i)
	return sc
}

// SetNillableSequence sets the "sequence" field if the given value is not nil.
func (sc *StopCreate) SetNillableSequence(i *int) *StopCreate {
	if i != nil {
		sc.SetSequence(*i)
	}
	return sc
}

// SetLocationID sets the "location_id" field.
func (sc *StopCreate) SetLocationID(u uuid.UUID) *StopCreate {
	sc.mutation.SetLocationID(u)
	return sc
}

// SetNillableLocationID sets the "location_id" field if the given value is not nil.
func (sc *StopCreate) SetNillableLocationID(u *uuid.UUID) *StopCreate {
	if u != nil {
		sc.SetLocationID(*u)
	}
	return sc
}

// SetPieces sets the "pieces" field.
func (sc *StopCreate) SetPieces(f float64) *StopCreate {
	sc.mutation.SetPieces(f)
	return sc
}

// SetNillablePieces sets the "pieces" field if the given value is not nil.
func (sc *StopCreate) SetNillablePieces(f *float64) *StopCreate {
	if f != nil {
		sc.SetPieces(*f)
	}
	return sc
}

// SetWeight sets the "weight" field.
func (sc *StopCreate) SetWeight(f float64) *StopCreate {
	sc.mutation.SetWeight(f)
	return sc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (sc *StopCreate) SetNillableWeight(f *float64) *StopCreate {
	if f != nil {
		sc.SetWeight(*f)
	}
	return sc
}

// SetAddressLine sets the "address_line" field.
func (sc *StopCreate) SetAddressLine(s string) *StopCreate {
	sc.mutation.SetAddressLine(s)
	return sc
}

// SetNillableAddressLine sets the "address_line" field if the given value is not nil.
func (sc *StopCreate) SetNillableAddressLine(s *string) *StopCreate {
	if s != nil {
		sc.SetAddressLine(*s)
	}
	return sc
}

// SetAppointmentStart sets the "appointment_start" field.
func (sc *StopCreate) SetAppointmentStart(t time.Time) *StopCreate {
	sc.mutation.SetAppointmentStart(t)
	return sc
}

// SetNillableAppointmentStart sets the "appointment_start" field if the given value is not nil.
func (sc *StopCreate) SetNillableAppointmentStart(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetAppointmentStart(*t)
	}
	return sc
}

// SetAppointmentEnd sets the "appointment_end" field.
func (sc *StopCreate) SetAppointmentEnd(t time.Time) *StopCreate {
	sc.mutation.SetAppointmentEnd(t)
	return sc
}

// SetNillableAppointmentEnd sets the "appointment_end" field if the given value is not nil.
func (sc *StopCreate) SetNillableAppointmentEnd(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetAppointmentEnd(*t)
	}
	return sc
}

// SetArrivalTime sets the "arrival_time" field.
func (sc *StopCreate) SetArrivalTime(t time.Time) *StopCreate {
	sc.mutation.SetArrivalTime(t)
	return sc
}

// SetNillableArrivalTime sets the "arrival_time" field if the given value is not nil.
func (sc *StopCreate) SetNillableArrivalTime(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetArrivalTime(*t)
	}
	return sc
}

// SetDepartureTime sets the "departure_time" field.
func (sc *StopCreate) SetDepartureTime(t time.Time) *StopCreate {
	sc.mutation.SetDepartureTime(t)
	return sc
}

// SetNillableDepartureTime sets the "departure_time" field if the given value is not nil.
func (sc *StopCreate) SetNillableDepartureTime(t *time.Time) *StopCreate {
	if t != nil {
		sc.SetDepartureTime(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StopCreate) SetID(u uuid.UUID) *StopCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StopCreate) SetNillableID(u *uuid.UUID) *StopCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (sc *StopCreate) SetBusinessUnit(b *BusinessUnit) *StopCreate {
	return sc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (sc *StopCreate) SetOrganization(o *Organization) *StopCreate {
	return sc.SetOrganizationID(o.ID)
}

// SetShipmentMove sets the "shipment_move" edge to the ShipmentMove entity.
func (sc *StopCreate) SetShipmentMove(s *ShipmentMove) *StopCreate {
	return sc.SetShipmentMoveID(s.ID)
}

// Mutation returns the StopMutation object of the builder.
func (sc *StopCreate) Mutation() *StopMutation {
	return sc.mutation
}

// Save creates the Stop in the database.
func (sc *StopCreate) Save(ctx context.Context) (*Stop, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StopCreate) SaveX(ctx context.Context) *Stop {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StopCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StopCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StopCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if stop.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized stop.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := stop.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if stop.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized stop.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := stop.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Version(); !ok {
		v := stop.DefaultVersion
		sc.mutation.SetVersion(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := stop.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Sequence(); !ok {
		v := stop.DefaultSequence
		sc.mutation.SetSequence(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if stop.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized stop.DefaultID (forgotten import ent/runtime?)")
		}
		v := stop.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *StopCreate) check() error {
	if _, ok := sc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "Stop.business_unit_id"`)}
	}
	if _, ok := sc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "Stop.organization_id"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Stop.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Stop.updated_at"`)}
	}
	if _, ok := sc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Stop.version"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Stop.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := stop.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Stop.status": %w`, err)}
		}
	}
	if _, ok := sc.mutation.ShipmentMoveID(); !ok {
		return &ValidationError{Name: "shipment_move_id", err: errors.New(`ent: missing required field "Stop.shipment_move_id"`)}
	}
	if _, ok := sc.mutation.StopType(); !ok {
		return &ValidationError{Name: "stop_type", err: errors.New(`ent: missing required field "Stop.stop_type"`)}
	}
	if v, ok := sc.mutation.StopType(); ok {
		if err := stop.StopTypeValidator(v); err != nil {
			return &ValidationError{Name: "stop_type", err: fmt.Errorf(`ent: validator failed for field "Stop.stop_type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`ent: missing required field "Stop.sequence"`)}
	}
	if v, ok := sc.mutation.Sequence(); ok {
		if err := stop.SequenceValidator(v); err != nil {
			return &ValidationError{Name: "sequence", err: fmt.Errorf(`ent: validator failed for field "Stop.sequence": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Pieces(); ok {
		if err := stop.PiecesValidator(v); err != nil {
			return &ValidationError{Name: "pieces", err: fmt.Errorf(`ent: validator failed for field "Stop.pieces": %w`, err)}
		}
	}
	if v, ok := sc.mutation.Weight(); ok {
		if err := stop.WeightValidator(v); err != nil {
			return &ValidationError{Name: "weight", err: fmt.Errorf(`ent: validator failed for field "Stop.weight": %w`, err)}
		}
	}
	if _, ok := sc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "Stop.business_unit"`)}
	}
	if _, ok := sc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "Stop.organization"`)}
	}
	if _, ok := sc.mutation.ShipmentMoveID(); !ok {
		return &ValidationError{Name: "shipment_move", err: errors.New(`ent: missing required edge "Stop.shipment_move"`)}
	}
	return nil
}

func (sc *StopCreate) sqlSave(ctx context.Context) (*Stop, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StopCreate) createSpec() (*Stop, *sqlgraph.CreateSpec) {
	var (
		_node = &Stop{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stop.Table, sqlgraph.NewFieldSpec(stop.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(stop.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(stop.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Version(); ok {
		_spec.SetField(stop.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(stop.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.StopType(); ok {
		_spec.SetField(stop.FieldStopType, field.TypeEnum, value)
		_node.StopType = value
	}
	if value, ok := sc.mutation.Sequence(); ok {
		_spec.SetField(stop.FieldSequence, field.TypeInt, value)
		_node.Sequence = value
	}
	if value, ok := sc.mutation.LocationID(); ok {
		_spec.SetField(stop.FieldLocationID, field.TypeUUID, value)
		_node.LocationID = &value
	}
	if value, ok := sc.mutation.Pieces(); ok {
		_spec.SetField(stop.FieldPieces, field.TypeFloat64, value)
		_node.Pieces = value
	}
	if value, ok := sc.mutation.Weight(); ok {
		_spec.SetField(stop.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := sc.mutation.AddressLine(); ok {
		_spec.SetField(stop.FieldAddressLine, field.TypeString, value)
		_node.AddressLine = value
	}
	if value, ok := sc.mutation.AppointmentStart(); ok {
		_spec.SetField(stop.FieldAppointmentStart, field.TypeTime, value)
		_node.AppointmentStart = &value
	}
	if value, ok := sc.mutation.AppointmentEnd(); ok {
		_spec.SetField(stop.FieldAppointmentEnd, field.TypeTime, value)
		_node.AppointmentEnd = &value
	}
	if value, ok := sc.mutation.ArrivalTime(); ok {
		_spec.SetField(stop.FieldArrivalTime, field.TypeTime, value)
		_node.ArrivalTime = &value
	}
	if value, ok := sc.mutation.DepartureTime(); ok {
		_spec.SetField(stop.FieldDepartureTime, field.TypeTime, value)
		_node.DepartureTime = &value
	}
	if nodes := sc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   stop.BusinessUnitTable,
			Columns: []string{stop.BusinessUnitColumn},
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
	if nodes := sc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   stop.OrganizationTable,
			Columns: []string{stop.OrganizationColumn},
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
	if nodes := sc.mutation.ShipmentMoveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stop.ShipmentMoveTable,
			Columns: []string{stop.ShipmentMoveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shipmentmove.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ShipmentMoveID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StopCreateBulk is the builder for creating many Stop entities in bulk.
type StopCreateBulk struct {
	config
	err      error
	builders []*StopCreate
}

// Save creates the Stop entities in the database.
func (scb *StopCreateBulk) Save(ctx context.Context) ([]*Stop, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stop, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StopMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StopCreateBulk) SaveX(ctx context.Context) []*Stop {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StopCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StopCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
