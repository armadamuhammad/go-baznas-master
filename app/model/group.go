package model

import (
	"strings"

	"github.com/spf13/viper"
)

// Group Group
type Group struct {
	Base
	DataOwner
	GroupAPI
}

// GroupAPI Group API
type GroupAPI struct {
	Name        *string `json:"name,omitempty" example:"SMA Negri 1" gorm:"type:varchar(256)"`                                 // Name
	Code        *string `json:"code,omitempty" example:"SMAN1S" gorm:"type:varchar(10);index:,unique,where:deleted_at is null"` // Code
	Description *string `json:"description,omitempty" gorm:"type:text"`                                                        // Description
}

func (s *Group) Seed() *[]Group {
	def := viper.GetString("DEF_GROUP")
	data := []Group{}
	items := []string{
		"Muzakki Umum|default Grup untuk registered user",
	}
	for i := range items {
		contents := strings.Split(items[i], "|")
		var name string = contents[0]
		// var code string = contents[1]
		var desc string = contents[1]
		data = append(data, Group{
			GroupAPI: GroupAPI{
				Name:        &name,
				Code:        &def,
				Description: &desc,
			},
		})
	}
	return &data
}
