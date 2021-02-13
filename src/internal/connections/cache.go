package connections

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type InMemoryCachedRepository struct {
	Delegate     Repository
	CacheManager *cache.Cache
	Ttl          string
}

func (r *InMemoryCachedRepository) StoreConnection(connection *Connection) error {
	err := r.Delegate.StoreConnection(connection)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(r.Ttl)
	r.CacheManager.Set(connection.Name, *connection, duration)

	return err
}

func (r *InMemoryCachedRepository) GetConnectionFor(id ConnectionId) (*Connection, error) {
	element, found := r.CacheManager.Get(id)
	if found {
		connection := element.(Connection)
		return &connection, nil
	} else {
		connection, err := r.Delegate.GetConnectionFor(id)
		if err != nil {
			return nil, err
		}

		duration, err := time.ParseDuration(r.Ttl)
		r.CacheManager.Set(connection.Name, *connection, duration)

		return connection, err
	}
}

func (r *InMemoryCachedRepository) GetConnections() (*[]Connection, error) {
	return r.Delegate.GetConnections()
}
