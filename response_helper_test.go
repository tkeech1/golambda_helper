package golambda_helper_test

import (
	"errors"
	"testing"

	"github.com/tkeech1/golambda_helper"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request          error
		errorReponseText string
		expect           int
		err              error
	}{
		{
			errorReponseText: "{\"message\":\"Some error\"}",
			request:          errors.New("Some error"),
			expect:           400,
			err:              nil,
		},
		{
			errorReponseText: "{\"message\":\"\"}",
			request:          errors.New(""),
			expect:           400,
			err:              nil,
		},
	}

	for _, test := range tests {
		response, err := golambda_helper.GenerateError(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.StatusCode)
		assert.Equal(t, test.errorReponseText, response.Body)
	}

}
