package balance

import (
	"api/app/model"
	"api/app/services"

	"github.com/google/uuid"
)

func IncomeBalance(trx *model.Transaction) *model.Balance {
	data := getBalance(trx.BalanceID)

	amount := *data.Amount
	total := *trx.Total
	*data.Amount = amount + total

	return data
}

func OutcomeBalance() {

}

func DeletedBalance() {

}

func getBalance(id *uuid.UUID) *model.Balance {
	db := services.DB
	var data model.Balance

	db.Model(&data).
		Where(db.Where(model.Balance{
			Base: model.Base{
				ID: id,
			},
		})).
		First(&data)
	return &data
}
