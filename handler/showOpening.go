package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrluizandre/gopportunities/schemas"
)

func ShowOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(context, "show-opening", opening)
}
