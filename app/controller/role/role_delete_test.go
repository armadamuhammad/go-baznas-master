package role

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteRole(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Delete("/roles/:id", DeleteRole)

	initial := model.Role{
		RoleAPI: model.RoleAPI{
			Name:        nil,
			Code:        nil,
			Description: nil,
		},
	}

	db.Create(&initial)

	uri := "/roles/" + initial.ID.String()
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
