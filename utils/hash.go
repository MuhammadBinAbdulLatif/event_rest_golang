package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainText string) (string, error) {
	hashedPassword, err:=bcrypt.GenerateFromPassword([]byte(plainText), 14)
	return string(hashedPassword), err
}


func AreSame(password string, hashedPassword string) bool {
	err :=bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return  err == nil
}