package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrluizandre/gopportunities/schemas"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-openings", openings)
}
