package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var context *gorm.DB

const connStr string = "./repositories/nion_test.db"

func InitializeDB() {
	gormDB, err := gorm.Open(sqlite.Open(connStr), getConfig())
	if err != nil {
		panic("failed to connect database")
	}

	context = gormDB
	migratesIfNoTable(gormDB)
}

func getConfig() *gorm.Config {
	return &gorm.Config{}
}

func migratesIfNoTable(gormDB *gorm.DB) {
	isExist := gormDB.Migrator().HasTable(&CustomerEntity{})
	if !isExist {
		gormDB.AutoMigrate(&CustomerEntity{})

	}
}
