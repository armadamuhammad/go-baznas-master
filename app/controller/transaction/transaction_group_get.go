package transaction

import (
	"api/app/controller/login"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTransactionUser godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param X-User-ID header string true "User ID"
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
	user := lib.GetXUserID(c)
	categories := login.GetCategory(user)
	groups := login.GetGroup(user)

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				UserID: &userID,
			},
		})).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
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
// @Param X-User-ID header string true "User ID"
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
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				CategoryID: &categoryID,
			},
		})).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
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
// @Param X-User-ID header string true "User ID"
// @Param id path string true "Balance ID"
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
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				BalanceID: &balanceID,
			},
		})).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})
	return lib.OK(c, page)
}

// GetTransactionAccount godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param X-User-ID header string true "User ID"
// @Param id path string true "Account ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/account/{id} [get]
// @Tags Transaction
func GetTransactionAccount(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	accountID, _ := uuid.Parse(c.Params("id"))
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)

	var data model.Transaction
	mod := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				AccountID: &accountID,
			},
		})).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})
	return lib.OK(c, page)
}
