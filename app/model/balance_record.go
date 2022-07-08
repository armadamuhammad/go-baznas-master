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
	Transaction *Transaction `json:"transaction" gorm:"foreignKey:TransactionID;references:ID"`
	Balance     *Balance     `json:"balance" gorm:"foreignKey:BalanceID;references:ID"`
}

// BalanceRecordAPI Balance Record API
type BalanceRecordAPI struct {
	Amount        *float64         `json:"amount,omitempty" example:"200000"`                                                  // Amount
	Type          *string          `json:"type,omitempty" example:"Income" gorm:"type:varchar(256)"`                           // Type
	Datetime      *strfmt.DateTime `json:"datetime,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // Datetime
	BalanceID     *uuid.UUID       `json:"balance_id,omitempty" swaggertype:"string" format:"uuid"`                            // Balance ID
	TransactionID *uuid.UUID       `json:"transaction_id,omitempty" swaggertype:"string" format:"uuid"`                        // Transaction ID
}
