package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// BalanceRecord Balance Record
type BalanceRecord struct {
	Base
	DataOwner
	BalanceRecordAPI
}

// BalanceRecordAPI Balance Record API
type BalanceRecordAPI struct {
	Amount        *float64         `json:"amount,omitempty" example:"200000"`                                                  // Amount
	Datetime      *strfmt.DateTime `json:"datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Datetime
	BalanceID     *uuid.UUID       `json:"balance_id,omitempty" swaggertype:"string" format:"uuid"`                            // Balance ID
	TransactionID *uuid.UUID       `json:"transaction_id,omitempty" swaggertype:"string" format:"uuid"`                        // Transaction ID
}
