package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrluizandre/gopportunities/schemas"
)

// @Basepath /api/v1

// @Summary Delete Opening
// @Description Delete job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening Identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(context, "delete-openging", opening)
}
