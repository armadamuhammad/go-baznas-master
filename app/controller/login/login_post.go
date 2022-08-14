package login

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Login godoc
// @Summary Login Session
// @Description Create new Session
// @Param data body model.UserLogin true "User Login"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.UserAfterLogin User data
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Router /login [post]
// @Tags Login
func Login(c *fiber.Ctx) error {
	salt := "salt"
	key := "CIPHER_SECRETKEY_MUST_HAVE_32BIT"
	db := services.DB
	api := new(model.UserLogin)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	email := api.Email
	raw := *api.Password
	// password := lib.PasswordEncrypt(raw, salt, key)
	// password := raw + "//" + salt + "//" + key

	var data model.User
	result := db.Model(&data).
		Where(db.Where(model.User{
			UserAPI: model.UserAPI{
				Email: email,
			},
		})).
		Joins("Role").
		Joins("Group").
		Joins("GroupAssigned").
		Joins("City").
		First(&data)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}
	hashed := *data.Password
	b := lib.PasswordCompare(hashed, raw, salt, key)

	if !b {
		return lib.ErrorBadRequest(c)
	}
	var res model.UserAfterLogin
	lib.Merge(data, &res)
	res.Password = nil
	res.CategoryView = GetCategory(res.ID)
	res.GroupView = GetGroup(res.ID)

	return lib.OK(c, &res)
}

func GetCategory(userID *uuid.UUID) *[]string {
	db := services.DB
	rows, err := db.Model(&model.ViewCategory{}).
		Where(db.Where(&model.ViewCategory{
			UserID: userID,
		})).Rows()
	result := []string{}
	if nil == err {
		for rows.Next() {
			each := model.ViewCategory{}
			db.ScanRows(rows, &each)
			var desc string
			if each.CategoryID != nil {
				desc = fmt.Sprint(*each.CategoryID)
			}
			result = append(result, desc)
		}
		defer rows.Close()
	}
	return &result
}

func GetGroup(userID *uuid.UUID) *[]string {
	db := services.DB
	rows, err := db.Model(&model.ViewGroup{}).
		Where(db.Where(&model.ViewGroup{
			UserID: userID,
		})).Rows()
	result := []string{}
	if nil == err {
		for rows.Next() {
			each := model.ViewGroup{}
			db.ScanRows(rows, &each)
			var desc string
			if each.GroupID != nil {
				desc = fmt.Sprint(*each.GroupID)
			}
			result = append(result, desc)
		}
		defer rows.Close()
	}
	return &result
}
