package golambda_helper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Recorder interface {
	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

type Handler struct {
	Svc Recorder
}

type DynamoDbRecord struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRecordById(id string, tableName string) (DynamoDbRecord, error) {
	dbRecord := DynamoDbRecord{}

	result, err := h.Svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(id),
					},
				},
			},
		},
	})
	if err != nil {
		return dbRecord, err
	}

	recordList := []DynamoDbRecord{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recordList)
	if err != nil {
		return dbRecord, err
	}

	if len(recordList) != 1 {
		return dbRecord, err
	}

	return recordList[0], nil
}
