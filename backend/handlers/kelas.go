package handlers

import (
	"jsdc-api/models"
	"jsdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllKelas returns all kelas
func GetAllKelas(c *gin.Context) {
	areaIDStr := c.Query("area_id")
	var areaID int
	if areaIDStr != "" {
		areaID, _ = strconv.Atoi(areaIDStr)
	}

	var kelasList []models.Kelas
	var err error

	if areaID > 0 {
		kelasList, err = models.GetKelasByArea(areaID)
	} else {
		kelasList, err = models.GetAllKelas()
	}

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch kelas")
		return
	}

	utils.SuccessResponse(c, "Kelas retrieved successfully", kelasList)
}

// GetKelas returns a single kelas
func GetKelas(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid kelas ID")
		return
	}

	kelas, err := models.GetKelasByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Kelas not found")
		return
	}

	utils.SuccessResponse(c, "Kelas retrieved successfully", kelas)
}

// CreateKelas creates a new kelas
func CreateKelas(c *gin.Context) {
	var kelas models.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateKelas(&kelas); err != nil {
		utils.InternalErrorResponse(c, "Failed to create kelas")
		return
	}

	utils.CreatedResponse(c, "Kelas created successfully", kelas)
}

// UpdateKelas updates a kelas
func UpdateKelas(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid kelas ID")
		return
	}

	var kelas models.Kelas
	if err := c.ShouldBindJSON(&kelas); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateKelas(id, &kelas); err != nil {
		utils.InternalErrorResponse(c, "Failed to update kelas")
		return
	}

	utils.SuccessResponse(c, "Kelas updated successfully", kelas)
}

// DeleteKelas deletes a kelas
func DeleteKelas(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid kelas ID")
		return
	}

	if err := models.DeleteKelas(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete kelas")
		return
	}

	utils.SuccessResponse(c, "Kelas deleted successfully", nil)
}
