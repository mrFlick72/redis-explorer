package connections

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDbRepository struct {
	TableName string
	Client    *dynamodb.DynamoDB
}

func (r *DynamoDbRepository) StoreConnection(connection *Connection) error {
	item, err := dynamodbattribute.MarshalMap(*connection)
	if err != nil {
		fmt.Printf("err %v", err)
		return err
	}
	fmt.Printf("item %v", item)

	response, err := r.Client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &r.TableName,
	})
	fmt.Printf("response %v", response)

	return err
}
