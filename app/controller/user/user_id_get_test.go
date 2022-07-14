package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetUserID(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/users/:id", GetUserID)

	initial := model.User{
		UserAPI: model.UserAPI{
			FirstName: nil,
			LastName:  nil,
			Surname:   nil,
			Email:     nil,
			Username:  nil,
			Address:   nil,
			Gender:    nil,
			RoleID:    nil,
			GroupID:   nil,
		},
	}

	db.Create(&initial)

	uri := "/users/" + initial.ID.String()
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, initial.ID.String(), body["id"], "getting response body")

	// test get non existing id
	uri = "/users/non-existing-id"
	response, _, err = lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
