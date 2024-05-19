package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/signup", signup)
	auth.POST("/login", login)
}
