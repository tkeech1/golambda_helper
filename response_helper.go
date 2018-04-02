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

func GenerateResponseShops(shop []Shop) (Response, error) {
	returnObject := ReturnObjectShops{shop}
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

func GenerateResponseShop(shop Shop) (Response, error) {
	returnObject := ReturnObjectShop{shop}
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
