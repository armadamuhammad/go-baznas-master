package user

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutUser(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Put("/users/:id", PutUser)

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
			RoleID:         nil,
			GroupID:        nil,
		},
	}

	initial2 := model.User{
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
			RoleID:         nil,
			GroupID:        nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/users/" + initial.ID.String()

	payload := `{
		"firstname": null,
		"lastname": null,
		"surname": null,
		"email": null,
		"username": null,
		"password": null,
		"address": null,
		"joindate": null,
		"gender": null,
		"status": null,
		"status_verified": null,
		"role": null,
		"groupid": null
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
	uri = "/users/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/users/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
