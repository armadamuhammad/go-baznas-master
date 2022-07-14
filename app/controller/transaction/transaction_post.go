package transaction

import (
	balance "api/app/controller/balance"
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
	api := new(model.TransactionAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	now := time.Now()

	db := services.DB

	var data model.Transaction
	lib.Merge(api, &data)
	if nil == data.BalanceID {
		id := data.CategoryID
		var cat model.Category
		db.Model(&cat).
			Where(db.Where(model.Category{
				Base: model.Base{
					ID: id,
				},
			})).
			First(&cat)
		data.BalanceID = cat.BalanceID
	}
	data.CreatorID = lib.GetXUserID(c)
	if nil == data.UserID {
		data.UserID = lib.GetXUserID(c)
	}
	data.Total = GetDiscount(*data.Amount, *data.Discount, *data.DiscountType)
	data.Total = GetTax(*data.Total, *data.Tax, *data.TaxType)
	data.Date = &now

	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	inv := "0000" + fmt.Sprint(data.Sort)
	data.InvoiceNumber = &inv

	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	if err := balance.UpdateBalance(data.BalanceID, data.Total, data.IsIncome, data.ID); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	return lib.OK(c, data)
}
