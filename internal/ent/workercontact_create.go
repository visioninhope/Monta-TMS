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
	"github.com/emoss08/trenova/internal/ent/worker"
	"github.com/emoss08/trenova/internal/ent/workercontact"
	"github.com/google/uuid"
)

// WorkerContactCreate is the builder for creating a WorkerContact entity.
type WorkerContactCreate struct {
	config
	mutation *WorkerContactMutation
	hooks    []Hook
}

// SetBusinessUnitID sets the "business_unit_id" field.
func (wcc *WorkerContactCreate) SetBusinessUnitID(u uuid.UUID) *WorkerContactCreate {
	wcc.mutation.SetBusinessUnitID(u)
	return wcc
}

// SetOrganizationID sets the "organization_id" field.
func (wcc *WorkerContactCreate) SetOrganizationID(u uuid.UUID) *WorkerContactCreate {
	wcc.mutation.SetOrganizationID(u)
	return wcc
}

// SetCreatedAt sets the "created_at" field.
func (wcc *WorkerContactCreate) SetCreatedAt(t time.Time) *WorkerContactCreate {
	wcc.mutation.SetCreatedAt(t)
	return wcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableCreatedAt(t *time.Time) *WorkerContactCreate {
	if t != nil {
		wcc.SetCreatedAt(*t)
	}
	return wcc
}

// SetUpdatedAt sets the "updated_at" field.
func (wcc *WorkerContactCreate) SetUpdatedAt(t time.Time) *WorkerContactCreate {
	wcc.mutation.SetUpdatedAt(t)
	return wcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableUpdatedAt(t *time.Time) *WorkerContactCreate {
	if t != nil {
		wcc.SetUpdatedAt(*t)
	}
	return wcc
}

// SetVersion sets the "version" field.
func (wcc *WorkerContactCreate) SetVersion(i int) *WorkerContactCreate {
	wcc.mutation.SetVersion(i)
	return wcc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableVersion(i *int) *WorkerContactCreate {
	if i != nil {
		wcc.SetVersion(*i)
	}
	return wcc
}

// SetWorkerID sets the "worker_id" field.
func (wcc *WorkerContactCreate) SetWorkerID(u uuid.UUID) *WorkerContactCreate {
	wcc.mutation.SetWorkerID(u)
	return wcc
}

// SetName sets the "name" field.
func (wcc *WorkerContactCreate) SetName(s string) *WorkerContactCreate {
	wcc.mutation.SetName(s)
	return wcc
}

// SetEmail sets the "email" field.
func (wcc *WorkerContactCreate) SetEmail(s string) *WorkerContactCreate {
	wcc.mutation.SetEmail(s)
	return wcc
}

// SetPhone sets the "phone" field.
func (wcc *WorkerContactCreate) SetPhone(s string) *WorkerContactCreate {
	wcc.mutation.SetPhone(s)
	return wcc
}

// SetRelationship sets the "relationship" field.
func (wcc *WorkerContactCreate) SetRelationship(s string) *WorkerContactCreate {
	wcc.mutation.SetRelationship(s)
	return wcc
}

// SetNillableRelationship sets the "relationship" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableRelationship(s *string) *WorkerContactCreate {
	if s != nil {
		wcc.SetRelationship(*s)
	}
	return wcc
}

// SetIsPrimary sets the "is_primary" field.
func (wcc *WorkerContactCreate) SetIsPrimary(b bool) *WorkerContactCreate {
	wcc.mutation.SetIsPrimary(b)
	return wcc
}

// SetNillableIsPrimary sets the "is_primary" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableIsPrimary(b *bool) *WorkerContactCreate {
	if b != nil {
		wcc.SetIsPrimary(*b)
	}
	return wcc
}

// SetID sets the "id" field.
func (wcc *WorkerContactCreate) SetID(u uuid.UUID) *WorkerContactCreate {
	wcc.mutation.SetID(u)
	return wcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wcc *WorkerContactCreate) SetNillableID(u *uuid.UUID) *WorkerContactCreate {
	if u != nil {
		wcc.SetID(*u)
	}
	return wcc
}

// SetBusinessUnit sets the "business_unit" edge to the BusinessUnit entity.
func (wcc *WorkerContactCreate) SetBusinessUnit(b *BusinessUnit) *WorkerContactCreate {
	return wcc.SetBusinessUnitID(b.ID)
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (wcc *WorkerContactCreate) SetOrganization(o *Organization) *WorkerContactCreate {
	return wcc.SetOrganizationID(o.ID)
}

// SetWorker sets the "worker" edge to the Worker entity.
func (wcc *WorkerContactCreate) SetWorker(w *Worker) *WorkerContactCreate {
	return wcc.SetWorkerID(w.ID)
}

// Mutation returns the WorkerContactMutation object of the builder.
func (wcc *WorkerContactCreate) Mutation() *WorkerContactMutation {
	return wcc.mutation
}

// Save creates the WorkerContact in the database.
func (wcc *WorkerContactCreate) Save(ctx context.Context) (*WorkerContact, error) {
	wcc.defaults()
	return withHooks(ctx, wcc.sqlSave, wcc.mutation, wcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wcc *WorkerContactCreate) SaveX(ctx context.Context) *WorkerContact {
	v, err := wcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcc *WorkerContactCreate) Exec(ctx context.Context) error {
	_, err := wcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcc *WorkerContactCreate) ExecX(ctx context.Context) {
	if err := wcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wcc *WorkerContactCreate) defaults() {
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		v := workercontact.DefaultCreatedAt()
		wcc.mutation.SetCreatedAt(v)
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		v := workercontact.DefaultUpdatedAt()
		wcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wcc.mutation.Version(); !ok {
		v := workercontact.DefaultVersion
		wcc.mutation.SetVersion(v)
	}
	if _, ok := wcc.mutation.IsPrimary(); !ok {
		v := workercontact.DefaultIsPrimary
		wcc.mutation.SetIsPrimary(v)
	}
	if _, ok := wcc.mutation.ID(); !ok {
		v := workercontact.DefaultID()
		wcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wcc *WorkerContactCreate) check() error {
	if _, ok := wcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit_id", err: errors.New(`ent: missing required field "WorkerContact.business_unit_id"`)}
	}
	if _, ok := wcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "WorkerContact.organization_id"`)}
	}
	if _, ok := wcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkerContact.created_at"`)}
	}
	if _, ok := wcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "WorkerContact.updated_at"`)}
	}
	if _, ok := wcc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "WorkerContact.version"`)}
	}
	if _, ok := wcc.mutation.WorkerID(); !ok {
		return &ValidationError{Name: "worker_id", err: errors.New(`ent: missing required field "WorkerContact.worker_id"`)}
	}
	if _, ok := wcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WorkerContact.name"`)}
	}
	if v, ok := wcc.mutation.Name(); ok {
		if err := workercontact.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WorkerContact.name": %w`, err)}
		}
	}
	if _, ok := wcc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "WorkerContact.email"`)}
	}
	if v, ok := wcc.mutation.Email(); ok {
		if err := workercontact.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "WorkerContact.email": %w`, err)}
		}
	}
	if _, ok := wcc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "WorkerContact.phone"`)}
	}
	if v, ok := wcc.mutation.Phone(); ok {
		if err := workercontact.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "WorkerContact.phone": %w`, err)}
		}
	}
	if _, ok := wcc.mutation.IsPrimary(); !ok {
		return &ValidationError{Name: "is_primary", err: errors.New(`ent: missing required field "WorkerContact.is_primary"`)}
	}
	if _, ok := wcc.mutation.BusinessUnitID(); !ok {
		return &ValidationError{Name: "business_unit", err: errors.New(`ent: missing required edge "WorkerContact.business_unit"`)}
	}
	if _, ok := wcc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "WorkerContact.organization"`)}
	}
	if _, ok := wcc.mutation.WorkerID(); !ok {
		return &ValidationError{Name: "worker", err: errors.New(`ent: missing required edge "WorkerContact.worker"`)}
	}
	return nil
}

func (wcc *WorkerContactCreate) sqlSave(ctx context.Context) (*WorkerContact, error) {
	if err := wcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wcc.driver, _spec); err != nil {
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
	wcc.mutation.id = &_node.ID
	wcc.mutation.done = true
	return _node, nil
}

func (wcc *WorkerContactCreate) createSpec() (*WorkerContact, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkerContact{config: wcc.config}
		_spec = sqlgraph.NewCreateSpec(workercontact.Table, sqlgraph.NewFieldSpec(workercontact.FieldID, field.TypeUUID))
	)
	if id, ok := wcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wcc.mutation.CreatedAt(); ok {
		_spec.SetField(workercontact.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wcc.mutation.UpdatedAt(); ok {
		_spec.SetField(workercontact.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := wcc.mutation.Version(); ok {
		_spec.SetField(workercontact.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := wcc.mutation.Name(); ok {
		_spec.SetField(workercontact.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wcc.mutation.Email(); ok {
		_spec.SetField(workercontact.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := wcc.mutation.Phone(); ok {
		_spec.SetField(workercontact.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := wcc.mutation.Relationship(); ok {
		_spec.SetField(workercontact.FieldRelationship, field.TypeString, value)
		_node.Relationship = value
	}
	if value, ok := wcc.mutation.IsPrimary(); ok {
		_spec.SetField(workercontact.FieldIsPrimary, field.TypeBool, value)
		_node.IsPrimary = value
	}
	if nodes := wcc.mutation.BusinessUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercontact.BusinessUnitTable,
			Columns: []string{workercontact.BusinessUnitColumn},
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
	if nodes := wcc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workercontact.OrganizationTable,
			Columns: []string{workercontact.OrganizationColumn},
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
	if nodes := wcc.mutation.WorkerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workercontact.WorkerTable,
			Columns: []string{workercontact.WorkerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(worker.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.WorkerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkerContactCreateBulk is the builder for creating many WorkerContact entities in bulk.
type WorkerContactCreateBulk struct {
	config
	err      error
	builders []*WorkerContactCreate
}

// Save creates the WorkerContact entities in the database.
func (wccb *WorkerContactCreateBulk) Save(ctx context.Context) ([]*WorkerContact, error) {
	if wccb.err != nil {
		return nil, wccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wccb.builders))
	nodes := make([]*WorkerContact, len(wccb.builders))
	mutators := make([]Mutator, len(wccb.builders))
	for i := range wccb.builders {
		func(i int, root context.Context) {
			builder := wccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkerContactMutation)
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
					_, err = mutators[i+1].Mutate(root, wccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wccb *WorkerContactCreateBulk) SaveX(ctx context.Context) []*WorkerContact {
	v, err := wccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wccb *WorkerContactCreateBulk) Exec(ctx context.Context) error {
	_, err := wccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wccb *WorkerContactCreateBulk) ExecX(ctx context.Context) {
	if err := wccb.Exec(ctx); err != nil {
		panic(err)
	}
}
