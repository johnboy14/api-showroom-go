package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Create(user User) (User, error) {
	return user, nil
}

var (
	ErrInvalidFirstName = errors.New("invalid first name")
	ErrInvalidLastName  = errors.New("invalid last name")
	ErrInvalidEmail     = errors.New("invalid email address")
	ErrInvalidPassword  = errors.New("invalid password")
)

func (u User) Validate(r *http.Request) error {
	if !govalidator.StringLength(u.FirstName, "2", "30") {
		return ErrInvalidFirstName
	}

	if !govalidator.StringLength(u.LastName, "2", "30") {
		return ErrInvalidLastName
	}

	if !govalidator.IsEmail(u.Email) {
		return ErrInvalidEmail
	}

	if !govalidator.StringLength(u.Password, "2", "30") {
		return ErrInvalidPassword
	}

	return nil
}
