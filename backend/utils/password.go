package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a bcrypt hashed password with a plain text password
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckMD5Password checks if password matches MD5 hash (for legacy migration)
func CheckMD5Password(password, md5Hash string) bool {
	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	return hash == md5Hash
}

// IsBcryptHash checks if a hash string is a bcrypt hash
func IsBcryptHash(hash string) bool {
	return len(hash) == 60 && hash[0] == '$'
}

// MigratePassword upgrades MD5 hash to bcrypt and updates the database
// Returns (bcryptHash, true) if migration was successful
func MigratePassword(db *sql.DB, userID, password, md5Hash string) (string, bool) {
	// Verify MD5 first
	if !CheckMD5Password(password, md5Hash) {
		return "", false
	}

	// Hash with bcrypt
	bcryptHash, err := HashPassword(password)
	if err != nil {
		return "", false
	}

	// Update database
	_, err = db.Exec("UPDATE mt_user SET user_pwd=? WHERE user_id=?", bcryptHash, userID)
	if err != nil {
		return "", false
	}

	return bcryptHash, true
}
