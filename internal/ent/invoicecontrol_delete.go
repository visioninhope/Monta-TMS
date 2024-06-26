// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/invoicecontrol"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// InvoiceControlDelete is the builder for deleting a InvoiceControl entity.
type InvoiceControlDelete struct {
	config
	hooks    []Hook
	mutation *InvoiceControlMutation
}

// Where appends a list predicates to the InvoiceControlDelete builder.
func (icd *InvoiceControlDelete) Where(ps ...predicate.InvoiceControl) *InvoiceControlDelete {
	icd.mutation.Where(ps...)
	return icd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (icd *InvoiceControlDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, icd.sqlExec, icd.mutation, icd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (icd *InvoiceControlDelete) ExecX(ctx context.Context) int {
	n, err := icd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (icd *InvoiceControlDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invoicecontrol.Table, sqlgraph.NewFieldSpec(invoicecontrol.FieldID, field.TypeUUID))
	if ps := icd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, icd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	icd.mutation.done = true
	return affected, err
}

// InvoiceControlDeleteOne is the builder for deleting a single InvoiceControl entity.
type InvoiceControlDeleteOne struct {
	icd *InvoiceControlDelete
}

// Where appends a list predicates to the InvoiceControlDelete builder.
func (icdo *InvoiceControlDeleteOne) Where(ps ...predicate.InvoiceControl) *InvoiceControlDeleteOne {
	icdo.icd.mutation.Where(ps...)
	return icdo
}

// Exec executes the deletion query.
func (icdo *InvoiceControlDeleteOne) Exec(ctx context.Context) error {
	n, err := icdo.icd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invoicecontrol.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (icdo *InvoiceControlDeleteOne) ExecX(ctx context.Context) {
	if err := icdo.Exec(ctx); err != nil {
		panic(err)
	}
}
