package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteCategory(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Delete("/categories/:id", DeleteCategory)

	initial := model.Category{
		CategoryAPI: model.CategoryAPI{
			Name:        nil,
			Code:        nil,
			Level:       nil,
			Category:    nil,
			IsPemasukan: nil,
			Description: nil,
			ParentID:    nil,
		},
	}

	db.Create(&initial)

	uri := "/categories/" + initial.ID.String()
	response, _, err := lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")

	// test delete with non existing id
	response, _, err = lib.DeleteTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
