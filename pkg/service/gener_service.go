package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type GenerService struct {
	repo repository.Gener
}

func NewGenerService(repo repository.Gener) *GenerService {
	return &GenerService{repo: repo}
}

func (repo *GenerService) CreateGener(gener todo.Gener) error {
	return repo.repo.CreateGener(gener)
}

func (repo *GenerService) GetGenerById(generId uint) (gener todo.GetGenerBook, err error) {
	return repo.repo.GetGenerById(generId)
}

func (repo *GenerService) GetGenerList() (generList []todo.GetGenerBook, err error) {
	return repo.repo.GetGenerList()
}

func (repo *GenerService) UpdateGenerInput(generId uint, input todo.UpdateGenerInput) error {
	return repo.repo.UpdateGenerInput(generId, input)
}

func (repo *GenerService) DeleteGener(generId uint) error {
	return repo.repo.DeleteGener(generId)
}
