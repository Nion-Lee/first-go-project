package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var context *gorm.DB

const connStr string = "./repositories/nion_test.db"

func InitializeDB() {
	db, err := gorm.Open(sqlite.Open(connStr), getConfig())
	if err != nil {
		panic("failed to connect database")
	}

	context = db
}

func getConfig() *gorm.Config {
	return &gorm.Config{}
}
