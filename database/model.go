package database

type Album struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"releaseDate"`
}
