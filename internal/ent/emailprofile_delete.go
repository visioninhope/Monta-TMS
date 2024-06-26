// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/emailprofile"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// EmailProfileDelete is the builder for deleting a EmailProfile entity.
type EmailProfileDelete struct {
	config
	hooks    []Hook
	mutation *EmailProfileMutation
}

// Where appends a list predicates to the EmailProfileDelete builder.
func (epd *EmailProfileDelete) Where(ps ...predicate.EmailProfile) *EmailProfileDelete {
	epd.mutation.Where(ps...)
	return epd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epd *EmailProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epd.sqlExec, epd.mutation, epd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epd *EmailProfileDelete) ExecX(ctx context.Context) int {
	n, err := epd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epd *EmailProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emailprofile.Table, sqlgraph.NewFieldSpec(emailprofile.FieldID, field.TypeUUID))
	if ps := epd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	epd.mutation.done = true
	return affected, err
}

// EmailProfileDeleteOne is the builder for deleting a single EmailProfile entity.
type EmailProfileDeleteOne struct {
	epd *EmailProfileDelete
}

// Where appends a list predicates to the EmailProfileDelete builder.
func (epdo *EmailProfileDeleteOne) Where(ps ...predicate.EmailProfile) *EmailProfileDeleteOne {
	epdo.epd.mutation.Where(ps...)
	return epdo
}

// Exec executes the deletion query.
func (epdo *EmailProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := epdo.epd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epdo *EmailProfileDeleteOne) ExecX(ctx context.Context) {
	if err := epdo.Exec(ctx); err != nil {
		panic(err)
	}
}
