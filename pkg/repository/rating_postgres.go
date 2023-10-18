package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/todo-app"
)

type RatingRepo struct {
	db *sqlx.DB
}

func NewRatingRepo(db *sqlx.DB) *RatingRepo {
	return &RatingRepo{db: db}
}

func (repo *RatingRepo) CreateRating(rating todo.RatingInput) error {
	db := repo.db
	row, err := db.Exec(`insert into rating (book_id, rating) values ($1, $2)`, rating.BookId, rating.Rating)
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

func (repo *RatingRepo) GetRatingList() (ratingList []todo.GetRatingBook, err error) {
	db := repo.db
	err = db.Select(&ratingList, `select name, author, year_of_book, "language", type_of_cover, rating, description, number_of_sheets from books inner join rating on rating.book_id = books.id`)
	if err != nil {
		return ratingList, err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return ratingList, err
		}
		return ratingList, err
	}
	return ratingList, err
}
