package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(ctx *gin.Context) (string, error) {
	id := ctx.Param("id")
	if id == "" {
		SendError(ctx, http.StatusBadRequest, ErrParamsIsRequired("id", "queryParameter").Error())
		return "", ErrParamsIsRequired("id", "queryParameter")
	}
	return id, nil
}
