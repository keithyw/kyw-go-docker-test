package utils

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(h), nil
}