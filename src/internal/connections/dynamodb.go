package connections

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
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
		TableName: aws.String(r.TableName),
	})
	fmt.Printf("response %v", response)

	return err
}

func (r *DynamoDbRepository) GetConnectionFor(name ConnectionName) (*Connection, error) {
	item, err := r.Client.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ConnectionName": {
				S: aws.String(name),
			},
		},
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		fmt.Sprintf("err: %v", err)
		return nil, err
	}
	itemMap := item.Item
	return &Connection{
		Name:        *itemMap["ConnectionName"].S,
		HostAndPort: itemStringValueFor(itemMap["HostAndPort"]),
		Username:    itemStringValueFor(itemMap["Username"]),
		Password:    itemStringValueFor(itemMap["Password"]),
	}, nil
}
func (r *DynamoDbRepository) GetConnections() (*[]Connection, error) { panic("TODO") }

func itemStringValueFor(itemMap *dynamodb.AttributeValue) string {
	if itemMap == nil {
		return ""
	} else {
		return *itemMap.S
	}
}
