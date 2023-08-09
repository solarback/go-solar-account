package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func PasswordMatches(pass, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err

		}
	}

	return true, nil
}
