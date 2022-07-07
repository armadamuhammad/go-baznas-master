package balancerecord

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestPutBalanceRecord(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Put("/balance-records/:id", PutBalanceRecord)

	initial := model.BalanceRecord{
		BalanceRecordAPI: model.BalanceRecordAPI{
			Amount:        nil,
			Datetime:      nil,
			BalanceID:     nil,
			TransactionID: nil,
		},
	}

	initial2 := model.BalanceRecord{
		BalanceRecordAPI: model.BalanceRecordAPI{
			Amount:        nil,
			Datetime:      nil,
			BalanceID:     nil,
			TransactionID: nil,
		},
	}

	db.Create(&initial)
	db.Create(&initial2)

	uri := "/balance-records/" + initial.ID.String()

	payload := `{
		"amount": null,
		"datetime": null,
		"balance_id": null,
		"transaction_id": null
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
	uri = "/balance-records/non-existing-id"
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 404, response.StatusCode, "getting response code")

	// test duplicate data
	uri = "/balance-records/" + initial2.ID.String()
	response, _, err = lib.PutTest(app, uri, headers, payload)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 409, response.StatusCode, "getting response code")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
