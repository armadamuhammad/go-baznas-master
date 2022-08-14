package city

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutCity godoc
// @Summary Update City by id
// @Description Update City by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "City ID"
// @Param data body model.CityAPI true "City data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.City "City data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /cities/{id} [put]
// @Tags City
func PutCity(c *fiber.Ctx) error {
	api := new(model.CityAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.City
	result := db.Model(&data).
		Where(db.Where(model.City{
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
