package products

import (
	"errors"
	"gin-gorm-ex/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetAll() ([]models.Product, error) {

	products := []models.Product{}

	if err := p.db.Where("is_deleted", "0").Find(&products).Error; err != nil {
		return products, errors.New("Failed get data")
	}
	return products, nil
}

func (p *ProductRepository) GetById(id string) (models.Product, error) {

	product := models.Product{}
	err := p.db.Where("id=?", id).Find(&product).Error

	if err != nil {
		return product, errors.New("Failed get data")
	}
	return product, nil
}

func (p *ProductRepository) Update(product models.Product) {

	updatedProduct := models.Product{}

	p.db.Where("id=?", product.ID).Find(&updatedProduct)

	p.db.Model(&updatedProduct).Updates(product)
}

func (p *ProductRepository) Delete(id string) {

	product := models.Product{}

	p.db.Where("id=?", id).Find(&product)
	p.db.Model(&product).Updates(models.Product{IsDeleted: true})
}

func (p *ProductRepository) Create(product models.Product) {

	_ = p.db.Create(&product)
}

func (p *ProductRepository) GetAllWithCategories() ([]models.Product, error) {

	var products []models.Product
	p.db.Model(&models.Product{}).Preload("Category").Find(&products)
	return products, nil
}

func (p *ProductRepository) GetByProductsByCategoryId(Id string) ([]models.Product, error) {

	var products []models.Product
	p.db.Model(&products).Preload("Categoery").Where("category_id=?", Id).Find(&products)

	return products, nil
}
