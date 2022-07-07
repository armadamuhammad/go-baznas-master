package group

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutGroup godoc
// @Summary Update Group by id
// @Description Update Group by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Group ID"
// @Param data body model.GroupAPI true "Group data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Group "Group data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /groups/{id} [put]
// @Tags Group
func PutGroup(c *fiber.Ctx) error {
	api := new(model.GroupAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

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

	lib.Merge(api, &data)
	data.ModifierID = lib.GetXUserID(c)

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
