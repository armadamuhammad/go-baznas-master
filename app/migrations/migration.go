package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Sample{},
	&model.Account{},
	&model.Payment{},
	&model.Category{},

}
