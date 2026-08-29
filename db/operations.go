package db

import (
	"uuid"
)

func CreateMovie(movie *Movie) (*Movie, error) {
	movie.Id = uuid.New().String()

	response := db.Create(&movie)
	if response.Error != nil {
		return nil, Error{ErrNotActioned}
	}

	return movie, nil
}

func GetAllMovies() ([]*Movie, error) {
	var allMovies []*Movie
	result := db.Find(&allMovies)

	if result.Error != nil {
		return nil, Error{ErrNotFound}
	}

	if len(allMovies) == 0 {
		return nil, Error{ErrNotFound}
	}
	return allMovies, nil
}

func GetMovie(id string) (*Movie, error) {
	var movie *Movie	
	response := db.First(&movie, "id=?", id)

	if response.RowsAffected == 0 {
		return nil, Error{ErrNotFound}
	}
	return movie, nil
}

func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Delete(&deletedMovie, "id = ?", id)

	if result.RowsAffected == 0 {
		return Error{ErrNotActioned}
	}
	return nil
}