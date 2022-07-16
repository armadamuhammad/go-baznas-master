package setting

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PostViewCategory godoc
// @Summary Create new view by caegory
// @Description Create new setting view by caegory
// @Param X-User-ID header string false "User ID"
// @Param data body model.ViewAPI true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/view/category [post]
// @Tags Setting
func PostViewCategory(c *fiber.Ctx) error {
	db := services.DB

	userID := lib.GetXUserID(c)
	var data model.ViewCategory
	db.Unscoped().Where(`user_id = ?`, userID).Delete(&data)

	api := new(model.ViewAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	// raw := []model.ViewCategory{}
	for _, raw := range *api.Items {
		id, _ := uuid.Parse(raw)
		each := model.ViewCategory{
			UserID:     userID,
			CategoryID: &id,
		}

		db.Create(&each)
	}

	return lib.OK(c)
}

// PostViewGroup godoc
// @Summary Create new view by Group
// @Description Create new setting view by Group
// @Param X-User-ID header string false "User ID"
// @Param data body model.ViewAPI true "User data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Response
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/view/group [post]
// @Tags Setting
func PostViewGroup(c *fiber.Ctx) error {
	db := services.DB

	userID := lib.GetXUserID(c)
	var data model.ViewGroup
	db.Unscoped().Where(`user_id = ?`, userID).Delete(&data)

	api := new(model.ViewAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	// raw := []model.ViewGroup{}
	for _, raw := range *api.Items {
		id, _ := uuid.Parse(raw)
		each := model.ViewGroup{
			UserID:  userID,
			GroupID: &id,
		}

		db.Create(&each)
	}

	return lib.OK(c)
}
