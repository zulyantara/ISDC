package handlers

import (
	"jsdc-api/models"
	"jsdc-api/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAllDaftar returns all registrations
func GetAllDaftar(c *gin.Context) {
	areaID, _ := c.Get("area_id")
	area, _ := areaID.(int)

	daftar, err := models.GetAllDaftar(area)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch pendaftaran")
		return
	}

	utils.SuccessResponse(c, "Pendaftaran retrieved successfully", daftar)
}

// GetDaftar returns a single registration
func GetDaftar(c *gin.Context) {
	pesertaID := c.Param("id")

	daftar, err := models.GetDaftarByID(pesertaID)
	if err != nil {
		utils.NotFoundResponse(c, "Pendaftaran not found")
		return
	}

	utils.SuccessResponse(c, "Pendaftaran retrieved successfully", daftar)
}

// CreateDaftar creates a new registration with auto-generated ID
func CreateDaftar(c *gin.Context) {
	var d models.Daftar
	if err := c.ShouldBindJSON(&d); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	// Get area info from JWT
	areaID, _ := c.Get("area_id")
	d.AreaID, _ = areaID.(int)

	// Get area kode for generating peserta_id
	area, err := models.GetAreaByID(d.AreaID)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid area")
		return
	}

	// Generate peserta_id: YYYY.MM.AREA_KODE.00001
	now := time.Now()
	yearMonth := fmt.Sprintf("%d.%02d", now.Year(), now.Month())

	// Get last peserta_id for this area and month
	lastID, err := models.GetLastPesertaID(area.AreaKode, yearMonth)
	var newNum int
	if err != nil {
		newNum = 1
	} else {
		// Extract last 5 digits and increment
		var num int
		fmt.Sscanf(lastID[len(lastID)-5:], "%d", &num)
		newNum = num + 1
	}

	d.PesertaID = fmt.Sprintf("%s.%s.%05d", yearMonth, area.AreaKode, newNum)
	d.TglDaftar = now.Format("2006-01-02")
	d.RefID = c.PostForm("ref_id")
	if d.RefID == "" {
		d.RefID = "J-000"
	}

	// Get tarif from kelas
	kelas, err := models.GetKelasByKode(d.KelasID)
	if err == nil {
		d.Harga = kelas.Tarif
		d.Biaya = kelas.Tarif
	}

	// Set user from JWT
	userID, _ := c.Get("user_id")
	d.UserID = userID.(string)

	if err := models.CreateDaftar(&d); err != nil {
		utils.InternalErrorResponse(c, "Failed to create pendaftaran")
		return
	}

	utils.CreatedResponse(c, "Pendaftaran created successfully", d)
}

// UpdateDaftar updates a registration
func UpdateDaftar(c *gin.Context) {
	pesertaID := c.Param("id")

	var d models.Daftar
	if err := c.ShouldBindJSON(&d); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateDaftar(pesertaID, &d); err != nil {
		utils.InternalErrorResponse(c, "Failed to update pendaftaran")
		return
	}

	utils.SuccessResponse(c, "Pendaftaran updated successfully", d)
}

// DeleteDaftar deletes a registration
func DeleteDaftar(c *gin.Context) {
	pesertaID := c.Param("id")

	if err := models.DeleteDaftar(pesertaID); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete pendaftaran")
		return
	}

	utils.SuccessResponse(c, "Pendaftaran deleted successfully", nil)
}

// GetCountDaftarToday returns today's registration count
func GetCountDaftarToday(c *gin.Context) {
	userID, _ := c.Get("user_id")

	count, err := models.GetCountDaftarToday(userID.(string))
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to count pendaftaran")
		return
	}

	utils.SuccessResponse(c, "Count retrieved", gin.H{
		"count": count,
	})
}
