package model

// Group Group
type Group struct {
	Base
	DataOwner
	GroupAPI
}

// GroupAPI Group API
type GroupAPI struct {
	Name        *string `json:"name,omitempty" example:"SMA Negri 1" gorm:"type:varchar(256)"` // Name
	Code        *string `json:"code,omitempty" example:"SMAN1S" gorm:"type:varchar(10)"`       // Code
	Description *string `json:"description,omitempty" gorm:"type:text"`                        // Description
}
