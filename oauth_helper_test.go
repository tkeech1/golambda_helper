package golambdahelper_test

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tkeech1/golambdahelper"
	"testing"
)

// https://medium.com/@matryer/meet-moq-easily-mock-interfaces-in-go-476444187d10
type NewV4erMock struct {
	NewV4Func func() (uuid.UUID, error)
}

func (mock NewV4erMock) NewV4() (uuid.UUID, error) {
	return mock.NewV4Func()
}

func TestHandlerShopify_GenerateState(t *testing.T) {

	tests := map[string]struct {
		Response     string
		NewV4Mock    NewV4erMock
		UuidResponse uuid.UUID
		UuidErr      error
	}{
		"success": {
			NewV4Mock: NewV4erMock{
				NewV4Func: func() (uuid.UUID, error) {
					return uuid.UUID{
						0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
					}, nil
				},
			},
			Response: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			UuidErr:  nil,
		},
		"failure_uuid": {
			NewV4Mock: NewV4erMock{
				NewV4Func: func() (uuid.UUID, error) {
					return uuid.UUID{}, errors.New("An error")
				},
			},
			Response: "",
			UuidErr:  errors.New("An error"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response, err := golambdahelper.GenerateState(test.NewV4Mock)
		assert.Equal(t, test.Response, response)
		assert.Equal(t, test.UuidErr, err)
	}
}
