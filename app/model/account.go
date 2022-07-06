package model

// Account Account
type Account struct {
	Base
	DataOwner
	AccountAPI
}

// AccountAPI Account API
type AccountAPI struct {
	Name        *string `json:"name,omitempty" example:"Bank BCA"`                                                                         // Name
	Code        *string `json:"code,omitempty" example:"BBCA" gorm:"type:varchar(10);index:Code,unique,where:deleted_at is null;not null"` // Code
	Number      *string `json:"number,omitempty" example:"3425-3234-2341" gorm:"type:varchar(256)"`                                        // Number
	Description *string `json:"description,omitempty" example:"lorem ipsum" gorm:"type:text"`                                              // Description
}
