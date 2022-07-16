package setting

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutUserGroup godoc
// @Summary Update Group User
// @Description Assingn Group User
// @Param X-User-ID header string false "User ID"
// @Param id path string true "User ID"
// @Param group_id path string true "Group ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/user/{id}/group/{group_id} [put]
// @Tags Setting
func PutUserGroup(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))
	groupID, _ := uuid.Parse(c.Params("group_id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

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
	data.GroupID = &groupID
	data.GroupAssignedID = userID

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutUserRole godoc
// @Summary Update Role User
// @Description Assingn Role User
// @Param X-User-ID header string false "User ID"
// @Param id path string true "User ID"
// @Param role_id path string true "Role ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/user/{id}/role/{role_id} [put]
// @Tags Setting
func PutUserRole(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))
	roleID, _ := uuid.Parse(c.Params("role_id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

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
	data.RoleID = &roleID
	// data.RoleAssignedID = userID

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}
