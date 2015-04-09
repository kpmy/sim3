package main

import (
	"log"
	//	"github.com/ivpusic/neo"
	//	"github.com/ivpusic/neo-cors"
	//"runtime"
	//	"sim3/api"
	"github.com/gopherjs/gopherjs/js"
	_ "sim3/ncl/extra"
	"sim3/ncl/tool"
	"sim3/portable"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func init() {
	tool.Src = portable.DataSource
	//runtime.GOMAXPROCS(1)
	wg.Add(1)
}

func load() {
	t := &tool.Solder{}
	t.F("counter.yml")
}

func main() {
	js.Global.Get("self")
	log.Println("sim3 started")

	/*	nw := func() {
			app := neo.App()
			app.Use(cors.Init())
			app.Get("/tri.json", api.Tri)
			app.Start()
		}
		go nw()
	*/
	load()
	wg.Wait()
	log.Println("sim3 closed")
}
