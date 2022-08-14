package transaction

import (
	"api/app/controller/login"
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
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)

	var data model.Transaction
	result := db.Model(&data).
		Where(db.Where(model.Transaction{
			Base: model.Base{
				ID: &id,
			},
		})).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		Joins("Account").
		Preload("User.Role").
		Preload("User.Group").
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	return lib.OK(c, data)
}
