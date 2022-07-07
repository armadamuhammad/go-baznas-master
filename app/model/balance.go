package model

import "github.com/google/uuid"

// Balance Balance
type Balance struct {
	Base
	DataOwner
	BalanceAPI
	Transaction *Transaction `json:"transaction" gorm:"foreignKey:TransactionID;references:ID"`
}

// BalanceAPI Balance API
type BalanceAPI struct {
	TransactionID *uuid.UUID `json:"transaction_id,omitempty" swaggertype:"string" format:"uuid"` // BalanceID
	Amount        *float64   `json:"amount,omitempty" example:"1000000"`                          // Saldo
	Name          *string    `json:"name,omitempty" example:"Kas Amil" gorm:"type:varchar(256)"`  // Name
	Code          *string    `json:"code,omitempty" example:"KAML" gorm:"type:varchar(10)"`       // Code
	Description   *string    `json:"description,omitempty" gorm:"type:text"`                      // Description
}
