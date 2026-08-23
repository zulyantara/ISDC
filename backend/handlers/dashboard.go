package handlers

import (
	"isdc-api/utils"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns server status
func HealthCheck(c *gin.Context) {
	utils.SuccessResponse(c, "ISDC API is running", gin.H{
		"version": "2.0.0",
		"status":  "healthy",
	})
}

// GetDashboard returns dashboard statistics
func GetDashboard(c *gin.Context) {
	utils.SuccessResponse(c, "Dashboard data", gin.H{
		"message": "Dashboard endpoint - implement stats queries here",
	})
}
