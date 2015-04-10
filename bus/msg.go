package bus

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/mitchellh/mapstructure"
	"ypk/assert"
)

type Msg struct {
	Typ string
}

type Handler func(m *Msg)

func Process(m *Msg) {
	assert.For(m != nil, 20)
	js.Global.Call("postMessage", m)
}

func Init(handler Handler) {
	js.Global.Set("onmessage", func(oEvent *js.Object) {
		data := oEvent.Get("data").Interface()
		m := &Msg{}
		err := mapstructure.Decode(data, m)
		assert.For(err == nil, 40)
		handler(m)
	})
}
