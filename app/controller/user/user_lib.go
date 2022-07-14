package user

import (
	"api/app/model"
	"api/app/services"

	"github.com/google/uuid"

)

func GetUserData(id *uuid.UUID) (*model.User, error) {
	db := services.DB
	var data model.User

	res := db.Model(&data).
		Where(db.Where(model.User{
			Base: model.Base{
				ID: id,
			},
		})).
		First(&data)
	if res.RowsAffected < 1 {
		return nil, res.Error
	}
	return &data, nil
}
