package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Transaction Transaction
type Transaction struct {
	Base
	DataOwner
	TransactionAPI
	Total         *float64 `json:"total,omitempty" example:"8000"`                                // Total
	Status        *string  `json:"status,omitempty" gorm:"type:varchar(256)"`                     // Status
	NoRef         *string  `json:"no_ref,omitempty" example:"INFQ00001" gorm:"type:varchar(256)"` // No Ref
	InvoiceNumber *int64   `json:"invoice_number,omitempty" gorm:"default:0"`                     // Invoice Number

}

// TransactionAPI Transaction API
type TransactionAPI struct {
	Name         *string    `json:"name,omitempty" example:"pembayaran infaq" gorm:"type:varchar(256)"` // Name
	Description  *string    `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`       // Description
	Type         *string    `json:"type,omitempty" example:"income" gorm:"type:varchar(256)"`           // Type
	Amount       *float64   `json:"amount,omitempty" example:"5000"`                                    // Income
	IsIncome     *int       `json:"is_income,omitempty" example:"1" gorm:"smallint"`                    // Outcome
	Note         *string    `json:"note,omitempty" example:"lorem ipsum" gorm:"type:text"`              // Note
	Tax          *float64   `json:"tax,omitempty" example:"3000"`                                       // Tax
	TaxType      *int       `json:"tax_type,omitempty" gorm:"type:smallint"`                            // Tax Type
	Contact      *string    `json:"contact,omitempty" example:"08423423432" gorm:"type:varchar(256)"`   // Contact
	ContactName  *string    `json:"contact_name,omitempty" example:"pak guru" gorm:"type:varchar(256)"` // Contact Name
	Discount     *float64   `json:"discount,omitempty" example:"2000"`                                  // Discount
	DiscountType *int       `json:"discount_type,omitempty" gorm:"type:smallint"`                       // 1 = percentage 2 = amount
	UserID       *uuid.UUID `json:"userid,omitempty" swaggertype:"string" format:"uuid"`                // UserID
	PaymentID    *uuid.UUID `json:"payment_id,omitempty" swaggertype:"string" format:"uuid"`            // PaymentID
	CategoryID   *uuid.UUID `json:"category_id,omitempty" swaggertype:"string" format:"uuid"`           // CategoryID
	BalanceID    *uuid.UUID `json:"balance_id,omitempty" swaggertype:"string" format:"uuid"`            // BalanceID
}

func (b *Transaction) BeforeCreate(tx *gorm.DB) error {
	_, e := uuid.NewRandom()
	// now := strfmt.DateTime(time.Now())
	var latest int64
	tx.Table(tx.Statement.Table).
		Select(`COALESCE(MAX(sort), 0) latest`).
		Where(`deleted_at IS NULL AND sort IS NOT NULL`).
		Scan(&latest)
	latest++
	b.InvoiceNumber = &latest
	return e
}
