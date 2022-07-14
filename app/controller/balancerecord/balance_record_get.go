package balancerecord

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetBalanceRecord godoc
// @Summary List of Balance Record
// @Description List of Balance Record
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Param X-User-ID header string false "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Page{items=[]model.BalanceRecord} List of Balance Record
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance-records [get]
// @Tags BalanceRecord
func GetBalanceRecord(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 {
		return lib.ErrorUnauthorized(c)
	}

	mod := db.Model(&model.BalanceRecord{}).
		Joins("Balance").
		Joins("Transaction")
	page := pg.With(mod).Request(c.Request()).Response(&[]model.BalanceRecord{})

	return lib.OK(c, page)
}
