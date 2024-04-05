package routes

import (
	"gin-gorm-ex/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(incomingRoutes *gin.Engine, ctrl *controllers.ProductsController) {

	incomingRoutes.GET("/product", ctrl.GetAllProducts())
	incomingRoutes.GET("/productwithc", ctrl.GetAllProductsWithCategories())
	incomingRoutes.GET("/productbycat/:catid", ctrl.GetByCategoryId())
	incomingRoutes.GET("/product/:id", ctrl.GetProductById())
	incomingRoutes.DELETE("/product/:id", ctrl.DeleteProduct())
	incomingRoutes.POST("/product", ctrl.CreateProduct())
	incomingRoutes.PUT("/product", ctrl.UpdateProduct())

}
