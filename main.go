package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()
	RegisterRoutes(router)
	router.Run(":8080")
}
