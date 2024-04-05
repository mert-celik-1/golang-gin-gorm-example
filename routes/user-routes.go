package routes

import (
	"gin-gorm-ex/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine, ctrl *controllers.UsersController) {

	incomingRoutes.POST("/users/signup", ctrl.SignUp())
	incomingRoutes.POST("/users/login", ctrl.Login())

}
