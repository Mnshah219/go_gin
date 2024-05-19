package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mnshah219/go_gin/auth"
	"github.com/mnshah219/go_gin/feature"
)

func RegisterRoutes(router *gin.Engine) {
	auth.RegisterRoutes(router)
	feature.RegisterRoutes(router)
}
