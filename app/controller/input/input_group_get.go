package input

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetInputUser godoc
// @Summary Get a Input by id
// @Description Get a Input by id
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Input Input data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /inputs/user/{id} [get]
// @Tags Input
func GetInputUser(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	userID, _ := uuid.Parse(c.Params("id"))

	var data model.Input
	mod := db.Model(&data).
		Where(db.Where(model.Input{
			InputAPI: model.InputAPI{
				UserID: &userID,
			},
		})).
		Joins("Group").
		Joins("User").
		Joins("Category")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Input{})

	return lib.OK(c, page)

}

// GetInputCategory godoc
// @Summary Get group input by category id
// @Description Get a Input by id
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Input Input data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /inputs/category/{id} [get]
// @Tags Input
func GetInputCategory(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	categoryID, _ := uuid.Parse(c.Params("id"))

	var data model.Input
	mod := db.Model(&data).
		Where(db.Where(model.Input{
			InputAPI: model.InputAPI{
				CategoryID: &categoryID,
			},
		})).
		Joins("Group").
		Joins("User").
		Joins("Category")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Input{})

	return lib.OK(c, page)

}

// GetInputGroup godoc
// @Summary Get group input by Group id
// @Description Get a Input by id
// @Param id path string true "Group ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Input Input data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /inputs/group/{id} [get]
// @Tags Input
func GetInputGroup(c *fiber.Ctx) error {
	db := services.DB
	pg := services.PG

	GroupID, _ := uuid.Parse(c.Params("id"))

	var data model.Input
	mod := db.Model(&data).
		Where(db.Where(model.Input{
			InputAPI: model.InputAPI{
				GroupID: &GroupID,
			},
		})).
		Joins("Group").
		Joins("User").
		Joins("Category")

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Input{})

	return lib.OK(c, page)

}
