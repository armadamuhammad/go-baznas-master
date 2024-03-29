package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteBalance(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Delete("/balances/:id", DeleteBalance)

	initial := model.Balance{
		BalanceAPI: model.BalanceAPI{
			Amount:      new(float64),
			Name:        new(string),
			Code:        new(string),
			Description: nil,
		},
	}

	db.Create(&initial)

	uri := "/balances/" + initial.ID.String()
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
