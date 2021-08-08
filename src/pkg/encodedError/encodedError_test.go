package encodedError

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertEqual checks if values are equal
/*
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
*/

type SendAndResponse struct {
	send     error
	response int
}

func TestGetCode(t *testing.T) {
	sendAndResponse := []SendAndResponse{
		{BadRequest, 400},
		{InternalServerError, 500},
		{errors.New("test"), 500},
	}

	for _, item := range sendAndResponse {
		assert.Equal(t, getCode(item.send), item.response)
	}
}
