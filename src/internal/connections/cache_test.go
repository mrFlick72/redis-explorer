package connections

import (
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestStoreAConnectionOnInMemoryCachedRepository(t *testing.T) {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	delegate := new(MockedDelegateObject)
	repo := InMemoryCachedRepository{delegate: delegate,
		cacheManager: cache,
		ttl:          "5m",
	}
	connection := aConnection()
	delegate.On("StoreConnection", connection).Return(nil)

	repo.StoreConnection(connection)

	actual, found := cache.Get(connection.Name)
	assert.True(t, found)
	assert.NotNil(t, actual)
}

func TestGetConnectionForConnectionNameOnInMemoryCachedRepository(t *testing.T) {
	cache := cache.New(5*time.Minute, 10*time.Minute)
	delegate := new(MockedDelegateObject)
	repo := InMemoryCachedRepository{delegate: delegate,
		cacheManager: cache,
		ttl:          "5m",
	}

	connection := aConnection()
	delegate.On("GetConnectionFor", "Test").Return(connection, nil)

	actual, err := repo.GetConnectionFor("Test")

	assert.NotNil(t, actual)
	assert.Nil(t, err)
}

func TestGetConnectionsOnInMemoryCachedRepository(t *testing.T) {

}

type MockedDelegateObject struct {
	mock.Mock
}

func (mock *MockedDelegateObject) StoreConnection(connection *Connection) error {
	args := mock.Called(connection)
	return args.Error(0)
}

func (mock *MockedDelegateObject) GetConnectionFor(name ConnectionName) (*Connection, error) {
	args := mock.Called(name)
	return args.Get(0).(*Connection), args.Error(1)
}

func (mock *MockedDelegateObject) GetConnections() (*[]Connection, error) {
	args := mock.Called()
	return args.Get(0).(*[]Connection), args.Error(1)
}
