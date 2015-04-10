package main

import (
	"log"
	"sim3/bus"
	_ "sim3/ncl/extra"
	"sim3/ncl/tool"
	"sim3/portable"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func init() {
	tool.Src = portable.DataSource
}

var busChan chan *bus.Msg

//этот хэндлер только пишет сообщения в канал главной горутины
func busHandler(m *bus.Msg) {
	busChan <- m
}

func load() {
	t := &tool.Solder{}
	t.F("counter.yml")
}

//этот хэндлер обрабатывает сообщения в рамках главной горутины
func handle(m *bus.Msg) {
	switch m.Typ {
	case "init":
		load()
	}
}

func main() {
	log.Println("sim3 started")
	bus.Init(busHandler)
	busChan = make(chan *bus.Msg)
	wg.Add(1)
	go func(wg *sync.WaitGroup, c chan *bus.Msg) {
		bus.Process(&bus.Msg{Typ: "init"})
		for {
			select {
			case m := <-c:
				handle(m)
			}
		}
	}(wg, busChan)
	wg.Wait()
	log.Println("sim3 closed")
}
