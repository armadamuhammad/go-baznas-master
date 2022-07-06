package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestGetTransaction(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Get("/transactions", GetTransaction)

	initial := model.Transaction{
		TransactionAPI: model.TransactionAPI{
			Name:         nil,
			Description:  nil,
			Type:         nil,
			Amount:       new(float64),
			IsIncome:     new(int),
			Note:         new(string),
			Tax:          new(float64),
			TaxType:      new(int),
			Contact:      new(string),
			ContactName:  new(string),
			Discount:     new(float64),
			DiscountType: new(int),

		},
	}

	db.Create(&initial)

	uri := "/transactions?page=0&size=1"
	response, body, err := lib.GetTest(app, uri, nil)
	utils.AssertEqual(t, nil, err, "sending request")
	utils.AssertEqual(t, 200, response.StatusCode, "getting response code")
	utils.AssertEqual(t, false, nil == body, "validate response body")
	utils.AssertEqual(t, float64(1), body["total"], "getting response body")

	sqlDB, _ := db.DB()
	sqlDB.Close()
}
