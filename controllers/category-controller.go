package controllers

import (
	"gin-gorm-ex/common/response"
	"gin-gorm-ex/repository/categories"
	"gin-gorm-ex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	Repo categories.CategoryRepositoryInterface
}

func NewCategoryControllers(cat categories.CategoryRepositoryInterface) *CategoriesController {
	return &CategoriesController{Repo: cat}
}

func (prep CategoriesController) GetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {

		result, err := prep.Repo.GetAll()

		response := response.ProductsResponseFormat{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		createResponse(response, err, utils.GetFunctionName(prep.GetAllCategories()), c)

	}
}
