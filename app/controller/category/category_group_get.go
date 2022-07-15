package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetCategoryGroup godoc
// @Summary List of Category on same level
// @Description List of Category
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Param id path string true "Category ID level to child"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Page{items=[]model.Category} List of Category
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories/group/{id} [get]
// @Tags Category
func GetCategoryGroup(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	level := c.Params("id")
	if level == "1" {
		mod := db.Model(&model.Category{}).
			Where(db.Where(model.Category{
				CategoryAPI: model.CategoryAPI{
					Level: lib.Intptr(1),
				},
			})).
			Joins("Balance")
		page := pg.With(mod).Request(c.Request()).Response(&[]model.Category{})

		return lib.OK(c, page)
	}

	parentID, _ := uuid.Parse(c.Params("id"))

	mod := db.Model(&model.Category{}).
		Where(db.Where(model.Category{
			CategoryAPI: model.CategoryAPI{
				ParentID: &parentID,
			},
		}))
	page := pg.With(mod).Request(c.Request()).Response(&[]model.Category{})

	return lib.OK(c, page)
}

// GetCategoryBalance godoc
// @Summary List of Category from balance ID
// @Description List of Category
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Param id path string true "balance ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Page{items=[]model.Category} List of Category
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /categories/Balance/{id} [get]
// @Tags Category
func GetCategoryBalance(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG
	balanceID, _ := uuid.Parse(c.Params("id"))

	mod := db.Model(&model.Category{}).
		Where(db.Where(model.Category{
			CategoryAPI: model.CategoryAPI{
				BalanceID: &balanceID,
			},
		})).
		Joins("Balance")
	page := pg.With(mod).Request(c.Request()).Response(&[]model.Category{})

	return lib.OK(c, page)
}
