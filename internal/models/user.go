package models

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
}

func (u *User) Validate() error {
	if len(u.FirstName) < 2 || len(u.FirstName) > 20 {
		return errors.New("first_name must be between 2 and 20 characters")
	}
	if len(u.LastName) < 2 || len(u.LastName) > 20 {
		return errors.New("last_name must be between 2 and 20 characters")
	}
	if len(u.Biography) < 20 || len(u.Biography) > 450 {
		return errors.New("biography must be between 20 and 450 characters")
	}
	return nil
}
