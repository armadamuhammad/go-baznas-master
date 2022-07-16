package model

import (
	"time"

	"github.com/google/uuid"
)

// Transaction Transaction
type Transaction struct {
	Base
	DataOwner
	TransactionAPI
	Total         *float64   `json:"total,omitempty" example:"8000"`                                // Total
	Status        *string    `json:"status,omitempty" gorm:"type:varchar(256)"`                     // Status
	NoRef         *string    `json:"no_ref,omitempty" example:"INFQ00001" gorm:"type:varchar(256)"` // No Ref
	InvoiceNumber *string    `json:"invoice_number,omitempty" gorm:"default:0"`                     // Invoice Number
	Date          *time.Time `json:"date,omitempty" gorm:"type:date"`
	User          *User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"`
	Payment       *Payment   `json:"payment,omitempty" gorm:"foreignKey:PaymentID;references:ID"`
	Category      *Category  `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Balance       *Balance   `json:"balance,omitempty" gorm:"foreignKey:BalanceID;references:ID"`
	Group         *Group     `json:"group,omitempty" gorm:"foreignKey:GroupID;references:ID"`
}

// TransactionAPI Transaction API
type TransactionAPI struct {
	Name         *string    `json:"name,omitempty" example:"pembayaran infaq" gorm:"type:varchar(256)"` // Name
	Description  *string    `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`       // Description
	Type         *string    `json:"type,omitempty" example:"income" gorm:"type:varchar(256)"`           // Type
	Amount       *float64   `json:"amount,omitempty" example:"5000"`                                    // Income
	IsIncome     *int       `json:"is_income,omitempty" example:"1" gorm:"smallint"`                    // 1 = income 0 = uoutcome
	Note         *string    `json:"note,omitempty" example:"lorem ipsum" gorm:"type:text"`              // Note
	Tax          *float64   `json:"tax,omitempty" example:"3000"`                                       // Tax
	TaxType      *int       `json:"tax_type,omitempty" gorm:"type:smallint"`                            // 1 = percentage, 2 = amount
	Contact      *string    `json:"contact,omitempty" example:"08423423432" gorm:"type:varchar(256)"`   // Contact
	ContactName  *string    `json:"contact_name,omitempty" example:"pak guru" gorm:"type:varchar(256)"` // Contact Name
	Discount     *float64   `json:"discount,omitempty" example:"2000"`                                  // Discount
	DiscountType *int       `json:"discount_type,omitempty" gorm:"type:smallint"`                       // 1 = percentage 2 = amount
	Media        *string    `json:"media,omitempty" gorm:"type:varchar(256)"`                           // Media
	UserID       *uuid.UUID `json:"user_id,omitempty" swaggertype:"string" format:"uuid"`               // UserID
	PaymentID    *uuid.UUID `json:"payment_id,omitempty" swaggertype:"string" format:"uuid"`            // PaymentID
	CategoryID   *uuid.UUID `json:"category_id,omitempty" swaggertype:"string" format:"uuid"`           // CategoryID
	BalanceID    *uuid.UUID `json:"balance_id,omitempty" swaggertype:"string" format:"uuid"`            // BalanceID
	GroupID      *uuid.UUID `json:"group_id,omitempty" swaggertype:"string" format:"uuid"`              // BalanceID
}

type TransactionMass struct {
	Items *[]TransactionAPI `json:"item"`
}
