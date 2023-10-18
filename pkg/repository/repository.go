package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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

type Repository struct {
	Authorization
	Book
	Gener
	Rating
	Library
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookRepo(db),
		Gener:         NewGenerRepo(db),
		Rating:        NewRatingRepo(db),
		Library:       NewLibraryRepo(db),
	}
}
