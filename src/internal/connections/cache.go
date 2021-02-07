package connections

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type InMemoryCachedRepository struct {
	delegate     Repository //todo to improve
	cacheManager *cache.Cache
	ttl          string
}

func (r *InMemoryCachedRepository) StoreConnection(connection *Connection) error {
	err := r.delegate.StoreConnection(connection)
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(r.ttl)
	r.cacheManager.Set(connection.Name, *connection, duration)

	return err
}

func (r *InMemoryCachedRepository) GetConnectionFor(name ConnectionName) (*Connection, error) {
	panic("TODO")
}

func (r *InMemoryCachedRepository) GetConnections() (*[]Connection, error) {
	panic("TODO")
}
