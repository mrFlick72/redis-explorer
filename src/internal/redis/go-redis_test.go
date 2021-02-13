package redis

import (
	"context"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGoRedisRepository_ConnectTo(t *testing.T) {
	cache := cache.New(5*time.Minute, 10*time.Minute)

	delegate := new(MockedDelegateObject)

	repository := GoRedisRepository{
		connectionsRepository: delegate,
		storage:               cache,
	}
	expected := aConnection()
	delegate.On("GetConnectionFor", "Test").Return(expected, nil)

	err := repository.ConnectTo("Test")

	connection, found := repository.storage.Get("Test")

	assert.Nil(t, err)
	assert.True(t, found)
	assert.NotNil(t, connection)
}

func TestGoRedisRepository_ConnectToShouldNotReplaceConnection(t *testing.T) {
	storage := cache.New(5*time.Minute, 10*time.Minute)

	delegate := new(MockedDelegateObject)
	connection := aGoRedisConnection()
	storage.Set("Test", &connection, cache.NoExpiration)

	repository := GoRedisRepository{
		connectionsRepository: delegate,
		storage:               storage,
	}

	err := repository.ConnectTo("Test")

	_, found := repository.storage.Get("Test")

	assert.Nil(t, err)
	assert.True(t, found)
	delegate.AssertNotCalled(t, "GetConnectionFor", "Test")
}

func TestGoRedisRepository_Save(t *testing.T) {
	storage := cache.New(5*time.Minute, 10*time.Minute)

	repository := GoRedisRepository{
		connectionsRepository: nil,
		storage:               storage,
	}
	ctx := context.TODO()
	client := go_redis.NewClient(&go_redis.Options{
		Addr:     "localhost:6379",
		Username: "", // no password set
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	connection := GoRedisConnection{
		client: client,
		cxt:    &ctx,
	}
	fmt.Printf("&redisConnection: %v\n", &connection)

	repository.storage.Set("Test", &connection, cache.NoExpiration)

	_, err := repository.Save("Test", &Object{
		content: map[string]string{"Value": "Value"},
		Id:      "key",
		Ttl:     time.Minute * 5,
	})

	assert.Nil(t, err)
}

type MockedDelegateObject struct {
	mock.Mock
}

func (mock *MockedDelegateObject) StoreConnection(connection *connections.Connection) error {
	args := mock.Called(connection)
	return args.Error(0)
}

func (mock *MockedDelegateObject) GetConnectionFor(connectionId connections.ConnectionId) (*connections.Connection, error) {
	args := mock.Called(connectionId)
	return args.Get(0).(*connections.Connection), args.Error(1)
}

func (mock *MockedDelegateObject) GetConnections() (*[]connections.Connection, error) {
	args := mock.Called()
	return args.Get(0).(*[]connections.Connection), args.Error(1)
}

func aConnection() *connections.Connection {
	return &connections.Connection{
		Id:          "Test",
		Name:        "Test",
		HostAndPort: "",
		Username:    "",
		Password:    "",
	}
}

func aGoRedisConnection() GoRedisConnection {
	ctx := context.TODO()
	client := go_redis.NewClient(&go_redis.Options{
		Addr:     "localhost:6379",
		Username: "", // no password set
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	connection := GoRedisConnection{
		client: client,
		cxt:    &ctx,
	}
	return connection
}
