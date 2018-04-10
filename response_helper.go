package golambda_helper

import (
	"encoding/json"
)

func GenerateError(err error) (Response, error) {
	errorResponse := ErrorResponse{err.Error()}
	responseBody, err := json.Marshal(errorResponse)
	return Response{
		Body:       string(responseBody),
		StatusCode: 400,
		Header: Header{
			ContentType:              "application/json",
			AccessControlAllowOrigin: "*",
		},
	}, nil
}

func GenerateRedirect(url string) (Response, error) {
	return Response{
		StatusCode: 302,
		Header: Header{
			Location: url,
		},
	}, nil
}

func GenerateResponseShops(shopname []ShopName) (Response, error) {
	returnObject := ReturnObjectShopNames{shopname}
	responseBody, err := json.Marshal(returnObject)
	if err == nil {
		return Response{
			Body:       string(responseBody),
			StatusCode: 200,
			Header: Header{
				ContentType:              "application/json",
				AccessControlAllowOrigin: "*",
			},
		}, nil
	}
	return GenerateError(err)
}

func GenerateResponseShop(shopname ShopName) (Response, error) {
	returnObject := ReturnObjectShopName{shopname}
	responseBody, err := json.Marshal(returnObject)
	if err == nil {
		return Response{
			Body:       string(responseBody),
			StatusCode: 200,
			Header: Header{
				ContentType:              "application/json",
				AccessControlAllowOrigin: "*",
			},
		}, nil
	}
	return GenerateError(err)
}
