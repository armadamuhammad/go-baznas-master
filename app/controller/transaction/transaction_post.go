package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostTransaction godoc
// @Summary Create new Transaction
// @Description Create new Transaction
// @Param X-User-ID header string false "User ID"
// @Param data body model.TransactionAPI true "Transaction data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction "Transaction data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions [post]
// @Tags Transaction
func PostTransaction(c *fiber.Ctx) error {
	api := new(model.TransactionAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Transaction
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)
	data.Total = GetDiscount(*data.Amount, *data.Discount, *data.DiscountType)
	data.Total = GetTax(*data.Total, *data.Tax, *data.TaxType)
	


	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
