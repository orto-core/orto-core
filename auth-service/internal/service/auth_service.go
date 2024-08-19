package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/orto-core/server/auth-service/internal/models"
	"github.com/orto-core/server/auth-service/internal/repository"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(*models.User) (string, error)
	Register(*models.User) (string, error)
}

type authService struct {
	repository  repository.AuthRepository
	mailService MailService
	jwtService  JwtService
	otpService  OtpService
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	dir, _ := os.Getwd()
	path := fmt.Sprintf("%s/templates/otp.html", dir)
	mailHost := viper.GetString("mail.host")
	mailUsername := viper.GetString("mail.username")
	mailPassword := viper.GetString("mail.password")

	return &authService{
		repository:  repository,
		jwtService:  NewJwtService(),
		otpService:  NewOtpService(),
		mailService: NewMailService(587, path, mailHost, mailUsername, mailPassword),
	}
}

func (s *authService) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func (s *authService) CompareHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *authService) Login(user *models.User) (string, error) {
	jwt := NewJwtService()

	existingUser, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if existingUser.ID == 0 {
		return "", errors.New("user does exists")
	}

	token, err := jwt.GenerateJWT(&existingUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Register(user *models.User) (string, error) {
	// Check if user already exists
	existingUser, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	fmt.Println(existingUser.ID != 0)
	if existingUser.ID != 0 {
		return "", errors.New("user already exists")
	}

	// Generate OTP secret
	secret, err := GenerateSecret(user.Email)
	if err != nil {
		return "", err
	}

	// Hash password
	hash, err := s.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = string(hash)

	// Generate OTP code
	code, err := s.otpService.GenerateOtp(secret.Secret())
	if err != nil {
		return "", err
	}

	// Send verification email
	if err := s.mailService.SendMail(user.Email, "Account Verification", code); err != nil {
		return "", err
	}

	// Create user
	user.Secret = secret.Secret()
	if err := s.repository.CreateUser(user); err != nil {
		return "", err
	}

	return "User account created successfully. Check your email for OTP code.", nil
}
