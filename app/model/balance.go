package model

import "github.com/google/uuid"

// Balance Balance
type Balance struct {
	Base
	BalanceAPI
}

// BalanceAPI Balance API
type BalanceAPI struct {
	Saldo         *float64   `json:"saldo,omitempty" example:"1000000"`                          // Saldo
	Category      *string    `json:"category,omitempty" gorm:"type:varchar(256)"`                // Category
	Income        *float64   `json:"income,omitempty" example:"50000"`                           // Income
	Outcome       *float64   `json:"outcome,omitempty" example:"40000"`                          // Outcome
	TransactionID *uuid.UUID `json:"transactionid,omitempty" swaggertype:"string" format:"uuid"` // TransactionID
	Description   *string    `json:"description,omitempty" gorm:"type:text"`                     // Description
}
