package main

import (
	"gin-gorm-ex/configs"
	"gin-gorm-ex/controllers"
	"gin-gorm-ex/middlewares"
	"gin-gorm-ex/repository/categories"
	"gin-gorm-ex/repository/products"
	"gin-gorm-ex/repository/users"
	"gin-gorm-ex/routes"
	"gin-gorm-ex/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDb(config)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	productRepo := products.NewProductsRepo(db)
	productCtrl := controllers.NewProductControllers(productRepo)

	categoryRepo := categories.NewCategoriesRepo(db)
	categoryCtrl := controllers.NewCategoryControllers(categoryRepo)

	userRepo := users.NewUsersRepo(db)
	userCtrl := controllers.NewUserControllers(userRepo)

	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router, userCtrl)

	router.Use(middlewares.Authentication())
	routes.ProductRoutes(router, productCtrl)
	routes.CategoryRoutes(router, categoryCtrl)

	router.Run(":" + port)

}

// logging
// validation
// todoları hallet
