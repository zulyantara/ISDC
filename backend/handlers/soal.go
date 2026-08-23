package handlers

import (
	"isdc-api/models"
	"isdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllSoal returns all exam questions
func GetAllSoal(c *gin.Context) {
	categoryStr := c.Query("category")
	var category int
	if categoryStr != "" {
		category, _ = strconv.Atoi(categoryStr)
	}

	var soalList []models.Soal
	var err error

	if category > 0 {
		soalList, err = models.GetSoalByCategory(category)
	} else {
		soalList, err = models.GetAllSoal()
	}

	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch soal")
		return
	}

	utils.SuccessResponse(c, "Soal retrieved successfully", soalList)
}

// GetSoal returns a single question
func GetSoal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid soal ID")
		return
	}

	soal, err := models.GetSoalByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Soal not found")
		return
	}

	utils.SuccessResponse(c, "Soal retrieved successfully", soal)
}

// CreateSoal creates a new question
func CreateSoal(c *gin.Context) {
	var soal models.Soal
	if err := c.ShouldBindJSON(&soal); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateSoal(&soal); err != nil {
		utils.InternalErrorResponse(c, "Failed to create soal")
		return
	}

	utils.CreatedResponse(c, "Soal created successfully", soal)
}

// UpdateSoal updates a question
func UpdateSoal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid soal ID")
		return
	}

	var soal models.Soal
	if err := c.ShouldBindJSON(&soal); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateSoal(id, &soal); err != nil {
		utils.InternalErrorResponse(c, "Failed to update soal")
		return
	}

	utils.SuccessResponse(c, "Soal updated successfully", soal)
}

// DeleteSoal deletes a question
func DeleteSoal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid soal ID")
		return
	}

	if err := models.DeleteSoal(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete soal")
		return
	}

	utils.SuccessResponse(c, "Soal deleted successfully", nil)
}
