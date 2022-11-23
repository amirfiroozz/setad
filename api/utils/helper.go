package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsWrongPassword(actual string, expected string) error {
	err := bcrypt.CompareHashAndPassword([]byte(expected), []byte(actual))
	if err != nil {
		return WrongPasswordError
	}
	return nil
}
