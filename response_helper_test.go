package golambdahelper_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkeech1/golambdahelper"
)

func TestHandlerResponseHelper_GenerateError(t *testing.T) {

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
		response, err := golambdahelper.GenerateError(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.StatusCode)
		assert.Equal(t, test.errorReponseText, response.Body)
	}

}

func TestHandlerResponseHelper_GenerateRedirect(t *testing.T) {

	tests := []struct {
		request      string
		responseText string
		expect       int
		err          error
	}{
		{
			responseText: "http://www.google.com",
			request:      "http://www.google.com",
			expect:       302,
			err:          nil,
		},
	}

	for _, test := range tests {
		response, err := golambdahelper.GenerateRedirect(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.StatusCode)
		assert.Equal(t, test.responseText, response.Header.Location)
	}

}

func TestHandlerResponseHelper_GenerateResponseShop(t *testing.T) {

	tests := []struct {
		shopname golambdahelper.ShopName
		response golambdahelper.Response
		err      error
	}{
		{
			shopname: golambdahelper.ShopName{ID: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
			response: golambdahelper.Response{},
			err:      nil,
		},
		{
			shopname: golambdahelper.ShopName{},
			response: golambdahelper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambdahelper.GenerateResponseShop(test.shopname)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambdahelper.ReturnObjectShopName{}
		err = json.Unmarshal([]byte(response.Body), body)
		assert.Equal(t, test.shopname.ID, body.ShopName.ID)
		assert.Equal(t, test.shopname.FriendlyName, body.ShopName.FriendlyName)
		assert.Equal(t, test.shopname.ShopName, body.ShopName.ShopName)
		assert.Equal(t, test.err, err)
	}

}

func TestHandlerResponseHelper_GenerateResponseShops(t *testing.T) {

	tests := []struct {
		shopname []golambdahelper.ShopName
		response golambdahelper.Response
		err      error
	}{
		{
			shopname: []golambdahelper.ShopName{
				{ID: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
				{ID: "someId2", FriendlyName: "Test Friendly Name2", ShopName: "Shop Name2"},
			},
			response: golambdahelper.Response{},
			err:      nil,
		},
		{
			shopname: []golambdahelper.ShopName{},
			response: golambdahelper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambdahelper.GenerateResponseShops(test.shopname)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambdahelper.ReturnObjectShopNames{}
		err = json.Unmarshal([]byte(response.Body), body)
		if len(body.ShopName) > 0 {
			assert.Equal(t, test.shopname[0].ID, body.ShopName[0].ID)
			assert.Equal(t, test.shopname[0].FriendlyName, body.ShopName[0].FriendlyName)
			assert.Equal(t, test.shopname[0].ShopName, body.ShopName[0].ShopName)
			assert.Equal(t, test.err, err)
		}
	}

}
