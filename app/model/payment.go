package model

// Payment Payment
type Payment struct {
	Base
	PaymentAPI
}

// PaymentAPI Payment API
type PaymentAPI struct {
	Name        *string `json:"name,omitempty" example:"Cash" gorm:"type:varchar(256)"`                                                   // Name
	Code        *string `json:"code,omitempty" example:"CSH" gorm:"type:varchar(10);index:Code,unique,where:deleted_at is null;not null"` // Code
	Description *string `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`                                             // Description
}
