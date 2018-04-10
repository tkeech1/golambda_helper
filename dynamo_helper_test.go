package golambda_helper_test

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"

	"github.com/tkeech1/golambda_helper"
	"github.com/tkeech1/golambda_helper/mocks"
)

func TestHandlerDynamoHelper_GetShopNameById(t *testing.T) {

	tests := map[string]struct {
		requestId     string
		queryResponse []map[string]*dynamodb.AttributeValue
		tableName     string
		err           error
	}{
		"success_0_records": {
			tableName:     "testTable",
			requestId:     "",
			queryResponse: []map[string]*dynamodb.AttributeValue{},
			err:           nil,
		},
		"success_1_record": {
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
		"success_2_records": {
			tableName: "testTable",
			requestId: "",
			queryResponse: []map[string]*dynamodb.AttributeValue{
				0: {
					"ids": {
						S: aws.String(""),
					},
				},
				1: {
					"ids": {
						S: aws.String(""),
					},
				},
			},
			err: nil,
		},
		"error": {
			tableName:     "testTable",
			requestId:     "",
			queryResponse: []map[string]*dynamodb.AttributeValue{},
			err:           errors.New("Some error"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		mockDynamoInterface := &mocks.DynamoInterface{}
		mockDynamoInterface.
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

		h := &golambda_helper.DynamoHandler{
			Svc: mockDynamoInterface,
		}

		response, err := h.GetShopNameById(test.requestId, test.tableName)
		assert.Equal(t, response.Id, test.requestId)
		assert.Equal(t, err, test.err)
		mockDynamoInterface.AssertExpectations(t)
	}
}

func TestHandlerDynamoHelper_GetShopFriendlyNamesByShopName(t *testing.T) {

	tests := map[string]struct {
		shopName      string
		queryResponse []map[string]*dynamodb.AttributeValue
		tableName     string
		err           error
		size          int
	}{
		"success_0_records": {
			tableName:     "testTable",
			shopName:      "testName",
			queryResponse: []map[string]*dynamodb.AttributeValue{},
			err:           nil,
			size:          0,
		},
		"success_1_record": {
			tableName: "testTable",
			shopName:  "testName",
			queryResponse: []map[string]*dynamodb.AttributeValue{
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
			err:  nil,
			size: 1,
		},
		"success_2_records": {
			tableName: "testTable",
			shopName:  "testName",
			queryResponse: []map[string]*dynamodb.AttributeValue{
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
			err:  nil,
			size: 2,
		},
		"error": {
			tableName:     "testTable",
			shopName:      "",
			queryResponse: []map[string]*dynamodb.AttributeValue{},
			err:           errors.New("Some error"),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		mockDynamoInterface := &mocks.DynamoInterface{}
		mockDynamoInterface.
			On("Query", &dynamodb.QueryInput{
				TableName: aws.String(test.tableName),
				IndexName: aws.String("shop_name-index"),
				KeyConditions: map[string]*dynamodb.Condition{
					"shop_name": {
						ComparisonOperator: aws.String("EQ"),
						AttributeValueList: []*dynamodb.AttributeValue{
							{
								S: aws.String(test.shopName),
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
			}).
			Return(&dynamodb.QueryOutput{
				Items: test.queryResponse,
			}, test.err).
			Once()

		h := &golambda_helper.DynamoHandler{
			Svc: mockDynamoInterface,
		}

		response, err := h.GetShopFriendlyNamesByShopName(test.shopName, test.tableName)
		assert.Equal(t, len(response), test.size)
		if len(response) > 0 {
			assert.Equal(t, response[0].ShopName, test.shopName)
		}
		assert.Equal(t, err, test.err)
		mockDynamoInterface.AssertExpectations(t)
	}
}

func TestHandlerDynamoHelper_PutShop(t *testing.T) {

	tests := map[string]struct {
		shopname      golambda_helper.ShopName
		queryResponse golambda_helper.ShopName
		tableName     string
		err           error
	}{
		"success": {
			tableName:     "testTable",
			shopname:      golambda_helper.ShopName{Id: "1234", ShopName: "Test"},
			queryResponse: golambda_helper.ShopName{Id: "1234", ShopName: "Test"},
			err:           nil,
		},
		"error": {
			tableName:     "testTable",
			shopname:      golambda_helper.ShopName{},
			queryResponse: golambda_helper.ShopName{},
			err:           errors.New("Could not insert shop "),
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		mockDynamoInterface := &mocks.DynamoInterface{}

		av, _ := dynamodbattribute.MarshalMap(test.shopname)
		mockDynamoInterface.
			On("PutItem", &dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String(test.tableName),
			}).
			Return(&dynamodb.PutItemOutput{}, test.err).
			Once()

		h := &golambda_helper.DynamoHandler{
			Svc: mockDynamoInterface,
		}

		response, err := h.PutShop(test.shopname, test.tableName)
		assert.Equal(t, response, test.queryResponse)
		assert.Equal(t, err, test.err)
		mockDynamoInterface.AssertExpectations(t)
	}
}
