package balance

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutBalance godoc
// @Summary Update Balance by id
// @Description Update Balance by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Balance ID"
// @Param data body model.BalanceAPI true "Balance data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Balance "Balance data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balances/{id} [put]
// @Tags Balance
func PutBalance(c *fiber.Ctx) error {
	db := services.DB
	api := new(model.BalanceAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	id, _ := uuid.Parse(c.Params("id"))
	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 && api.Amount != nil{
		return lib.ErrorUnauthorized(c)
	}
	
	var data model.Balance
	result := db.Model(&data).
		Where(db.Where(model.Balance{
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
