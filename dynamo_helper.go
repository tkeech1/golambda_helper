package golambdahelper

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func createShopFriendlyNamesByShopNameQuery(shopName string, tableName string) *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
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
			":deleted_v": {
				S: aws.String("false"),
			},
		},
		FilterExpression: aws.String("deleted = :deleted_v"),
	}
}

// GetShopFriendlyNamesByShopName gets shops by friendly name
func GetShopFriendlyNamesByShopName(shopName string, tableName string, queryer DynamoQueryer) ([]ShopName, error) {
	queryInput := createShopFriendlyNamesByShopNameQuery(shopName, tableName)

	result, err := queryer.Query(queryInput)
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

func createPutItemInput(item interface{}, tableName string) (*dynamodb.PutItemInput, error) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	return &dynamodb.PutItemInput{
		Item:         av,
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_NEW"),
	}, nil
}

// Put puts items to a Dynamo DB
func Put(item interface{}, tableName string, puter DynamoPutItemer) (*dynamodb.PutItemOutput, error) {

	input, err := createPutItemInput(item, tableName)
	if err != nil {
		return nil, err
	}

	return puter.PutItem(input)
}

func createGetByIDQuery(idName, idValue, tableName string) *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
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
	}
}

// GetByID gets items from a DynamoDB by ID
func GetByID(idName, idValue, tableName string, v interface{}, queryer DynamoQueryer) error {
	query := createGetByIDQuery(idName, idValue, tableName)
	result, err := queryer.Query(query)
	if err != nil {
		return err
	}

	if len(result.Items) == 1 {
		return dynamodbattribute.UnmarshalMap(result.Items[0], &v)
	}

	return errors.New("an error occurred during processing")

}

// Query queries against a DynamoDB
func (d *Dynamo) Query(queryInput *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("ENV_AWS_REGION")),
	})
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(sess)
	return svc.Query(queryInput)
}

// PutItem puts an item into a DynamoDB
func (d *Dynamo) PutItem(putItem *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("ENV_AWS_REGION")),
	})
	if err != nil {
		return nil, err
	}
	svc := dynamodb.New(sess)
	return svc.PutItem(putItem)
}

// Dynamo is used a receiever for DynamoDB methods
type Dynamo struct {
}
