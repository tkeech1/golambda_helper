package golambda_helper

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

func (h *DynamoHandler) Query(queryInput *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("ENV_AWS_REGION")),
	})
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(sess)
	return svc.Query(queryInput)
}

func (h *DynamoHandler) PutItem(putItem *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("ENV_AWS_REGION")),
	})
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(sess)
	return svc.PutItem(putItem)
}
