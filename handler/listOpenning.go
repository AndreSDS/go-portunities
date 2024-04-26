package handler

import (
	"net/http"

	models "github.com/andresds/go-portunities/handler/models"
	utils "github.com/andresds/go-portunities/handler/utils"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {
	opennings := []models.OpenningResponse{}

	if err := db.Table("opennings").Find(&opennings).Error; err != nil {
		utils.SendError(c, http.StatusInternalServerError, "error fetching opennings")
		return
	}

	utils.SendSuccess(c, "listopennings", opennings)
}
