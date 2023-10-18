package repository

import (
	"database/sql"
	"errors"
	"github.com/zhashkevych/todo-app"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (repo *BookRepo) CreateBook(book todo.CreateBook) error {

	db := repo.db
	row, err := db.Exec(`insert into Books (name, number_of_sheets, author, year_of_book, type_of_cover, language, description) values ($1, $2, $3, $4 , $5 , $6, $7)`, book.Name, book.NumberOfSheets, book.Author, book.YearOfBook, book.TypeOfCover, book.Language, book.Description)
	if err != nil {
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		return errors.New("not update")
	}
	return nil
}

func (repo *BookRepo) GetBookList() (bookList []todo.GetBookList, err error) {
	db := repo.db
	err = db.Select(&bookList, `select id, name, author, year_of_book, "language", type_of_cover, description, number_of_sheets from books`)
	if err != nil {
		return bookList, err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return bookList, err
		}
		return bookList, err
	}
	return bookList, err
}

func (repo *BookRepo) GetBookById(bookId uint) (book todo.GetBookList, err error) {
	db := repo.db
	err = db.Get(&book, `select id, name, author, year_of_book, "language", type_of_cover, description, number_of_sheets from Books where id = $1`, bookId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return book, err
		}
		return book, err
	}
	return book, err
}

func (repo *BookRepo) UpdateBookInput(bookId uint, input todo.UpdateBookInput) error {
	db := repo.db
	row, err := db.Exec(`update Books set name = $1 , number_of_sheets = $2 , author = $3 , year_of_book = $4 , type_of_cover = $5 , language = $6 , description = $7 , genre_id = $8 where id = $9`, input.Name, input.NumberOfSheets, input.Author, input.YearOfBook, input.TypeOfCover, input.Language, input.Description, input.GenreId, bookId)
	if err != nil {
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("not working")
	}
	return nil
}

func (repo *BookRepo) DeleteBook(bookId uint) error {
	db := repo.db
	row, err := db.Exec(`delete from Books where id = $1`, bookId)
	if err != nil {
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		return errors.New("not update")
	}
	return nil
}
