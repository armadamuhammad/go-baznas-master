package input

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostInput godoc
// @Summary Create new Input
// @Description Create new Input
// @Param X-User-ID header string false "User ID"
// @Param data body model.InputAPI true "Input data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Input "Input data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /inputs [post]
// @Tags Input
func PostInput(c *fiber.Ctx) error {
	api := new(model.InputAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Input
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
