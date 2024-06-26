// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/emoss08/trenova/internal/ent/equipmentmanufactuer"
	"github.com/emoss08/trenova/internal/ent/predicate"
)

// EquipmentManufactuerDelete is the builder for deleting a EquipmentManufactuer entity.
type EquipmentManufactuerDelete struct {
	config
	hooks    []Hook
	mutation *EquipmentManufactuerMutation
}

// Where appends a list predicates to the EquipmentManufactuerDelete builder.
func (emd *EquipmentManufactuerDelete) Where(ps ...predicate.EquipmentManufactuer) *EquipmentManufactuerDelete {
	emd.mutation.Where(ps...)
	return emd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (emd *EquipmentManufactuerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, emd.sqlExec, emd.mutation, emd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (emd *EquipmentManufactuerDelete) ExecX(ctx context.Context) int {
	n, err := emd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (emd *EquipmentManufactuerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(equipmentmanufactuer.Table, sqlgraph.NewFieldSpec(equipmentmanufactuer.FieldID, field.TypeUUID))
	if ps := emd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, emd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	emd.mutation.done = true
	return affected, err
}

// EquipmentManufactuerDeleteOne is the builder for deleting a single EquipmentManufactuer entity.
type EquipmentManufactuerDeleteOne struct {
	emd *EquipmentManufactuerDelete
}

// Where appends a list predicates to the EquipmentManufactuerDelete builder.
func (emdo *EquipmentManufactuerDeleteOne) Where(ps ...predicate.EquipmentManufactuer) *EquipmentManufactuerDeleteOne {
	emdo.emd.mutation.Where(ps...)
	return emdo
}

// Exec executes the deletion query.
func (emdo *EquipmentManufactuerDeleteOne) Exec(ctx context.Context) error {
	n, err := emdo.emd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{equipmentmanufactuer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (emdo *EquipmentManufactuerDeleteOne) ExecX(ctx context.Context) {
	if err := emdo.Exec(ctx); err != nil {
		panic(err)
	}
}
