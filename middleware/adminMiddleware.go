package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Warn("Missing Authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) != 2 || splitToken[0] != "Bearer" {
			log.Warn("Invalid Authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})
			return
		}

		tokenStr := splitToken[1]
		token, err := VerifyToken(tokenStr)
		if err != nil || !token.Valid {
			log.Warnf("Invalid or expired token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Warn("Invalid token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		// Ambil ID dan role
		role, okRole := claims["role"].(string)
		idFloat, okID := claims["user_id"].(float64)
		userID := "unknown"
		if okID {
			userID = fmt.Sprintf("%d", int(idFloat))
			c.Set("userID", int(idFloat))
		}

		if !okRole || role != "admin" {
			log.Warnf("Unauthorized access attempt by user ID %v with role: %v", userID, role)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Access denied, admin only",
			})
			return
		}

		c.Next()
	}
}
