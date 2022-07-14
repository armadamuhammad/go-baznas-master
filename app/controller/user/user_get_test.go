package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetUser(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/users", GetUser)

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

	uri := "/users?page=0&size=1"
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, float64(1), body["total"], "getting response body")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
