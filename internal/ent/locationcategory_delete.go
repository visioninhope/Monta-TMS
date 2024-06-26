// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/locationcategory"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// LocationCategoryDelete is the builder for deleting a LocationCategory entity.
type LocationCategoryDelete struct {
	config
	hooks    []Hook
	mutation *LocationCategoryMutation
}

// Where appends a list predicates to the LocationCategoryDelete builder.
func (lcd *LocationCategoryDelete) Where(ps ...predicate.LocationCategory) *LocationCategoryDelete {
	lcd.mutation.Where(ps...)
	return lcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lcd *LocationCategoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, lcd.sqlExec, lcd.mutation, lcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (lcd *LocationCategoryDelete) ExecX(ctx context.Context) int {
	n, err := lcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lcd *LocationCategoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(locationcategory.Table, sqlgraph.NewFieldSpec(locationcategory.FieldID, field.TypeUUID))
	if ps := lcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	lcd.mutation.done = true
	return affected, err
}

// LocationCategoryDeleteOne is the builder for deleting a single LocationCategory entity.
type LocationCategoryDeleteOne struct {
	lcd *LocationCategoryDelete
}

// Where appends a list predicates to the LocationCategoryDelete builder.
func (lcdo *LocationCategoryDeleteOne) Where(ps ...predicate.LocationCategory) *LocationCategoryDeleteOne {
	lcdo.lcd.mutation.Where(ps...)
	return lcdo
}

// Exec executes the deletion query.
func (lcdo *LocationCategoryDeleteOne) Exec(ctx context.Context) error {
	n, err := lcdo.lcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{locationcategory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (lcdo *LocationCategoryDeleteOne) ExecX(ctx context.Context) {
	if err := lcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
