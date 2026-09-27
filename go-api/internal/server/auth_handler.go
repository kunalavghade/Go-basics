package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
)

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

func (s *Server) GetProfile(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := s.userService.GetUserProfile(userID)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to get user profile", err)
		return
	}
	utils.SucessResponse(c, "User profile fetched successfully", user)
}

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
