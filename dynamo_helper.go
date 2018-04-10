package golambda_helper

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (h *DynamoHandler) GetShopFriendlyNamesByShopName(shopName string, tableName string) ([]ShopName, error) {
	queryInput := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		IndexName: aws.String("shop_name-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"shop_name": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(shopName),
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
	}

	result, err := h.Svc.Query(queryInput)
	if err != nil {
		return []ShopName{}, err
	}

	shopList := []ShopName{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &shopList)
	if err != nil {
		return []ShopName{}, err
	}

	return shopList, err
}

func (h *DynamoHandler) Put(item interface{}, tableName string) error {

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = h.Svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

func (h *DynamoHandler) GetById(idName string, idValue, tableName string, v interface{}) error {
	result, err := h.Svc.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			idName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(idValue),
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	if len(result.Items) == 1 {
		return dynamodbattribute.UnmarshalMap(result.Items[0], &v)
	}

	return errors.New("An error occurred during processing.")

}
