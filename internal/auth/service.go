package auth

import (
	"errors"

	"github.com/vyantik/auth-jwt-service/internal/jwt"
	"github.com/vyantik/auth-jwt-service/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService *user.Service
	jwtService  *jwt.Service
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewService(userService *user.Service, jwtService *jwt.Service) *Service {
	return &Service{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (s *Service) Register(req RegisterRequest) (*user.User, error) {
	user := user.User{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}

	existingUser, _ := s.userService.GetByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	createdUser, err := s.userService.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) Login(req LoginRequest) (*TokenPair, error) {
	user, err := s.userService.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshTokens(refreshToken string) (*TokenPair, error) {
	claims, err := s.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.jwtService.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}