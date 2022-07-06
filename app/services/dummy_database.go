//go:build !production
// +build !production

package services

import (
	"api/app/migrations"

	"github.com/morkid/paginate"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DBConnectTest test
func DBConnectTest(database ...string) *gorm.DB {
	dbPath := "file::memory:"
	if len(database) > 0 {
		dbPath = database[0]
		if dbPath == "" {
			dbPath = "file::memory:"
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
	})
	if nil != err {
		panic(err)
	}

	err = db.AutoMigrate(migrations.ModelMigrations...)
	if nil != err {
		panic(err)
	}
	DB = db
	PG = paginate.New()

	return db
}
