package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountingControl struct {
	TimeStampedModel
	BusinessUnitID               uuid.UUID                  `gorm:"type:uuid;not null;index"                              json:"businessUnitId"`
	OrganizationID               uuid.UUID                  `gorm:"type:uuid;not null;index"                              json:"organizationId"`
	RecThreshold                 int64                      `gorm:"type:int;not null;default:50"                          json:"recThreshold"            validate:"required"`
	DefaultRevenueAccountID      *uuid.UUID                 `gorm:"type:uuid"                                             json:"defaultRevenueAccountId" validate:"omitempty"`
	DefaultExpenseAccountID      *uuid.UUID                 `gorm:"type:uuid"                                             json:"defaultExpenseAccountId" validate:"omitempty"`
	BusinessUnit                 BusinessUnit               `json:"-" validate:"omitempty"`
	JournalEntryCriteria         *AutomaticJournalEntryType `gorm:"type:varchar(50);default:'ON_SHIPMENT_BILL'"           json:"journalEntryCriteria"    validate:"omitempty,oneof=ON_SHIPMENT_BILL ON_RECEIPT_OF_PAYMENT ON_EXPENSE_RECOGNITION"`
	RecThresholdAction           ThresholdActionType        `gorm:"type:ac_threshold_action_type;not null;default:'HALT'" json:"recThresholdAction"      validate:"required,oneof=HALT WARN"`
	DefaultRevenueAccount        *GeneralLedgerAccount      `gorm:"foreignKey:DefaultRevenueAccountID;references:ID"      json:"-"                       validate:"omitempty"`
	DefaultExpenseAccount        *GeneralLedgerAccount      `gorm:"foreignKey:DefaultExpenseAccountID;references:ID"      json:"-"                       validate:"omitempty"`
	AutoCreateJournalEntries     bool                       `gorm:"type:boolean;not null;default:false"                   json:"autoCreateJournalEntries"     validate:"omitempty"`
	RestrictManualJournalEntries bool                       `gorm:"type:boolean;not null;default:false"                   json:"restrictManualJournalEntries" validate:"omitempty"`
	RequireJournalEntryApproval  bool                       `gorm:"type:boolean;not null;default:false"                   json:"requireJournalEntryApproval"  validate:"omitempty"`
	EnableRecNotifications       bool                       `gorm:"type:boolean;not null;default:true"                    json:"enableRecNotifications"       validate:"omitempty"`
	HaltOnPendingRec             bool                       `gorm:"type:boolean;not null;default:false"                   json:"haltOnPendingRec"             validate:"omitempty"`
}

var ErrExpenseAccount = errors.New("default expense account must be an expense account")
var ErrRevenueAccount = errors.New("default revenue account must be a revenue account")

func (ac *AccountingControl) validateAccountingControl() error {
	if ac.DefaultExpenseAccountID != nil && ac.DefaultExpenseAccount.AccountType != Exp {
		return ErrExpenseAccount
	}

	if ac.DefaultRevenueAccountID != nil && ac.DefaultRevenueAccount.AccountType != Rev {
		return ErrRevenueAccount
	}

	return nil
}

func (ac *AccountingControl) BeforeCreate(_ *gorm.DB) error {
	return ac.validateAccountingControl()
}

func (ac *AccountingControl) BeforeUpdate(_ *gorm.DB) error {
	return ac.validateAccountingControl()
}
