package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestDeleteTransaction(t *testing.T) {
	db := services.DBConnectTest()
	app := fiber.New()
	app.Delete("/transactions/:id", DeleteTransaction)

	initial := model.Transaction{
		TransactionAPI: model.TransactionAPI{
			Name:          nil,
			Description:   nil,
			InvoiceNumber: nil,
			NoRef:         nil,
			Type:          nil,
			Income:        nil,
			Outcome:       nil,
			Total:         nil,
		},
	}

	db.Create(&initial)

	uri := "/transactions/" + initial.ID.String()
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
