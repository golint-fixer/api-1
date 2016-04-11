package common

import (
	"testing"

	"github.com/gin-gonic/gin"
)

//////////////////////////
// Tests for RequestID. //
//////////////////////////
func TestRequestIDSetsExpectedRequestID(t *testing.T) {
	context, _, _ := gin.CreateTestContext()

	RequestID(context)
	value, exists := context.Get("request_id")

	if !exists {
		t.Error("Key is expected to exist.")
		t.Fail()
	}

	length := len(value.(string))
	if length != 36 { // Length of UUID4.
		t.Errorf("Length of ID should be 36; got %d", length)
		t.Fail()
	}
}

func TestRequestIDSetsExpectedHeader(t *testing.T) {
	context, _, _ := gin.CreateTestContext()

	RequestID(context)
	header := context.Writer.Header().Get("X-Request-ID")

	length := len(header)
	if length != 36 { // Length of UUID4.
		t.Errorf("Length of ID should be 36; got %d", length)
		t.Fail()
	}
}
