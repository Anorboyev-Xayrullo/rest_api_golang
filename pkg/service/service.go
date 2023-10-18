package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Book interface {
	CreateBook(book todo.CreateBook) error
	GetBookList() (bookList []todo.GetBookList, err error)
	GetBookById(bookId uint) (book todo.GetBookList, err error)
	UpdateBookInput(bookId uint, input todo.UpdateBookInput) error
	DeleteBook(bookId uint) error
}

type Gener interface {
	CreateGener(gener todo.Gener) error
	GetGenerList() (generList []todo.GetGenerBook, err error)
	GetGenerById(generId uint) (gener todo.GetGenerBook, err error)
	UpdateGenerInput(generId uint, input todo.UpdateGenerInput) error
	DeleteGener(generId uint) error
}

type Rating interface {
	CreateRating(gener todo.RatingInput) error
	GetRatingList() (ratingList []todo.GetRatingBook, err error)
}

type Library interface {
	GetLibraryList() (libraryList []todo.Library, err error)
}

type Service struct {
	Authorization
	Book
	Gener
	Rating
	Library
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book),
		Gener:         NewGenerService(repos.Gener),
		Rating:        NewRatingService(repos.Rating),
		Library:       NewLibraryService(repos.Library),
	}
}
