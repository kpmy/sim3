package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	routine := func(x int, g *sync.WaitGroup) {
		fmt.Println(x)
		g.Done()
	}
	wg := new(sync.WaitGroup)
	for i := 0; i < 150000; i++ {
		wg.Add(1)
		go routine(i, wg)
	}
	wg.Wait()
}
