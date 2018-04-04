package golambda_helper

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (h *Handler) GetShopNameById(id string, tableName string) (Shop, error) {
	shop := Shop{}

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
		return shop, err
	}

	recordList := []Shop{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recordList)
	if err != nil {
		return shop, err
	}

	if len(recordList) != 1 {
		return shop, err
	}

	return recordList[0], nil
}

func (h *Handler) GetShopFriendlyNamesByShopName(shopName string, tableName string) ([]Shop, error) {
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
		return []Shop{}, err
	}

	shopList := []Shop{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &shopList)
	if err != nil {
		return []Shop{}, err
	}

	return shopList, err
}

func (h *Handler) PutShop(shop Shop, tableName string) (Shop, error) {
	returnVal := Shop{}

	av, err := dynamodbattribute.MarshalMap(shop)
	if err != nil {
		return returnVal, errors.New("Could not marshal shop " + shop.ShopName)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = h.Svc.PutItem(input)

	if err != nil {
		return returnVal, errors.New("Could not insert shop " + shop.ShopName)
	}

	return shop, nil
}
