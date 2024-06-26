// Code generated by entc, DO NOT EDIT.

package accessorialcharge

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldVersion, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldDescription, v))
}

// IsDetention applies equality check predicate on the "is_detention" field. It's identical to IsDetentionEQ.
func IsDetention(v bool) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldIsDetention, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldAmount, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldStatus, vs...))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldContainsFold(FieldCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldContainsFold(FieldDescription, v))
}

// IsDetentionEQ applies the EQ predicate on the "is_detention" field.
func IsDetentionEQ(v bool) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldIsDetention, v))
}

// IsDetentionNEQ applies the NEQ predicate on the "is_detention" field.
func IsDetentionNEQ(v bool) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldIsDetention, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v Method) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v Method) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...Method) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...Method) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldMethod, vs...))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.FieldLTE(FieldAmount, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipmentCharges applies the HasEdge predicate on the "shipment_charges" edge.
func HasShipmentCharges() predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShipmentChargesTable, ShipmentChargesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentChargesWith applies the HasEdge predicate on the "shipment_charges" edge with a given conditions (other predicates).
func HasShipmentChargesWith(preds ...predicate.ShipmentCharges) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(func(s *sql.Selector) {
		step := newShipmentChargesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessorialCharge) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessorialCharge) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessorialCharge) predicate.AccessorialCharge {
	return predicate.AccessorialCharge(sql.NotPredicates(p))
}
