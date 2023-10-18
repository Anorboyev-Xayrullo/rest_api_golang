package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
)

type LibraryRepo struct {
	db *sqlx.DB
}

func NewLibraryRepo(db *sqlx.DB) *LibraryRepo {
	return &LibraryRepo{db: db}
}

func (repo *LibraryRepo) GetLibraryList() (libraryList []todo.Library, err error) {
	db := repo.db
	err = db.Select(&libraryList, `select id, name, author, year_of_book, "language", type_of_cover, gener, rating, description, number_of_sheets from books inner join gener on gener.book_id = books.id inner join rating on rating.book_id = books.id`)
	if err != nil {
		return libraryList, err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return libraryList, err
		}
		return libraryList, err
	}
	return libraryList, err
}
