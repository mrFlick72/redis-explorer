package redis

import (
	"context"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
)

type GoRedisRepository struct {
	connectionsRepository *connections.Repository
}

func (repository *GoRedisRepository) ConnectTo(connectionName connections.ConnectionName) error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) DisconnectFrom(connectionName connections.ConnectionName) error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) GetDatabases(connectionName connections.ConnectionName) (*[]Database, error) {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushAllDatabases(connectionName connections.ConnectionName) error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushDatabaseFor(connectionName connections.ConnectionName, id DatabaseId) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) Save(connectionName connections.ConnectionName, object Object) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) GetObjectsFor(connectionName connections.ConnectionName, id DatabaseId, page int, pageSize int) (*[]Object, error) {
	panic("tobe defined")
}
func (repository *GoRedisRepository) DeleteObjectFor(connectionName connections.ConnectionName, id ObjetsId) (*Object, error) {
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
