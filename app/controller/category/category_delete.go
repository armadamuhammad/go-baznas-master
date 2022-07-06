package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteCategory godoc
// @Summary Delete Category by id
// @Description Delete Category by id
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories/{id} [delete]
// @Tags Category
func DeleteCategory(c *fiber.Ctx) error {
	db := services.DB

	var data model.Category
	result1 := db.Model(&data).Where("id = ?", c.Params("id")).First(&data)
	if result1.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&data)

	return lib.OK(c)
}
