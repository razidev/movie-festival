package repository

import (
	"github.com/google/uuid"
	"github.com/razidev/movie-festival/src/configs"
	"github.com/razidev/movie-festival/src/models"
)

type UserVoteRepository interface {
	FindCurrentVote(movieUniueId uuid.UUID, userUniqueId uuid.UUID) (models.UserVotes, error)
	CreateVote(vote models.UserVotes) error
}

type userVoteRepository struct{}

func NewUserVoteRepository() UserVoteRepository {
	return &userVoteRepository{}
}

func (ur *userVoteRepository) FindCurrentVote(movieUniqueId uuid.UUID, userUniqueId uuid.UUID) (models.UserVotes, error) {
	var vote models.UserVotes
	if err := configs.DB.Where("movie_unique_id =? AND user_unique_id =?", movieUniqueId, userUniqueId).First(&vote).Error; err != nil {
		return vote, err
	}

	return vote, nil
}

func (ur *userVoteRepository) CreateVote(vote models.UserVotes) error {
	if err := configs.DB.Create(&vote).Error; err != nil {
		return err
	}

	return nil
}
