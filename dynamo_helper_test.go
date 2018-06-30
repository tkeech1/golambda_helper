package golambda_helper

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
)

type QueryerMock struct {
	QueryFunc func(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

func (mock QueryerMock) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return mock.QueryFunc(input)
}

type PutItemerMock struct {
	PutItemFunc func(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func (mock PutItemerMock) PutItem(item *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return mock.PutItemFunc(item)
}

func TestHandlerDynamoHelper_GetById(t *testing.T) {

	tests := map[string]struct {
		idValue     string
		idName      string
		tableName   string
		queryer     QueryerMock
		errResponse error
	}{
		"error_0_records": {
			tableName: "testTable",
			idValue:   "THEVALUE",
			idName:    "THEID",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{},
					}, nil
				},
			},
			errResponse: errors.New("An error occurred during processing."),
		},
		"success_1_record": {
			tableName: "testTable",
			idValue:   "THEVALUE",
			idName:    "THEID",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							0: {
								"id": {
									S: aws.String("THEVALUE"),
								},
							},
						},
					}, nil
				},
			},
			errResponse: nil,
		},
		"error_2_records": {
			tableName: "testTable",
			idValue:   "THEVALUE",
			idName:    "THEID",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							0: {
								"id": {
									S: aws.String("THEVALUE"),
								},
							},
							1: {
								"id": {
									S: aws.String("THEVALUE"),
								},
							},
						},
					}, nil
				},
			},
			errResponse: errors.New("An error occurred during processing."),
		},
		"error": {
			tableName: "testTable",
			idValue:   "THEVALUE",
			idName:    "THEID",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return nil, errors.New("Some error")
				},
			},
			errResponse: errors.New("Some error"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		var response ShopName
		err := GetById(test.idName, test.idValue, test.tableName, &response, test.queryer)
		if err != nil {
			assert.Equal(t, err, test.errResponse)
		} else {
			assert.Equal(t, response.Id, test.idValue)
		}

	}
}

func TestHandlerDynamoHelper_createGetByIDQuery(t *testing.T) {

	tests := map[string]struct {
		tableName string
		idName    string
		idValue   string
		response  *dynamodb.QueryInput
	}{
		"0": {
			tableName: "testTable",
			idName:    "someID",
			idValue:   "someValue",
			response: &dynamodb.QueryInput{
				TableName: aws.String("testTable"),
				KeyConditions: map[string]*dynamodb.Condition{
					"someID": {
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							{
								S: aws.String("someValue"),
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response := createGetByIDQuery(test.idName, test.idValue, test.tableName)
		assert.Equal(t, response, test.response)
	}
}

func TestHandlerDynamoHelper_createPutItemOutput(t *testing.T) {

	tests := map[string]struct {
		shopname  ShopName
		tableName string
	}{
		"1": {
			tableName: "testTable",
			shopname:  ShopName{Id: "1234", ShopName: "Test"},
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		av, err := dynamodbattribute.MarshalMap(test.shopname)
		if err != nil {
			t.Errorf("Failed to marshal item")
		}
		response, err := createPutItemInput(test.shopname, test.tableName)
		if err != nil {
			t.Errorf("Failed to marshal item")
		}
		assert.Equal(t, response.TableName, aws.String(test.tableName))
		assert.Equal(t, response.ReturnValues, aws.String("ALL_NEW"))
		assert.Equal(t, response.Item, av)
	}
}

func TestHandlerDynamoHelper_createShopFriendlyNamesByShopNameQuery(t *testing.T) {

	tests := map[string]struct {
		shopName  string
		response  *dynamodb.QueryInput
		tableName string
	}{
		"1": {
			tableName: "testTable",
			shopName:  "testName",
			response: &dynamodb.QueryInput{
				TableName: aws.String("testTable"),
				IndexName: aws.String("shop_name-index"),
				KeyConditions: map[string]*dynamodb.Condition{
					"shop_name": {
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							{
								S: aws.String("testName"),
							},
						},
					},
				},
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":deleted_v": { // Required
						S: aws.String("false"),
					},
				},
				FilterExpression: aws.String("deleted = :deleted_v"),
			},
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response := createShopFriendlyNamesByShopNameQuery(test.shopName, test.tableName)
		assert.Equal(t, response, test.response)
	}
}

func TestHandlerDynamoHelper_GetShopFriendlyNamesByShopName(t *testing.T) {

	tests := map[string]struct {
		shopName    string
		queryer     QueryerMock
		errResponse error
		size        int
	}{
		"success_0_records": {
			shopName: "testName",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{},
					}, nil
				},
			},
			errResponse: nil,
			size:        0,
		},
		"success_1_record": {
			shopName: "testName",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							0: {
								"id": {
									S: aws.String("0"),
								},
								"friendly_name": {
									S: aws.String("friendlyName"),
								},
								"shop_name": {
									S: aws.String("testName"),
								},
							},
						},
					}, nil
				},
			},
			errResponse: nil,
			size:        1,
		},
		"success_2_records": {
			shopName: "testName",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return &dynamodb.QueryOutput{
						Items: []map[string]*dynamodb.AttributeValue{
							0: {
								"id": {
									S: aws.String("0"),
								},
								"friendly_name": {
									S: aws.String("friendlyName"),
								},
								"shop_name": {
									S: aws.String("testName"),
								},
							},
							1: {
								"id": {
									S: aws.String("0"),
								},
								"friendly_name": {
									S: aws.String("friendlyName"),
								},
								"shop_name": {
									S: aws.String("testName"),
								},
							},
						},
					}, nil
				},
			},
			errResponse: nil,
			size:        2,
		},
		"error": {
			shopName: "",
			queryer: QueryerMock{
				QueryFunc: func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
					return nil, errors.New("Some error")
				},
			},
			errResponse: errors.New("Some error"),
			size:        0,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response, err := GetShopFriendlyNamesByShopName(test.shopName, "", test.queryer)
		assert.Equal(t, len(response), test.size)
		if len(response) > 0 {
			assert.Equal(t, response[0].ShopName, test.shopName)
		}
		assert.Equal(t, err, test.errResponse)
	}
}

func TestHandlerDynamoHelper_PutItem(t *testing.T) {

	tests := map[string]struct {
		shopname    ShopName
		tableName   string
		putitemer   PutItemerMock
		response    *dynamodb.PutItemOutput
		errResponse error
	}{
		"success": {
			tableName: "testTable",
			shopname:  ShopName{Id: "1234", ShopName: "Test"},
			putitemer: PutItemerMock{
				PutItemFunc: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
					return &dynamodb.PutItemOutput{
						Attributes: map[string]*dynamodb.AttributeValue{
							"Id": {
								S: aws.String("1234"),
							},
							"ShopName": {
								S: aws.String("Test"),
							},
						},
					}, nil
				},
			},
			response: &dynamodb.PutItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"Id": {
						S: aws.String("1234"),
					},
					"ShopName": {
						S: aws.String("Test"),
					},
				},
			},
			errResponse: nil,
		},
		"error": {
			tableName: "testTable",
			shopname:  ShopName{},
			putitemer: PutItemerMock{
				PutItemFunc: func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
					return nil, errors.New("Could not insert shop")
				},
			},
			response:    &dynamodb.PutItemOutput{},
			errResponse: errors.New("Could not insert shop"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		output, err := Put(test.shopname, test.tableName, test.putitemer)
		if err != nil {
			assert.Equal(t, err, test.errResponse)
		} else {
			assert.Equal(t, output, test.response)
		}
	}
}
