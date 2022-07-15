package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

// UpdateBalance func
func UpdateBalance(id *uuid.UUID, amount *float64, isIncome *int, trx *uuid.UUID) error {
	data, err := getBalance(id)
	if nil != err {
		return err
	}
	db := services.DB
	var types string

	if *isIncome == 1 {
		trxAmount := *amount
		total := *data.Amount
		x := total + trxAmount
		data.Amount = &x
		types = "income"

	} else {
		trxAmount := *amount
		total := *data.Amount
		x := total - trxAmount
		data.Amount = &x
		types = "outcome"
	}

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return err
	}
	BalanceRecord(data.ID, data.Amount, trx, &types)
	return nil
}

// DeletedBalance func for delete and put transaction
func DeletedBalance(trx *uuid.UUID, amount *float64, isIncome *int) error {
	data, err := getBalance(trx)
	if nil != err {
		return err
	}
	db := services.DB

	if *isIncome == 1 {
		trxAmount := *amount
		total := *data.Amount
		x := total - trxAmount
		data.Amount = &x
	} else {
		trxAmount := *amount
		total := *data.Amount
		x := total + trxAmount
		data.Amount = &x
	}

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return err
	}
	BalanceRecord(data.ID, data.Amount, trx, lib.Strptr("delete trx"))

	return nil
}

// getBalance by id
func getBalance(id *uuid.UUID) (*model.Balance, error) {
	db := services.DB
	var data model.Balance

	res := db.Model(&data).
		Where(db.Where(model.Balance{
			Base: model.Base{
				ID: id,
			},
		})).
		First(&data)
	if res.RowsAffected < 1 {
		return nil, res.Error
	}
	return &data, nil
}

func CheckBalance(id *uuid.UUID, amount *float64, isIncome *int) bool {
	if *isIncome == 1 {
		return true
	}
	valid, _ := getBalance(id)
	return (*valid.Amount - *amount) >= 0

}

// BalanceRecord create balance record in every action
func BalanceRecord(balanceID *uuid.UUID, amount *float64, trxID *uuid.UUID, tipe *string) {
	db := services.DB
	now := strfmt.DateTime(time.Now())

	data := model.BalanceRecord{
		BalanceRecordAPI: model.BalanceRecordAPI{
			Amount:        amount,
			Type:          tipe,
			Datetime:      &now,
			BalanceID:     balanceID,
			TransactionID: trxID,
		},
	}

	db.Create(&data)
}
