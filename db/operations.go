package db

import (
	"errors"
	"fmt"
	"uuid"
)

func CreateMovie(movie *Movie) (*Movie, error) {
	movie.Id = uuid.New().String()

	response := db.Create(&movie)
	if response.Error != nil {
		return nil, response.Error
	}

	return movie, nil
}

func GetAllMovies() ([]*Movie, error) {
	var allMovies []*Movie
	result := db.Find(&allMovies)

	if result.Error != nil {
		return nil, errors.New("Could not recover any movies")
	}
	return allMovies, nil
}

func GetMovie(id string) (*Movie, error) {
	var movie *Movie	
	response := db.First(&movie, "id=?", id)

	if response.RowsAffected == 0 {
		return nil, &InvalidIdError{id}
	}
	return movie, nil
}

func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Delete(&deletedMovie, "id = ?", id)

	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("No movie with id %s found", id))
	}
	return nil
}