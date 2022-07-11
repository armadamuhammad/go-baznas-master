package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutCategory(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Put("/categories/:id", PutCategory)

	initial := model.Category{
		CategoryAPI: model.CategoryAPI{
			Name:        nil,
			Code:        nil,
			Level:       nil,
			Category:    nil,
			IsIncome: nil,
			Description: nil,
			ParentID:    nil,
		},
	}

	initial2 := model.Category{
		CategoryAPI: model.CategoryAPI{
			Name:        nil,
			Code:        nil,
			Level:       nil,
			Category:    nil,
			IsIncome: nil,
			Description: nil,
			ParentID:    nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/categories/" + initial.ID.String()

	payload := `{
		"name": null,
		"code": null,
		"level": null,
		"category": null,
		"IsIncome": null,
		"description": null,
		"parentid": null
	}`

	headers := map[string]string{
		"Content-Type": "application/json",
		"X-User-ID":    "70fb0d70-05dd-46e2-b113-b56437a8b694",
	}

	response, body, err := lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, headers["X-User-ID"], body["modifier_id"], "Modifier ID")

	// test invalid json body
	response, _, err = lib.PutTest(app, uri, headers, "invalid json format")
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 400, response.StatusCode, "getting response code")

	// test update with non existing id
	uri = "/categories/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/categories/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
