package categories

import (
	"errors"
	"gin-gorm-ex/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c *CategoryRepository) GetAll() ([]models.Category, error) {

	categories := []models.Category{}

	if err := c.db.Find(&categories).Error; err != nil {
		return categories, errors.New("Failed get data")
	}

	return categories, nil
}
