package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostCategory godoc
// @Summary Create new Category
// @Description Create new Category
// @Param X-User-ID header string false "User ID"
// @Param data body model.CategoryAPI true "Category data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories [post]
// @Tags Category
func PostCategory(c *fiber.Ctx) error {
	api := new(model.CategoryAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Category
	lib.Merge(api, &data)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
