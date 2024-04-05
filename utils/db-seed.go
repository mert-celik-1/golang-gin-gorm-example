package utils

import (
	"errors"
	"gin-gorm-ex/models"

	"gorm.io/gorm"
)

func seedProduct() {
	if err := db.First(&models.Product{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&models.Product{Name: "Product 1", Quantity: 15, IsDeleted: false, CategoryId: 1})
		db.Create(&models.Product{Name: "Product 2", Quantity: 2, IsDeleted: false, CategoryId: 2})
		db.Create(&models.Product{Name: "Product 3", Quantity: 5, IsDeleted: false, CategoryId: 2})
		db.Create(&models.Product{Name: "Product 4", Quantity: 12, IsDeleted: false, CategoryId: 1})
		db.Create(&models.Product{Name: "Product 5", Quantity: 83, IsDeleted: false, CategoryId: 1})
	}
}

func seedUser() {
	if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		token, refreshToken, _ := GenerateAllTokens("mert@gmail.com")
		db.Create(&models.User{Username: "mertcelik", Email: "mert@gmail.com", Password: HashPassword("12345"), Token: token, Refresh_Token: refreshToken})
	}
}

func seedCategory() {
	if err := db.First(&models.Category{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		db.Create(&models.Category{Name: "Category 1", ID: 1})
		db.Create(&models.Category{Name: "Category 2", ID: 2})
	}
}
