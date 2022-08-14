package model

import (
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// Category Category
type Category struct {
	Base
	DataOwner
	CategoryAPI
	Balance *Balance  `json:"balance,omitempty" gorm:"foreignKey:BalanceID;references:ID"`
	Parent  *Category `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:ID"`
	// Child   *[]Category `json:"child,omitempty" gorm:"foreignKey:ID;references:ParentID"`
}

// CategoryAPI Category API
type CategoryAPI struct {
	Name        *string    `json:"name,omitempty" example:"infaq" gorm:"type:varchar(256)"`                                     // Name
	Code        *string    `json:"code,omitempty" example:"12" gorm:"type:varchar(256);index:,unique,where:deleted_at is null"` // Code
	Level       *int       `json:"level,omitempty" gorm:"type:smallint"`                                                        // Level
	Category    *string    `json:"category,omitempty" gorm:"type:varchar(256)"`                                                 // Category
	IsIncome    *int       `json:"is_income,omitempty" example:"1"`                                                             // 1 = income 0 = outcome
	Anonym      *int       `json:"Anonym,omitempty" example:"1"`                                                                // 1 = can be anonym transaction 0 = should login
	Description *string    `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`                                // Description
	ParentID    *uuid.UUID `json:"parent_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`              // ParentID
	BalanceID   *uuid.UUID `json:"balance_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`             // BalanceID
}

func (s *Category) Seed() *[]Category {
	data := []Category{}
	items := []string{
		"Aset|1|Kategori Aset",
		"Kewajiban|2|Kategori Kewajiban",
		"Saldo|3|Kategori Saldo",
		"Penerimaan|4|Kategori Penerimaan",
		"Penyaluran|5|Kategori Penyaluran",
	}
	for i := range items {
		parent, _ := uuid.Parse(viper.GetString("PARENT_LV1"))
		contents := strings.Split(items[i], "|")
		var name string = contents[0]
		var code string = contents[1]
		var desc string = contents[2]
		data = append(data, Category{
			CategoryAPI: CategoryAPI{
				Name:        &name,
				Code:        &code,
				Description: &desc,
				ParentID:    &parent,
			},
		})
	}
	return &data
}
