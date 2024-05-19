package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mnshah219/go_gin/auth/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authorization")
	if !strings.HasPrefix(authToken, "Bearer ") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized! Bearer token not found"})

	}
	jwt := strings.TrimPrefix(authToken, "Bearer ")
	userID, err := utils.VerifyJWT(jwt)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	ctx.Set("userID", userID)
	ctx.Next()
}
