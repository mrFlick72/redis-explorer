package redis

import (
	"context"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
	"github.com/patrickmn/go-cache"
)

type GoRedisRepository struct {
	storage               *cache.Cache
	connectionsRepository *connections.Repository
}

func (repository *GoRedisRepository) ConnectTo(connectionId connections.ConnectionId) error {
	connectionFor, err := repository.connectionsRepository.Operations.GetConnectionFor(connectionId)
	if err == nil {
		rdb := go_redis.NewClient(&go_redis.Options{
			Addr:     connectionFor.HostAndPort,
			Username: connectionFor.Username, // no password set
			Password: connectionFor.Password, // no password set
			DB:       0,                      // use default DB
		})
		repository.storage.Set(connectionId, rdb, cache.NoExpiration)
	}
	return err
}

func (repository *GoRedisRepository) DisconnectFrom(connectionId connections.ConnectionId) error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) GetDatabases(connectionId connections.ConnectionId) (*[]Database, error) {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushAllDatabases(connectionId connections.ConnectionId) error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushDatabaseFor(connectionId connections.ConnectionId, id DatabaseId) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) Save(connectionId connections.ConnectionId, object Object) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) GetObjectsFor(connectionId connections.ConnectionId, id DatabaseId, page int, pageSize int) (*[]Object, error) {
	panic("tobe defined")
}
func (repository *GoRedisRepository) DeleteObjectFor(connectionId connections.ConnectionId, id ObjetsId) (*Object, error) {
	panic("tobe defined")
}

var ctx = context.Background()

func ExampleClient() {
	rdb := go_redis.NewClient(&go_redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == go_redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
