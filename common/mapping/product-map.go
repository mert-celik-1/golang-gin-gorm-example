package mapping

import (
	"gin-gorm-ex/common/response"
	"gin-gorm-ex/models"
)

func MapProductToProductGetResponse(productModel []models.Product) (productResponse []response.ProductGetResponse) {

	for _, s := range productModel {
		productResponse = append(productResponse, response.ProductGetResponse{
			Name:     s.Name,
			Quantity: s.Quantity,
		})
	}

	return productResponse
}
