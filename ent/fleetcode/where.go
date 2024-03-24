// Code generated by ent, DO NOT EDIT.

package fleetcode

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldDescription, v))
}

// RevenueGoal applies equality check predicate on the "revenue_goal" field. It's identical to RevenueGoalEQ.
func RevenueGoal(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldRevenueGoal, v))
}

// DeadheadGoal applies equality check predicate on the "deadhead_goal" field. It's identical to DeadheadGoalEQ.
func DeadheadGoal(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldDeadheadGoal, v))
}

// MileageGoal applies equality check predicate on the "mileage_goal" field. It's identical to MileageGoalEQ.
func MileageGoal(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldMileageGoal, v))
}

// ManagerID applies equality check predicate on the "manager_id" field. It's identical to ManagerIDEQ.
func ManagerID(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldManagerID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldStatus, vs...))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldContainsFold(FieldCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldContainsFold(FieldDescription, v))
}

// RevenueGoalEQ applies the EQ predicate on the "revenue_goal" field.
func RevenueGoalEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldRevenueGoal, v))
}

// RevenueGoalNEQ applies the NEQ predicate on the "revenue_goal" field.
func RevenueGoalNEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldRevenueGoal, v))
}

// RevenueGoalIn applies the In predicate on the "revenue_goal" field.
func RevenueGoalIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldRevenueGoal, vs...))
}

// RevenueGoalNotIn applies the NotIn predicate on the "revenue_goal" field.
func RevenueGoalNotIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldRevenueGoal, vs...))
}

// RevenueGoalGT applies the GT predicate on the "revenue_goal" field.
func RevenueGoalGT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldRevenueGoal, v))
}

// RevenueGoalGTE applies the GTE predicate on the "revenue_goal" field.
func RevenueGoalGTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldRevenueGoal, v))
}

// RevenueGoalLT applies the LT predicate on the "revenue_goal" field.
func RevenueGoalLT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldRevenueGoal, v))
}

// RevenueGoalLTE applies the LTE predicate on the "revenue_goal" field.
func RevenueGoalLTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldRevenueGoal, v))
}

// RevenueGoalIsNil applies the IsNil predicate on the "revenue_goal" field.
func RevenueGoalIsNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIsNull(FieldRevenueGoal))
}

// RevenueGoalNotNil applies the NotNil predicate on the "revenue_goal" field.
func RevenueGoalNotNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotNull(FieldRevenueGoal))
}

// DeadheadGoalEQ applies the EQ predicate on the "deadhead_goal" field.
func DeadheadGoalEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldDeadheadGoal, v))
}

// DeadheadGoalNEQ applies the NEQ predicate on the "deadhead_goal" field.
func DeadheadGoalNEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldDeadheadGoal, v))
}

// DeadheadGoalIn applies the In predicate on the "deadhead_goal" field.
func DeadheadGoalIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldDeadheadGoal, vs...))
}

// DeadheadGoalNotIn applies the NotIn predicate on the "deadhead_goal" field.
func DeadheadGoalNotIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldDeadheadGoal, vs...))
}

// DeadheadGoalGT applies the GT predicate on the "deadhead_goal" field.
func DeadheadGoalGT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldDeadheadGoal, v))
}

// DeadheadGoalGTE applies the GTE predicate on the "deadhead_goal" field.
func DeadheadGoalGTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldDeadheadGoal, v))
}

// DeadheadGoalLT applies the LT predicate on the "deadhead_goal" field.
func DeadheadGoalLT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldDeadheadGoal, v))
}

// DeadheadGoalLTE applies the LTE predicate on the "deadhead_goal" field.
func DeadheadGoalLTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldDeadheadGoal, v))
}

// DeadheadGoalIsNil applies the IsNil predicate on the "deadhead_goal" field.
func DeadheadGoalIsNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIsNull(FieldDeadheadGoal))
}

// DeadheadGoalNotNil applies the NotNil predicate on the "deadhead_goal" field.
func DeadheadGoalNotNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotNull(FieldDeadheadGoal))
}

// MileageGoalEQ applies the EQ predicate on the "mileage_goal" field.
func MileageGoalEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldMileageGoal, v))
}

// MileageGoalNEQ applies the NEQ predicate on the "mileage_goal" field.
func MileageGoalNEQ(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldMileageGoal, v))
}

// MileageGoalIn applies the In predicate on the "mileage_goal" field.
func MileageGoalIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldMileageGoal, vs...))
}

// MileageGoalNotIn applies the NotIn predicate on the "mileage_goal" field.
func MileageGoalNotIn(vs ...float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldMileageGoal, vs...))
}

// MileageGoalGT applies the GT predicate on the "mileage_goal" field.
func MileageGoalGT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGT(FieldMileageGoal, v))
}

// MileageGoalGTE applies the GTE predicate on the "mileage_goal" field.
func MileageGoalGTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldGTE(FieldMileageGoal, v))
}

// MileageGoalLT applies the LT predicate on the "mileage_goal" field.
func MileageGoalLT(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLT(FieldMileageGoal, v))
}

// MileageGoalLTE applies the LTE predicate on the "mileage_goal" field.
func MileageGoalLTE(v float64) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldLTE(FieldMileageGoal, v))
}

// MileageGoalIsNil applies the IsNil predicate on the "mileage_goal" field.
func MileageGoalIsNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIsNull(FieldMileageGoal))
}

// MileageGoalNotNil applies the NotNil predicate on the "mileage_goal" field.
func MileageGoalNotNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotNull(FieldMileageGoal))
}

// ManagerIDEQ applies the EQ predicate on the "manager_id" field.
func ManagerIDEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldEQ(FieldManagerID, v))
}

// ManagerIDNEQ applies the NEQ predicate on the "manager_id" field.
func ManagerIDNEQ(v uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNEQ(FieldManagerID, v))
}

// ManagerIDIn applies the In predicate on the "manager_id" field.
func ManagerIDIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIn(FieldManagerID, vs...))
}

// ManagerIDNotIn applies the NotIn predicate on the "manager_id" field.
func ManagerIDNotIn(vs ...uuid.UUID) predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotIn(FieldManagerID, vs...))
}

// ManagerIDIsNil applies the IsNil predicate on the "manager_id" field.
func ManagerIDIsNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldIsNull(FieldManagerID))
}

// ManagerIDNotNil applies the NotNil predicate on the "manager_id" field.
func ManagerIDNotNil() predicate.FleetCode {
	return predicate.FleetCode(sql.FieldNotNull(FieldManagerID))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasManager applies the HasEdge predicate on the "manager" edge.
func HasManager() predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ManagerTable, ManagerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasManagerWith applies the HasEdge predicate on the "manager" edge with a given conditions (other predicates).
func HasManagerWith(preds ...predicate.User) predicate.FleetCode {
	return predicate.FleetCode(func(s *sql.Selector) {
		step := newManagerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FleetCode) predicate.FleetCode {
	return predicate.FleetCode(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FleetCode) predicate.FleetCode {
	return predicate.FleetCode(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FleetCode) predicate.FleetCode {
	return predicate.FleetCode(sql.NotPredicates(p))
}