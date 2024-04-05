package categories

import "gin-gorm-ex/models"

type CategoryRepositoryInterface interface {
	GetAll() ([]models.Category, error)
}
