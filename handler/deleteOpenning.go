package handler

import (
	"fmt"
	"net/http"

	models "github.com/andresds/go-portunities/handler/models"
	"github.com/andresds/go-portunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		models.SendError(ctx, http.StatusBadRequest, models.ErrParamsIsRequired("id", "queryParameter").Error())
		return
	}

	openning := schemas.Openning{}
	if err := db.First(&openning, id).Error; err != nil {
		models.SendError(ctx, http.StatusNotFound, fmt.Sprintf("openning with id: %s not found", id))
		return
	}

	if err := db.Delete(&openning).Error; err != nil {
		models.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opoening with id: %s", id))
		return
	}

	models.SendSuccess(ctx, "deleting-opning", openning)
}
