package handlers

import (
	"jsdc-api/models"
	"jsdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllNilaiLulus returns all passing scores
func GetAllNilaiLulus(c *gin.Context) {
	list, err := models.GetAllNilaiLulus()
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch nilai lulus")
		return
	}

	utils.SuccessResponse(c, "Nilai lulus retrieved successfully", list)
}

// GetNilaiLulus returns a single passing score
func GetNilaiLulus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	nl, err := models.GetNilaiLulusByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Nilai lulus not found")
		return
	}

	utils.SuccessResponse(c, "Nilai lulus retrieved successfully", nl)
}

// CreateNilaiLulus creates a new passing score
func CreateNilaiLulus(c *gin.Context) {
	var nl models.NilaiLulus
	if err := c.ShouldBindJSON(&nl); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateNilaiLulus(&nl); err != nil {
		utils.InternalErrorResponse(c, "Failed to create nilai lulus")
		return
	}

	utils.CreatedResponse(c, "Nilai lulus created successfully", nl)
}

// UpdateNilaiLulus updates a passing score
func UpdateNilaiLulus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	var nl models.NilaiLulus
	if err := c.ShouldBindJSON(&nl); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateNilaiLulus(id, &nl); err != nil {
		utils.InternalErrorResponse(c, "Failed to update nilai lulus")
		return
	}

	utils.SuccessResponse(c, "Nilai lulus updated successfully", nl)
}

// DeleteNilaiLulus deletes a passing score
func DeleteNilaiLulus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	if err := models.DeleteNilaiLulus(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete nilai lulus")
		return
	}

	utils.SuccessResponse(c, "Nilai lulus deleted successfully", nil)
}
