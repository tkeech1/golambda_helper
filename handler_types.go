package golambda_helper

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

type NewV4er interface {
	NewV4() (uuid.UUID, error)
}

type DynamoQueryer interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

type DynamoPutItemer interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
