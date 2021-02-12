package main

import (
	"github.com/mrflick72/redis-explorer/src/configuration/application"
	"sync"
)

func main() {
	initApplicationServer()
}

func initApplicationServer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go application.NewApplicationServer(wg)
	wg.Wait()
}
