package model

import "github.com/google/uuid"

// Input Input
type Input struct {
	Base
	DataOwner
	InputAPI
	Group    *Group    `json:"group" gorm:"foreignKey:GroupID;references:ID"`
	User     *User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Category *Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
}

// InputAPI Input API
type InputAPI struct {
	GroupID    *uuid.UUID `json:"group_id,omitempty" swaggertype:"string" format:"uuid"`    // Group ID
	UserID     *uuid.UUID `json:"user_id,omitempty" swaggertype:"string" format:"uuid"`     // User ID
	CategoryID *uuid.UUID `json:"category_id,omitempty" swaggertype:"string" format:"uuid"` // Category ID
}
