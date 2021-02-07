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
		Name:     "Test",
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
	})

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
