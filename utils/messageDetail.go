package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessMessage creates a standardized success response
func SuccessMessage(c *gin.Context, responseData interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    responseData,
	})
}

// ErrorMessage creates a standardized error response
func ErrorMessage(c *gin.Context, message string, errorDetail string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
		"error":   errorDetail,
	})
}
