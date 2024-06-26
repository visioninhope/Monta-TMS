// Code generated by entc, DO NOT EDIT.

package dispatchcontrol

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeadheadTarget applies equality check predicate on the "deadhead_target" field. It's identical to DeadheadTargetEQ.
func DeadheadTarget(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldDeadheadTarget, v))
}

// MaxShipmentWeightLimit applies equality check predicate on the "max_shipment_weight_limit" field. It's identical to MaxShipmentWeightLimitEQ.
func MaxShipmentWeightLimit(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldMaxShipmentWeightLimit, v))
}

// GracePeriod applies equality check predicate on the "grace_period" field. It's identical to GracePeriodEQ.
func GracePeriod(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldGracePeriod, v))
}

// EnforceWorkerAssign applies equality check predicate on the "enforce_worker_assign" field. It's identical to EnforceWorkerAssignEQ.
func EnforceWorkerAssign(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldEnforceWorkerAssign, v))
}

// TrailerContinuity applies equality check predicate on the "trailer_continuity" field. It's identical to TrailerContinuityEQ.
func TrailerContinuity(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldTrailerContinuity, v))
}

// DupeTrailerCheck applies equality check predicate on the "dupe_trailer_check" field. It's identical to DupeTrailerCheckEQ.
func DupeTrailerCheck(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldDupeTrailerCheck, v))
}

// MaintenanceCompliance applies equality check predicate on the "maintenance_compliance" field. It's identical to MaintenanceComplianceEQ.
func MaintenanceCompliance(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldMaintenanceCompliance, v))
}

// RegulatoryCheck applies equality check predicate on the "regulatory_check" field. It's identical to RegulatoryCheckEQ.
func RegulatoryCheck(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldRegulatoryCheck, v))
}

// PrevShipmentOnHold applies equality check predicate on the "prev_shipment_on_hold" field. It's identical to PrevShipmentOnHoldEQ.
func PrevShipmentOnHold(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldPrevShipmentOnHold, v))
}

// WorkerTimeAwayRestriction applies equality check predicate on the "worker_time_away_restriction" field. It's identical to WorkerTimeAwayRestrictionEQ.
func WorkerTimeAwayRestriction(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldWorkerTimeAwayRestriction, v))
}

// TractorWorkerFleetConstraint applies equality check predicate on the "tractor_worker_fleet_constraint" field. It's identical to TractorWorkerFleetConstraintEQ.
func TractorWorkerFleetConstraint(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldTractorWorkerFleetConstraint, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldUpdatedAt, v))
}

// RecordServiceIncidentEQ applies the EQ predicate on the "record_service_incident" field.
func RecordServiceIncidentEQ(v RecordServiceIncident) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldRecordServiceIncident, v))
}

// RecordServiceIncidentNEQ applies the NEQ predicate on the "record_service_incident" field.
func RecordServiceIncidentNEQ(v RecordServiceIncident) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldRecordServiceIncident, v))
}

// RecordServiceIncidentIn applies the In predicate on the "record_service_incident" field.
func RecordServiceIncidentIn(vs ...RecordServiceIncident) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldRecordServiceIncident, vs...))
}

// RecordServiceIncidentNotIn applies the NotIn predicate on the "record_service_incident" field.
func RecordServiceIncidentNotIn(vs ...RecordServiceIncident) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldRecordServiceIncident, vs...))
}

// DeadheadTargetEQ applies the EQ predicate on the "deadhead_target" field.
func DeadheadTargetEQ(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldDeadheadTarget, v))
}

// DeadheadTargetNEQ applies the NEQ predicate on the "deadhead_target" field.
func DeadheadTargetNEQ(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldDeadheadTarget, v))
}

// DeadheadTargetIn applies the In predicate on the "deadhead_target" field.
func DeadheadTargetIn(vs ...float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldDeadheadTarget, vs...))
}

// DeadheadTargetNotIn applies the NotIn predicate on the "deadhead_target" field.
func DeadheadTargetNotIn(vs ...float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldDeadheadTarget, vs...))
}

// DeadheadTargetGT applies the GT predicate on the "deadhead_target" field.
func DeadheadTargetGT(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldDeadheadTarget, v))
}

// DeadheadTargetGTE applies the GTE predicate on the "deadhead_target" field.
func DeadheadTargetGTE(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldDeadheadTarget, v))
}

// DeadheadTargetLT applies the LT predicate on the "deadhead_target" field.
func DeadheadTargetLT(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldDeadheadTarget, v))
}

// DeadheadTargetLTE applies the LTE predicate on the "deadhead_target" field.
func DeadheadTargetLTE(v float64) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldDeadheadTarget, v))
}

// MaxShipmentWeightLimitEQ applies the EQ predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitEQ(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldMaxShipmentWeightLimit, v))
}

// MaxShipmentWeightLimitNEQ applies the NEQ predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitNEQ(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldMaxShipmentWeightLimit, v))
}

// MaxShipmentWeightLimitIn applies the In predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitIn(vs ...int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldMaxShipmentWeightLimit, vs...))
}

// MaxShipmentWeightLimitNotIn applies the NotIn predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitNotIn(vs ...int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldMaxShipmentWeightLimit, vs...))
}

// MaxShipmentWeightLimitGT applies the GT predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitGT(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldMaxShipmentWeightLimit, v))
}

// MaxShipmentWeightLimitGTE applies the GTE predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitGTE(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldMaxShipmentWeightLimit, v))
}

// MaxShipmentWeightLimitLT applies the LT predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitLT(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldMaxShipmentWeightLimit, v))
}

// MaxShipmentWeightLimitLTE applies the LTE predicate on the "max_shipment_weight_limit" field.
func MaxShipmentWeightLimitLTE(v int32) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldMaxShipmentWeightLimit, v))
}

// GracePeriodEQ applies the EQ predicate on the "grace_period" field.
func GracePeriodEQ(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldGracePeriod, v))
}

// GracePeriodNEQ applies the NEQ predicate on the "grace_period" field.
func GracePeriodNEQ(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldGracePeriod, v))
}

// GracePeriodIn applies the In predicate on the "grace_period" field.
func GracePeriodIn(vs ...uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldIn(FieldGracePeriod, vs...))
}

// GracePeriodNotIn applies the NotIn predicate on the "grace_period" field.
func GracePeriodNotIn(vs ...uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNotIn(FieldGracePeriod, vs...))
}

// GracePeriodGT applies the GT predicate on the "grace_period" field.
func GracePeriodGT(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGT(FieldGracePeriod, v))
}

// GracePeriodGTE applies the GTE predicate on the "grace_period" field.
func GracePeriodGTE(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldGTE(FieldGracePeriod, v))
}

// GracePeriodLT applies the LT predicate on the "grace_period" field.
func GracePeriodLT(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLT(FieldGracePeriod, v))
}

// GracePeriodLTE applies the LTE predicate on the "grace_period" field.
func GracePeriodLTE(v uint8) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldLTE(FieldGracePeriod, v))
}

// EnforceWorkerAssignEQ applies the EQ predicate on the "enforce_worker_assign" field.
func EnforceWorkerAssignEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldEnforceWorkerAssign, v))
}

// EnforceWorkerAssignNEQ applies the NEQ predicate on the "enforce_worker_assign" field.
func EnforceWorkerAssignNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldEnforceWorkerAssign, v))
}

// TrailerContinuityEQ applies the EQ predicate on the "trailer_continuity" field.
func TrailerContinuityEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldTrailerContinuity, v))
}

// TrailerContinuityNEQ applies the NEQ predicate on the "trailer_continuity" field.
func TrailerContinuityNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldTrailerContinuity, v))
}

// DupeTrailerCheckEQ applies the EQ predicate on the "dupe_trailer_check" field.
func DupeTrailerCheckEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldDupeTrailerCheck, v))
}

// DupeTrailerCheckNEQ applies the NEQ predicate on the "dupe_trailer_check" field.
func DupeTrailerCheckNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldDupeTrailerCheck, v))
}

// MaintenanceComplianceEQ applies the EQ predicate on the "maintenance_compliance" field.
func MaintenanceComplianceEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldMaintenanceCompliance, v))
}

// MaintenanceComplianceNEQ applies the NEQ predicate on the "maintenance_compliance" field.
func MaintenanceComplianceNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldMaintenanceCompliance, v))
}

// RegulatoryCheckEQ applies the EQ predicate on the "regulatory_check" field.
func RegulatoryCheckEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldRegulatoryCheck, v))
}

// RegulatoryCheckNEQ applies the NEQ predicate on the "regulatory_check" field.
func RegulatoryCheckNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldRegulatoryCheck, v))
}

// PrevShipmentOnHoldEQ applies the EQ predicate on the "prev_shipment_on_hold" field.
func PrevShipmentOnHoldEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldPrevShipmentOnHold, v))
}

// PrevShipmentOnHoldNEQ applies the NEQ predicate on the "prev_shipment_on_hold" field.
func PrevShipmentOnHoldNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldPrevShipmentOnHold, v))
}

// WorkerTimeAwayRestrictionEQ applies the EQ predicate on the "worker_time_away_restriction" field.
func WorkerTimeAwayRestrictionEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldWorkerTimeAwayRestriction, v))
}

// WorkerTimeAwayRestrictionNEQ applies the NEQ predicate on the "worker_time_away_restriction" field.
func WorkerTimeAwayRestrictionNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldWorkerTimeAwayRestriction, v))
}

// TractorWorkerFleetConstraintEQ applies the EQ predicate on the "tractor_worker_fleet_constraint" field.
func TractorWorkerFleetConstraintEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldEQ(FieldTractorWorkerFleetConstraint, v))
}

// TractorWorkerFleetConstraintNEQ applies the NEQ predicate on the "tractor_worker_fleet_constraint" field.
func TractorWorkerFleetConstraintNEQ(v bool) predicate.DispatchControl {
	return predicate.DispatchControl(sql.FieldNEQ(FieldTractorWorkerFleetConstraint, v))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.DispatchControl {
	return predicate.DispatchControl(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.DispatchControl {
	return predicate.DispatchControl(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.DispatchControl {
	return predicate.DispatchControl(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.DispatchControl {
	return predicate.DispatchControl(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DispatchControl) predicate.DispatchControl {
	return predicate.DispatchControl(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DispatchControl) predicate.DispatchControl {
	return predicate.DispatchControl(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DispatchControl) predicate.DispatchControl {
	return predicate.DispatchControl(sql.NotPredicates(p))
}
