package account

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostAccount godoc
// @Summary Create new Account
// @Description Create new Account
// @Param X-User-ID header string false "User ID"
// @Param data body model.AccountAPI true "Account data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Account "Account data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /accounts [post]
// @Tags Account
func PostAccount(c *fiber.Ctx) error {
	api := new(model.AccountAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	
	db := services.DB

	var data model.Account
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
