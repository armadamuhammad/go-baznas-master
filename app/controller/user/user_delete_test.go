package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteUser(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Delete("/users/:id", DeleteUser)

	initial := model.User{
		UserAPI: model.UserAPI{
			FirstName:      nil,
			LastName:       nil,
			Surname:        nil,
			Email:          nil,
			Username:       nil,
			Password:       nil,
			Address:        nil,
			JoinDate:       nil,
			Gender:         nil,
			Status:         nil,
			StatusVerified: nil,
			RoleID:           nil,
			GroupID:        nil,
		},
	}

	db.Create(&initial)

	uri := "/users/" + initial.ID.String()
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
