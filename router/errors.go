package router

import (
	"errors"
	"net/http"

	db "local.learn.go/db"
)

// sentient errors
var (
	ErrMissingParameter = errors.New("Expected parameter not found.")
	ErrBadRequest = errors.New("The request was formatted incorrectly.")
)

type Error struct {
	Status int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

// Formats the router error based on the sentient error
// returned from the database operation
func FromError(e error) Error {
	var routerError Error

	var dbError db.Error
	routerError.Message = e.Error()
	
	if errors.As(e, &dbError) {
		switch dbError.DbError {
			case db.ErrInvalidParameter:
				routerError.Status = http.StatusBadRequest
				break
			case db.ErrNotActioned:
				routerError.Status = http.StatusInternalServerError
				break
			case db.ErrNotFound:
				routerError.Status = http.StatusNotFound
				break
			default:
				routerError.Status = http.StatusInternalServerError
		}
	}

	return routerError
}