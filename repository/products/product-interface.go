package products

import "gin-gorm-ex/models"

type ProductRepositoryInterface interface {
	GetAll() ([]models.Product, error)
	GetById(id string) (models.Product, error)
	Delete(id string)
	GetAllWithCategories() ([]models.Product, error)
	GetByProductsByCategoryId(Id string) ([]models.Product, error)
	Update(models.Product)
	Create(models.Product) // TODO : Burda bir sorun var geriye değer döndüremiyorum
}
