package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/halt"
)

type not struct {
	I ncl.In
	O ncl.Out
}

type buffer struct {
	I ncl.In
	O ncl.Out
}

func (n *not) Pin(c ncl.PinCode) ncl.Pin {
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

func (b *buffer) Pin(c ncl.PinCode) ncl.Pin {
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

func Not() ncl.Element {
	n := &not{I: newIn(), O: newOut()}
	go func(n *not) {
		ncl.Step(n, func() {
			ok, val := n.I.Select()
			if ok {
				n.O.Validate(true, tri.Not(val))
			} else {
				n.O.Validate(false)
			}
		})
	}(n)
	return n
}

func Buffer() ncl.Element {
	b := &buffer{I: newIn(), O: newOut()}
	go func(b *buffer) {
		ncl.Step(b, func() {
			ok, val := b.I.Select()
			b.O.Validate(ok, val)
		})
	}(b)
	return b
}
