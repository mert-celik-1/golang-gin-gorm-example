package controllers

import (
	"gin-gorm-ex/common/exceptions"
	"gin-gorm-ex/common/mapping"
	"gin-gorm-ex/common/response"
	"gin-gorm-ex/models"
	"gin-gorm-ex/repository/products"
	"gin-gorm-ex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	Repo products.ProductRepositoryInterface
}

func NewProductControllers(prep products.ProductRepositoryInterface) *ProductsController {
	return &ProductsController{Repo: prep}
}

func (prep ProductsController) GetAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := prep.Repo.GetAll()

		var productResponse []response.ProductGetResponse

		productResponse = mapping.MapProductToProductGetResponse(result)

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    productResponse,
		}

		createResponse(response, err, utils.GetFunctionName(prep.GetAllProducts()), c)
	}

}

func (prep ProductsController) GetAllProductsWithCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := prep.Repo.GetAllWithCategories()

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		createResponse(response, err, utils.GetFunctionName(prep.GetAllProductsWithCategories()), c)
	}

}

func (prep ProductsController) GetByCategoryId() gin.HandlerFunc {
	return func(c *gin.Context) {

		catId := c.Param("catid")

		result, err := prep.Repo.GetByProductsByCategoryId(catId)

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		createResponse(response, err, utils.GetFunctionName(prep.GetAllProductsWithCategories()), c)
	}

}

func (prep ProductsController) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {

		productId := c.Param("id")

		result, err := prep.Repo.GetById(productId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, exceptions.NotFoundException{}.CreateNew(err))
		}

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		createResponse(response, err, utils.GetFunctionName(prep.GetProductById()), c)
	}
}

func (prep ProductsController) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		productId := c.Param("id")

		prep.Repo.Delete(productId)

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
		}

		createResponse(response, nil, utils.GetFunctionName(prep.DeleteProduct()), c)
	}
}

func (prep ProductsController) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var updateProduct response.ProductUpdateResponse

		if err := c.Bind(&updateProduct); err != nil {
			c.JSON(http.StatusBadRequest, exceptions.BadRequestException{}.CreateNew(err))
		}

		Product := models.Product{
			Name:     updateProduct.Name,
			Quantity: updateProduct.Quantity,
			ID:       updateProduct.Id,
		}

		prep.Repo.Update(Product)

		createResponse(nil, nil, utils.GetFunctionName(prep.UpdateProduct()), c)

	}
}

func (prep ProductsController) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createProduct response.ProductGetResponse

		if err := c.Bind(&createProduct); err != nil {
			c.JSON(http.StatusBadRequest, exceptions.BadRequestException{}.CreateNew(err))
		}

		Product := models.Product{
			Name:      createProduct.Name,
			Quantity:  createProduct.Quantity,
			IsDeleted: false,
		}

		prep.Repo.Create(Product)

		createResponse(nil, nil, utils.GetFunctionName(prep.CreateProduct()), c)
	}
}
