package api

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
)

type ConnectionEndpoints struct {
	Repo *connections.Repository
}

func (endpoint *ConnectionEndpoints) RegisterEndpoint(application *iris.Application) {
	application.Get("/connections", endpoint.getConnectionsEndpoint)
	application.Get("/connections/{id}", endpoint.getConnectionForEndpoint)
	application.Put("/connections", endpoint.storeConnectionForEndpoint)
}

func (endpoint *ConnectionEndpoints) storeConnectionForEndpoint(ctx iris.Context) {
	body, err := ctx.GetBody()
	if err == nil {
		connection := connections.Connection{}
		json.Unmarshal(body, &connection)
		endpoint.Repo.Operations.StoreConnection(&connection)
		ctx.StatusCode(iris.StatusNoContent)
	} else {
		ctx.StatusCode(iris.StatusInternalServerError)
	}
}

func (endpoint *ConnectionEndpoints) getConnectionsEndpoint(ctx iris.Context) {
	find, _ := (*endpoint.Repo).Operations.GetConnections()
	ctx.JSON(find)
	ctx.StatusCode(iris.StatusOK)
}

func (endpoint *ConnectionEndpoints) getConnectionForEndpoint(ctx iris.Context) {
	id := urlParam(ctx, "id", "")
	find, _ := (*endpoint.Repo).Operations.GetConnectionFor(id)
	ctx.JSON(find)
	ctx.StatusCode(iris.StatusOK)
}

func urlParam(ctx iris.Context, paramName string, defaultValue string) string {
	lang := ctx.URLParam(paramName)
	if &lang == nil {
		return defaultValue
	}
	return lang
}
