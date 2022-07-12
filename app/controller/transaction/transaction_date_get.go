package transaction


// GetTransactionDate godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "Transaction ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Transaction Transaction data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /transactions/{id} [get]
// @Tags Transaction