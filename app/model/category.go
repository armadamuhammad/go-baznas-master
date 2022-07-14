package model

import "github.com/google/uuid"

// Category Category
type Category struct {
	Base
	DataOwner
	CategoryAPI
	Balance *Balance `json:"balance,omitempty" gorm:"foreignKey:BalanceID;references:ID"`
	// Child   *[]Category `json:"child,omitempty" gorm:"foreignKey:ID;references:ParentID"`
}

// CategoryAPI Category API
type CategoryAPI struct {
	Name        *string    `json:"name,omitempty" example:"infaq" gorm:"type:varchar(256)"`                         // Name
	Code        *string    `json:"code,omitempty" example:"12" gorm:"type:varchar(256)"`                            // Code
	Level       *int       `json:"level,omitempty" gorm:"type:smallint"`                                            // Level
	Category    *string    `json:"category,omitempty" gorm:"type:varchar(256)"`                                     // Category
	IsIncome    *int       `json:"is_income,omitempty" example:"1"`                                                 // 1 = income 0 = outcome
	Anonym      *int       `json:"Anonym,omitempty" example:"1"`                                                    // 1 = can be anonym transaction 0 = should login
	Description *string    `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`                    // Description
	ParentID    *uuid.UUID `json:"parent_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`  // ParentID
	BalanceID   *uuid.UUID `json:"balance_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"` // BalanceID
}
