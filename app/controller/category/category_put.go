package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutCategory godoc
// @Summary Update Category by id
// @Description Update Category by id
// @Param X-User-ID header string false "User ID"
// @Param Accept-Language header string false "2 character language code"
// @Param id path string true "Category ID"
// @Param data body model.CategoryAPI true "Category data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories/{id} [put]
// @Tags Category
func PutCategory(c *fiber.Ctx) error {
	api := new(model.CategoryAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Category
	result := db.Model(&data).
		Where(db.Where(model.Category{
			Base: model.Base{
				ID: &id,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
