package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// PostUser godoc
// @Summary Create new User
// @Description Create new User
// @Param X-User-ID header string false "User ID"
// @Param data body model.UserAPI true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users [post]
// @Tags User
func PostUser(c *fiber.Ctx) error {
	api := new(model.UserAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	raw := viper.GetString("DEF_PASS")
	now := strfmt.DateTime(time.Now())

	db := services.DB

	var data model.User
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)
	data.GroupAssignedID = lib.GetXUserID(c)
	pass := lib.PasswordEncrypt(raw, salt, key)
	data.Password = &pass
	data.Status = lib.Intptr(1)
	if nil == data.StatusVerified {
		data.StatusVerified = lib.Intptr(1)
	}
	data.IsAdmin = lib.Intptr(0)
	data.JoinDate = &now
	data.Super = lib.Intptr(0)

	if data.RoleID == nil {
		defRole := getRoleDefault()
		data.RoleID = defRole.ID
	}
	if data.GroupID == nil {
		defGroup := getGroupDefault()
		data.GroupID = defGroup.ID
	}

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}

func getRoleDefault() *model.Role {
	db := services.DB
	code := viper.GetString("DEF_ROLE")
	var data model.Role
	result := db.Model(&data).
		Where(db.Where(model.Role{
			RoleAPI: model.RoleAPI{
				Code: &code,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return nil
	}
	return &data

}

func getGroupDefault() *model.Group {
	db := services.DB
	code := viper.GetString("DEF_GROUP")
	var data model.Group
	result := db.Model(&data).
		Where(db.Where(model.Group{
			GroupAPI: model.GroupAPI{
				Code: &code,
			},
		})).
		First(&data)
	if result.RowsAffected < 1 {
		return nil
	}
	return &data

}
