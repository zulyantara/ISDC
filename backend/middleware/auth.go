package middleware

import (
	"jsdc-api/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthRequired validates JWT token from Authorization header
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Check Bearer token format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Invalid authorization format, use: Bearer <token>",
			})
			c.Abort()
			return
		}

		// Validate token
		claims, err := config.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("user_name", claims.UserName)
		c.Set("user_level", claims.UserLevel)
		c.Set("area_id", claims.AreaID)

		c.Next()
	}
}

// AdminOnly restricts access to admin level users
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		level, exists := c.Get("user_level")
		if !exists || level.(int) > 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Admin access required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// KasirOnly restricts access to kasir level users
func KasirOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		level, exists := c.Get("user_level")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Access denied",
			})
			c.Abort()
			return
		}
		// Level 1=admin, 2=superadmin, 3=kasir
		userLevel := level.(int)
		if userLevel > 3 {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": "Kasir access required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
