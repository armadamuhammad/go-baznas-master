package model

// City City
type City struct {
	Base
	DataOwner
	CityAPI
}

// CityAPI City API
type CityAPI struct {
	Name       *string `json:"name,omitempty" gorm:"type:varchar(256)"`                                                              // Name
	Code       *string `json:"code,omitempty" example:"SKH" gorm:"type:varchar(10);index:,unique,where:deleted_at is null;not null"` // Code
	PostalCode *string `json:"postalcode,omitempty" example:"57465" gorm:"type:varchar(10)"`                                         // PostalCode
	Latitude   *string `json:"latitude,omitempty" gorm:"type:varchar(256)"`                                                          // Latitude
	Longitude  *string `json:"longitude,omitempty" gorm:"type:varchar(256)"`                                                         // Longitude
}
