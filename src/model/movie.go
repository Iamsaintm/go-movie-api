package model

type Movie struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	ImdbID      string  `json:"imdbID" gorm:"column:imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero" gorm:"column:isSuperHero"`
}

func (Movie) TableName() string {
	return "goimdb"
}
