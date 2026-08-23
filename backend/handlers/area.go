package handlers

import (
	"isdc-api/models"
	"isdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllArea returns all areas
func GetAllArea(c *gin.Context) {
	areas, err := models.GetAllArea()
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch areas")
		return
	}

	utils.SuccessResponse(c, "Areas retrieved successfully", areas)
}

// GetArea returns a single area
func GetArea(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid area ID")
		return
	}

	area, err := models.GetAreaByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Area not found")
		return
	}

	utils.SuccessResponse(c, "Area retrieved successfully", area)
}

// CreateArea creates a new area
func CreateArea(c *gin.Context) {
	var area models.Area
	if err := c.ShouldBindJSON(&area); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateArea(&area); err != nil {
		utils.InternalErrorResponse(c, "Failed to create area")
		return
	}

	utils.CreatedResponse(c, "Area created successfully", area)
}

// UpdateArea updates an area
func UpdateArea(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid area ID")
		return
	}

	var area models.Area
	if err := c.ShouldBindJSON(&area); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateArea(id, &area); err != nil {
		utils.InternalErrorResponse(c, "Failed to update area")
		return
	}

	utils.SuccessResponse(c, "Area updated successfully", area)
}

// DeleteArea deletes an area
func DeleteArea(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid area ID")
		return
	}

	if err := models.DeleteArea(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete area")
		return
	}

	utils.SuccessResponse(c, "Area deleted successfully", nil)
}
