package golambda_helper_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"

	"github.com/tkeech1/golambda_helper"
	"github.com/tkeech1/golambda_helper/mocks"
)

func TestHandlerDynamoHelper(t *testing.T) {

	tests := map[string]struct {
		requestId     string
		queryResponse []map[string]*dynamodb.AttributeValue
		tableName     string
		err           error
	}{
		"success": {
			tableName: "testTable",
			requestId: "testID",
			queryResponse: []map[string]*dynamodb.AttributeValue{
				0: {
					"id": {
						S: aws.String("testID"),
					},
				},
			},
			err: nil,
		},
		"error": {
			tableName: "testTable",
			requestId: "",
			queryResponse: []map[string]*dynamodb.AttributeValue{
				0: {
					"id": {
						S: aws.String(""),
					},
				},
			},
			err: errors.New("Some error"),
		},
		"multiple return records": {
			tableName: "testTable",
			requestId: "",
			queryResponse: []map[string]*dynamodb.AttributeValue{
				0: {
					"id": {
						S: aws.String(""),
					},
				},
				1: {
					"id": {
						S: aws.String(""),
					},
				},
			},
			err: nil,
		},
		"missing array": {
			tableName:     "testTable",
			requestId:     "",
			queryResponse: []map[string]*dynamodb.AttributeValue{},
			err:           nil,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		mockRecorder := &mocks.Recorder{}
		mockRecorder.
			On("Query", &dynamodb.QueryInput{
				TableName: aws.String(test.tableName),
				KeyConditions: map[string]*dynamodb.Condition{
					"id": {
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							{
								S: aws.String(test.requestId),
							},
						},
					},
				},
			}).
			Return(&dynamodb.QueryOutput{
				Items: test.queryResponse,
			}, test.err).
			Once()

		h := &golambda_helper.Handler{
			Svc: mockRecorder,
		}

		response, err := h.GetRecordById(test.requestId, test.tableName)
		assert.Equal(t, response.Id, test.requestId)
		assert.Equal(t, err, test.err)
		mockRecorder.AssertExpectations(t)
	}
}
