package golambda_helper

import (
	"encoding/json"
)

type Header struct {
	ContentType              string `json:"Content-Type"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Header     Header `json:"headers"`
}

var headers = Header{
	ContentType:              "application/json",
	AccessControlAllowOrigin: "*",
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func GenerateError(err error) (Response, error) {
	errorResponse := ErrorResponse{err.Error()}
	responseBody, err := json.Marshal(errorResponse)
	return Response{
		Body:       string(responseBody),
		StatusCode: 400,
		Header:     headers,
	}, nil
}
