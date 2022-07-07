package input

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutInput(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Put("/inputs/:id", PutInput)

	initial := model.Input{
		InputAPI: model.InputAPI{
			GroupID:    nil,
			UserID:     nil,
			CategoryID: nil,
		},
	}

	initial2 := model.Input{
		InputAPI: model.InputAPI{
			GroupID:    nil,
			UserID:     nil,
			CategoryID: nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/inputs/" + initial.ID.String()

	payload := `{
		"group_id": null,
		"user_id": null,
		"category_id": null
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
	uri = "/inputs/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/inputs/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
