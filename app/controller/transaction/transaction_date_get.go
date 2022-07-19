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
// @Param date path string true "Transaction date"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/date/{date} [get]
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
		Where(`date = ?`, t).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		Preload("User.Role").
		Preload("User.Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, &page)
}

// GetTransactionFromDateTo godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param from path string true "Transaction date from"
// @Param to path string true "Transaction date to"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/{from}/date/{to} [get]
// @Tags Transaction
func GetTransactionFromDateTo(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	layout := "2006-01-02"
	from := c.Params("from")
	to := c.Params("to")
	fromDate, _ := time.Parse(layout, from)
	toDate, _ := time.Parse(layout, to)
	// fmt.Println(t)
	// t.Format("2006-01-02")
	// d := strfmt.Date(string(date))

	data := model.Transaction{}
	mod := db.Model(&data).
		Where(`date >= ? and date <= ?`, fromDate, toDate).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		Preload("User.Role").
		Preload("User.Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, &page)
}

// GetTransactionMonth godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param month path string true "Transaction month"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/month/{month} [get]
// @Tags Transaction

func GetTransactionMonth(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	layout := "2006-01-02"
	month := c.Params("month")
	t, _ := time.Parse(layout, month)
	fmt.Println(t)
	// t.Format("2006-01-02")
	// d := strfmt.Date(string(month))

	data := model.Transaction{}
	mod := db.Model(&data).
		Where(`month(date) = ?`, t).
		Joins("User").
		Joins("Payment").
		Joins("Category").
		Joins("Balance").
		Joins("Group").
		Preload("User.Role").
		Preload("User.Group")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Transaction{})

	return lib.OK(c, &page)
}
