package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetBalance(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/balances", GetBalance)

	initial := model.Balance{
		BalanceAPI: model.BalanceAPI{
			Amount:      new(float64),
			Name:        new(string),
			Code:        new(string),
			Description: nil,
		},
	}

	db.Create(&initial)

	uri := "/balances?page=0&size=1"
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, float64(1), body["total"], "getting response body")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
