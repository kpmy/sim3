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

func doAny2(e *any2) {
	ncl.Step(e, func() {
		ok, val := e.I.Select()
		if ok {
			e.O.Validate(true, e.fn(val))
		} else {
			e.O.Validate(false)
		}
	})
}

func Not() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.Not(p)
		return
	}
	e := &any2{typ: "¬", I: newIn(), O: newOut(), fn: f}
	go doAny2(e)
	return e
}

func Buffer() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = p
		return
	}
	e := &any2{typ: "BUF", I: newIn(), O: newOut(), fn: f}
	go doAny2(e)
	return e
}

func CycleRight() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.CNot(p)
		return
	}
	e := &any2{typ: "→", I: newIn(), O: newOut(), fn: f}
	go doAny2(e)
	return e
}

func CycleLeft() ncl.Element {
	f := func(p tri.Trit) (q tri.Trit) {
		q = tri.CNot(tri.CNot(p))
		return
	}
	e := &any2{typ: "←", I: newIn(), O: newOut(), fn: f}
	go doAny2(e)
	return e
}
