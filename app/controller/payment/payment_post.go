package payment

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostPayment godoc
// @Summary Create new Payment
// @Description Create new Payment
// @Param X-User-ID header string false "User ID"
// @Param data body model.PaymentAPI true "Payment data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Payment "Payment data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /payments [post]
// @Tags Payment
func PostPayment(c *fiber.Ctx) error {
	api := new(model.PaymentAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Payment
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
