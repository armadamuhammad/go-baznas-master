package role

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetRole(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/roles", GetRole)

	initial := model.Role{
		RoleAPI: model.RoleAPI{
			Name:        nil,
			Code:        nil,
			Description: nil,
		},
	}

	db.Create(&initial)

	uri := "/roles?page=0&size=1"
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, float64(1), body["total"], "getting response body")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
