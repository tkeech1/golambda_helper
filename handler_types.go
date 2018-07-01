package golambdahelper

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/satori/go.uuid"
)

// NewV4er wraps NewV4
type NewV4er interface {
	NewV4() (uuid.UUID, error)
}

// DynamoQueryer defines a DynamoDB query
type DynamoQueryer interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

// DynamoPutItemer defines a DynamoDB putitem operation
type DynamoPutItemer interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
