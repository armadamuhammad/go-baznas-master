package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetBalanceRecordBalance godoc
// @Summary Get Gorup Balance Record by id_balance
// @Description Get a Balance Record by id
// @Param id path string true "Balance ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceRecord Balance Record data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records/balance/{id} [get]
// @Tags BalanceRecord
func GetBalanceRecordBalance(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	balanceID, _ := uuid.Parse(c.Params("id"))

	var data model.BalanceRecord
	result := db.Model(&data).
		Where(db.Where(model.BalanceRecord{
			BalanceRecordAPI: model.BalanceRecordAPI{
				BalanceID: &balanceID,
			},
		})).
		Joins("Balance").
		Joins("Transaction")
	page := pg.With(result).Request(c.Request()).Response(&[]model.BalanceRecord{})

	return lib.OK(c, page)
}

// GetBalanceRecordTransaction godoc
// @Summary Get Gorup Balance Record by transaction ID
// @Description Get a Balance Record by id
// @Param id path string true "Transaction ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceRecord Balance Record data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records/transaction/{id} [get]
// @Tags BalanceRecord
func GetBalanceRecordTransaction(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	transactionID, _ := uuid.Parse(c.Params("id"))

	var data model.BalanceRecord
	result := db.Model(&data).
		Where(db.Where(model.BalanceRecord{
			BalanceRecordAPI: model.BalanceRecordAPI{
				TransactionID: &transactionID,
			},
		})).
		Joins("Balance").
		Joins("Transaction")
	page := pg.With(result).Request(c.Request()).Response(&[]model.BalanceRecord{})

	return lib.OK(c, page)
}
