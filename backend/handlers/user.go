package handlers

import (
	"isdc-api/models"
	"isdc-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllUsers returns all users
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch users")
		return
	}

	utils.SuccessResponse(c, "Users retrieved successfully", users)
}

// GetUser returns a single user
func GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := models.GetUserByID(userID)
	if err != nil {
		utils.NotFoundResponse(c, "User not found")
		return
	}

	utils.SuccessResponse(c, "User retrieved successfully", user)
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	// Check if user_id already exists
	exists, _ := models.UserExists(user.UserID)
	if exists {
		utils.BadRequestResponse(c, "User ID already exists")
		return
	}

	// Password from request body
	password := c.PostForm("user_pwd")
	if password == "" {
		password = "password123" // default password
	}

	if err := models.CreateUser(&user, password); err != nil {
		utils.InternalErrorResponse(c, "Failed to create user")
		return
	}

	utils.CreatedResponse(c, "User created successfully", user)
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c, "Invalid request body")
		return
	}

	if err := models.UpdateUser(userID, &user); err != nil {
		utils.InternalErrorResponse(c, "Failed to update user")
		return
	}

	// Update password if provided
	newPassword := c.PostForm("user_pwd")
	if newPassword != "" {
		_ = models.UpdatePassword(userID, newPassword)
	}

	utils.SuccessResponse(c, "User updated successfully", user)
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	if err := models.DeleteUser(userID); err != nil {
		utils.InternalErrorResponse(c, "Failed to delete user")
		return
	}

	utils.SuccessResponse(c, "User deleted successfully", nil)
}

// GetUsersByArea returns users filtered by area
func GetUsersByArea(c *gin.Context) {
	areaIDStr := c.Param("area_id")
	areaID, err := strconv.Atoi(areaIDStr)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid area_id")
		return
	}

	users, err := models.GetUsersByArea(areaID)
	if err != nil {
		utils.InternalErrorResponse(c, "Failed to fetch users")
		return
	}

	utils.SuccessResponse(c, "Users retrieved successfully", users)
}
