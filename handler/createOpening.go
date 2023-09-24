package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrluizandre/gopportunities/schemas"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Location,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creaging opening: %v", err)
		sendError(context, http.StatusInternalServerError, "error creaging opening on database")
		return
	}

	sendSuccess(context, "create-opening", opening)
}
