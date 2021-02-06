package redis

import (
	"context"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
)

type GoRedisRepository struct {
}

func (repository *GoRedisRepository) GetDatabases() (*[]Database, error) {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushAllDatabases() error {
	panic("tobe defined")
}

func (repository *GoRedisRepository) FlushDatabaseFor(id DatabaseId) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) Save(object Object) error {
	panic("tobe defined")
}
func (repository *GoRedisRepository) GetObjectsFor(id DatabaseId, page int, pageSize int) (*[]Object, error) {
	panic("tobe defined")
}
func (repository *GoRedisRepository) DeleteObjectFor(id ObjetsId) (*Object, error) {
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
