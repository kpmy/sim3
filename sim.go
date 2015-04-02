package main

import (
	"github.com/ivpusic/neo"
	"github.com/ivpusic/neo-cors"
	"runtime"
	"sim3/api"
	"sim3/ncl"
	"sim3/ncl/std"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)

	nw := func() {
		app := neo.App()
		app.Use(cors.Init())
		app.Get("/tri.json", api.Tri)
		app.Start()
	}
	go nw()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	board := std.Board()
	not := std.NewNot()
	board.Point("+").Solder(std.NewProbe("+").Pin(ncl.I))
	board.Point("+").Solder(not.Pin(ncl.I))
	board.Point("not+").Solder(not.Pin(ncl.O))
	buf := std.NewBuffer()
	board.Point("not+").Solder(buf.Pin(ncl.I))
	board.Point("not+buf").Solder(buf.Pin(ncl.O), std.NewProbe("not+buf").Pin(ncl.I))
	wg.Wait()
}
