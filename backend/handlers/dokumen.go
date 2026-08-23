package handlers

import (
	"jsdc-api/models"
	"jsdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllJenisDokumen returns all document types
func GetAllJenisDokumen(c *gin.Context) {
	list, err := models.GetAllJenisDokumen()
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch jenis dokumen")
		return
	}

	utils.SuccessResponse(c, "Jenis dokumen retrieved successfully", list)
}

// GetJenisDokumen returns a single document type
func GetJenisDokumen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	jd, err := models.GetJenisDokumenByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Jenis dokumen not found")
		return
	}

	utils.SuccessResponse(c, "Jenis dokumen retrieved successfully", jd)
}

// CreateJenisDokumen creates a new document type
func CreateJenisDokumen(c *gin.Context) {
	var jd models.JenisDokumen
	if err := c.ShouldBindJSON(&jd); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateJenisDokumen(&jd); err != nil {
		utils.InternalErrorResponse(c, "Failed to create jenis dokumen")
		return
	}

	utils.CreatedResponse(c, "Jenis dokumen created successfully", jd)
}

// UpdateJenisDokumen updates a document type
func UpdateJenisDokumen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	var jd models.JenisDokumen
	if err := c.ShouldBindJSON(&jd); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateJenisDokumen(id, &jd); err != nil {
		utils.InternalErrorResponse(c, "Failed to update jenis dokumen")
		return
	}

	utils.SuccessResponse(c, "Jenis dokumen updated successfully", jd)
}

// DeleteJenisDokumen deletes a document type
func DeleteJenisDokumen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	if err := models.DeleteJenisDokumen(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete jenis dokumen")
		return
	}

	utils.SuccessResponse(c, "Jenis dokumen deleted successfully", nil)
}

// GetAllDaftarDokumen returns all document entries
func GetAllDaftarDokumen(c *gin.Context) {
	list, err := models.GetAllDaftarDokumen()
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch daftar dokumen")
		return
	}

	utils.SuccessResponse(c, "Daftar dokumen retrieved successfully", list)
}

// CreateDaftarDokumen creates a new document entry
func CreateDaftarDokumen(c *gin.Context) {
	var d models.DaftarDokumen
	if err := c.ShouldBindJSON(&d); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.CreateDaftarDokumen(&d); err != nil {
		utils.InternalErrorResponse(c, "Failed to create daftar dokumen")
		return
	}

	utils.CreatedResponse(c, "Daftar dokumen created successfully", d)
}

// UpdateDaftarDokumen updates a document entry
func UpdateDaftarDokumen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	var d models.DaftarDokumen
	if err := c.ShouldBindJSON(&d); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateDaftarDokumen(id, &d); err != nil {
		utils.InternalErrorResponse(c, "Failed to update daftar dokumen")
		return
	}

	utils.SuccessResponse(c, "Daftar dokumen updated successfully", d)
}

// DeleteDaftarDokumen deletes a document entry
func DeleteDaftarDokumen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid ID")
		return
	}

	if err := models.DeleteDaftarDokumen(id); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete daftar dokumen")
		return
	}

	utils.SuccessResponse(c, "Daftar dokumen deleted successfully", nil)
}
