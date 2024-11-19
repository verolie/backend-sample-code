package utils

import (
	"net/http"

	"github.com/code-sample/model/modelResponse"
	"github.com/gin-gonic/gin"
)

func SuccessMessage(c *gin.Context, responseData interface{}, message string) {
	response := modelResponse.SuccessResponse{
		Success: true,
		Message: message,
		Data:    responseData,
	}
	c.JSON(http.StatusOK, response)
}

func ErrorMessage(c *gin.Context, message string, errorDetail string, statusCode int) {
	response := modelResponse.ErrorResponse{
		Success: false,
		Message: message,
		Error:   errorDetail,
	}
	c.JSON(statusCode, response)
}
