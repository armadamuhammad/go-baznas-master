package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTransactionID godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "Transaction ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/{id} [get]
// @Tags Transaction
func GetTransactionID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Transaction
	result := db.Model(&data).
		Where(db.Where(model.Transaction{
			Base: model.Base{
				ID: &id,
			},
		})).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
