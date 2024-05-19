package feature

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnshah219/go_gin/common/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	feature := router.Group("/feature", middleware.AuthMiddleware)
	feature.GET("/ping", func(ctx *gin.Context) {
		userID, _ := ctx.Get("userID")
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Authenticated pong to userID %s", userID)})
	})
}
