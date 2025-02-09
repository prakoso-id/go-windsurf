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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and name
// @Tags users
// @Accept json
// @Produce json
// @Param user body registerRequest true "User registration details"
// @Success 201 {object} response.Response{data=models.User} "User registered successfully"
// @Failure 400 {object} response.Response "Invalid request parameters"
// @Router /api/v1/users/register [post]
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

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=models.User} "User profile retrieved successfully"
// @Failure 401 {object} response.Response "Unauthorized"
// @Failure 404 {object} response.Response "User not found"
// @Router /api/v1/users/profile [get]
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

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update the profile of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body updateProfileRequest true "User profile update details"
// @Success 200 {object} response.Response "Profile updated successfully"
// @Failure 400 {object} response.Response "Invalid request parameters"
// @Failure 401 {object} response.Response "Unauthorized"
// @Router /api/v1/users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", "user not authenticated")
		return
	}

	type updateProfileRequest struct {
		Email string `json:"email" binding:"required,email"`
		Name  string `json:"name" binding:"required"`
	}

	var req updateProfileRequest
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

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param passwords body changePasswordRequest true "Password change details"
// @Success 200 {object} response.Response "Password changed successfully"
// @Failure 400 {object} response.Response "Invalid request parameters"
// @Failure 401 {object} response.Response "Unauthorized"
// @Router /api/v1/users/change-password [post]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized", "user not authenticated")
		return
	}

	type changePasswordRequest struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword    string `json:"new_password" binding:"required,min=6"`
	}

	var req changePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	// First verify the current password
	_, err := h.userService.Login(userID.(string), req.CurrentPassword)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid current password", err.Error())
		return
	}

	// Update the password
	if err := h.userService.UpdatePassword(userID.(string), req.NewPassword); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to change password", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Password changed successfully", nil)
}
