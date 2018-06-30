package golambda_helper

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

type NewV4er interface {
	NewV4() (uuid.UUID, error)
}

type Queryer interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

type PutItemer interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

type Puter interface {
	Put(interface{}, string) error
}

type GetByIDer interface {
	GetById(string, string, string, interface{}) error
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
