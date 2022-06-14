package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model
type Base struct {
	ID        *uuid.UUID       `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid"`        // model ID
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // created at automatically inserted on post
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // updated at automatically changed on put or add on post
	DeletedAt gorm.DeletedAt   `json:"-" gorm:"index" swaggerignore:"true"`
}

// BeforeCreate Data
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if nil != b.ID {
		return nil
	}
	id, e := uuid.NewRandom()
	now := strfmt.DateTime(time.Now())
	b.ID = &id
	b.CreatedAt = &now
	b.UpdatedAt = &now
	return e
}

// BeforeUpdate Data
func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	if nil != b.ID {
		return nil
	}
	now := strfmt.DateTime(time.Now())
	b.UpdatedAt = &now
	return nil
}
