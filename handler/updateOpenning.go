package handler

import (
	"fmt"
	"net/http"

	models "github.com/andresds/go-portunities/handler/models"
	utils "github.com/andresds/go-portunities/handler/utils"
	"github.com/andresds/go-portunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id, err := utils.GetIdFromParam(ctx)
	if err != nil {
		return
	}

	openning := schemas.Openning{}
	if err := db.First(&openning, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("openning with id: %s not found", id))
		return
	}

	request := models.OpenningRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, fmt.Sprintf("malformad request body: %s", err.Error()))
		return
	}

	if err := utils.Validate(&request); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// cast the request to a iterable value
	// to update the openning
	if err := db.Model(&openning).Updates(request).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opoening with id: %s", id))
		return
	}

	utils.SendSuccess(ctx, "updating-opning", openning)
}
