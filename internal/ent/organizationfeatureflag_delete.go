// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/organizationfeatureflag"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// OrganizationFeatureFlagDelete is the builder for deleting a OrganizationFeatureFlag entity.
type OrganizationFeatureFlagDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationFeatureFlagMutation
}

// Where appends a list predicates to the OrganizationFeatureFlagDelete builder.
func (offd *OrganizationFeatureFlagDelete) Where(ps ...predicate.OrganizationFeatureFlag) *OrganizationFeatureFlagDelete {
	offd.mutation.Where(ps...)
	return offd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (offd *OrganizationFeatureFlagDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, offd.sqlExec, offd.mutation, offd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (offd *OrganizationFeatureFlagDelete) ExecX(ctx context.Context) int {
	n, err := offd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (offd *OrganizationFeatureFlagDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(organizationfeatureflag.Table, sqlgraph.NewFieldSpec(organizationfeatureflag.FieldID, field.TypeUUID))
	if ps := offd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, offd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	offd.mutation.done = true
	return affected, err
}

// OrganizationFeatureFlagDeleteOne is the builder for deleting a single OrganizationFeatureFlag entity.
type OrganizationFeatureFlagDeleteOne struct {
	offd *OrganizationFeatureFlagDelete
}

// Where appends a list predicates to the OrganizationFeatureFlagDelete builder.
func (offdo *OrganizationFeatureFlagDeleteOne) Where(ps ...predicate.OrganizationFeatureFlag) *OrganizationFeatureFlagDeleteOne {
	offdo.offd.mutation.Where(ps...)
	return offdo
}

// Exec executes the deletion query.
func (offdo *OrganizationFeatureFlagDeleteOne) Exec(ctx context.Context) error {
	n, err := offdo.offd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationfeatureflag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (offdo *OrganizationFeatureFlagDeleteOne) ExecX(ctx context.Context) {
	if err := offdo.Exec(ctx); err != nil {
		panic(err)
	}
}
