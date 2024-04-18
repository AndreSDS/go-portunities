package router

import (
	"github.com/gin-gonic/gin"

	"github.com/andresds/go-portunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openin", handler.ListOpeningsHandler)
		v1.GET("/openin/:id", handler.ShowOpeningHandler)
		v1.POST("/openin", handler.CreateOpenningHandler)
		v1.PUT("/openin/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/openin/:id", handler.DeleteOpeningHandler)
	}
}
