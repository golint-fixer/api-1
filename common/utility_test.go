package common

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

/////////////////////
// Tests for Hash. //
/////////////////////
func TestHash(t *testing.T) {
	plainText := "plain-text-password"

	output := Hash(plainText)

	// Ensure relationship between plain text and hash can be computed.
	if bcrypt.CompareHashAndPassword([]byte(output), []byte(plainText)) != nil {
		t.Fail()
	}
}

//////////////////////////
// Tests for CheckHash. //
//////////////////////////
func TestCheckHashReturnsTrueWhenTrue(t *testing.T) {
	plainText := "plain-text-password"
	hash := Hash(plainText)

	output := CheckHash(hash, plainText) // Should be true.

	if output != true {
		t.Fail()
	}
}

func TestCheckHashReturnsFalseWhenFalse(t *testing.T) {
	plainText := "plain-text-password"
	hash := Hash("some-other-string")

	output := CheckHash(hash, plainText) // Should be false.

	if output != false {
		t.Fail()
	}
}
