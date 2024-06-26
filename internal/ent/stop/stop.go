// Code generated by entc, DO NOT EDIT.

package stop

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the stop type in the database.
	Label = "stop"
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
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldShipmentMoveID holds the string denoting the shipment_move_id field in the database.
	FieldShipmentMoveID = "shipment_move_id"
	// FieldStopType holds the string denoting the stop_type field in the database.
	FieldStopType = "stop_type"
	// FieldSequence holds the string denoting the sequence field in the database.
	FieldSequence = "sequence"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldPieces holds the string denoting the pieces field in the database.
	FieldPieces = "pieces"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldAddressLine holds the string denoting the address_line field in the database.
	FieldAddressLine = "address_line"
	// FieldAppointmentStart holds the string denoting the appointment_start field in the database.
	FieldAppointmentStart = "appointment_start"
	// FieldAppointmentEnd holds the string denoting the appointment_end field in the database.
	FieldAppointmentEnd = "appointment_end"
	// FieldArrivalTime holds the string denoting the arrival_time field in the database.
	FieldArrivalTime = "arrival_time"
	// FieldDepartureTime holds the string denoting the departure_time field in the database.
	FieldDepartureTime = "departure_time"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeShipmentMove holds the string denoting the shipment_move edge name in mutations.
	EdgeShipmentMove = "shipment_move"
	// Table holds the table name of the stop in the database.
	Table = "stops"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "stops"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "stops"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// ShipmentMoveTable is the table that holds the shipment_move relation/edge.
	ShipmentMoveTable = "stops"
	// ShipmentMoveInverseTable is the table name for the ShipmentMove entity.
	// It exists in this package in order to avoid circular dependency with the "shipmentmove" package.
	ShipmentMoveInverseTable = "shipment_moves"
	// ShipmentMoveColumn is the table column denoting the shipment_move relation/edge.
	ShipmentMoveColumn = "shipment_move_id"
)

// Columns holds all SQL columns for stop fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldShipmentMoveID,
	FieldStopType,
	FieldSequence,
	FieldLocationID,
	FieldPieces,
	FieldWeight,
	FieldAddressLine,
	FieldAppointmentStart,
	FieldAppointmentEnd,
	FieldArrivalTime,
	FieldDepartureTime,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/emoss08/trenova/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// DefaultSequence holds the default value on creation for the "sequence" field.
	DefaultSequence int
	// SequenceValidator is a validator for the "sequence" field. It is called by the builders before save.
	SequenceValidator func(int) error
	// PiecesValidator is a validator for the "pieces" field. It is called by the builders before save.
	PiecesValidator func(float64) error
	// WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	WeightValidator func(float64) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusNew is the default value of the Status enum.
const DefaultStatus = StatusNew

// Status values.
const (
	StatusNew        Status = "New"
	StatusInProgress Status = "InProgress"
	StatusCompleted  Status = "Completed"
	StatusVoided     Status = "Voided"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusNew, StatusInProgress, StatusCompleted, StatusVoided:
		return nil
	default:
		return fmt.Errorf("stop: invalid enum value for status field: %q", s)
	}
}

// StopType defines the type for the "stop_type" enum field.
type StopType string

// StopType values.
const (
	StopTypePickup      StopType = "Pickup"
	StopTypeSplitPickup StopType = "SplitPickup"
	StopTypeSplitDrop   StopType = "SplitDrop"
	StopTypeDelivery    StopType = "Delivery"
	StopTypeDropOff     StopType = "DropOff"
)

func (st StopType) String() string {
	return string(st)
}

// StopTypeValidator is a validator for the "stop_type" field enum values. It is called by the builders before save.
func StopTypeValidator(st StopType) error {
	switch st {
	case StopTypePickup, StopTypeSplitPickup, StopTypeSplitDrop, StopTypeDelivery, StopTypeDropOff:
		return nil
	default:
		return fmt.Errorf("stop: invalid enum value for stop_type field: %q", st)
	}
}

// OrderOption defines the ordering options for the Stop queries.
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

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByShipmentMoveID orders the results by the shipment_move_id field.
func ByShipmentMoveID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShipmentMoveID, opts...).ToFunc()
}

// ByStopType orders the results by the stop_type field.
func ByStopType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStopType, opts...).ToFunc()
}

// BySequence orders the results by the sequence field.
func BySequence(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSequence, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByPieces orders the results by the pieces field.
func ByPieces(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPieces, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByAddressLine orders the results by the address_line field.
func ByAddressLine(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine, opts...).ToFunc()
}

// ByAppointmentStart orders the results by the appointment_start field.
func ByAppointmentStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppointmentStart, opts...).ToFunc()
}

// ByAppointmentEnd orders the results by the appointment_end field.
func ByAppointmentEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppointmentEnd, opts...).ToFunc()
}

// ByArrivalTime orders the results by the arrival_time field.
func ByArrivalTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArrivalTime, opts...).ToFunc()
}

// ByDepartureTime orders the results by the departure_time field.
func ByDepartureTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartureTime, opts...).ToFunc()
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

// ByShipmentMoveField orders the results by shipment_move field.
func ByShipmentMoveField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentMoveStep(), sql.OrderByField(field, opts...))
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
func newShipmentMoveStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentMoveInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ShipmentMoveTable, ShipmentMoveColumn),
	)
}
