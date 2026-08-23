package handlers

import (
	"fmt"
	"isdc-api/models"
	"isdc-api/utils"

	"github.com/gin-gonic/gin"
)

// GetAllPeserta returns all participants
func GetAllPeserta(c *gin.Context) {
	areaID, _ := c.Get("area_id")
	area, _ := areaID.(int)

	peserta, err := models.GetAllPeserta(area)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch peserta")
		return
	}

	utils.SuccessResponse(c, "Peserta retrieved successfully", peserta)
}

// GetPeserta returns a single participant
func GetPeserta(c *gin.Context) {
	pesertaID := c.Param("id")

	peserta, err := models.GetPesertaByID(pesertaID)
	if err != nil {
		utils.NotFoundResponse(c, "Peserta not found")
		return
	}

	utils.SuccessResponse(c, "Peserta retrieved successfully", peserta)
}

// GetPesertaSoal returns exam questions for a participant
func GetPesertaSoal(c *gin.Context) {
	pesertaID := c.Param("id")

	// Check if participant exists
	exists, err := models.CheckPesertaExists(pesertaID)
	if err != nil || !exists {
		utils.NotFoundResponse(c, "Peserta not found")
		return
	}

	// Get participant data
	peserta, err := models.GetPesertaByID(pesertaID)
	if err != nil {
		utils.NotFoundResponse(c, "Peserta not found")
		return
	}

	// Get questions based on participant's class category (teori_id)
	soalList, err := models.GetSoalByCategory(peserta.TeoriID)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch soal")
		return
	}

	utils.SuccessResponse(c, "Soal retrieved successfully", gin.H{
		"peserta": peserta,
		"soal":    soalList,
	})
}

// SearchPeserta searches participants
func SearchPeserta(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		utils.BadRequestResponse(c, "Search keyword (q) is required")
		return
	}

	areaID, _ := c.Get("area_id")
	area, _ := areaID.(int)

	peserta, err := models.SearchPeserta(keyword, area)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to search peserta")
		return
	}

	utils.SuccessResponse(c, "Peserta search results", peserta)
}

// SubmitPraktek submits practical exam results
func SubmitPraktek(c *gin.Context) {
	pesertaID := c.Param("id")

	var input struct {
		Results []struct {
			SoalID int `json:"soal_id"`
			Hasil  int `json:"hasil"`
		} `json:"results"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	// Convert to UjiPraktek slice
	var praktekList []models.UjiPraktek
	for _, r := range input.Results {
		praktekList = append(praktekList, models.UjiPraktek{
			SoalID:    r.SoalID,
			PesertaID: pesertaID,
			Hasil:     r.Hasil,
		})
	}

	// Upsert all results
	if err := models.UpsertPraktek(praktekList, "web"); err != nil {
		utils.InternalErrorResponse(c, "Failed to submit praktek")
		return
	}

	// Calculate average and update peserta
	avg, err := models.GetRataRataPraktek(pesertaID)
	if err == nil {
		_ = models.UpdatePraktekHasil(pesertaID, fmt.Sprintf("%.0f", avg))
	}

	utils.SuccessResponse(c, "Praktek submitted successfully", nil)
}

// GetPraktekByPeserta returns practical exam results for a participant
func GetPraktekByPeserta(c *gin.Context) {
	pesertaID := c.Param("id")

	results, err := models.GetPraktekByPeserta(pesertaID)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch praktek results")
		return
	}

	utils.SuccessResponse(c, "Praktek results retrieved successfully", results)
}

// SubmitComment submits comments for a participant
func SubmitComment(c *gin.Context) {
	pesertaID := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	comment.PesertaID = pesertaID

	if err := models.UpsertComment(&comment); err != nil {
		utils.InternalErrorResponse(c, "Failed to submit comment")
		return
	}

	utils.SuccessResponse(c, "Comment submitted successfully", comment)
}

// GetCommentByPeserta returns comments for a participant
func GetCommentByPeserta(c *gin.Context) {
	pesertaID := c.Param("id")

	comment, err := models.GetCommentByPeserta(pesertaID)
	if err != nil {
		utils.NotFoundResponse(c, "Comment not found")
		return
	}

	utils.SuccessResponse(c, "Comment retrieved successfully", comment)
}
