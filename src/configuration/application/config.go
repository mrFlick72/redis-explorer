package application

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kataras/iris/v12"
	"github.com/mrflick72/redis-explorer/src/api"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
	"github.com/patrickmn/go-cache"
	"time"
)

func dynamoDbClient() *dynamodb.DynamoDB {
	return dynamodb.New(session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})))
}

func ConfigureConnectionsRepository() *connections.Repository {
	cache := cache.New(5*time.Minute, 10*time.Minute)

	dynamoDbRepository := connections.Repository{
		Repo: &connections.DynamoDbRepository{
			TableName: "RedisExplorerConnections",
			Client:    dynamoDbClient(),
		}}

	cachedRepository := connections.InMemoryCachedRepository{
		Delegate:     &dynamoDbRepository,
		CacheManager: cache,
		Ttl:          "1m",
	}
	return &connections.Repository{
		Repo: &cachedRepository,
	}
}

func ConfigureMessageEndpoints(repository *connections.Repository, app *iris.Application) {
	endpoints := api.ConnectionEndpoints{Repo: repository}
	endpoints.RegisterEndpoint(app)
}
