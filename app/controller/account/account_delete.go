package account

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// DeleteAccount godoc
// @Summary Delete Account by id
// @Description Delete Account by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Account ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /accounts/{id} [delete]
// @Tags Account
func DeleteAccount(c *fiber.Ctx) error {
	db := services.DB

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Account
	result1 := db.Model(&data).Where("id = ?", c.Params("id")).First(&data)
	if result1.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	db.Delete(&data)

	return lib.OK(c)
}
