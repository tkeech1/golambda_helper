package golambda_helper

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (h *DynamoHandler) GetShopNameById(id string, tableName string) (ShopName, error) {
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
		return ShopName{}, err
	}

	recordList := []ShopName{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &recordList)
	if err != nil {
		return ShopName{}, err
	}

	if len(recordList) != 1 {
		return ShopName{}, err
	}

	return recordList[0], nil
}

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

func (h *DynamoHandler) PutShop(shop ShopName, tableName string) (ShopName, error) {
	returnVal := ShopName{}

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

/*func (h *DynamoHandler) PutItem(item interface{}, tableName string) error {

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
}*/

// call this like this
// https://stackoverflow.com/questions/40895901/golang-passing-in-a-type-variable-into-function
// var x = []Shop{}
// h.GetById("someid","shop-dev",&x)
func (h *DynamoHandler) GetById(id string, tableName string, v interface{}) error {
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
		return err
	}

	if len(result.Items) == 1 {
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &v)
		if err != nil {
			return err
		}
	}

	return errors.New("An error occurred during processing.")

}
