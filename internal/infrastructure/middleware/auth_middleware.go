package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/interfaces/http/response"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization required", "missing authorization header")
			c.Abort()
			return
		}

		// Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization format", "invalid authorization header format")
			c.Abort()
			return
		}

		userID, err := authService.ValidateToken(parts[1])
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token", err.Error())
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("userID", userID)
		c.Next()
	}
}
