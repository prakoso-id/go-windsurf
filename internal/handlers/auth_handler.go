package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/interfaces/http/response"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthHandler(authService services.AuthService, userService services.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	type loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Login failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization required", "authorization header required")
			c.Abort()
			return
		}

		userID, err := authService.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "Invalid token", err.Error())
			c.Abort()
			return
		}

		if userID == "" {
			response.Error(c, http.StatusUnauthorized, "Invalid token", "token is not valid")
			c.Abort()
			return
		}

		// Store the validated userID in the context for later use
		c.Set("userID", userID)

		c.Next()
	}
}
