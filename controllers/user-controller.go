package controllers

import (
	"gin-gorm-ex/models"
	"gin-gorm-ex/repository/users"
	"gin-gorm-ex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	Repo users.UserRepositoryInterface
}

func NewUserControllers(u users.UserRepositoryInterface) *UsersController {
	return &UsersController{Repo: u}
}

func (u UsersController) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		password := utils.HashPassword(user.Password)
		user.Password = password

		token, refreshToken, _ := utils.GenerateAllTokens(user.Email)
		user.Token = token
		user.Refresh_Token = refreshToken

		result, err := u.Repo.CreateUser(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
		return

	}

}

func (u UsersController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var ctx, _ = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		foundUser, err := u.Repo.GetUserByEmail(user.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems to be incorrect"})
			return
		}

		passwordIsValid, msg := utils.VerifyPassword(user.Password, foundUser.Password)

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := utils.GenerateAllTokens(foundUser.Email)

		//update tokens - token and refersh token
		utils.UpdateAllTokens(token, refreshToken, foundUser.ID)

		//return statusOK
		c.JSON(http.StatusOK, foundUser)
	}

}
