package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type RatingService struct {
	repo repository.Rating
}

func NewRatingService(repo repository.Rating) *RatingService {
	return &RatingService{repo: repo}
}

func (repo *RatingService) CreateRating(gener todo.RatingInput) error {
	return repo.repo.CreateRating(gener)
}

func (repo *RatingService) GetRatingList() (ratingList []todo.GetRatingBook, err error) {
	return repo.repo.GetRatingList()
}
