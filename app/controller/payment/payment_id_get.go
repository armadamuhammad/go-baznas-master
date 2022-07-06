package payment

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetPaymentID godoc
// @Summary Get a Payment by id
// @Description Get a Payment by id
// @Param Accept-Language header string false "2 character language code"
// @Param id path string true "Payment ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Payment Payment data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /payments/{id} [get]
// @Tags Payment
func GetPaymentID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Payment
	result := db.Model(&data).
		Where(db.Where(model.Payment{
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
