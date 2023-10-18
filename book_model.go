package todo

type Gener struct {
	BookId int    `json:"book_id" db:"book_id"`
	Gener  string `json:"gener" db:"gener"`
}

type GetGenerBook struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
	NumberOfViews  int    `json:"number_of_views" db:"number_of_views"`
	Gener          string `json:"gener" db:"gener"`
}

type GetRatingBook struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
	NumberOfViews  int    `json:"number_of_views" db:"number_of_views"`
	Rating         int    `json:"rating" db:"rating"`
}

type CreateBook struct {
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
}

type GetBookList struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
}

type UpdateBookInput struct {
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
	GenreId        int    `json:"genreId" db:"genre_id"`
}

type SuccessData struct {
	Data interface{}
}

type UpdateGenerInput struct {
	Gener string `json:"gener" db:"gener"`
}

type RatingInput struct {
	BookId int `json:"book_id" db:"book_id"`
	Rating int `json:"rating" db:"rating"`
}

type Library struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	NumberOfSheets int    `json:"numberOfSheets" db:"number_of_sheets"`
	Author         string `json:"author" db:"author"`
	YearOfBook     int    `json:"yearOfBook" db:"year_of_book"`
	TypeOfCover    bool   `json:"typeOfCover" db:"type_of_cover"`
	Language       string `json:"language" db:"language"`
	Description    string `json:"description" db:"description"`
	Rating         int    `json:"rating" db:"rating"`
	Gener          string `json:"gener" db:"gener"`
}
