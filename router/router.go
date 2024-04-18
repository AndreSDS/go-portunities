package router

import (
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":8080")
}
