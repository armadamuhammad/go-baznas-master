package transaction

import (
	"api/app/controller/login"
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
// @Param X-User-ID header string true "User ID"
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
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)

	layout := "2006-01-02"
	date := c.Params("date")
	t, _ := time.Parse(layout, date)
	fmt.Println(t)
	// t.Format("2006-01-02")
	// d := strfmt.Date(string(date))

	data := model.Transaction{}
	mod := db.Model(&data).
		Where(`date = ?`, t).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
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
// @Param X-User-ID header string true "User ID"
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
	userID := lib.GetXUserID(c)
	categories := login.GetCategory(userID)
	groups := login.GetGroup(userID)
	// fmt.Println(t)
	// t.Format("2006-01-02")
	// d := strfmt.Date(string(date))

	data := model.Transaction{}
	mod := db.Model(&data).
		Where(`date >= ? and date <= ?`, fromDate, toDate).
		Where(`"transaction".category_id IN ? OR "transaction".group_id IN ?`, *categories, *groups).
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