package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionIncome godoc
// @Summary Get a total income Transaction
// @Description Get a Transaction total income
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transaction/income/ [get]
// @Tags Transaction
func GetTransactionIncome(c *fiber.Ctx) error {
	db := services.DB

	var total float64
	var data model.Transaction

	row := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				IsIncome: lib.Intptr(1),
			},
		})).
		Select(`SUM(total)`).
		Row()
	row.Scan(&total)

	return lib.OK(c, total)
}

// GetTransactionOutcome godoc
// @Summary Get a total income Transaction
// @Description Get a Transaction total income
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transaction/outcome/ [get]
// @Tags Transaction
func GetTransactionOutcome(c *fiber.Ctx) error {
	db := services.DB

	var total float64
	var data model.Transaction

	row := db.Model(&data).
		Where(db.Where(model.Transaction{
			TransactionAPI: model.TransactionAPI{
				IsIncome: lib.Intptr(0),
			},
		})).
		Select(`SUM(total)`).
		Row()
	row.Scan(&total)

	return lib.OK(c, total)
}
