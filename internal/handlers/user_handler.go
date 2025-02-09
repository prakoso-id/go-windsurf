package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/interfaces/http/response"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	type registerRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Name     string `json:"name" binding:"required"`
	}

	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	user, err := h.userService.Register(req.Email, req.Password, req.Name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Registration failed", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", "user not authenticated")
		return
	}

	user, err := h.userService.GetUserByID(userID.(string))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get user profile", err.Error())
		return
	}
	if user == nil {
		response.Error(c, http.StatusNotFound, "User not found", "user does not exist")
		return
	}

	response.Success(c, http.StatusOK, "User profile retrieved successfully", user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", "user not authenticated")
		return
	}

	type updateRequest struct {
		Email string `json:"email" binding:"required,email"`
		Name  string `json:"name" binding:"required"`
	}

	var req updateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	err := h.userService.UpdateUser(userID.(string), req.Email, req.Name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to update profile", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Profile updated successfully", nil)
}
