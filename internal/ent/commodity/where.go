// Code generated by entc, DO NOT EDIT.

package commodity

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/internal/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldVersion, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldName, v))
}

// IsHazmat applies equality check predicate on the "is_hazmat" field. It's identical to IsHazmatEQ.
func IsHazmat(v bool) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldIsHazmat, v))
}

// UnitOfMeasure applies equality check predicate on the "unit_of_measure" field. It's identical to UnitOfMeasureEQ.
func UnitOfMeasure(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldUnitOfMeasure, v))
}

// MinTemp applies equality check predicate on the "min_temp" field. It's identical to MinTempEQ.
func MinTemp(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldMinTemp, v))
}

// MaxTemp applies equality check predicate on the "max_temp" field. It's identical to MaxTempEQ.
func MaxTemp(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldMaxTemp, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldDescription, v))
}

// HazardousMaterialID applies equality check predicate on the "hazardous_material_id" field. It's identical to HazardousMaterialIDEQ.
func HazardousMaterialID(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldHazardousMaterialID, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldStatus, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContainsFold(FieldName, v))
}

// IsHazmatEQ applies the EQ predicate on the "is_hazmat" field.
func IsHazmatEQ(v bool) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldIsHazmat, v))
}

// IsHazmatNEQ applies the NEQ predicate on the "is_hazmat" field.
func IsHazmatNEQ(v bool) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldIsHazmat, v))
}

// UnitOfMeasureEQ applies the EQ predicate on the "unit_of_measure" field.
func UnitOfMeasureEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldUnitOfMeasure, v))
}

// UnitOfMeasureNEQ applies the NEQ predicate on the "unit_of_measure" field.
func UnitOfMeasureNEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldUnitOfMeasure, v))
}

// UnitOfMeasureIn applies the In predicate on the "unit_of_measure" field.
func UnitOfMeasureIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldUnitOfMeasure, vs...))
}

// UnitOfMeasureNotIn applies the NotIn predicate on the "unit_of_measure" field.
func UnitOfMeasureNotIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldUnitOfMeasure, vs...))
}

// UnitOfMeasureGT applies the GT predicate on the "unit_of_measure" field.
func UnitOfMeasureGT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldUnitOfMeasure, v))
}

// UnitOfMeasureGTE applies the GTE predicate on the "unit_of_measure" field.
func UnitOfMeasureGTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldUnitOfMeasure, v))
}

// UnitOfMeasureLT applies the LT predicate on the "unit_of_measure" field.
func UnitOfMeasureLT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldUnitOfMeasure, v))
}

// UnitOfMeasureLTE applies the LTE predicate on the "unit_of_measure" field.
func UnitOfMeasureLTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldUnitOfMeasure, v))
}

// UnitOfMeasureContains applies the Contains predicate on the "unit_of_measure" field.
func UnitOfMeasureContains(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContains(FieldUnitOfMeasure, v))
}

// UnitOfMeasureHasPrefix applies the HasPrefix predicate on the "unit_of_measure" field.
func UnitOfMeasureHasPrefix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasPrefix(FieldUnitOfMeasure, v))
}

// UnitOfMeasureHasSuffix applies the HasSuffix predicate on the "unit_of_measure" field.
func UnitOfMeasureHasSuffix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasSuffix(FieldUnitOfMeasure, v))
}

// UnitOfMeasureIsNil applies the IsNil predicate on the "unit_of_measure" field.
func UnitOfMeasureIsNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldIsNull(FieldUnitOfMeasure))
}

// UnitOfMeasureNotNil applies the NotNil predicate on the "unit_of_measure" field.
func UnitOfMeasureNotNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldNotNull(FieldUnitOfMeasure))
}

// UnitOfMeasureEqualFold applies the EqualFold predicate on the "unit_of_measure" field.
func UnitOfMeasureEqualFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEqualFold(FieldUnitOfMeasure, v))
}

// UnitOfMeasureContainsFold applies the ContainsFold predicate on the "unit_of_measure" field.
func UnitOfMeasureContainsFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContainsFold(FieldUnitOfMeasure, v))
}

// MinTempEQ applies the EQ predicate on the "min_temp" field.
func MinTempEQ(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldMinTemp, v))
}

// MinTempNEQ applies the NEQ predicate on the "min_temp" field.
func MinTempNEQ(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldMinTemp, v))
}

// MinTempIn applies the In predicate on the "min_temp" field.
func MinTempIn(vs ...int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldMinTemp, vs...))
}

// MinTempNotIn applies the NotIn predicate on the "min_temp" field.
func MinTempNotIn(vs ...int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldMinTemp, vs...))
}

// MinTempGT applies the GT predicate on the "min_temp" field.
func MinTempGT(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldMinTemp, v))
}

// MinTempGTE applies the GTE predicate on the "min_temp" field.
func MinTempGTE(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldMinTemp, v))
}

// MinTempLT applies the LT predicate on the "min_temp" field.
func MinTempLT(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldMinTemp, v))
}

// MinTempLTE applies the LTE predicate on the "min_temp" field.
func MinTempLTE(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldMinTemp, v))
}

// MinTempIsNil applies the IsNil predicate on the "min_temp" field.
func MinTempIsNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldIsNull(FieldMinTemp))
}

// MinTempNotNil applies the NotNil predicate on the "min_temp" field.
func MinTempNotNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldNotNull(FieldMinTemp))
}

// MaxTempEQ applies the EQ predicate on the "max_temp" field.
func MaxTempEQ(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldMaxTemp, v))
}

// MaxTempNEQ applies the NEQ predicate on the "max_temp" field.
func MaxTempNEQ(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldMaxTemp, v))
}

// MaxTempIn applies the In predicate on the "max_temp" field.
func MaxTempIn(vs ...int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldMaxTemp, vs...))
}

// MaxTempNotIn applies the NotIn predicate on the "max_temp" field.
func MaxTempNotIn(vs ...int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldMaxTemp, vs...))
}

// MaxTempGT applies the GT predicate on the "max_temp" field.
func MaxTempGT(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldMaxTemp, v))
}

// MaxTempGTE applies the GTE predicate on the "max_temp" field.
func MaxTempGTE(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldMaxTemp, v))
}

// MaxTempLT applies the LT predicate on the "max_temp" field.
func MaxTempLT(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldMaxTemp, v))
}

// MaxTempLTE applies the LTE predicate on the "max_temp" field.
func MaxTempLTE(v int8) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldMaxTemp, v))
}

// MaxTempIsNil applies the IsNil predicate on the "max_temp" field.
func MaxTempIsNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldIsNull(FieldMaxTemp))
}

// MaxTempNotNil applies the NotNil predicate on the "max_temp" field.
func MaxTempNotNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldNotNull(FieldMaxTemp))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Commodity {
	return predicate.Commodity(sql.FieldContainsFold(FieldDescription, v))
}

// HazardousMaterialIDEQ applies the EQ predicate on the "hazardous_material_id" field.
func HazardousMaterialIDEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldEQ(FieldHazardousMaterialID, v))
}

// HazardousMaterialIDNEQ applies the NEQ predicate on the "hazardous_material_id" field.
func HazardousMaterialIDNEQ(v uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNEQ(FieldHazardousMaterialID, v))
}

// HazardousMaterialIDIn applies the In predicate on the "hazardous_material_id" field.
func HazardousMaterialIDIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldIn(FieldHazardousMaterialID, vs...))
}

// HazardousMaterialIDNotIn applies the NotIn predicate on the "hazardous_material_id" field.
func HazardousMaterialIDNotIn(vs ...uuid.UUID) predicate.Commodity {
	return predicate.Commodity(sql.FieldNotIn(FieldHazardousMaterialID, vs...))
}

// HazardousMaterialIDIsNil applies the IsNil predicate on the "hazardous_material_id" field.
func HazardousMaterialIDIsNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldIsNull(FieldHazardousMaterialID))
}

// HazardousMaterialIDNotNil applies the NotNil predicate on the "hazardous_material_id" field.
func HazardousMaterialIDNotNil() predicate.Commodity {
	return predicate.Commodity(sql.FieldNotNull(FieldHazardousMaterialID))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHazardousMaterial applies the HasEdge predicate on the "hazardous_material" edge.
func HasHazardousMaterial() predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, HazardousMaterialTable, HazardousMaterialColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHazardousMaterialWith applies the HasEdge predicate on the "hazardous_material" edge with a given conditions (other predicates).
func HasHazardousMaterialWith(preds ...predicate.HazardousMaterial) predicate.Commodity {
	return predicate.Commodity(func(s *sql.Selector) {
		step := newHazardousMaterialStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Commodity) predicate.Commodity {
	return predicate.Commodity(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Commodity) predicate.Commodity {
	return predicate.Commodity(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Commodity) predicate.Commodity {
	return predicate.Commodity(sql.NotPredicates(p))
}
