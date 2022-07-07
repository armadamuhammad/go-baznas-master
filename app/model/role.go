package model

// Role Role
type Role struct {
	Base
	DataOwner
	RoleAPI
}

// RoleAPI Role API
type RoleAPI struct {
	Name        *string `json:"name,omitempty" gorm:"type:varchar(256)"` // Name
	Code        *string `json:"code,omitempty" gorm:"type:varchar(10)"`  // Code
	Description *string `json:"description,omitempty" gorm:"type:text"`  // Description
}
