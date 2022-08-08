package transaction

import (
	"api/app/controller/login"
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionFromDateToGroup godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id {account, category, balance, group} format date {2022-12-30}
// @Param from path string true "Transaction date from"
// @Param to path string true "Transaction date to"
// @Param group path string true "group filter"
// @Param id path string true "ID"
// @Param X-User-ID header string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/{from}/date/{to}/group/{group}/{id} [get]
// @Tags Transaction
func GetTransactionFromDateToFilter(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	layout := "2006-01-02"
	from := c.Params("from")
	to := c.Params("to")
	group := c.Params("group")
	id := c.Params("id")
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
		Where(`"transaction".`+group+`_id = ?`, id).
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
