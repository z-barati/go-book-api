package services

import (
	"context"
	"errors"

	"go-book-api/internal/models"
	"go-book-api/internal/repositories"
	"go-book-api/pkg/auth"
)

// userService implements UserService
type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register handles user registration
func (s *userService) Register(ctx context.Context, req RegisterRequest) (*models.User, string, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByUsernameOrEmail(ctx, req.Username, req.Email)
	if err == nil && existingUser != nil {
		return nil, "", errors.New("username or email already exists")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, "", errors.New("failed to create user")
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

// Login handles user login
func (s *userService) Login(ctx context.Context, req LoginRequest) (*models.User, string, error) {
	// Find user by username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if !auth.CheckPassword(req.Password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, "", errors.New("account is deactivated")
	}

	// Generate token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

// RefreshToken handles token refresh
func (s *userService) RefreshToken(ctx context.Context, userID uint) (string, error) {
	// Verify user exists
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check if user is active
	if !user.IsActive {
		return "", errors.New("account is deactivated")
	}

	// Generate new token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	return s.userRepo.GetByID(ctx, userID)
} 