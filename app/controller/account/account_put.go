package account

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutAccount godoc
// @Summary Update Account by id
// @Description Update Account by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Account ID"
// @Param data body model.AccountAPI true "Account data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Account "Account data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /accounts/{id} [put]
// @Tags Account
func PutAccount(c *fiber.Ctx) error {
	api := new(model.AccountAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 {
		return lib.ErrorUnauthorized(c)
	}

	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Account
	result := db.Model(&data).
		Where(db.Where(model.Account{
			Base: model.Base{
				ID: &id,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)
	data.ModifierID = lib.GetXUserID(c)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
