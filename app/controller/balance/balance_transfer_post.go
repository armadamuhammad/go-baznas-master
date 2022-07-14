package balance

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PostBalance godoc
// @Summary Create new Balance
// @Description Create new Balance
// @Param X-User-ID header string false "User ID"
// @Param data body model.BalanceTransfer true "Balance data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.BalanceTransfer "Balance data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /balance/transfer [post]
// @Tags Balance
func PostBalanceTransfer(c *fiber.Ctx) error {
	db := services.DB
	api := new(model.BalanceTransfer)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	userID := lib.GetXUserID(c)

	trx := model.Transaction{
		TransactionAPI: model.TransactionAPI{
			Name:         lib.Strptr("Transfer Balance"),
			Description:  nil,
			Type:         lib.Strptr("transfer"),
			Amount:       api.Amount,
			IsIncome:     lib.Intptr(1),
			Note:         nil,
			Tax:          nil,
			TaxType:      nil,
			Contact:      nil,
			ContactName:  nil,
			Discount:     nil,
			DiscountType: nil,
			UserID:       userID,
			CategoryID:   nil,
			BalanceID:    api.To,
		},
	}
	db.Create(&trx)
	UpdateBalance(api.From, api.Amount, lib.Intptr(0), trx.ID)
	UpdateBalance(api.To, api.Amount, lib.Intptr(1), trx.ID)

	return lib.OK(c, api)
}
