package handlers

import (
	"jsdc-api/config"
	"jsdc-api/models"
	"jsdc-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles user authentication
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "user_id and user_pwd are required")
		return
	}

	// Check if password is still the default password
	isDefaultPassword := (req.UserPwd == config.AppConfig.JWT.DefaultPassword)

	// Authenticate user
	user, err := models.AuthenticateUser(req.UserID, req.UserPwd)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials or user inactive")
		return
	}

	// Generate JWT token
	token, err := config.GenerateToken(user.UserID, user.UserName, user.UserLevel, user.AreaID)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, "Login successful", gin.H{
		"user_id":             user.UserID,
		"user_name":           user.UserName,
		"user_level":          user.UserLevel,
		"area_id":             user.AreaID,
		"token":               token,
		"must_change_password": isDefaultPassword,
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	userID := c.GetString("user_id")

	// Reset login flag
	_ = models.UpdateFlag(userID, 0)

	utils.SuccessResponse(c, "Logout successful", nil)
}

// GetCurrentUser returns the authenticated user's info
func GetCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")

	user, err := models.GetUserByID(userID)
	if err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, "User data retrieved", user)
}

// ChangePassword forces user to change from default password
func ChangePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "old_password and new_password are required")
		return
	}

	// Validate new password strength
	if len(req.NewPassword) < 8 {
		utils.BadRequestResponse(c, "Password baru minimal 8 karakter")
		return
	}

	if req.NewPassword == config.AppConfig.JWT.DefaultPassword {
		utils.BadRequestResponse(c, "Password baru tidak boleh sama dengan password default")
		return
	}

	if req.NewPassword == req.OldPassword {
		utils.BadRequestResponse(c, "Password baru harus berbeda dari password lama")
		return
	}

	userID := c.GetString("user_id")

	// Verify old password
	_, err := models.AuthenticateUser(userID, req.OldPassword)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Password lama salah")
		return
	}

	// Update to new password (bcrypt)
	if err := models.UpdatePassword(userID, req.NewPassword); err != nil {
		utils.InternalErrorResponse(c, "Gagal mengubah password")
		return
	}

	utils.SuccessResponse(c, "Password berhasil diubah", nil)
}
