package main

import (
	"runtime"
	"sim3/ncl"
	"sim3/ncl/std"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	probe := std.NewProbe()
	board := std.Board()
	board.Point("+").Solder(probe.Pin(ncl.I))
	time.Sleep(time.Second)
}
