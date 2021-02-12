package redis

import (
	"github.com/mrflick72/redis-explorer/src/internal/connections"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGoRedisRepository_ConnectTo(t *testing.T) {
	delegate := new(MockedDelegateObject)
	delegateWrapper := connections.Repository{Repo: delegate}

	repository := GoRedisRepository{
		connectionsRepository: &delegateWrapper,
	}
	repository.ConnectTo("TEST")
	assert.True(t, false)
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

func (mock *MockedDelegateObject) GetConnectionFor(name connections.ConnectionName) (*connections.Connection, error) {
	args := mock.Called(name)
	return args.Get(0).(*connections.Connection), args.Error(1)
}

func (mock *MockedDelegateObject) GetConnections() (*[]connections.Connection, error) {
	args := mock.Called()
	return args.Get(0).(*[]connections.Connection), args.Error(1)
}
