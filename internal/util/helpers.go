package util

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPass(password string) (string, error) {
	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate hash")
	}
	return string(hash), nil
}

func VerifyPassword(hashedPassword, password string) error {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return errors.Wrap(err, "passwords do not match")
	}
	return nil
}
