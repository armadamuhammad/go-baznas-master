package group

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetGroupID godoc
// @Summary Get a Group by id
// @Description Get a Group by id
// @Param id path string true "Group ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Group Group data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /groups/{id} [get]
// @Tags Group
func GetGroupID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Group
	result := db.Model(&data).
		Where(db.Where(model.Group{
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
