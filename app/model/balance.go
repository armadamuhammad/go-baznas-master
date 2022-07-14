package model

import "github.com/google/uuid"

// Balance Balance
type Balance struct {
	Base
	DataOwner
	BalanceAPI
}

// BalanceAPI Balance API
type BalanceAPI struct {
	Amount      *float64 `json:"amount,omitempty" example:"1000000"`                         // Saldo
	Name        *string  `json:"name,omitempty" example:"Kas Amil" gorm:"type:varchar(256)"` // Name
	Code        *string  `json:"code,omitempty" example:"KAML" gorm:"type:varchar(10)"`      // Code
	Description *string  `json:"description,omitempty" gorm:"type:text"`                     // Description
}

// BalanceTotal struct
type BalanceTotal struct {
	Amount *float64 `json:"amount,omitempty" example:"1000000"` // Saldo
}

// BalanceTransfer struct
type BalanceTransfer struct {
	From   *uuid.UUID `json:"from,omitempty" swaggertype:"string" format:"uuid"` // From balance
	To     *uuid.UUID `json:"to,omitempty" swaggertype:"string" format:"uuid"`   // To balance
	Amount *float64   `json:"amount,omitempty" example:"1000000"`                // Amount

}
