package golambdahelper

import (
	"encoding/json"
)

// GenerateError is an AWS Lmabda response wrapped in a 400
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

// GenerateRedirect is an AWS Lmabda response wrapped in a 302
func GenerateRedirect(url string) (Response, error) {
	return Response{
		StatusCode: 302,
		Header: Header{
			Location: url,
		},
	}, nil
}

// GenerateResponseShops is an AWS Lmabda response that wraps slice of shopnames
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

// GenerateResponseShop is an AWS Lmabda response that wraps a single shop
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
