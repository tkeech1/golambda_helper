package golambda_helper

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

type UUIDInterface interface {
	NewV4() (uuid.UUID, error)
	GenerateState() (string, error)
}

type UuidHandler struct {
	Uuid UUIDInterface
}

type DynamoInterface interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
	Put(interface{}, string) error
	GetById(string, string, string, interface{}) error
}

type DynamoHandler struct {
	Svc DynamoInterface
}
