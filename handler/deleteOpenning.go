package handler

import (
	"fmt"
	"net/http"

	utils "github.com/andresds/go-portunities/handler/utils"

	"github.com/andresds/go-portunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id, err := utils.GetIdFromParam(ctx)
	if err != nil {
		return
	}

	openning := schemas.Openning{}
	if err := db.First(&openning, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("openning with id: %s not found", id))
		return
	}

	if err := db.Delete(&openning).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opoening with id: %s", id))
		return
	}

	utils.SendSuccess(ctx, "deleting-opning", openning)
}
