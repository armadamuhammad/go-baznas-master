package transaction

import (
	"api/app/controller/login"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionViewCategory godoc
// @Summary Get a Transaction by category
// @Description Get a Transaction by category
// @Param X-User-ID header string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/view/category [get]
// @Tags Transaction
func GetTransactionViewCategory(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)

	var data model.Transaction
	mod := db.Model(&data).
		Where(`"transaction".category_id IN ?`, *categories).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}

// GetTransactionViewGroup godoc
// @Summary Get Transaction view by group user view
// @Description Get a Transaction by group user view
// @Param X-User-ID header string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/view/group [get]
// @Tags Transaction
func GetTransactionViewGroup(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	userID := lib.GetXUserID(c)
	groups := login.GetGroup(userID)

	var data model.Transaction
	mod := db.Model(&data).
		Where(`"transaction".group_id IN ?`, *groups).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}
