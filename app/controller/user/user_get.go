package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetUser godoc
// @Summary List of User
// @Description List of User
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Page{items=[]model.User} List of User
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /users [get]
// @Tags User
func GetUser(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	mod := db.Model(&model.User{}).
		Joins("Role").
		Joins("Group").
		Joins("GroupAssigned").
		Joins("City")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.User{})

	return lib.OK(c, page)
}
