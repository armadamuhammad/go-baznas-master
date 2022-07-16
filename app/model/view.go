package model

import "github.com/google/uuid"

type ViewCategory struct {
	Base
	UserID     *uuid.UUID `json:"user_id,omitempty"`
	CategoryID *uuid.UUID `json:"category,omitempty"`
}

type ViewGroup struct {
	Base
	UserID  *uuid.UUID `json:"user_id,omitempty" swaggertype:"string" format:"uuid"`
	GroupID *uuid.UUID `json:"group,omitempty" swaggertype:"string" format:"uuid"`
}

type ViewAPI struct {
	// UserID *uuid.UUID `json:"user_id,omitempty" swaggertype:"string" format:"uuid"`
	Items  *[]string  `json:"items,omitempty" format:"uuid"`
}
