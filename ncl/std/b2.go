package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/halt"
)

type any2 struct {
	typ string
	I   ncl.In
	O   ncl.Out
	fn  func(tri.Trit) tri.Trit
}

func (e *any2) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.I:
		return e.I
	case ncl.O:
		return e.O
	default:
		halt.As(100)
	}
	panic(0)
}

func (e *any2) init() {
	e.I = NewIn(e)
	e.O = NewOut(e)
}

func (e *any2) Do() {
	if sig := e.I.Select(); sig != nil {
		val := e.fn(*sig)
		e.O.Update(&val)
	} else {
		e.O.Update(nil)
	}
}

func Not() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.Not(p)
		return
	}
	e := &any2{typ: "¬", fn: f}
	e.init()
	return e
}

func Buffer() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = p
		return
	}
	e := &any2{typ: "BUF", fn: f}
	e.init()
	return e
}

func CycleRight() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.CNot(p)
		return
	}
	e := &any2{typ: "→", fn: f}
	e.init()
	return e
}

func CycleLeft() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.CNot(tri.CNot(p))
		return
	}
	e := &any2{typ: "←", fn: f}
	e.init()
	return e
}
