package login

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ChangePassword godoc
// @Summary Update Password user
// @Description Update User by id
// @Param X-User-ID header string false "User ID"
// @Param data body model.UserPassword true "User data"
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /change-password/{id} [put]
// @Tags Login
func ChangePassword(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))
	api := new(model.UserPassword)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	old := *api.PasswordOld

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
	hashed := *data.Password
	if b := lib.PasswordCompare(hashed, old, salt, key); !b {
		return lib.ErrorBadRequest(c)
	}
	raw := *api.PasswordNew
	newPassword := lib.PasswordEncrypt(raw, salt, key)
	data.Password = &newPassword

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}
