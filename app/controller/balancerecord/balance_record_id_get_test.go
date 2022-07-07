package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetBalanceRecordID(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/balance-records/:id", GetBalanceRecordID)

	initial := model.BalanceRecord{
		BalanceRecordAPI: model.BalanceRecordAPI{
			Amount:        nil,
			Datetime:      nil,
			BalanceID:     nil,
			TransactionID: nil,
		},
	}

	db.Create(&initial)

	uri := "/balance-records/" + initial.ID.String()
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, initial.ID.String(), body["id"], "getting response body")

	// test get non existing id
	uri = "/balance-records/non-existing-id"
	response, _, err = lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
