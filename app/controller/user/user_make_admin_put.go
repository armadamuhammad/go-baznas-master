package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutUserMakeAdmin godoc
// @Summary Update status User by id
// @Description Update User by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users/{id}/make-admin [post]
// @Tags User
func PutUserMakeAdmin(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.User
	result := db.Model(&data).
		Where(db.Where(model.User{
			Base: model.Base{
				ID: &id,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}
	data.IsAdmin = lib.Intptr(1)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}
