package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
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
	Super          *int             `json:"-" gorm:"type:smallint"`                                                              // is Super Admin
	Password       *string          `json:"password,omitempty" gorm:"type:varchar(256)"`                                         // Password
	Role           *Role            `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
	Group          *Group           `json:"group,omitempty" gorm:"foreignKey:GroupID;references:ID"`
}

// UserAPI User API
type UserAPI struct {
	FirstName *string    `json:"first_name,omitempty" gorm:"type:varchar(256)"`                                                                                 // FirstName
	LastName  *string    `json:"last_name,omitempty" gorm:"type:varchar(256)"`                                                                                  // LastName
	Surname   *string    `json:"surname,omitempty" gorm:"type:varchar(256)"`                                                                                    // Surname
	Email     *string    `json:"email,omitempty" example:"user@mail.com" gorm:"type:varchar(256);index:email,unique,where:deleted_at is null;not null"`         // Email
	Username  *string    `json:"username,omitempty" example:"armada_muhammad" gorm:"type:varchar(256);index:Username,unique,where:deleted_at is null;not null"` // Username
	Address   *string    `json:"address,omitempty" example:"jl. Klaten" gorm:"type:text"`                                                                       // Address
	Gender    *string    `json:"gender,omitempty" example:"male" gorm:"type:varchar(256)"`                                                                      // Gender
	RoleID    *uuid.UUID `json:"role_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                  // Role
	GroupID   *uuid.UUID `json:"group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                 // GroupID
}

type UserLogin struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserPassword struct {
	PasswordOld *string `json:"password_old,omitempty"`
	PasswordNew *string `json:"password_new,omitempty"`
}
