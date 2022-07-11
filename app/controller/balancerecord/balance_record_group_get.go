package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBalanceRecordGroup godoc
// @Summary Get Gorup Balance Record by id_balanc
// @Description Get a Balance Record by id
// @Param id path string true "Balance Record ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceRecord Balance Record data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records/{id}/group [get]
// @Tags BalanceRecord
func GetBalanceRecordGroup(c *fiber.Ctx) error {
	db := services.DB
	balanceID, _ := uuid.Parse(c.Params("id"))

	var data model.BalanceRecord
	result := db.Model(&data).
		Where(db.Where(model.BalanceRecord{
			BalanceRecordAPI: model.BalanceRecordAPI{
				BalanceID: &balanceID,
			},
		})).
		Find(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
