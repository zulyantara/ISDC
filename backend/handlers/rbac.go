package handlers

import (
	"isdc-api/models"
	"isdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPermissions returns menu permissions for the logged-in user's role
func GetPermissions(c *gin.Context) {
	levelID, _ := c.Get("user_level")
	level, _ := levelID.(int)

	permissions, err := models.GetPermissionsByRole(level)
	if err != nil {
		utils.InternalErrorResponse(c, "Gagal mengambil data permissions")
		return
	}

	utils.SuccessResponse(c, "Permissions retrieved", permissions)
}

// GetAllMenus returns all menus (for admin management)
func GetAllMenus(c *gin.Context) {
	menus, err := models.GetAllMenus()
	if err != nil {
		utils.InternalErrorResponse(c, "Gagal mengambil data menu")
		return
	}
	utils.SuccessResponse(c, "Menus retrieved", menus)
}

// GetAllLevels returns all user levels
func GetAllLevels(c *gin.Context) {
	levels, err := models.GetAllLevels()
	if err != nil {
		utils.InternalErrorResponse(c, "Gagal mengambil data level")
		return
	}
	utils.SuccessResponse(c, "Levels retrieved", levels)
}

// CreateLevel creates a new level/role
func CreateLevel(c *gin.Context) {
	var l models.Level
	if err := c.ShouldBindJSON(&l); err != nil {
		utils.BadRequestResponse(c, "level_desc is required")
		return
	}
	if err := models.CreateLevel(&l); err != nil {
		utils.InternalErrorResponse(c, "Gagal membuat level")
		return
	}
	utils.SuccessResponse(c, "Level berhasil dibuat", l)
}

// UpdateLevel updates a level/role
func UpdateLevel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var l models.Level
	if err := c.ShouldBindJSON(&l); err != nil {
		utils.BadRequestResponse(c, "level_desc is required")
		return
	}
	if err := models.UpdateLevel(id, &l); err != nil {
		utils.InternalErrorResponse(c, "Gagal update level")
		return
	}
	utils.SuccessResponse(c, "Level berhasil diupdate", l)
}

// DeleteLevel deletes a level/role
func DeleteLevel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteLevel(id); err != nil {
		utils.InternalErrorResponse(c, "Gagal hapus level")
		return
	}
	utils.SuccessResponse(c, "Level berhasil dihapus", nil)
}

// CreateMenu creates a new menu
func CreateMenu(c *gin.Context) {
	var m models.Menu
	if err := c.ShouldBindJSON(&m); err != nil {
		utils.BadRequestResponse(c, "menu data is required")
		return
	}
	if err := models.CreateMenu(&m); err != nil {
		utils.InternalErrorResponse(c, "Gagal membuat menu")
		return
	}
	utils.SuccessResponse(c, "Menu berhasil dibuat", m)
}

// UpdateMenu updates a menu
func UpdateMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var m models.Menu
	if err := c.ShouldBindJSON(&m); err != nil {
		utils.BadRequestResponse(c, "menu data is required")
		return
	}
	if err := models.UpdateMenu(id, &m); err != nil {
		utils.InternalErrorResponse(c, "Gagal update menu")
		return
	}
	utils.SuccessResponse(c, "Menu berhasil diupdate", m)
}

// DeleteMenu deletes a menu
func DeleteMenu(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteMenu(id); err != nil {
		utils.InternalErrorResponse(c, "Gagal hapus menu")
		return
	}
	utils.SuccessResponse(c, "Menu berhasil dihapus", nil)
}

// GetAllHakAkses returns all permission entries
func GetAllHakAkses(c *gin.Context) {
	haList, err := models.GetAllHakAkses()
	if err != nil {
		utils.InternalErrorResponse(c, "Gagal mengambil data hak akses")
		return
	}
	utils.SuccessResponse(c, "Hak akses retrieved", haList)
}

// GetHakAksesByRole returns permissions for a specific role
func GetHakAksesByRole(c *gin.Context) {
	levelID, _ := strconv.Atoi(c.Param("level_id"))
	haList, err := models.GetHakAksesByRole(levelID)
	if err != nil {
		utils.InternalErrorResponse(c, "Gagal mengambil data hak akses")
		return
	}
	utils.SuccessResponse(c, "Hak akses retrieved", haList)
}

// SaveHakAksesForRole bulk-replaces all permissions for a role
func SaveHakAksesForRole(c *gin.Context) {
	levelID, _ := strconv.Atoi(c.Param("level_id"))
	var permissions []models.HakAkses
	if err := c.ShouldBindJSON(&permissions); err != nil {
		utils.BadRequestResponse(c, "permissions array is required")
		return
	}
	if err := models.UpsertHakAksesForRole(levelID, permissions); err != nil {
		utils.InternalErrorResponse(c, "Gagal menyimpan hak akses")
		return
	}
	utils.SuccessResponse(c, "Hak akses berhasil disimpan", nil)
}
