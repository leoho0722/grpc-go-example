package database

func Create(album *Album) (err error) {
	err = DB.Create(album).Error
	return
}

func Get(id string) (result *Album, err error) {
	result = &Album{}
	err = DB.Where("id = ?", id).First(result).Error
	return
}

func GetAll() ([]Album, error) {
	var albums []Album
	result := DB.Find(&albums)
	return albums, result.Error
}
