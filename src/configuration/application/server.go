package application

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"sync"
)

func newWebServer() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}

func NewApplicationServer(wg *sync.WaitGroup) {
	app := newWebServer()
	connectionsRepository := ConfigureConnectionsRepository()
	ConfigureMessageEndpoints(connectionsRepository, app)

	// Listen and serve on 0.0.0.0:8080
	app.Listen(":8080")
	wg.Done()
}
