package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", errors.New("failed to encrypt password")
	}

	return string(passHash), nil
}
