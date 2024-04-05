package middlewares

import (
	"gin-gorm-ex/common/exceptions"
	"gin-gorm-ex/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		clientToken := ctx.Request.Header.Get("token")

		if clientToken == "" {
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			ctx.JSON(http.StatusInternalServerError, exceptions.AuthException{}.CreateNew(nil))
			ctx.Abort()
			return

		}

		claims, err := utils.ValidateToken(clientToken)

		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)

	}
}
