package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type LibraryService struct {
	repo repository.Library
}

func NewLibraryService(repo repository.Library) *LibraryService {
	return &LibraryService{repo: repo}
}

func (repo *LibraryService) GetLibraryList() (libraryList []todo.Library, err error) {
	return repo.repo.GetLibraryList()
}
