package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/andresds/go-portunities/handler/models"
	utils "github.com/andresds/go-portunities/handler/utils"
	"github.com/andresds/go-portunities/schemas"
)

func CreateOpenningHandler(ctx *gin.Context) {
	request := models.OpenningRequest{}

	ctx.BindJSON(&request)
	if err := utils.Validate(&request); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// unmarshal request to Openning
	openning := schemas.Openning{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&openning).Error; err != nil {
		logger.Errorf("Error creating openning: %v", err.Error())
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "create_openning", openning)
}
