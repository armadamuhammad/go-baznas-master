package input

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetInputID godoc
// @Summary Get a Input by id
// @Description Get a Input by id
// @Param id path string true "Input ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Input Input data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /inputs/{id} [get]
// @Tags Input
func GetInputID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Input
	result := db.Model(&data).
		Where(db.Where(model.Input{
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
