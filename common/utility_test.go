package common

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUtilityHashPassword(t *testing.T) {
	plainTextPassword := "testpass"

	output := HashPassword(plainTextPassword)

	// Ensure relationship between password and hash can be computed.
	if bcrypt.CompareHashAndPassword([]byte(output), []byte(plainTextPassword)) != nil {
		t.Fail()
	}
}
