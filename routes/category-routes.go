package routes

import (
	"gin-gorm-ex/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(incomingRoutes *gin.Engine, ctrl *controllers.CategoriesController) {

	incomingRoutes.GET("/category", ctrl.GetAllCategories())

}
