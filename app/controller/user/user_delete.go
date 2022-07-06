package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteUser godoc
// @Summary Delete User by id
// @Description Delete User by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id} [delete]
// @Tags User
func DeleteUser(c *fiber.Ctx) error {
	db := services.DB

	var data model.User
	result1 := db.Model(&data).Where("id = ?", c.Params("id")).First(&data)
	if result1.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	data.Status = lib.Intptr(0)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	db.Delete(&data)

	return lib.OK(c)
}
