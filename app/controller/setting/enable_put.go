package setting

import (
	"api/app/controller/user"
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// PutBalanceEnable godoc
// @Summary Update status Balance by id
// @Description Update Balance by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Balance ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Balance "Balance data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/balance/{id}/enable [put]
// @Tags Setting
func PutBalanceEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.Super != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Balance
	result := db.Model(&data).
		Where(db.Where(model.Balance{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutCategoryEnable godoc
// @Summary Update status Category by id
// @Description Update Category by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Category ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Category "Category data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/category/{id}/enable [put]
// @Tags Setting
func PutCategoryEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Category
	result := db.Model(&data).
		Where(db.Where(model.Category{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutAccountEnable godoc
// @Summary Update status Account by id
// @Description Update Account by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Account ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Account "Account data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/account/{id}/enable [put]
// @Tags Setting
func PutAccountEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Account
	result := db.Model(&data).
		Where(db.Where(model.Account{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutGroupEnable godoc
// @Summary Update status Group by id
// @Description Update Group by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Group ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Group "Group data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/group/{id}/enable [put]
// @Tags Setting
func PutGroupEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Group
	result := db.Model(&data).
		Where(db.Where(model.Group{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutPaymentEnable godoc
// @Summary Update status Payment by id
// @Description Update Payment by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "Payment ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Payment "Payment data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/payment/{id}/enable [put]
// @Tags Setting
func PutPaymentEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.Payment
	result := db.Model(&data).
		Where(db.Where(model.Payment{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}

// PutUserEnable godoc
// @Summary Update status User by id
// @Description Update User by id
// @Param X-User-ID header string false "User ID"
// @Param id path string true "User ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.User "User data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /settings/user/{id}/enable [put]
// @Tags Setting
func PutUserEnable(c *fiber.Ctx) error {
	db := services.DB
	id, _ := uuid.Parse(c.Params("id"))

	userID := lib.GetXUserID(c)
	ver, _ := user.GetUserData(userID)
	if *ver.IsAdmin != 1 {
		return lib.ErrorUnauthorized(c)
	}

	var data model.User
	result := db.Model(&data).
		Where(db.Where(model.User{
			Base: model.Base{
				ID: &id,
			},
		})).First(&data)

	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	if *data.Status == 1 {
		data.Status = lib.Intptr(0)
	} else {
		data.Status = lib.Intptr(1)
	}
	if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}
	return lib.OK(c, data)
}
