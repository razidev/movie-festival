package repository

import (
	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	ListUser(uniqueId []uuid.UUID) ([]models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := configs.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	user.UniqueID = uuid.New()
	if err := configs.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) ListUser(uniqueId []uuid.UUID) ([]models.User, error) {
	var users []models.User
	if err := configs.DB.Where("unique_id IN ?", uniqueId).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
