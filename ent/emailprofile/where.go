// Code generated by ent, DO NOT EDIT.

package emailprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldEmail, v))
}

// Host applies equality check predicate on the "host" field. It's identical to HostEQ.
func Host(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldHost, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldPort, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldPassword, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldIsDefault, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContainsFold(FieldEmail, v))
}

// ProtocolEQ applies the EQ predicate on the "protocol" field.
func ProtocolEQ(v Protocol) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldProtocol, v))
}

// ProtocolNEQ applies the NEQ predicate on the "protocol" field.
func ProtocolNEQ(v Protocol) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldProtocol, v))
}

// ProtocolIn applies the In predicate on the "protocol" field.
func ProtocolIn(vs ...Protocol) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldProtocol, vs...))
}

// ProtocolNotIn applies the NotIn predicate on the "protocol" field.
func ProtocolNotIn(vs ...Protocol) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldProtocol, vs...))
}

// ProtocolIsNil applies the IsNil predicate on the "protocol" field.
func ProtocolIsNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIsNull(FieldProtocol))
}

// ProtocolNotNil applies the NotNil predicate on the "protocol" field.
func ProtocolNotNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotNull(FieldProtocol))
}

// HostEQ applies the EQ predicate on the "host" field.
func HostEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldHost, v))
}

// HostNEQ applies the NEQ predicate on the "host" field.
func HostNEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldHost, v))
}

// HostIn applies the In predicate on the "host" field.
func HostIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldHost, vs...))
}

// HostNotIn applies the NotIn predicate on the "host" field.
func HostNotIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldHost, vs...))
}

// HostGT applies the GT predicate on the "host" field.
func HostGT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldHost, v))
}

// HostGTE applies the GTE predicate on the "host" field.
func HostGTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldHost, v))
}

// HostLT applies the LT predicate on the "host" field.
func HostLT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldHost, v))
}

// HostLTE applies the LTE predicate on the "host" field.
func HostLTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldHost, v))
}

// HostContains applies the Contains predicate on the "host" field.
func HostContains(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContains(FieldHost, v))
}

// HostHasPrefix applies the HasPrefix predicate on the "host" field.
func HostHasPrefix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasPrefix(FieldHost, v))
}

// HostHasSuffix applies the HasSuffix predicate on the "host" field.
func HostHasSuffix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasSuffix(FieldHost, v))
}

// HostIsNil applies the IsNil predicate on the "host" field.
func HostIsNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIsNull(FieldHost))
}

// HostNotNil applies the NotNil predicate on the "host" field.
func HostNotNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotNull(FieldHost))
}

// HostEqualFold applies the EqualFold predicate on the "host" field.
func HostEqualFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEqualFold(FieldHost, v))
}

// HostContainsFold applies the ContainsFold predicate on the "host" field.
func HostContainsFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContainsFold(FieldHost, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int16) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldPort, v))
}

// PortIsNil applies the IsNil predicate on the "port" field.
func PortIsNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIsNull(FieldPort))
}

// PortNotNil applies the NotNil predicate on the "port" field.
func PortNotNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotNull(FieldPort))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldContainsFold(FieldPassword, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.EmailProfile {
	return predicate.EmailProfile(sql.FieldNEQ(FieldIsDefault, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.EmailProfile {
	return predicate.EmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.EmailProfile {
	return predicate.EmailProfile(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.EmailProfile {
	return predicate.EmailProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.EmailProfile {
	return predicate.EmailProfile(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailProfile) predicate.EmailProfile {
	return predicate.EmailProfile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailProfile) predicate.EmailProfile {
	return predicate.EmailProfile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EmailProfile) predicate.EmailProfile {
	return predicate.EmailProfile(sql.NotPredicates(p))
}
