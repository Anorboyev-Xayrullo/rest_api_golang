package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (repo *BookService) CreateBook(book todo.CreateBook) error {
	return repo.repo.CreateBook(book)
}

func (repo *BookService) GetBookById(bookId uint) (book todo.GetBookList, err error) {
	return repo.repo.GetBookById(bookId)
}

func (repo *BookService) GetBookList() (bookList []todo.GetBookList, err error) {
	return repo.repo.GetBookList()
}

func (repo *BookService) UpdateBookInput(bookId uint, input todo.UpdateBookInput) error {
	return repo.repo.UpdateBookInput(bookId, input)
}

func (repo *BookService) DeleteBook(bookId uint) error {
	return repo.repo.DeleteBook(bookId)
}
