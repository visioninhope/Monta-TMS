// Code generated by entc, DO NOT EDIT.

package session

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the session type in the database.
	Label = "session"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// Table holds the table name of the session in the database.
	Table = "sessions"
)

// Columns holds all SQL columns for session fields.
var Columns = []string{
	FieldID,
	FieldData,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldExpiresAt,
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
	// DataValidator is a validator for the "data" field. It is called by the builders before save.
	DataValidator func(string) error
)

// OrderOption defines the ordering options for the Session queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}
