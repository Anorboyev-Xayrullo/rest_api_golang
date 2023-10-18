package repository

import (
	"database/sql"
	"errors"
	"github.com/zhashkevych/todo-app"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type GenerRepo struct {
	db *sqlx.DB
}

func NewGenerRepo(db *sqlx.DB) *GenerRepo {
	return &GenerRepo{db: db}
}

func (repo *GenerRepo) CreateGener(gener todo.Gener) error {

	db := repo.db
	row, err := db.Exec(`insert into gener (book_id, gener) values ($1, $2)`, gener.BookId, gener.Gener)
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

func (repo *GenerRepo) GetGenerList() (generList []todo.GetGenerBook, err error) {
	db := repo.db
	err = db.Select(&generList, `select id, name, author, year_of_book, "language", type_of_cover, gener, description, number_of_sheets from books inner join gener on gener.book_id = books.id`)
	if err != nil {
		return generList, err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return generList, err
		}
		return generList, err
	}
	return generList, err
}

func (repo *GenerRepo) GetGenerById(generId uint) (gener todo.GetGenerBook, err error) {
	db := repo.db
	err = db.Get(&gener, `select id, name, author, year_of_book, "language", type_of_cover, gener, description, number_of_sheets from books inner join gener on gener.book_id = books.id where book_id = $1`, generId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return gener, err
		}
		return gener, err
	}
	return gener, err
}

func (repo *GenerRepo) UpdateGenerInput(generId uint, input todo.UpdateGenerInput) error {
	db := repo.db
	row, err := db.Exec(`update gener set gener = $1 where book_id = $2`, input.Gener, generId)
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

func (repo *GenerRepo) DeleteGener(generId uint) error {
	db := repo.db
	row, err := db.Exec(`delete from gener where book_id = $1`, generId)
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
