package redis

import (
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
	delegateWrapper := connections.Repository{Operations: delegate}

	repository := GoRedisRepository{
		connectionsRepository: &delegateWrapper,
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

func TestGoRedisRepository_DeleteObjectFor(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_DisconnectFrom(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_FlushAllDatabases(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_FlushDatabaseFor(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_GetDatabases(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_GetObjectsFor(t *testing.T) {
	panic("TODO")
}

func TestGoRedisRepository_Save(t *testing.T) {
	panic("TODO")
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
