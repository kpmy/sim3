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

func load() {
	board := std.Board()
	probe := std.Probe("NOT(~)")
	not := std.Not()
	board.Point("~").Solder(std.Probe("~").Pin(ncl.I), std.Generator(tri.TRUE, tri.NIL, tri.FALSE, tri.NIL).Pin(ncl.O), not.Pin(ncl.I))
	board.Point("!~").Solder(not.Pin(ncl.O), probe.Pin(ncl.I))
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
