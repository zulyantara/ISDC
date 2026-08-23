package models

import (
	"isdc-api/config"
	"isdc-api/utils"
	"fmt"
)

type User struct {
	IDUser    int    `json:"id_user"`
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLevel int    `json:"user_level"`
	AppID     int    `json:"app_id"`
	Flag      int    `json:"flag"`
	Aktif     int    `json:"aktif"`
	AreaID    int    `json:"area_id"`
}

type LoginRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	UserPwd string `json:"user_pwd" binding:"required"`
}

type LoginResponse struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLevel int    `json:"user_level"`
	AreaID    int    `json:"area_id"`
	Token     string `json:"token"`
}

// GetAllUsers returns all users
func GetAllUsers() ([]User, error) {
	query := "SELECT id_user, user_id, user_name, user_level, app_id, flag, aktif, area_id FROM mt_user ORDER BY user_id"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.IDUser, &u.UserID, &u.UserName, &u.UserLevel, &u.AppID, &u.Flag, &u.Aktif, &u.AreaID)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// GetUserByID returns a single user by user_id
func GetUserByID(userID string) (*User, error) {
	query := "SELECT id_user, user_id, user_name, user_level, app_id, flag, aktif, area_id FROM mt_user WHERE user_id = ?"
	var u User
	err := config.DB.QueryRow(query, userID).Scan(&u.IDUser, &u.UserID, &u.UserName, &u.UserLevel, &u.AppID, &u.Flag, &u.Aktif, &u.AreaID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// AuthenticateUser validates user credentials using bcrypt with MD5 migration support.
// If the stored hash is MD5 and the password matches, it automatically upgrades to bcrypt.
func AuthenticateUser(userID, password string) (*User, error) {
	// 1. Fetch user + password hash from DB
	var u User
	var storedHash string
	query := "SELECT id_user, user_id, user_name, user_level, app_id, flag, aktif, area_id, user_pwd FROM mt_user WHERE user_id = ? AND aktif = -1"
	err := config.DB.QueryRow(query, userID).Scan(
		&u.IDUser, &u.UserID, &u.UserName, &u.UserLevel, &u.AppID, &u.Flag, &u.Aktif, &u.AreaID, &storedHash,
	)
	if err != nil {
		return nil, fmt.Errorf("user not found or inactive")
	}

	// 2. Check if hash is bcrypt
	if utils.IsBcryptHash(storedHash) {
		// Verify with bcrypt
		if !utils.CheckPassword(password, storedHash) {
			return nil, fmt.Errorf("invalid password")
		}
		return &u, nil
	}

	// 3. Legacy MD5 hash — try verify and auto-migrate to bcrypt
	if _, ok := utils.MigratePassword(config.DB, userID, password, storedHash); ok {
		return &u, nil
	}

	return nil, fmt.Errorf("invalid password")
}

// CreateUser creates a new user with bcrypt hashed password
func CreateUser(u *User, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := "INSERT INTO mt_user (user_id, user_name, user_pwd, user_level, app_id, flag, aktif, area_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = config.DB.Exec(query, u.UserID, u.UserName, hashedPassword, u.UserLevel, u.AppID, u.Flag, u.Aktif, u.AreaID)
	return err
}

// UpdateUser updates user data
func UpdateUser(userID string, u *User) error {
	query := "UPDATE mt_user SET user_name=?, user_level=?, app_id=?, flag=?, aktif=?, area_id=? WHERE user_id=?"
	_, err := config.DB.Exec(query, u.UserName, u.UserLevel, u.AppID, u.Flag, u.Aktif, u.AreaID, userID)
	return err
}

// UpdatePassword updates user password with bcrypt hash
func UpdatePassword(userID, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := "UPDATE mt_user SET user_pwd=? WHERE user_id=?"
	_, err = config.DB.Exec(query, hashedPassword, userID)
	return err
}

// DeleteUser deletes a user
func DeleteUser(userID string) error {
	query := "DELETE FROM mt_user WHERE user_id=?"
	_, err := config.DB.Exec(query, userID)
	return err
}

// GetUsersByArea returns users filtered by area
func GetUsersByArea(areaID int) ([]User, error) {
	query := "SELECT id_user, user_id, user_name, user_level, app_id, flag, aktif, area_id FROM mt_user WHERE area_id=? ORDER BY user_id"
	rows, err := config.DB.Query(query, areaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.IDUser, &u.UserID, &u.UserName, &u.UserLevel, &u.AppID, &u.Flag, &u.Aktif, &u.AreaID)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// UpdateFlag updates user login flag
func UpdateFlag(userID string, flag int) error {
	query := "UPDATE mt_user SET flag=? WHERE user_id=?"
	_, err := config.DB.Exec(query, flag, userID)
	return err
}

// CountUsers returns total user count
func CountUsers() (int64, error) {
	var count int64
	err := config.DB.QueryRow("SELECT COUNT(*) FROM mt_user").Scan(&count)
	return count, err
}

// UserExists checks if user_id already exists
func UserExists(userID string) (bool, error) {
	var count int64
	err := config.DB.QueryRow("SELECT COUNT(*) FROM mt_user WHERE user_id=?", userID).Scan(&count)
	return count > 0, err
}
