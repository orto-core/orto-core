package repository

import (
	"errors"
	"log"

	"github.com/orto-core/server/auth-service/internal/models"
	"github.com/orto-core/server/auth-service/internal/store"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(*models.User) error
	GetUserById(uint) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByEmail(string) (models.User, error)
	UpdateUser(*models.User) error
	DeleteUserById(uint) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: store.DB,
	}
}

func (r *authRepository) CreateUser(user *models.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, nil
		}
		log.Printf("Error retrieving user by email: %v", err)
		return models.User{}, err
	}
	return user, nil
}

func (r *authRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *authRepository) UpdateUser(user *models.User) error {
	if err := r.db.Model(&user).Updates(models.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) DeleteUserById(id uint) error {
	if err := r.db.Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
