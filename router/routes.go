package router

import (
	"github.com/gin-gonic/gin"

	handler "github.com/andresds/go-portunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	// initialize the handlers
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpenningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
	}
}
