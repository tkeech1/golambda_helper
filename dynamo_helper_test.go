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

func TestHandler2(t *testing.T) {

	tests := map[string]struct {
		recordErr error
		response  string
		tableName string
		id        string
		err       error
	}{
		"successful": {
			recordErr: errors.New("Some error"),
			response:  "i dont know",
			tableName: "testTable",
			id:        "someID",
			err:       nil,
		},
		"with error": {
			recordErr: errors.New("Some error"),
			response:  "i dont know",
			tableName: "testTable",
			id:        "someID",
			err:       nil,
		},
	}

	/*sess, err := session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("ENV_AWS_REGION"))},
		)
		if err != nil {
			return dbRecord, err
	}

	svc := dynamodb.New(sess)*/

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
								S: aws.String(test.id),
							},
						},
					},
				},
			}).
			Return(&dynamodb.QueryOutput{
				Items: []map[string]*dynamodb.AttributeValue{
					0: {
						"id": {
							S: aws.String(test.id),
						},
					},
				},
			}, nil).
			Once()
		h := &golambda_helper.Handler{
			Svc: mockRecorder,
		}
		response, err := h.GetRecordById(test.id, test.tableName)
		assert.Equal(t, response.Id, test.id)
		assert.Equal(t, err, test.err)
		mockRecorder.AssertExpectations(t)
	}
}
