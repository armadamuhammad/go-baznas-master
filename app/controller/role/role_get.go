package role

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetRole godoc
// @Summary List of Role
// @Description List of Role
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Page{items=[]model.Role} List of Role
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /roles [get]
// @Tags Role
func GetRole(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	mod := db.Model(&model.Role{})
	page := pg.With(mod).Request(c.Request()).Response(&[]model.Role{})

	return lib.OK(c, page)
}
