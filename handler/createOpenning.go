package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createOpenningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "createOpportunity",
	})
}
