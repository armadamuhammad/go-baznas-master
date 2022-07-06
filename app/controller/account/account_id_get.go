package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAccountID godoc
// @Summary Get a Account by id
// @Description Get a Account by id
// @Param id path string true "Account ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Account Account data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /accounts/{id} [get]
// @Tags Account
func GetAccountID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Account
	result := db.Model(&data).
		Where(db.Where(model.Account{
			Base: model.Base{
				ID: &id,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
