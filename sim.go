package main

import (
	"github.com/ivpusic/neo"
	"runtime"
	"sim3/ncl"
	"sim3/ncl/std"
	"time"
)

func main() {
	nw := func() {
		app := neo.App()

		app.Get("/", func(ctx *neo.Ctx) {
			ctx.Res.Text("I am Neo Programmer", 200)
		})

		app.Start()
	}
	go nw()
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	board := std.Board()
	not := std.NewNot()
	board.Point("+").Solder(std.NewProbe("+").Pin(ncl.I))
	board.Point("+").Solder(not.Pin(ncl.I))
	board.Point("not+").Solder(not.Pin(ncl.O))
	buf := std.NewBuffer()
	board.Point("not+").Solder(buf.Pin(ncl.I))
	board.Point("not+buf").Solder(buf.Pin(ncl.O), std.NewProbe("not+buf").Pin(ncl.I))
	time.Sleep(time.Duration(time.Second * 2))
}
