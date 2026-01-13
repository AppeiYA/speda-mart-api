package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), cost)
	if hashErr != nil {
		log.Println("Hash error: ", hashErr)
		return "", hashErr
	}
	return string(hash), nil
}

func CompareHashAndPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
