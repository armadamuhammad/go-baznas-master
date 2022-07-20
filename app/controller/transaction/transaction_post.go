package transaction

import (
	balance "api/app/controller/balance"
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// PostTransaction godoc
// @Summary Create new Transaction
// @Description Create new Transaction
// @Param X-User-ID header string false "User ID"
// @Param data body model.TransactionAPI true "Transaction data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction "Transaction data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions [post]
// @Tags Transaction
func PostTransaction(c *fiber.Ctx) error {
	db := services.DB
	api := new(model.TransactionAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	var data model.Transaction
	lib.Merge(api, &data)

	id := data.CategoryID
	var cat model.Category
	db.Model(&cat).
		Where(db.Where(model.Category{
			Base: model.Base{
				ID: id,
			},
		})).
		First(&cat)

	now := time.Now()
	userID := lib.GetXUserID(c)
	userDefault := user.GetUserDefault()

	ver, _ := user.GetUserData(userID)
	if nil == userID {
		if cat.Anonym == lib.Intptr(1) {
			data.UserID = userDefault
		} else {
			return lib.ErrorUnauthorized(c)
		}
	}
	if *ver.StatusVerified != 1 {
		return lib.ErrorUnauthorized(c)
	}

	if nil == data.BalanceID {
		data.BalanceID = cat.BalanceID
	}
	data.CreatorID = userID
	if nil == data.UserID {
		data.UserID = userID
	}
	data.Total = GetDiscount(*data.Amount, *data.Discount, *data.DiscountType)
	data.Total = GetTax(*data.Total, *data.Tax, *data.TaxType)
	if !balance.CheckBalance(data.BalanceID, data.Total, data.IsIncome) {
		return lib.ErrorBadRequest(c)
	}
	data.Date = &now
	userGroup, _ := user.GetUserData(data.UserID)
	if nil == data.GroupID {
		data.GroupID = userGroup.GroupID
	}

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	inv := "0000" + fmt.Sprint(*data.Sort)
	data.InvoiceNumber = &inv

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	if err := balance.UpdateBalance(data.BalanceID, data.Total, data.IsIncome, data.ID); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	return lib.OK(c, data)
}
