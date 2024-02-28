package util

import (
	"github.com/inancgumus/screen"
	"golang.org/x/crypto/bcrypt"
)

func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func CreateHashPass(pasword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pasword), bcrypt.MinCost)
	return string(hash), err
}
func CompareHashPass(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
