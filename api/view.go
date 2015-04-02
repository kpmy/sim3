package api

import (
	"github.com/ivpusic/neo"
	"sim3/tri"
	"time"
)

type Item struct {
	Timestamp int64
	Name      string
	Type      string
	Meta      bool
	Signal    tri.Trit
}

var LogChannel chan *Item = make(chan *Item, 1024)

func Log(i *Item) {
	i.Timestamp = time.Now().Unix()
	LogChannel <- i
}

func Tri(ctx *neo.Ctx) {
	var il []*Item
	for ok := true; ok; {
		select {
		case i := <-LogChannel:
			il = append(il, i)
		default:
			ok = false
		}
	}
	if il != nil {
		ctx.Res.Json(il, 200)
	} else {
		ctx.Res.Text("[]", 200)
	}
}
