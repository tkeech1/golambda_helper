package golambda_helper_test

import (
	"errors"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tkeech1/golambda_helper"
	"github.com/tkeech1/golambda_helper/mocks"
)

func TestHandlerShopify_GenerateState(t *testing.T) {

	tests := map[string]struct {
		Response     string
		UuidResponse uuid.UUID
		UuidErr      error
	}{
		"success": {
			Response: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			UuidResponse: uuid.UUID{
				0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
			},
			UuidErr: nil,
		},
		"failure_uuid": {
			Response: "",
			UuidErr:  errors.New("An error"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		mockUuidInterface := &mocks.UUIDInterface{}
		mockUuidInterface.
			On("NewV4").
			Return(test.UuidResponse, test.UuidErr).
			Once()

		h := &golambda_helper.UuidHandler{
			Uuid: mockUuidInterface,
		}
		response, err := h.GenerateState()
		assert.Equal(t, test.Response, response)
		assert.Equal(t, test.UuidErr, err)
	}
}

func TestHandlerShopify_NewV4(t *testing.T) {

	tests := map[string]struct {
		Expected  string
		UuidInput uuid.UUID
		UuidErr   error
	}{
		"success": {
			Expected: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
			UuidInput: uuid.UUID{
				0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8,
			},
			UuidErr: nil,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		uuidHandler := &golambda_helper.UuidHandler{}
		h := golambda_helper.UuidHandler{
			Uuid: uuidHandler,
		}
		response, err := h.NewV4()
		assert.NotEqual(t, test.Expected, response.String())
		assert.Equal(t, len(test.Expected), len(response.String()))
		assert.Equal(t, test.UuidErr, err)
	}
}
