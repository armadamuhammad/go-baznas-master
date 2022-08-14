package city

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

// PostCity godoc
// @Summary Create new City
// @Description Create new City
// @Param X-User-ID header string false "User ID"
// @Param data body model.CityAPI true "City data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.City "City data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /cities [post]
// @Tags City
func PostCity(c *fiber.Ctx) error {
	api := new(model.CityAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	cityCode := *api.Code
	length := len([]rune(cityCode))
	log.Println(length)

	db := services.DB

	var data model.City
	lib.Merge(api, &data)
	data.CreatorID = lib.GetXUserID(c)

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

	return lib.OK(c, data)
}
