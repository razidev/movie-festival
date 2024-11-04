package repository

import (
	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type UserRepository interface {
	FindByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
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
