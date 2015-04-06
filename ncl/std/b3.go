package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/halt"
)

type any3 struct {
	typ string
	A   ncl.In
	B   ncl.In
	O   ncl.Out
	fn  func(tri.Trit, tri.Trit) tri.Trit
}

func (e *any3) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.A:
		return e.A
	case ncl.O:
		return e.O
	case ncl.B:
		return e.B
	default:
		halt.As(100)
	}
	panic(0)
}

func doAny3(e *any3) {
	ncl.Step(e, func() {
		ok, a := e.A.Select()
		da, b := e.B.Select()
		if ok && da {
			e.O.Validate(true, e.fn(a, b))
		} else {
			e.O.Validate(false)
		}
	})
}

func AndNot() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Not(tri.And(a, b))
	}
	e := &any3{typ: "¬&", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func OrNot() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Not(tri.Or(a, b))
	}
	e := &any3{typ: "¬|", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Cmp() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		if a == b {
			return tri.NIL
		} else if tri.Ord(a) > tri.Ord(b) {
			return tri.TRUE
		} else if tri.Ord(a) < tri.Ord(b) {
			return tri.FALSE
		}
		panic(0)
	}
	e := &any3{typ: "<=>", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Sum3() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Sum3(a, b)
	}
	e := &any3{typ: "SUM3", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Sum3r() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Sum3r(a, b)
	}
	e := &any3{typ: "SUM3r", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Mul3() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Mul3(a, b)
	}
	e := &any3{typ: "MUL3", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Mul3r() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Mul3r(a, b)
	}
	e := &any3{typ: "MUL3r", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Car3s() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarryS(a, b)
	}
	e := &any3{typ: "CAR3s", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Car3sr() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarrySr(a, b)
	}
	e := &any3{typ: "CAR3sr", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}

func Car3m() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarryM(a, b)
	}
	e := &any3{typ: "CAR3m", A: NewIn(), B: NewIn(), O: NewOut(), fn: f}
	go doAny3(e)
	return e
}
