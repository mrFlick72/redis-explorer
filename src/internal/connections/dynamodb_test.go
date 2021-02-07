package connections

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreAConnection(t *testing.T) {
	repository := repository()

	err := repository.StoreConnection(&Connection{
		Name:        "Test",
		HostAndPort: "",
		Username:    "",
		Password:    "",
	})

	assert.Nil(t, err)
}

func TestGetConnectionForConnectionName(t *testing.T) {
	repository := repository()
	connection, err := repository.GetConnectionFor("Test")

	assert.Equal(t, connection, &Connection{
		Name:        "Test",
		HostAndPort: "",
		Username:    "",
		Password:    "",
	})
	assert.Nil(t, err)
}

func TestGetConnections(t *testing.T) {
	repository := repository()
	connections, err := repository.GetConnections()
	connection := Connection{
		Name:        "Test",
		HostAndPort: "",
		Username:    "",
		Password:    "",
	}
	expected := make([]Connection, 1)
	expected = append(expected, connection)

	assert.Equal(t, connections, &expected)
	assert.Nil(t, err)
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
