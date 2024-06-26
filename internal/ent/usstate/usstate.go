// Code generated by entc, DO NOT EDIT.

package usstate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the usstate type in the database.
	Label = "us_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAbbreviation holds the string denoting the abbreviation field in the database.
	FieldAbbreviation = "abbreviation"
	// FieldCountryName holds the string denoting the country_name field in the database.
	FieldCountryName = "country_name"
	// FieldCountryIso3 holds the string denoting the country_iso3 field in the database.
	FieldCountryIso3 = "country_iso3"
	// Table holds the table name of the usstate in the database.
	Table = "us_states"
)

// Columns holds all SQL columns for usstate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldAbbreviation,
	FieldCountryName,
	FieldCountryIso3,
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
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AbbreviationValidator is a validator for the "abbreviation" field. It is called by the builders before save.
	AbbreviationValidator func(string) error
	// DefaultCountryName holds the default value on creation for the "country_name" field.
	DefaultCountryName string
	// CountryNameValidator is a validator for the "country_name" field. It is called by the builders before save.
	CountryNameValidator func(string) error
	// DefaultCountryIso3 holds the default value on creation for the "country_iso3" field.
	DefaultCountryIso3 string
	// CountryIso3Validator is a validator for the "country_iso3" field. It is called by the builders before save.
	CountryIso3Validator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the UsState queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAbbreviation orders the results by the abbreviation field.
func ByAbbreviation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAbbreviation, opts...).ToFunc()
}

// ByCountryName orders the results by the country_name field.
func ByCountryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryName, opts...).ToFunc()
}

// ByCountryIso3 orders the results by the country_iso3 field.
func ByCountryIso3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountryIso3, opts...).ToFunc()
}
