package group

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostGroup godoc
// @Summary Create new Group
// @Description Create new Group
// @Param X-User-ID header string false "User ID"
// @Param data body model.GroupAPI true "Group data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Group "Group data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /groups [post]
// @Tags Group
func PostGroup(c *fiber.Ctx) error {
	api := new(model.GroupAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Group
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
