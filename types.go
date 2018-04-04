package golambda_helper

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DynamoInterface interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

type Handler struct {
	Svc DynamoInterface
}

type Header struct {
	ContentType              string `json:"Content-Type"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Header     Header `json:"headers"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ReturnObjectShop struct {
	Shop Shop `json:"shopname"`
}

type ReturnObjectShops struct {
	Shop []Shop `json:"shopnames"`
}

type Shop struct {
	Id           string `json:"id"`
	FriendlyName string `json:"friendly_name"`
	ShopName     string `json:"shop_name"`
	CreateDate   string `json:"create_date"`
	Deleted      string `json:"deleted"`
}
