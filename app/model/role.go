package model

import "strings"

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

func (s *Role) Seed() *[]Role {
	data := []Role{}
	items := []string{
		"User|USER",
		"Admin|ADMN",
		"UPZ|UUPZ",
	}
	for i := range items {
		contents := strings.Split(items[i], "|")
		var name string = contents[0]
		var code string = contents[1]
		data = append(data, Role{
			RoleAPI: RoleAPI{
				Name: &name,
				Code: &code,
			},
		})
	}
	return &data
}
