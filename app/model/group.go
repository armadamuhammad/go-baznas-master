package model

import "strings"

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

func (s *Group) Seed() *[]Group {
	data := []Group{}
	items := []string{
		"Muzakki Umum|MZKU|default Grup untuk registered user",
	}
	for i := range items {
		contents := strings.Split(items[i], "|")
		var name string = contents[0]
		var code string = contents[1]
		var desc string = contents[2]
		data = append(data, Group{
			GroupAPI: GroupAPI{
				Name:        &name,
				Code:        &code,
				Description: &desc,
			},
		})
	}
	return &data
}
