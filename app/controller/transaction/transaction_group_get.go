package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTransactionUser godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/user/{id} [get]
// @Tags Transaction
func GetTransactionUser(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	userID, _ := uuid.Parse(c.Params("id"))

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				UserID: &userID,
			},
		})).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}

// GetTransactionCategory godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/category/{id} [get]
// @Tags Transaction
func GetTransactionCategory(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	categoryID, _ := uuid.Parse(c.Params("id"))

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				CategoryID: &categoryID,
			},
		})).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		Preload("User.Role").
		Preload("User.Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}

// GetTransactionBalance godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/balance/{id} [get]
// @Tags Transaction
func GetTransactionBalance(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	balanceID, _ := uuid.Parse(c.Params("id"))

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				BalanceID: &balanceID,
			},
		})).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, page)
}
