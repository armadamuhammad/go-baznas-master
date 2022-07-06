package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// User User
type User struct {
	Base
	UserAPI
}

// UserAPI User API
type UserAPI struct {
	FirstName      *string          `json:"first_name,omitempty" gorm:"type:varchar(256)"`                                                                                   // FirstName
	LastName       *string          `json:"last_name,omitempty" gorm:"type:varchar(256)"`                                                                                    // LastName
	Surname        *string          `json:"surname,omitempty" gorm:"type:varchar(256)"`                                                                                     // Surname
	Email          *string          `json:"email,omitempty" example:"user@mail.com" gorm:"type:varchar(256);index:email,unique,where:deleted_at is null;not null"`          // Email
	Username       *string          `json:"username,omitempty" example:"armadam_muhammad" gorm:"type:varchar(256);index:Username,unique,where:deleted_at is null;not null"` // Username
	Password       *string          `json:"password,omitempty" gorm:"type:varchar(256)"`                                                                                    // Password
	Address        *string          `json:"address,omitempty" example:"jl. Klaten" gorm:"type:text"`                                                                        // Address
	JoinDate       *strfmt.DateTime `json:"join_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz"`                                             // JoinDate
	Gender         *string          `json:"gender,omitempty" example:"male" gorm:"type:varchar(256)"`                                                                       // Gender
	Status         *bool            `json:"status,omitempty" example:"true"`                                                                                                // Status
	StatusVerified *bool            `json:"status_verified,omitempty" example:"false"`                                                                                      // Status Verified
	RoleID         *uuid.UUID       `json:"role_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                                         // Role
	GroupID        *uuid.UUID       `json:"group_id,omitempty" gorm:"type:varchar(36)" swaggertype:"string" format:"uuid"`                                                                           // GroupID
}
