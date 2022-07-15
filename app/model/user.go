package model

import (
	"api/app/lib"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// User User
type User struct {
	Base
	DataOwner
	UserAPI
	Status         *int             `json:"status,omitempty" example:"0" gorm:"type:smallint"`                                   // Status
	StatusVerified *int             `json:"status_verified,omitempty" example:"0" gorm:"type:smallint"`                          // Status Verified
	IsAdmin        *int             `json:"is_admin" gorm:"type:smallint"`                                                       // is admin
	JoinDate       *strfmt.DateTime `json:"join_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"` // JoinDate
	Super          *int             `json:"super" gorm:"type:smallint"`                                                          // is Super Admin
	Password       *string          `json:"password,omitempty" gorm:"type:varchar(256)"`                                         // Password
	Role           *Role            `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
	Group          *Group           `json:"group,omitempty" gorm:"foreignKey:GroupID;references:ID"`
}

// UserAPI User API
type UserAPI struct {
	FirstName *string    `json:"first_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // FirstName
	LastName  *string    `json:"last_name,omitempty" gorm:"type:varchar(256)"`                                                                                  // LastName
	Surname   *string    `json:"surname,omitempty" gorm:"type:varchar(256)"`                                                                                    // Surname
	Email     *string    `json:"email,omitempty" example:"user@mail.com" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"`         // Email
	Username  *string    `json:"username,omitempty" example:"armada_muhammad" gorm:"type:varchar(256);index:,unique,where:deleted_at is null;not null"` // Username
	Address   *string    `json:"address,omitempty" example:"jl. Klaten" gorm:"type:text"`                                                                       // Address
	Gender    *string    `json:"gender,omitempty" example:"male" gorm:"type:varchar(256)"`                                                                      // Gender
	RoleID    *uuid.UUID `json:"role_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                  // Role
	GroupID   *uuid.UUID `json:"group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                 // GroupID
}

// UserLogin struct
type UserLogin struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UserPassword struct
type UserPassword struct {
	PasswordOld *string `json:"password_old,omitempty"`
	PasswordNew *string `json:"password_new,omitempty"`
}

func (s *User) Seed() *[]User {
	data := []User{}
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	hamba := lib.PasswordEncrypt("password", salt, key)
	super := lib.PasswordEncrypt("super", salt, key)

	seed := []User{
		{
			Base:      Base{},
			DataOwner: DataOwner{},
			UserAPI: UserAPI{
				FirstName: lib.Strptr("Hamba"),
				LastName:  lib.Strptr("Allah"),
				Surname:   nil,
				Email:     lib.Strptr("hambaallah@baznas.com"),
				Username:  lib.Strptr(viper.GetString("USER_ANON")),
				Address:   lib.Strptr("bumi"),
				Gender:    nil,
				RoleID:    nil,
				GroupID:   nil,
			},
			Status:         lib.Intptr(1),
			StatusVerified: lib.Intptr(1),
			IsAdmin:        lib.Intptr(0),
			JoinDate:       nil,
			Super:          lib.Intptr(0),
			Password:       &hamba,
			Role:           nil,
			Group:          nil,
		},
		{
			Base:      Base{},
			DataOwner: DataOwner{},
			UserAPI: UserAPI{
				FirstName: lib.Strptr("super"),
				LastName:  lib.Strptr("admin"),
				Surname:   nil,
				Email:     lib.Strptr("superadmin@baznas.com"),
				Username:  lib.Strptr("superadmin"),
				Address:   lib.Strptr("kantor baznas surakarta"),
				Gender:    nil,
				RoleID:    nil,
				GroupID:   nil,
			},
			Status:         lib.Intptr(1),
			StatusVerified: lib.Intptr(1),
			IsAdmin:        lib.Intptr(1),
			JoinDate:       nil,
			Super:          lib.Intptr(1),
			Password:       &super,
			Role:           nil,
			Group:          nil,
		},
	}
	lib.Merge(seed, &data)
	return &data
}
