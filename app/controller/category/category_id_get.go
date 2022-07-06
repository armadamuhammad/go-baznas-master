package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetCategoryID godoc
// @Summary Get a Category by id
// @Description Get a Category by id
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category Category data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories/{id} [get]
// @Tags Category
func GetCategoryID(c *fiber.Ctx) error {
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

	return lib.OK(c, data)
}
