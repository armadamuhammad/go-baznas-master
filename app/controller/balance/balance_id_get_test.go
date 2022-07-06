package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetBalanceID(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/balances/:id", GetBalanceID)

	initial := model.Balance{
		BalanceAPI: model.BalanceAPI{
			Saldo:         nil,
			Category:      nil,
			Income:        nil,
			Outcome:       nil,
			TransactionID: nil,
			Description:   nil,
		},
	}

	db.Create(&initial)

	uri := "/balances/" + initial.ID.String()
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, initial.ID.String(), body["id"], "getting response body")

	// test get non existing id
	uri = "/balances/non-existing-id"
	response, _, err = lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
