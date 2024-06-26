// Code generated by entc, DO NOT EDIT.

package hazardousmaterial

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldVersion, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldName, v))
}

// ErgNumber applies equality check predicate on the "erg_number" field. It's identical to ErgNumberEQ.
func ErgNumber(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldErgNumber, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldDescription, v))
}

// PackingGroup applies equality check predicate on the "packing_group" field. It's identical to PackingGroupEQ.
func PackingGroup(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldPackingGroup, v))
}

// ProperShippingName applies equality check predicate on the "proper_shipping_name" field. It's identical to ProperShippingNameEQ.
func ProperShippingName(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldProperShippingName, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContainsFold(FieldName, v))
}

// HazardClassEQ applies the EQ predicate on the "hazard_class" field.
func HazardClassEQ(v HazardClass) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldHazardClass, v))
}

// HazardClassNEQ applies the NEQ predicate on the "hazard_class" field.
func HazardClassNEQ(v HazardClass) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldHazardClass, v))
}

// HazardClassIn applies the In predicate on the "hazard_class" field.
func HazardClassIn(vs ...HazardClass) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldHazardClass, vs...))
}

// HazardClassNotIn applies the NotIn predicate on the "hazard_class" field.
func HazardClassNotIn(vs ...HazardClass) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldHazardClass, vs...))
}

// ErgNumberEQ applies the EQ predicate on the "erg_number" field.
func ErgNumberEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldErgNumber, v))
}

// ErgNumberNEQ applies the NEQ predicate on the "erg_number" field.
func ErgNumberNEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldErgNumber, v))
}

// ErgNumberIn applies the In predicate on the "erg_number" field.
func ErgNumberIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldErgNumber, vs...))
}

// ErgNumberNotIn applies the NotIn predicate on the "erg_number" field.
func ErgNumberNotIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldErgNumber, vs...))
}

// ErgNumberGT applies the GT predicate on the "erg_number" field.
func ErgNumberGT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldErgNumber, v))
}

// ErgNumberGTE applies the GTE predicate on the "erg_number" field.
func ErgNumberGTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldErgNumber, v))
}

// ErgNumberLT applies the LT predicate on the "erg_number" field.
func ErgNumberLT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldErgNumber, v))
}

// ErgNumberLTE applies the LTE predicate on the "erg_number" field.
func ErgNumberLTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldErgNumber, v))
}

// ErgNumberContains applies the Contains predicate on the "erg_number" field.
func ErgNumberContains(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContains(FieldErgNumber, v))
}

// ErgNumberHasPrefix applies the HasPrefix predicate on the "erg_number" field.
func ErgNumberHasPrefix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasPrefix(FieldErgNumber, v))
}

// ErgNumberHasSuffix applies the HasSuffix predicate on the "erg_number" field.
func ErgNumberHasSuffix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasSuffix(FieldErgNumber, v))
}

// ErgNumberIsNil applies the IsNil predicate on the "erg_number" field.
func ErgNumberIsNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIsNull(FieldErgNumber))
}

// ErgNumberNotNil applies the NotNil predicate on the "erg_number" field.
func ErgNumberNotNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotNull(FieldErgNumber))
}

// ErgNumberEqualFold applies the EqualFold predicate on the "erg_number" field.
func ErgNumberEqualFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEqualFold(FieldErgNumber, v))
}

// ErgNumberContainsFold applies the ContainsFold predicate on the "erg_number" field.
func ErgNumberContainsFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContainsFold(FieldErgNumber, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContainsFold(FieldDescription, v))
}

// PackingGroupEQ applies the EQ predicate on the "packing_group" field.
func PackingGroupEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldPackingGroup, v))
}

// PackingGroupNEQ applies the NEQ predicate on the "packing_group" field.
func PackingGroupNEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldPackingGroup, v))
}

// PackingGroupIn applies the In predicate on the "packing_group" field.
func PackingGroupIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldPackingGroup, vs...))
}

// PackingGroupNotIn applies the NotIn predicate on the "packing_group" field.
func PackingGroupNotIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldPackingGroup, vs...))
}

// PackingGroupGT applies the GT predicate on the "packing_group" field.
func PackingGroupGT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldPackingGroup, v))
}

// PackingGroupGTE applies the GTE predicate on the "packing_group" field.
func PackingGroupGTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldPackingGroup, v))
}

// PackingGroupLT applies the LT predicate on the "packing_group" field.
func PackingGroupLT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldPackingGroup, v))
}

// PackingGroupLTE applies the LTE predicate on the "packing_group" field.
func PackingGroupLTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldPackingGroup, v))
}

// PackingGroupContains applies the Contains predicate on the "packing_group" field.
func PackingGroupContains(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContains(FieldPackingGroup, v))
}

// PackingGroupHasPrefix applies the HasPrefix predicate on the "packing_group" field.
func PackingGroupHasPrefix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasPrefix(FieldPackingGroup, v))
}

// PackingGroupHasSuffix applies the HasSuffix predicate on the "packing_group" field.
func PackingGroupHasSuffix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasSuffix(FieldPackingGroup, v))
}

// PackingGroupIsNil applies the IsNil predicate on the "packing_group" field.
func PackingGroupIsNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIsNull(FieldPackingGroup))
}

// PackingGroupNotNil applies the NotNil predicate on the "packing_group" field.
func PackingGroupNotNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotNull(FieldPackingGroup))
}

// PackingGroupEqualFold applies the EqualFold predicate on the "packing_group" field.
func PackingGroupEqualFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEqualFold(FieldPackingGroup, v))
}

// PackingGroupContainsFold applies the ContainsFold predicate on the "packing_group" field.
func PackingGroupContainsFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContainsFold(FieldPackingGroup, v))
}

// ProperShippingNameEQ applies the EQ predicate on the "proper_shipping_name" field.
func ProperShippingNameEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEQ(FieldProperShippingName, v))
}

// ProperShippingNameNEQ applies the NEQ predicate on the "proper_shipping_name" field.
func ProperShippingNameNEQ(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNEQ(FieldProperShippingName, v))
}

// ProperShippingNameIn applies the In predicate on the "proper_shipping_name" field.
func ProperShippingNameIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIn(FieldProperShippingName, vs...))
}

// ProperShippingNameNotIn applies the NotIn predicate on the "proper_shipping_name" field.
func ProperShippingNameNotIn(vs ...string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotIn(FieldProperShippingName, vs...))
}

// ProperShippingNameGT applies the GT predicate on the "proper_shipping_name" field.
func ProperShippingNameGT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGT(FieldProperShippingName, v))
}

// ProperShippingNameGTE applies the GTE predicate on the "proper_shipping_name" field.
func ProperShippingNameGTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldGTE(FieldProperShippingName, v))
}

// ProperShippingNameLT applies the LT predicate on the "proper_shipping_name" field.
func ProperShippingNameLT(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLT(FieldProperShippingName, v))
}

// ProperShippingNameLTE applies the LTE predicate on the "proper_shipping_name" field.
func ProperShippingNameLTE(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldLTE(FieldProperShippingName, v))
}

// ProperShippingNameContains applies the Contains predicate on the "proper_shipping_name" field.
func ProperShippingNameContains(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContains(FieldProperShippingName, v))
}

// ProperShippingNameHasPrefix applies the HasPrefix predicate on the "proper_shipping_name" field.
func ProperShippingNameHasPrefix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasPrefix(FieldProperShippingName, v))
}

// ProperShippingNameHasSuffix applies the HasSuffix predicate on the "proper_shipping_name" field.
func ProperShippingNameHasSuffix(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldHasSuffix(FieldProperShippingName, v))
}

// ProperShippingNameIsNil applies the IsNil predicate on the "proper_shipping_name" field.
func ProperShippingNameIsNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldIsNull(FieldProperShippingName))
}

// ProperShippingNameNotNil applies the NotNil predicate on the "proper_shipping_name" field.
func ProperShippingNameNotNil() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldNotNull(FieldProperShippingName))
}

// ProperShippingNameEqualFold applies the EqualFold predicate on the "proper_shipping_name" field.
func ProperShippingNameEqualFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldEqualFold(FieldProperShippingName, v))
}

// ProperShippingNameContainsFold applies the ContainsFold predicate on the "proper_shipping_name" field.
func ProperShippingNameContainsFold(v string) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.FieldContainsFold(FieldProperShippingName, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.HazardousMaterial {
	return predicate.HazardousMaterial(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HazardousMaterial) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HazardousMaterial) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HazardousMaterial) predicate.HazardousMaterial {
	return predicate.HazardousMaterial(sql.NotPredicates(p))
}
