package common

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Hash returns a hash of the given plain text.
func Hash(plainText string) string {
	textBytes := []byte(plainText)
	hash, _ := bcrypt.GenerateFromPassword(textBytes, 10)
	return string(hash[:])
}

// CheckHash returns a boolean indicating if the given hash was computed from the given plain text.
func CheckHash(hash, plainText string) bool {
	hashBytes := []byte(hash)
	textBytes := []byte(plainText)
	if err := bcrypt.CompareHashAndPassword(hashBytes, textBytes); err != nil {
		return false
	}
	return true
}

// GetObjectID get an ObjectId from the given string, or error.
func GetObjectID(id string) (oid bson.ObjectId, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Invalid ObjectId.")
		}
	}()
	oid = bson.ObjectIdHex(id)
	return oid, err
}
