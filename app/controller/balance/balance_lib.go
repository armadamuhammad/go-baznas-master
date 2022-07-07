package balance

import (
	"api/app/model"
	"api/app/services"

	"github.com/google/uuid"
)

func ChangeBalance(trx model.Transaction) *model.Balance {
	data := getBalance(trx.BalanceID)

	// data.Amount = data.Amount - trx.Total
	// if trx.IsIncome == lib.Intptr(1) {

	// }

	return &data
}

func getBalance(id *uuid.UUID) model.Balance {
	db := services.DB
	var data model.Balance

	db.Model(&data).
		Where(db.Where(model.Balance{
			Base: model.Base{
				ID: id,
			},
		})).
		First(&data)
	return data
}