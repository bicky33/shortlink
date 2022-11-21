package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	HashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(HashPassword), err
}

func MatchPassword(hasPassword, currentPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasPassword), []byte(currentPassword))
	return err == nil
}
