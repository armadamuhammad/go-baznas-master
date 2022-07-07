package group

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteGroup godoc
// @Summary Delete Group by id
// @Description Delete Group by id
// @Param id path string true "Group ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /groups/{id} [delete]
// @Tags Group
func DeleteGroup(c *fiber.Ctx) error {
	db := services.DB

	var data model.Group
	result1 := db.Model(&data).Where("id = ?", c.Params("id")).First(&data)
	if result1.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&data)

	return lib.OK(c)
}
