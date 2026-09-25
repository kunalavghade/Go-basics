package services

import (
	"errors"
	"time"

	"github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/dto"
	"github.com/kunalavghade/Go-basics/go-api/internal/models"
	"github.com/kunalavghade/Go-basics/go-api/internal/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	db     *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB, config *config.Config) *AuthService {
	return &AuthService{
		db:     db,
		config: config,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthenticationResponse, error) {
	// Check if user already exists
	var user models.User
	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err == nil {
		return nil, errors.New("a user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Hash Password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create User
	user = models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Role:      models.UserRoleCustomer,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	cart := models.Cart{
		UserID: user.ID,
	}
	if err := s.db.Create(&cart).Error; err != nil {
		return nil, err
	}

	return s.generateAuthResponse(&user)
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthenticationResponse, error) {
	var user models.User
	if err := s.db.Where("email = ? AND is_active = ? ", req.Email, true).First(&user).Error; err != nil {
		return nil, err
	}
	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, errors.New("Invalid credentials")
	}
	return s.generateAuthResponse(&user)
}

func (s *AuthService) RefreshToken(req *dto.RequestTokenRequest) (*dto.AuthenticationResponse, error) {
	claims, err := utils.ValidateToken(req.RefreshToken, s.config.JWT.Secret)
	if err != nil {
		return nil, errors.New("Invalid refresh token")
	}
	var refreshToken models.RefreshToken
	if err := s.db.Where("token = ? AND expires_at > ?", req.RefreshToken, time.Now()).First(&refreshToken).Error; err != nil {
		return nil, errors.New("refresh token not found or expired")
	}
	var user models.User
	if err := s.db.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return nil, errors.New("user not found or invalid user")
	}

	s.db.Delete(&refreshToken)
	return s.generateAuthResponse(&user)
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.db.Where("token = ?", refreshToken).Delete(&models.RefreshToken{}).Error
}

func (s *AuthService) generateAuthResponse(user *models.User) (*dto.AuthenticationResponse, error) {
	accessToken, refreshToken, err := utils.GenerateTockenPair(s.config, user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	refreshTokenModel := models.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(s.config.JWT.RefreshTokenExpiresIn),
	}

	if err := s.db.Create(&refreshTokenModel).Error; err != nil {
		return nil, err
	}

	return &dto.AuthenticationResponse{
		User: dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Phone:     user.Phone,
			Role:      string(user.Role),
			IsActive:  user.IsActive,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshTokenModel.Token,
	}, nil
}
