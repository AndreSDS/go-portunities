package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "github.com/andresds/go-portunities/handler/models"
	"github.com/andresds/go-portunities/schemas"
)

func CreateOpenningHandler(ctx *gin.Context) {
	request := models.OpenningRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		models.SendError(ctx, http.StatusBadRequest, err.Error())
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
		models.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	models.SendSuccess(ctx, "create_openning", openning)
}

// https://www.youtube.com/watch?v=L6gk7FHBNkM 3:20
