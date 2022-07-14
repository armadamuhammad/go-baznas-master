package balance

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostBalance godoc
// @Summary Create new Balance
// @Description Create new Balance
// @Param X-User-ID header string false "User ID"
// @Param data body model.BalanceAPI true "Balance data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Balance "Balance data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balances [post]
// @Tags Balance
func PostBalance(c *fiber.Ctx) error {
	api := new(model.BalanceAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 {
		return lib.ErrorUnauthorized(c)
	}

	db := services.DB

	var data model.Balance
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
