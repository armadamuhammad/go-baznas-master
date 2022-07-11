package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostBalanceRecord godoc
// @Summary Create new Balance Record
// @Description Create new Balance Record
// @Param X-User-ID header string false "User ID"
// @Param data body model.BalanceRecordAPI true "Balance Record data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceRecord "Balance Record data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records [post]
// @Tags BalanceRecord

func PostBalanceRecord(c *fiber.Ctx) error {
	api := new(model.BalanceRecordAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.BalanceRecord
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
