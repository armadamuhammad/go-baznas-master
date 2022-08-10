package city

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetCityID godoc
// @Summary Get a City by id
// @Description Get a City by id
// @Param id path string true "City ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.City City data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /cities/{id} [get]
// @Tags City
func GetCityID(c *fiber.Ctx) error {
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

	return lib.OK(c, data)
}
