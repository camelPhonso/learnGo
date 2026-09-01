package db

import (
	"errors"
)

// sentient errors
var (
	ErrInvalidParameter = errors.New("The parameter is of an invalid type.")
	ErrNotFound = errors.New("The query did not match any existing records.")
	ErrNotActioned = errors.New("The operation could not be completed.")
	ErrBadFile = errors.New("The file could not be read.")
)

// Implements the error interface
type Error struct {
	DbError error
}

func (e Error) Error() string {
	return e.DbError.Error()
}