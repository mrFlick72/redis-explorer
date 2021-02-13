package redis

import (
	"context"
	json "encoding/json"
	"errors"
	"fmt"
	go_redis "github.com/go-redis/redis/v8"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
	"github.com/patrickmn/go-cache"
)

type GoRedisConnection struct {
	client *go_redis.Client
	cxt    *context.Context
}

type GoRedisRepository struct {
	storage               *cache.Cache
	connectionsRepository *connections.Repository
}

func (repository *GoRedisRepository) ConnectTo(connectionId connections.ConnectionId) error {
	connectionFor, err := repository.connectionsRepository.Operations.GetConnectionFor(connectionId)
	if err == nil {
		connectionContext := context.Background()

		rdb := go_redis.NewClient(&go_redis.Options{
			Addr:     connectionFor.HostAndPort,
			Username: connectionFor.Username, // no password set
			Password: connectionFor.Password, // no password set
			DB:       0,                      // use default DB
		})
		repository.storage.Set(connectionId, GoRedisConnection{
			client: rdb,
			cxt:    &connectionContext,
		}, cache.NoExpiration)
	}
	return err
}

func (repository *GoRedisRepository) Save(connectionId connections.ConnectionId, object *Object) (*ObjectId, error) {
	storedConnection, found := repository.storage.Get(connectionId)
	redisConnection := storedConnection.(GoRedisConnection)
	if found {
		fmt.Printf("&redisConnection: %v\n", &redisConnection)
		content, _ := json.Marshal(object.content)
		client, context := redisClientFor(&redisConnection)
		ttl := object.Ttl
		set := client.Set(*context, object.Id, content, ttl)
		fmt.Printf("&set: %v\n", set)

		return &object.Id, nil
	} else {
		return nil, errors.New(fmt.Sprintf("connection %v not found", connectionId))
	}
}

func redisClientFor(connection *GoRedisConnection) (*go_redis.Client, *context.Context) {
	var redisClient = connection.client
	var context = connection.cxt
	return redisClient, context
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
