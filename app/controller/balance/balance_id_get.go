package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// GetBalanceID godoc
// @Summary Get a Balance by id
// @Description Get a Balance by id
// @Param id path string true "Balance ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Balance Balance data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balances/{id} [get]
// @Tags Balance
func GetBalanceID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Balance
	result := db.Model(&data).
		Where(db.Where(model.Balance{
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

func GetBalanceDefault() *uuid.UUID {
	db := services.DB

	var data model.Balance
	db.Model(&data).
		Where(db.Where(model.Balance{
			BalanceAPI: model.BalanceAPI{
				Code: lib.Strptr(viper.GetString("DEF_BALANCE")),
			},
		})).First(&data)
	return data.ID
}
