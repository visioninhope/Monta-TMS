// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/featureflag"
	"github.com/emoss08/trenova/internal/ent/organizationfeatureflag"
	"github.com/google/uuid"
)

// FeatureFlagCreate is the builder for creating a FeatureFlag entity.
type FeatureFlagCreate struct {
	config
	mutation *FeatureFlagMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ffc *FeatureFlagCreate) SetCreatedAt(t time.Time) *FeatureFlagCreate {
	ffc.mutation.SetCreatedAt(t)
	return ffc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ffc *FeatureFlagCreate) SetNillableCreatedAt(t *time.Time) *FeatureFlagCreate {
	if t != nil {
		ffc.SetCreatedAt(*t)
	}
	return ffc
}

// SetUpdatedAt sets the "updated_at" field.
func (ffc *FeatureFlagCreate) SetUpdatedAt(t time.Time) *FeatureFlagCreate {
	ffc.mutation.SetUpdatedAt(t)
	return ffc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ffc *FeatureFlagCreate) SetNillableUpdatedAt(t *time.Time) *FeatureFlagCreate {
	if t != nil {
		ffc.SetUpdatedAt(*t)
	}
	return ffc
}

// SetName sets the "name" field.
func (ffc *FeatureFlagCreate) SetName(s string) *FeatureFlagCreate {
	ffc.mutation.SetName(s)
	return ffc
}

// SetCode sets the "code" field.
func (ffc *FeatureFlagCreate) SetCode(s string) *FeatureFlagCreate {
	ffc.mutation.SetCode(s)
	return ffc
}

// SetBeta sets the "beta" field.
func (ffc *FeatureFlagCreate) SetBeta(b bool) *FeatureFlagCreate {
	ffc.mutation.SetBeta(b)
	return ffc
}

// SetNillableBeta sets the "beta" field if the given value is not nil.
func (ffc *FeatureFlagCreate) SetNillableBeta(b *bool) *FeatureFlagCreate {
	if b != nil {
		ffc.SetBeta(*b)
	}
	return ffc
}

// SetDescription sets the "description" field.
func (ffc *FeatureFlagCreate) SetDescription(s string) *FeatureFlagCreate {
	ffc.mutation.SetDescription(s)
	return ffc
}

// SetPreviewPictureURL sets the "preview_picture_url" field.
func (ffc *FeatureFlagCreate) SetPreviewPictureURL(s string) *FeatureFlagCreate {
	ffc.mutation.SetPreviewPictureURL(s)
	return ffc
}

// SetNillablePreviewPictureURL sets the "preview_picture_url" field if the given value is not nil.
func (ffc *FeatureFlagCreate) SetNillablePreviewPictureURL(s *string) *FeatureFlagCreate {
	if s != nil {
		ffc.SetPreviewPictureURL(*s)
	}
	return ffc
}

// SetID sets the "id" field.
func (ffc *FeatureFlagCreate) SetID(u uuid.UUID) *FeatureFlagCreate {
	ffc.mutation.SetID(u)
	return ffc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ffc *FeatureFlagCreate) SetNillableID(u *uuid.UUID) *FeatureFlagCreate {
	if u != nil {
		ffc.SetID(*u)
	}
	return ffc
}

// AddOrganizationFeatureFlagIDs adds the "organization_feature_flag" edge to the OrganizationFeatureFlag entity by IDs.
func (ffc *FeatureFlagCreate) AddOrganizationFeatureFlagIDs(ids ...uuid.UUID) *FeatureFlagCreate {
	ffc.mutation.AddOrganizationFeatureFlagIDs(ids...)
	return ffc
}

// AddOrganizationFeatureFlag adds the "organization_feature_flag" edges to the OrganizationFeatureFlag entity.
func (ffc *FeatureFlagCreate) AddOrganizationFeatureFlag(o ...*OrganizationFeatureFlag) *FeatureFlagCreate {
	ids := make([]uuid.UUID, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ffc.AddOrganizationFeatureFlagIDs(ids...)
}

// Mutation returns the FeatureFlagMutation object of the builder.
func (ffc *FeatureFlagCreate) Mutation() *FeatureFlagMutation {
	return ffc.mutation
}

// Save creates the FeatureFlag in the database.
func (ffc *FeatureFlagCreate) Save(ctx context.Context) (*FeatureFlag, error) {
	ffc.defaults()
	return withHooks(ctx, ffc.sqlSave, ffc.mutation, ffc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ffc *FeatureFlagCreate) SaveX(ctx context.Context) *FeatureFlag {
	v, err := ffc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ffc *FeatureFlagCreate) Exec(ctx context.Context) error {
	_, err := ffc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ffc *FeatureFlagCreate) ExecX(ctx context.Context) {
	if err := ffc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ffc *FeatureFlagCreate) defaults() {
	if _, ok := ffc.mutation.CreatedAt(); !ok {
		v := featureflag.DefaultCreatedAt()
		ffc.mutation.SetCreatedAt(v)
	}
	if _, ok := ffc.mutation.UpdatedAt(); !ok {
		v := featureflag.DefaultUpdatedAt()
		ffc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ffc.mutation.Beta(); !ok {
		v := featureflag.DefaultBeta
		ffc.mutation.SetBeta(v)
	}
	if _, ok := ffc.mutation.ID(); !ok {
		v := featureflag.DefaultID()
		ffc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ffc *FeatureFlagCreate) check() error {
	if _, ok := ffc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FeatureFlag.created_at"`)}
	}
	if _, ok := ffc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FeatureFlag.updated_at"`)}
	}
	if _, ok := ffc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FeatureFlag.name"`)}
	}
	if v, ok := ffc.mutation.Name(); ok {
		if err := featureflag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "FeatureFlag.name": %w`, err)}
		}
	}
	if _, ok := ffc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "FeatureFlag.code"`)}
	}
	if v, ok := ffc.mutation.Code(); ok {
		if err := featureflag.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "FeatureFlag.code": %w`, err)}
		}
	}
	if _, ok := ffc.mutation.Beta(); !ok {
		return &ValidationError{Name: "beta", err: errors.New(`ent: missing required field "FeatureFlag.beta"`)}
	}
	if _, ok := ffc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "FeatureFlag.description"`)}
	}
	if v, ok := ffc.mutation.Description(); ok {
		if err := featureflag.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "FeatureFlag.description": %w`, err)}
		}
	}
	return nil
}

func (ffc *FeatureFlagCreate) sqlSave(ctx context.Context) (*FeatureFlag, error) {
	if err := ffc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ffc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ffc.driver, _spec); err != nil {
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
	ffc.mutation.id = &_node.ID
	ffc.mutation.done = true
	return _node, nil
}

func (ffc *FeatureFlagCreate) createSpec() (*FeatureFlag, *sqlgraph.CreateSpec) {
	var (
		_node = &FeatureFlag{config: ffc.config}
		_spec = sqlgraph.NewCreateSpec(featureflag.Table, sqlgraph.NewFieldSpec(featureflag.FieldID, field.TypeUUID))
	)
	if id, ok := ffc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ffc.mutation.CreatedAt(); ok {
		_spec.SetField(featureflag.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ffc.mutation.UpdatedAt(); ok {
		_spec.SetField(featureflag.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ffc.mutation.Name(); ok {
		_spec.SetField(featureflag.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ffc.mutation.Code(); ok {
		_spec.SetField(featureflag.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := ffc.mutation.Beta(); ok {
		_spec.SetField(featureflag.FieldBeta, field.TypeBool, value)
		_node.Beta = value
	}
	if value, ok := ffc.mutation.Description(); ok {
		_spec.SetField(featureflag.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ffc.mutation.PreviewPictureURL(); ok {
		_spec.SetField(featureflag.FieldPreviewPictureURL, field.TypeString, value)
		_node.PreviewPictureURL = value
	}
	if nodes := ffc.mutation.OrganizationFeatureFlagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   featureflag.OrganizationFeatureFlagTable,
			Columns: []string{featureflag.OrganizationFeatureFlagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationfeatureflag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FeatureFlagCreateBulk is the builder for creating many FeatureFlag entities in bulk.
type FeatureFlagCreateBulk struct {
	config
	err      error
	builders []*FeatureFlagCreate
}

// Save creates the FeatureFlag entities in the database.
func (ffcb *FeatureFlagCreateBulk) Save(ctx context.Context) ([]*FeatureFlag, error) {
	if ffcb.err != nil {
		return nil, ffcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ffcb.builders))
	nodes := make([]*FeatureFlag, len(ffcb.builders))
	mutators := make([]Mutator, len(ffcb.builders))
	for i := range ffcb.builders {
		func(i int, root context.Context) {
			builder := ffcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeatureFlagMutation)
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
					_, err = mutators[i+1].Mutate(root, ffcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ffcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ffcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ffcb *FeatureFlagCreateBulk) SaveX(ctx context.Context) []*FeatureFlag {
	v, err := ffcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ffcb *FeatureFlagCreateBulk) Exec(ctx context.Context) error {
	_, err := ffcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ffcb *FeatureFlagCreateBulk) ExecX(ctx context.Context) {
	if err := ffcb.Exec(ctx); err != nil {
		panic(err)
	}
}
