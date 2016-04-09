package common

import "golang.org/x/crypto/bcrypt"

// HashPassword returns the hash of the given plain text password.
func HashPassword(plainTxtPasswd string) string {
	passwordBytes := []byte(plainTxtPasswd)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordBytes, 10)
	return string(hashedPassword[:])
}
