// Code generated by entc, DO NOT EDIT.

package shipmentdocumentation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldVersion, v))
}

// ShipmentID applies equality check predicate on the "shipment_id" field. It's identical to ShipmentIDEQ.
func ShipmentID(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldShipmentID, v))
}

// DocumentURL applies equality check predicate on the "document_url" field. It's identical to DocumentURLEQ.
func DocumentURL(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldDocumentURL, v))
}

// DocumentClassificationID applies equality check predicate on the "document_classification_id" field. It's identical to DocumentClassificationIDEQ.
func DocumentClassificationID(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldDocumentClassificationID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLTE(FieldVersion, v))
}

// ShipmentIDEQ applies the EQ predicate on the "shipment_id" field.
func ShipmentIDEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldShipmentID, v))
}

// ShipmentIDNEQ applies the NEQ predicate on the "shipment_id" field.
func ShipmentIDNEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldShipmentID, v))
}

// ShipmentIDIn applies the In predicate on the "shipment_id" field.
func ShipmentIDIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldShipmentID, vs...))
}

// ShipmentIDNotIn applies the NotIn predicate on the "shipment_id" field.
func ShipmentIDNotIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldShipmentID, vs...))
}

// DocumentURLEQ applies the EQ predicate on the "document_url" field.
func DocumentURLEQ(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldDocumentURL, v))
}

// DocumentURLNEQ applies the NEQ predicate on the "document_url" field.
func DocumentURLNEQ(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldDocumentURL, v))
}

// DocumentURLIn applies the In predicate on the "document_url" field.
func DocumentURLIn(vs ...string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldDocumentURL, vs...))
}

// DocumentURLNotIn applies the NotIn predicate on the "document_url" field.
func DocumentURLNotIn(vs ...string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldDocumentURL, vs...))
}

// DocumentURLGT applies the GT predicate on the "document_url" field.
func DocumentURLGT(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGT(FieldDocumentURL, v))
}

// DocumentURLGTE applies the GTE predicate on the "document_url" field.
func DocumentURLGTE(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldGTE(FieldDocumentURL, v))
}

// DocumentURLLT applies the LT predicate on the "document_url" field.
func DocumentURLLT(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLT(FieldDocumentURL, v))
}

// DocumentURLLTE applies the LTE predicate on the "document_url" field.
func DocumentURLLTE(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldLTE(FieldDocumentURL, v))
}

// DocumentURLContains applies the Contains predicate on the "document_url" field.
func DocumentURLContains(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldContains(FieldDocumentURL, v))
}

// DocumentURLHasPrefix applies the HasPrefix predicate on the "document_url" field.
func DocumentURLHasPrefix(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldHasPrefix(FieldDocumentURL, v))
}

// DocumentURLHasSuffix applies the HasSuffix predicate on the "document_url" field.
func DocumentURLHasSuffix(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldHasSuffix(FieldDocumentURL, v))
}

// DocumentURLEqualFold applies the EqualFold predicate on the "document_url" field.
func DocumentURLEqualFold(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEqualFold(FieldDocumentURL, v))
}

// DocumentURLContainsFold applies the ContainsFold predicate on the "document_url" field.
func DocumentURLContainsFold(v string) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldContainsFold(FieldDocumentURL, v))
}

// DocumentClassificationIDEQ applies the EQ predicate on the "document_classification_id" field.
func DocumentClassificationIDEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldEQ(FieldDocumentClassificationID, v))
}

// DocumentClassificationIDNEQ applies the NEQ predicate on the "document_classification_id" field.
func DocumentClassificationIDNEQ(v uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNEQ(FieldDocumentClassificationID, v))
}

// DocumentClassificationIDIn applies the In predicate on the "document_classification_id" field.
func DocumentClassificationIDIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldIn(FieldDocumentClassificationID, vs...))
}

// DocumentClassificationIDNotIn applies the NotIn predicate on the "document_classification_id" field.
func DocumentClassificationIDNotIn(vs ...uuid.UUID) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.FieldNotIn(FieldDocumentClassificationID, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipment applies the HasEdge predicate on the "shipment" edge.
func HasShipment() predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShipmentTable, ShipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentWith applies the HasEdge predicate on the "shipment" edge with a given conditions (other predicates).
func HasShipmentWith(preds ...predicate.Shipment) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := newShipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDocumentClassification applies the HasEdge predicate on the "document_classification" edge.
func HasDocumentClassification() predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DocumentClassificationTable, DocumentClassificationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentClassificationWith applies the HasEdge predicate on the "document_classification" edge with a given conditions (other predicates).
func HasDocumentClassificationWith(preds ...predicate.DocumentClassification) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(func(s *sql.Selector) {
		step := newDocumentClassificationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ShipmentDocumentation) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ShipmentDocumentation) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ShipmentDocumentation) predicate.ShipmentDocumentation {
	return predicate.ShipmentDocumentation(sql.NotPredicates(p))
}