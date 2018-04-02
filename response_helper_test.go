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

func TestHandlerResponseHelper_GenerateResponseShop(t *testing.T) {

	tests := []struct {
		shop     golambda_helper.Shop
		response golambda_helper.Response
		err      error
	}{
		{
			shop:     golambda_helper.Shop{Id: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
			response: golambda_helper.Response{},
			err:      nil,
		},
		{
			shop:     golambda_helper.Shop{},
			response: golambda_helper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambda_helper.GenerateResponseShop(test.shop)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambda_helper.ReturnObjectShop{}
		err = json.Unmarshal([]byte(response.Body), body)
		assert.Equal(t, test.shop.Id, body.Shop.Id)
		assert.Equal(t, test.shop.FriendlyName, body.Shop.FriendlyName)
		assert.Equal(t, test.shop.ShopName, body.Shop.ShopName)
	}

}

func TestHandlerResponseHelper_GenerateResponseShops(t *testing.T) {

	tests := []struct {
		shop     []golambda_helper.Shop
		response golambda_helper.Response
		err      error
	}{
		{
			shop: []golambda_helper.Shop{
				{Id: "someId", FriendlyName: "Test Friendly Name", ShopName: "Shop Name"},
				{Id: "someId2", FriendlyName: "Test Friendly Name2", ShopName: "Shop Name2"},
			},
			response: golambda_helper.Response{},
			err:      nil,
		},
		{
			shop:     []golambda_helper.Shop{},
			response: golambda_helper.Response{},
			err:      nil,
		},
	}

	for _, test := range tests {
		response, err := golambda_helper.GenerateResponseShops(test.shop)
		assert.IsType(t, test.err, err)
		assert.IsType(t, test.response, response)
		body := &golambda_helper.ReturnObjectShops{}
		err = json.Unmarshal([]byte(response.Body), body)
		if len(body.Shop) > 0 {
			assert.Equal(t, test.shop[0].Id, body.Shop[0].Id)
			assert.Equal(t, test.shop[0].FriendlyName, body.Shop[0].FriendlyName)
			assert.Equal(t, test.shop[0].ShopName, body.Shop[0].ShopName)
		}
	}

}
