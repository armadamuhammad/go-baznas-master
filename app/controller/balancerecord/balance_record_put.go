package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutBalanceRecord godoc
// @Summary Update Balance Record by id
// @Description Update Balance Record by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Balance Record ID"
// @Param data body model.BalanceRecordAPI true "Balance Record data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceRecord "Balance Record data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records/{id} [put]
// @Tags BalanceRecord
func PutBalanceRecord(c *fiber.Ctx) error {
	api := new(model.BalanceRecordAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.BalanceRecord
	result := db.Model(&data).
		Where(db.Where(model.BalanceRecord{
			Base: model.Base{
				ID: &id,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)
	data.ModifierID = lib.GetXUserID(c)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
