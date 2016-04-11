package common

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

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

////////////////////////////
// Tests for GetObjectID. //
////////////////////////////
func TestGetObjectIDReturnsExpectedObjectID(t *testing.T) {
	oid := bson.NewObjectId().Hex()

	output, err := GetObjectID(oid)

	if err != nil || output.Hex() != oid {
		t.Error("Bad output")
		t.Fail()
	}
}

func TestGetObjectIDReturnsErrorForInvalidObjectID(t *testing.T) {
	oid := "invalidOID"

	_, err := GetObjectID(oid)

	if err == nil || err.Error() != "Invalid ObjectId." {
		t.Error("Bad output")
		t.Fail()
	}
}
