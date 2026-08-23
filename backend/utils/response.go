package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse sends a success JSON response
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  false,
		"message": message,
	})
}

// PaginatedResponse sends a paginated JSON response
func PaginatedResponse(c *gin.Context, message string, data interface{}, total int64, page, perPage int) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
		"pagination": gin.H{
			"total":    total,
			"page":     page,
			"per_page": perPage,
		},
	})
}

// CreatedResponse sends a 201 Created JSON response
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

// BadRequestResponse sends a 400 Bad Request JSON response
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

// NotFoundResponse sends a 404 Not Found JSON response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message)
}

// InternalErrorResponse sends a 500 Internal Server Error JSON response
func InternalErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}
