// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/billingcontrol"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// BillingControlDelete is the builder for deleting a BillingControl entity.
type BillingControlDelete struct {
	config
	hooks    []Hook
	mutation *BillingControlMutation
}

// Where appends a list predicates to the BillingControlDelete builder.
func (bcd *BillingControlDelete) Where(ps ...predicate.BillingControl) *BillingControlDelete {
	bcd.mutation.Where(ps...)
	return bcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bcd *BillingControlDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bcd.sqlExec, bcd.mutation, bcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bcd *BillingControlDelete) ExecX(ctx context.Context) int {
	n, err := bcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bcd *BillingControlDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(billingcontrol.Table, sqlgraph.NewFieldSpec(billingcontrol.FieldID, field.TypeUUID))
	if ps := bcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bcd.mutation.done = true
	return affected, err
}

// BillingControlDeleteOne is the builder for deleting a single BillingControl entity.
type BillingControlDeleteOne struct {
	bcd *BillingControlDelete
}

// Where appends a list predicates to the BillingControlDelete builder.
func (bcdo *BillingControlDeleteOne) Where(ps ...predicate.BillingControl) *BillingControlDeleteOne {
	bcdo.bcd.mutation.Where(ps...)
	return bcdo
}

// Exec executes the deletion query.
func (bcdo *BillingControlDeleteOne) Exec(ctx context.Context) error {
	n, err := bcdo.bcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{billingcontrol.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bcdo *BillingControlDeleteOne) ExecX(ctx context.Context) {
	if err := bcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
