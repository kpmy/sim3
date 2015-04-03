package main

import (
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
	"runtime"
	"sim3/api"
	"sim3/ncl"
	"sim3/ncl/std"
	"sim3/tri"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	wg.Add(1)
}

func loader() {
	board := std.Board()
	board.Point("~").Solder(std.Probe("~").Pin(ncl.I), std.Generator(tri.TRUE, tri.NIL, tri.FALSE, tri.NIL).Pin(ncl.O))
}

func main() {
	nw := func() {
		app := neo.App()
		app.Use(cors.Init())
		app.Get("/tri.json", api.Tri)
		app.Start()
	}
	go nw()
	go loader()
	wg.Wait()
}
