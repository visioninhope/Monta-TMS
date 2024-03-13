// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/emoss08/trenova/ent/tablechangealert"
)

// TableChangeAlertDelete is the builder for deleting a TableChangeAlert entity.
type TableChangeAlertDelete struct {
	config
	hooks    []Hook
	mutation *TableChangeAlertMutation
}

// Where appends a list predicates to the TableChangeAlertDelete builder.
func (tcad *TableChangeAlertDelete) Where(ps ...predicate.TableChangeAlert) *TableChangeAlertDelete {
	tcad.mutation.Where(ps...)
	return tcad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcad *TableChangeAlertDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tcad.sqlExec, tcad.mutation, tcad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcad *TableChangeAlertDelete) ExecX(ctx context.Context) int {
	n, err := tcad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcad *TableChangeAlertDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tablechangealert.Table, sqlgraph.NewFieldSpec(tablechangealert.FieldID, field.TypeUUID))
	if ps := tcad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcad.mutation.done = true
	return affected, err
}

// TableChangeAlertDeleteOne is the builder for deleting a single TableChangeAlert entity.
type TableChangeAlertDeleteOne struct {
	tcad *TableChangeAlertDelete
}

// Where appends a list predicates to the TableChangeAlertDelete builder.
func (tcado *TableChangeAlertDeleteOne) Where(ps ...predicate.TableChangeAlert) *TableChangeAlertDeleteOne {
	tcado.tcad.mutation.Where(ps...)
	return tcado
}

// Exec executes the deletion query.
func (tcado *TableChangeAlertDeleteOne) Exec(ctx context.Context) error {
	n, err := tcado.tcad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tablechangealert.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcado *TableChangeAlertDeleteOne) ExecX(ctx context.Context) {
	if err := tcado.Exec(ctx); err != nil {
		panic(err)
	}
}
