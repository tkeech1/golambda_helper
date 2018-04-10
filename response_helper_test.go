package golambda_helper_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tkeech1/golambda_helper"
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
		response, err := golambda_helper.GenerateError(test.request)
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
		response, err := golambda_helper.GenerateRedirect(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.StatusCode)
		assert.Equal(t, test.responseText, response.Header.Location)
	}

}

func TestHandlerResponseHelper_GenerateResponseShop(t *testing.T) {

	tests := []struct {
		shopname golambda_helper.ShopName
		response golambda_helper.Response
		err      error
	}{
		{
			shopname: golambda_helper.ShopName{Id: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
			response: golambda_helper.Response{},
			err:      nil,
		},
		{
			shopname: golambda_helper.ShopName{},
			response: golambda_helper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambda_helper.GenerateResponseShop(test.shopname)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambda_helper.ReturnObjectShopName{}
		err = json.Unmarshal([]byte(response.Body), body)
		assert.Equal(t, test.shopname.Id, body.ShopName.Id)
		assert.Equal(t, test.shopname.FriendlyName, body.ShopName.FriendlyName)
		assert.Equal(t, test.shopname.ShopName, body.ShopName.ShopName)
	}

}

func TestHandlerResponseHelper_GenerateResponseShops(t *testing.T) {

	tests := []struct {
		shopname []golambda_helper.ShopName
		response golambda_helper.Response
		err      error
	}{
		{
			shopname: []golambda_helper.ShopName{
				{Id: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
				{Id: "someId2", FriendlyName: "Test Friendly Name2", ShopName: "Shop Name2"},
			},
			response: golambda_helper.Response{},
			err:      nil,
		},
		{
			shopname: []golambda_helper.ShopName{},
			response: golambda_helper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambda_helper.GenerateResponseShops(test.shopname)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambda_helper.ReturnObjectShopNames{}
		err = json.Unmarshal([]byte(response.Body), body)
		if len(body.ShopName) > 0 {
			assert.Equal(t, test.shopname[0].Id, body.ShopName[0].Id)
			assert.Equal(t, test.shopname[0].FriendlyName, body.ShopName[0].FriendlyName)
			assert.Equal(t, test.shopname[0].ShopName, body.ShopName[0].ShopName)
		}
	}

}
