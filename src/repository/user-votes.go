package repository

import (
	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type UserVoteRepository interface {
	FindCurrentVote(movieUniueId uuid.UUID, userUniqueId uuid.UUID) (models.UserVotes, error)
	CreateVote(vote models.UserVotes) error
	UpdateVote(id uint, status string) error
}

type userVoteRepository struct{}

func NewUserVoteRepository() UserVoteRepository {
	return &userVoteRepository{}
}

func (r *userVoteRepository) FindCurrentVote(movieUniqueId uuid.UUID, userUniqueId uuid.UUID) (models.UserVotes, error) {
	var vote models.UserVotes
	if err := configs.DB.Where("movie_unique_id =? AND user_unique_id =?", movieUniqueId, userUniqueId).First(&vote).Error; err != nil {
		return vote, err
	}

	return vote, nil
}

func (r *userVoteRepository) CreateVote(vote models.UserVotes) error {
	if err := configs.DB.Create(&vote).Error; err != nil {
		return err
	}

	return nil
}

func (r *userVoteRepository) UpdateVote(id uint, status string) error {
	if err := configs.DB.Model(&models.UserVotes{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
