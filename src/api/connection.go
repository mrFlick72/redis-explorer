package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mrflick72/redis-explorer/src/internal/connections"
)

type ConnectionEndpoints struct {
	Repo *connections.Repository
}

func (endpoint *ConnectionEndpoints) RegisterEndpoint(application *iris.Application) {
	application.Get("/connections", endpoint.getConnectionsEndpoint)
	application.Get("/connections/{id}", endpoint.getConnectionForEndpoint)
}

func (endpoint *ConnectionEndpoints) getConnectionsEndpoint(ctx iris.Context) {
	find, _ := (*endpoint.Repo).Repo.GetConnections()
	ctx.JSON(find)
	ctx.StatusCode(iris.StatusOK)
}

func (endpoint *ConnectionEndpoints) getConnectionForEndpoint(ctx iris.Context) {
	id := urlParam(ctx, "id", "")
	find, _ := (*endpoint.Repo).Repo.GetConnectionFor(id)
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
