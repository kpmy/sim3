package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/halt"
)

type Not struct {
	I ncl.In
	O ncl.Out
}

type Buffer struct {
	I ncl.In
	O ncl.Out
}

func (n *Not) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.I:
		return n.I
	case ncl.O:
		return n.O
	default:
		halt.As(100)
	}
	panic(0)
}

func (b *Buffer) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.I:
		return b.I
	case ncl.O:
		return b.O
	default:
		halt.As(100)
	}
	panic(0)
}

func NewNot() ncl.Element {
	n := &Not{I: newIn(), O: newOut()}
	go func(n *Not) {
		ncl.Step(n, func() {
			ok, val := n.I.Select()
			if ok {
				n.O.Validate(true, notMap[val])
			} else {
				n.O.Validate(false)
			}
		})
	}(n)
	return n
}

func NewBuffer() ncl.Element {
	b := &Buffer{I: newIn(), O: newOut()}
	go func(b *Buffer) {
		ncl.Step(b, func() {
			ok, val := b.I.Select()
			b.O.Validate(ok, val)
		})
	}(b)
	return b
}

var notMap map[tri.Trit]tri.Trit

func init() {
	// таблицы истинности
	notMap = make(map[tri.Trit]tri.Trit)
	notMap[tri.TRUE] = tri.FALSE
	notMap[tri.FALSE] = tri.TRUE
	notMap[tri.NIL] = tri.NIL
}
