package main

import (
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
	"runtime"
	"sim3/api"
	_ "sim3/ncl/extra"
	"sim3/ncl/tool"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func init() {
	runtime.GOMAXPROCS(1)
	wg.Add(1)
}

func load() {
	t := &tool.Solder{}
	t.F("counter.yml")
}

func main() {
	nw := func() {
		app := neo.App()
		app.Use(cors.Init())
		app.Get("/tri.json", api.Tri)
		app.Start()
	}
	go nw()
	load()
	wg.Wait()
}
