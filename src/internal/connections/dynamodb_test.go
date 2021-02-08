package connections

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreAConnectionOnDynamoDB(t *testing.T) {
	repository := repository()

	err := repository.StoreConnection(aConnection())

	assert.Nil(t, err)
}

func TestGetConnectionForConnectionNameOnDynamoDB(t *testing.T) {
	repository := repository()
	connection, err := repository.GetConnectionFor("Test")

	assert.Equal(t, connection, aConnection())
	assert.Nil(t, err)
}

func TestGetConnectionsOnDynamoDB(t *testing.T) {
	repository := repository()
	connections, err := repository.GetConnections()
	connection := aConnection()
	expected := make([]Connection, 1)
	expected = append(expected, *connection)

	assert.Equal(t, &expected, connections)
	assert.Nil(t, err)
}

func aConnection() *Connection {
	return &Connection{
		Name:        "Test",
		HostAndPort: "",
		Username:    "",
		Password:    "",
	}
}

func repository() DynamoDbRepository {
	return DynamoDbRepository{
		Client:    client(),
		TableName: "RedisExplorerConnections",
	}
}

func client() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}
