package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and personal details
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User Registration Info"
// @Success 201 {object} utils.Response{data=dto.AuthenticationResponse}
// @Failure 400 {object} utils.Response
// @Router /api/auth/register [post]
func (s *Server) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	response, err := s.authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to register user", err)
		return
	}

	utils.CreatedResponse(c, "User registered successfully", response)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return access & refresh tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login Credentials"
// @Success 200 {object} utils.Response{data=dto.AuthenticationResponse}
// @Failure 400 {object} utils.Response
// @Router /api/auth/login [post]
func (s *Server) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	response, err := s.authService.Login(&req)
	if err != nil {
		utils.BadRequestResponse(c, "Failed to login user", err)
		return
	}

	utils.SucessResponse(c, "User logged in successfully", response)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Obtain a new access token using a refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RequestTokenRequest true "Refresh Token"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /api/auth/refresh [post]
func (s *Server) RefreshToken(c *gin.Context) {
	var req dto.RequestTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	response, err := s.authService.RefreshToken(&req)
	if err != nil {
		utils.UnauthorizedResponse(c, "Failed to refresh token", err)
		return
	}

	utils.SucessResponse(c, "Token refreshed successfully", response)
}

// Logout godoc
// @Summary Logout user
// @Description Invalidate the user refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.RequestTokenRequest true "Refresh Token"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /api/auth/logout [post]
func (s *Server) Logout(c *gin.Context) {
	var req dto.RequestTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}

	err := s.authService.Logout(req.RefreshToken)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to logout user", err)
		return
	}

	utils.SucessResponse(c, "User logged out successfully", nil)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get profile of the currently authenticated user
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Response{data=dto.UserResponse}
// @Failure 401 {object} utils.Response
// @Router /api/users/profile [get]
func (s *Server) GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := s.userService.GetUserProfile(userID)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get user profile", err)
		return
	}
	utils.SucessResponse(c, "User profile fetched successfully", user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update profile details of the currently authenticated user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateProfileRequest true "Profile details"
// @Success 200 {object} utils.Response{data=dto.UserResponse}
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /api/users/profile [put]
func (s *Server) UpdateProfile(c *gin.Context) {
	userID := c.GetInt("user_id")
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request data", err)
		return
	}
	user, err := s.userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to update user profile", err)
		return
	}
	utils.SucessResponse(c, "User profile updated successfully", user)
}
