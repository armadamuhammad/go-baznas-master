package login

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary Login Session
// @Description Create new Session
// @Param data body model.UserLogin true "User Login"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User User data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /login [post]
// @Tags Login
func Login(c *fiber.Ctx) error {
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	db := services.DB
	api := new(model.UserLogin)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	email := api.Email
	raw := *api.Password
	// password := lib.PasswordEncrypt(raw, salt, key)
	// password := raw + "//" + salt + "//" + key

	var data model.User
	result := db.Model(&data).
		Where(db.Where(model.User{
			UserAPI: model.UserAPI{
				Email: email,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}
	hashed := *data.Password
	b := lib.PasswordCompare(hashed, raw, salt, key)

	if !b {
		return lib.ErrorBadRequest(c)
	}
	return lib.OK(c, &data)
}
