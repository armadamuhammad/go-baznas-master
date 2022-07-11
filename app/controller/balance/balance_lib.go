package balance

import (
	"api/app/model"
	"api/app/services"

	"github.com/google/uuid"
)

func UpdateBalance(trx *uuid.UUID, amount *float64, isIncome *int) error{
	data := getBalance(trx)

	if *isIncome == 1 {
		trxAmount := *amount
		total := *data.Amount
		x := total + trxAmount
		data.Amount = &x
	} else {
		trxAmount := *amount
		total := *data.Amount
		x := total - trxAmount
		data.Amount = &x
	}

	
	return nil
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
