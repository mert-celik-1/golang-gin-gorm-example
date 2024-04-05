package utils

import (
	"gin-gorm-ex/configs"
	"gin-gorm-ex/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDb(config *configs.AppConfig) *gorm.DB {

	Connect()
	InitialMigration(db)

	return db
}

func GetDB() *gorm.DB {
	return db
}

func InitialMigration(db *gorm.DB) {
	db.Statement.RaiseErrorOnNotFound = true
	db.AutoMigrate(models.Product{}, models.User{})
	seedCategory()
	seedProduct()
	seedUser()
}

func Connect() {

	d, err := gorm.Open(mysql.Open("root:root123@tcp(localhost:3306)/test?parseTime_true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}
