package model

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

type BalanceTotal struct {
	Amount *float64 `json:"amount,omitempty" example:"1000000"` // Saldo
}
