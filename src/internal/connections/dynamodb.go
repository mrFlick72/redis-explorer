package connections

import (
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
		return err
	}
	_, err = r.Client.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.TableName),
	})

	return err
}

func (r *DynamoDbRepository) GetConnectionFor(id ConnectionId) (*Connection, error) {
	item, err := r.Client.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ConnectionId": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		return nil, err
	}
	itemMap := item.Item
	return newConnectionFor(itemMap), nil
}

func (r *DynamoDbRepository) GetConnections() (*[]Connection, error) {
	connections := make([]Connection, 0)

	items, err := r.Client.Scan(&dynamodb.ScanInput{
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		return nil, err
	}

	for _, item := range items.Items {
		connections = append(connections, *newConnectionFor(item))
	}

	return &connections, nil
}

func newConnectionFor(itemMap map[string]*dynamodb.AttributeValue) *Connection {
	return &Connection{
		Id:          *itemMap["ConnectionId"].S,
		Name:        *itemMap["ConnectionName"].S,
		HostAndPort: itemStringValueFor(itemMap["HostAndPort"]),
		Username:    itemStringValueFor(itemMap["Username"]),
		Password:    itemStringValueFor(itemMap["Password"]),
	}
}

func itemStringValueFor(itemMap *dynamodb.AttributeValue) string {
	if itemMap == nil {
		return ""
	} else {
		return *itemMap.S
	}
}
