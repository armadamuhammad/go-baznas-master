package role

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// GetRoleID godoc
// @Summary Get a Role by id
// @Description Get a Role by id
// @Param id path string true "Role ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Role Role data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /roles/{id} [get]
// @Tags Role
func GetRoleID(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	var data model.Role
	result := db.Model(&data).
		Where(db.Where(model.Role{
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

func GetRoleDefault() *uuid.UUID {
	db := services.DB

	var data model.Role
	db.Model(&data).
		Where(db.Where(model.Role{
			RoleAPI: model.RoleAPI{
				Code: lib.Strptr(viper.GetString("DEF_ROLE")),
			},
		})).First(&data)
	return data.ID
}
