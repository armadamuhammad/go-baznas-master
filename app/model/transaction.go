package model

import "github.com/google/uuid"

// Transaction Transaction
type Transaction struct {
	Base
	TransactionAPI
}

// TransactionAPI Transaction API
type TransactionAPI struct {
	Name          *string    `json:"name,omitempty" example:"pembayaran infaq" gorm:"type:varchar(256)"` // Name
	Description   *string    `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`       // Description
	InvoiceNumber *int       `json:"invoice_number,omitempty" gorm:"type:smallint"`                      // Invoice Number
	NoRef         *string    `json:"no_ref,omitempty" example:"INFQ00001" gorm:"type:varchar(256)"`      // No Ref
	Type          *string    `json:"type,omitempty" example:"income" gorm:"type:varchar(256)"`           // Type
	Income        *float64   `json:"income,omitempty" example:"5000"`                                    // Income
	Outcome       *float64   `json:"outcome,omitempty" example:"3000"`                                   // Outcome
	Total         *float64   `json:"total,omitempty" example:"8000"`                                     // Total
	Note          *string    `json:"note,omitempty" example:"lorem ipsum" gorm:"type:text"`              // Note
	Tax           *float64   `json:"tax,omitempty" example:"3000"`                                       // Tax
	TaxType       *int       `json:"tax_type,omitempty" gorm:"type:smallint"`                            // Tax Type
	Contact       *string    `json:"contact,omitempty" example:"08423423432" gorm:"type:varchar(256)"`   // Contact
	ContactName   *string    `json:"contact_name,omitempty" example:"pak guru" gorm:"type:varchar(256)"` // Contact Name
	Discount      *float64   `json:"discount,omitempty" example:"2000"`                                  // Discount
	DiscountType  *int       `json:"discount_type,omitempty" gorm:"type:smallint"`                       // 1 = percentage 2 = amount
	Status        *string    `json:"status,omitempty" gorm:"type:varchar(256)"`                          // Status
	UserID        *uuid.UUID `json:"userid,omitempty" swaggertype:"string" format:"uuid"`                // UserID
	PaymentID     *uuid.UUID `json:"paymentid,omitempty" swaggertype:"string" format:"uuid"`             // PaymentID
	CategoryID    *uuid.UUID `json:"categoryid,omitempty" swaggertype:"string" format:"uuid"`            // CategoryID
}
