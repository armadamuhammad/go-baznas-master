package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionDate godoc
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
// @Router /transactions/{date}/date [get]
// @Tags Transaction
func GetTransactionDate(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	layout := "2006-01-02"
	date := c.Params("date")
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	// t.Format("2006-01-02")
	// d := strfmt.Date(string(date))

	data := model.Transaction{}
	mod := db.Model(&data).
		Where(`date = ?`, t)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, &page)
}
