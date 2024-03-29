package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Account{},
	&model.Payment{},
	&model.Category{},
	&model.City{},
	&model.User{},
	&model.Transaction{},
	&model.Balance{},
	&model.BalanceRecord{},
	&model.Group{},
	&model.Input{},
	&model.Role{},
	&model.ViewCategory{},
	&model.ViewGroup{},
}
