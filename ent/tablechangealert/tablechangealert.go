// Code generated by ent, DO NOT EDIT.

package tablechangealert

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tablechangealert type in the database.
	Label = "table_change_alert"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBusinessUnitID holds the string denoting the business_unit_id field in the database.
	FieldBusinessUnitID = "business_unit_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDatabaseAction holds the string denoting the database_action field in the database.
	FieldDatabaseAction = "database_action"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldTableName holds the string denoting the table_name field in the database.
	FieldTableName = "table_name"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCustomSubject holds the string denoting the custom_subject field in the database.
	FieldCustomSubject = "custom_subject"
	// FieldFunctionName holds the string denoting the function_name field in the database.
	FieldFunctionName = "function_name"
	// FieldTriggerName holds the string denoting the trigger_name field in the database.
	FieldTriggerName = "trigger_name"
	// FieldListenerName holds the string denoting the listener_name field in the database.
	FieldListenerName = "listener_name"
	// FieldEmailRecipients holds the string denoting the email_recipients field in the database.
	FieldEmailRecipients = "email_recipients"
	// FieldConditionalLogic holds the string denoting the conditional_logic field in the database.
	FieldConditionalLogic = "conditional_logic"
	// FieldEffectiveDate holds the string denoting the effective_date field in the database.
	FieldEffectiveDate = "effective_date"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the tablechangealert in the database.
	Table = "table_change_alerts"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "table_change_alerts"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "table_change_alerts"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
)

// Columns holds all SQL columns for tablechangealert fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldDatabaseAction,
	FieldSource,
	FieldTableName,
	FieldTopic,
	FieldDescription,
	FieldCustomSubject,
	FieldFunctionName,
	FieldTriggerName,
	FieldListenerName,
	FieldEmailRecipients,
	FieldConditionalLogic,
	FieldEffectiveDate,
	FieldExpirationDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TableNameValidator is a validator for the "table_name" field. It is called by the builders before save.
	TableNameValidator func(string) error
	// TopicValidator is a validator for the "topic" field. It is called by the builders before save.
	TopicValidator func(string) error
	// CustomSubjectValidator is a validator for the "custom_subject" field. It is called by the builders before save.
	CustomSubjectValidator func(string) error
	// FunctionNameValidator is a validator for the "function_name" field. It is called by the builders before save.
	FunctionNameValidator func(string) error
	// TriggerNameValidator is a validator for the "trigger_name" field. It is called by the builders before save.
	TriggerNameValidator func(string) error
	// ListenerNameValidator is a validator for the "listener_name" field. It is called by the builders before save.
	ListenerNameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusA is the default value of the Status enum.
const DefaultStatus = StatusA

// Status values.
const (
	StatusA Status = "A"
	StatusI Status = "I"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusA, StatusI:
		return nil
	default:
		return fmt.Errorf("tablechangealert: invalid enum value for status field: %q", s)
	}
}

// DatabaseAction defines the type for the "database_action" enum field.
type DatabaseAction string

// DatabaseAction values.
const (
	DatabaseActionInsert DatabaseAction = "Insert"
	DatabaseActionUpdate DatabaseAction = "Update"
	DatabaseActionDelete DatabaseAction = "Delete"
	DatabaseActionAll    DatabaseAction = "All"
)

func (da DatabaseAction) String() string {
	return string(da)
}

// DatabaseActionValidator is a validator for the "database_action" field enum values. It is called by the builders before save.
func DatabaseActionValidator(da DatabaseAction) error {
	switch da {
	case DatabaseActionInsert, DatabaseActionUpdate, DatabaseActionDelete, DatabaseActionAll:
		return nil
	default:
		return fmt.Errorf("tablechangealert: invalid enum value for database_action field: %q", da)
	}
}

// Source defines the type for the "source" enum field.
type Source string

// Source values.
const (
	SourceKafka Source = "Kafka"
	SourceDb    Source = "Db"
)

func (s Source) String() string {
	return string(s)
}

// SourceValidator is a validator for the "source" field enum values. It is called by the builders before save.
func SourceValidator(s Source) error {
	switch s {
	case SourceKafka, SourceDb:
		return nil
	default:
		return fmt.Errorf("tablechangealert: invalid enum value for source field: %q", s)
	}
}

// OrderOption defines the ordering options for the TableChangeAlert queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBusinessUnitID orders the results by the business_unit_id field.
func ByBusinessUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessUnitID, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDatabaseAction orders the results by the database_action field.
func ByDatabaseAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDatabaseAction, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByTableName orders the results by the table_name field.
func ByTableName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTableName, opts...).ToFunc()
}

// ByTopic orders the results by the topic field.
func ByTopic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTopic, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCustomSubject orders the results by the custom_subject field.
func ByCustomSubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomSubject, opts...).ToFunc()
}

// ByFunctionName orders the results by the function_name field.
func ByFunctionName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFunctionName, opts...).ToFunc()
}

// ByTriggerName orders the results by the trigger_name field.
func ByTriggerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggerName, opts...).ToFunc()
}

// ByListenerName orders the results by the listener_name field.
func ByListenerName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldListenerName, opts...).ToFunc()
}

// ByEmailRecipients orders the results by the email_recipients field.
func ByEmailRecipients(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmailRecipients, opts...).ToFunc()
}

// ByEffectiveDate orders the results by the effective_date field.
func ByEffectiveDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEffectiveDate, opts...).ToFunc()
}

// ByExpirationDate orders the results by the expiration_date field.
func ByExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationDate, opts...).ToFunc()
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
	)
}
