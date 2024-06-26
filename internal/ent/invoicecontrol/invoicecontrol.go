// Code generated by entc, DO NOT EDIT.

package invoicecontrol

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invoicecontrol type in the database.
	Label = "invoice_control"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldInvoiceNumberPrefix holds the string denoting the invoice_number_prefix field in the database.
	FieldInvoiceNumberPrefix = "invoice_number_prefix"
	// FieldCreditMemoNumberPrefix holds the string denoting the credit_memo_number_prefix field in the database.
	FieldCreditMemoNumberPrefix = "credit_memo_number_prefix"
	// FieldInvoiceTerms holds the string denoting the invoice_terms field in the database.
	FieldInvoiceTerms = "invoice_terms"
	// FieldInvoiceFooter holds the string denoting the invoice_footer field in the database.
	FieldInvoiceFooter = "invoice_footer"
	// FieldInvoiceLogoURL holds the string denoting the invoice_logo_url field in the database.
	FieldInvoiceLogoURL = "invoice_logo_url"
	// FieldInvoiceDateFormat holds the string denoting the invoice_date_format field in the database.
	FieldInvoiceDateFormat = "invoice_date_format"
	// FieldInvoiceDueAfterDays holds the string denoting the invoice_due_after_days field in the database.
	FieldInvoiceDueAfterDays = "invoice_due_after_days"
	// FieldInvoiceLogoWidth holds the string denoting the invoice_logo_width field in the database.
	FieldInvoiceLogoWidth = "invoice_logo_width"
	// FieldShowAmountDue holds the string denoting the show_amount_due field in the database.
	FieldShowAmountDue = "show_amount_due"
	// FieldAttachPdf holds the string denoting the attach_pdf field in the database.
	FieldAttachPdf = "attach_pdf"
	// FieldShowInvoiceDueDate holds the string denoting the show_invoice_due_date field in the database.
	FieldShowInvoiceDueDate = "show_invoice_due_date"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// Table holds the table name of the invoicecontrol in the database.
	Table = "invoice_controls"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "invoice_controls"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "invoice_controls"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
)

// Columns holds all SQL columns for invoicecontrol fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldInvoiceNumberPrefix,
	FieldCreditMemoNumberPrefix,
	FieldInvoiceTerms,
	FieldInvoiceFooter,
	FieldInvoiceLogoURL,
	FieldInvoiceDateFormat,
	FieldInvoiceDueAfterDays,
	FieldInvoiceLogoWidth,
	FieldShowAmountDue,
	FieldAttachPdf,
	FieldShowInvoiceDueDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "invoice_controls"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_unit_id",
	"organization_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultInvoiceNumberPrefix holds the default value on creation for the "invoice_number_prefix" field.
	DefaultInvoiceNumberPrefix string
	// InvoiceNumberPrefixValidator is a validator for the "invoice_number_prefix" field. It is called by the builders before save.
	InvoiceNumberPrefixValidator func(string) error
	// DefaultCreditMemoNumberPrefix holds the default value on creation for the "credit_memo_number_prefix" field.
	DefaultCreditMemoNumberPrefix string
	// CreditMemoNumberPrefixValidator is a validator for the "credit_memo_number_prefix" field. It is called by the builders before save.
	CreditMemoNumberPrefixValidator func(string) error
	// DefaultInvoiceDueAfterDays holds the default value on creation for the "invoice_due_after_days" field.
	DefaultInvoiceDueAfterDays uint8
	// InvoiceDueAfterDaysValidator is a validator for the "invoice_due_after_days" field. It is called by the builders before save.
	InvoiceDueAfterDaysValidator func(uint8) error
	// DefaultInvoiceLogoWidth holds the default value on creation for the "invoice_logo_width" field.
	DefaultInvoiceLogoWidth uint16
	// InvoiceLogoWidthValidator is a validator for the "invoice_logo_width" field. It is called by the builders before save.
	InvoiceLogoWidthValidator func(uint16) error
	// DefaultShowAmountDue holds the default value on creation for the "show_amount_due" field.
	DefaultShowAmountDue bool
	// DefaultAttachPdf holds the default value on creation for the "attach_pdf" field.
	DefaultAttachPdf bool
	// DefaultShowInvoiceDueDate holds the default value on creation for the "show_invoice_due_date" field.
	DefaultShowInvoiceDueDate bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// InvoiceDateFormat defines the type for the "invoice_date_format" enum field.
type InvoiceDateFormat string

// InvoiceDateFormatInvoiceDateFormatMDY is the default value of the InvoiceDateFormat enum.
const DefaultInvoiceDateFormat = InvoiceDateFormatInvoiceDateFormatMDY

// InvoiceDateFormat values.
const (
	InvoiceDateFormatInvoiceDateFormatMDY InvoiceDateFormat = "InvoiceDateFormatMDY"
	InvoiceDateFormatInvoiceDateFormatDMY InvoiceDateFormat = "InvoiceDateFormatDMY"
	InvoiceDateFormatInvoiceDateFormatYMD InvoiceDateFormat = "InvoiceDateFormatYMD"
	InvoiceDateFormatInvoiceDateFormatYDM InvoiceDateFormat = "InvoiceDateFormatYDM"
)

func (idf InvoiceDateFormat) String() string {
	return string(idf)
}

// InvoiceDateFormatValidator is a validator for the "invoice_date_format" field enum values. It is called by the builders before save.
func InvoiceDateFormatValidator(idf InvoiceDateFormat) error {
	switch idf {
	case InvoiceDateFormatInvoiceDateFormatMDY, InvoiceDateFormatInvoiceDateFormatDMY, InvoiceDateFormatInvoiceDateFormatYMD, InvoiceDateFormatInvoiceDateFormatYDM:
		return nil
	default:
		return fmt.Errorf("invoicecontrol: invalid enum value for invoice_date_format field: %q", idf)
	}
}

// OrderOption defines the ordering options for the InvoiceControl queries.
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

// ByInvoiceNumberPrefix orders the results by the invoice_number_prefix field.
func ByInvoiceNumberPrefix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceNumberPrefix, opts...).ToFunc()
}

// ByCreditMemoNumberPrefix orders the results by the credit_memo_number_prefix field.
func ByCreditMemoNumberPrefix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreditMemoNumberPrefix, opts...).ToFunc()
}

// ByInvoiceTerms orders the results by the invoice_terms field.
func ByInvoiceTerms(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceTerms, opts...).ToFunc()
}

// ByInvoiceFooter orders the results by the invoice_footer field.
func ByInvoiceFooter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceFooter, opts...).ToFunc()
}

// ByInvoiceLogoURL orders the results by the invoice_logo_url field.
func ByInvoiceLogoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceLogoURL, opts...).ToFunc()
}

// ByInvoiceDateFormat orders the results by the invoice_date_format field.
func ByInvoiceDateFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceDateFormat, opts...).ToFunc()
}

// ByInvoiceDueAfterDays orders the results by the invoice_due_after_days field.
func ByInvoiceDueAfterDays(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceDueAfterDays, opts...).ToFunc()
}

// ByInvoiceLogoWidth orders the results by the invoice_logo_width field.
func ByInvoiceLogoWidth(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvoiceLogoWidth, opts...).ToFunc()
}

// ByShowAmountDue orders the results by the show_amount_due field.
func ByShowAmountDue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShowAmountDue, opts...).ToFunc()
}

// ByAttachPdf orders the results by the attach_pdf field.
func ByAttachPdf(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttachPdf, opts...).ToFunc()
}

// ByShowInvoiceDueDate orders the results by the show_invoice_due_date field.
func ByShowInvoiceDueDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShowInvoiceDueDate, opts...).ToFunc()
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
