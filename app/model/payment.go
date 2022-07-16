package model

import "strings"

// Payment Payment
type Payment struct {
	Base
	DataOwner
	PaymentAPI
}

// PaymentAPI Payment API
type PaymentAPI struct {
	Name        *string `json:"name,omitempty" example:"Cash" gorm:"type:varchar(256)"`                                               // Name
	Code        *string `json:"code,omitempty" example:"CSH" gorm:"type:varchar(10);index:,unique,where:deleted_at is null;not null"` // Code
	Description *string `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`                                         // Description
}

func (s *Payment) Seed() *[]Payment {
	data := []Payment{}
	items := []string{
		"Cash|CASH|pembayaran tunai",
		"Transfer Bank|TRFB|metode pembayaran transfer melalui bank",
	}
	for i := range items {
		contents := strings.Split(items[i], "|")
		var name string = contents[0]
		var code string = contents[1]
		var desc string = contents[2]
		data = append(data, Payment{
			PaymentAPI: PaymentAPI{
				Name:        &name,
				Code:        &code,
				Description: &desc,
			},
		})
	}
	return &data
}
