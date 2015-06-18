package std

import (
	"fmt"
	"github.com/kpmy/trigo"
	"github.com/kpmy/ypk/halt"
	"sim3/ncl"
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

func (e *any3) Do() {
	a := e.A.Select()
	b := e.B.Select()
	if a != nil && b != nil {
		val := e.fn(*a, *b)
		e.O.Update(&val)
	} else {
		e.O.Update(nil)
	}
}

func (e *any3) String() string {
	return fmt.Sprint(e.typ)
}

func (e *any3) init() {
	e.A = NewIn(e)
	e.B = NewIn(e)
	e.O = NewOut(e)

}

func AndNot() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Not(tri.And(a, b))
	}
	e := &any3{typ: "¬&", fn: f}
	e.init()
	return e
}

func OrNot() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Not(tri.Or(a, b))
	}
	e := &any3{typ: "¬|", fn: f}
	e.init()
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
	e := &any3{typ: "<=>", fn: f}
	e.init()
	return e
}

func Sum3() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Sum3(a, b)
	}
	e := &any3{typ: "SUM3", fn: f}
	e.init()
	return e
}

func Sum3r() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Sum3r(a, b)
	}
	e := &any3{typ: "SUM3r", fn: f}
	e.init()
	return e
}

func Mul3() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Mul3(a, b)
	}
	e := &any3{typ: "MUL3", fn: f}
	e.init()
	return e
}

func Mul3r() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.Mul3r(a, b)
	}
	e := &any3{typ: "MUL3r", fn: f}
	e.init()
	return e
}

func Car3s() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarryS(a, b)
	}
	e := &any3{typ: "CAR3s", fn: f}
	e.init()
	return e
}

func Car3sr() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarrySr(a, b)
	}
	e := &any3{typ: "CAR3sr", fn: f}
	e.init()
	return e
}

func Car3m() ncl.Element {
	f := func(a tri.Trit, b tri.Trit) tri.Trit {
		return tri.CarryM(a, b)
	}
	e := &any3{typ: "CAR3m", fn: f}
	e.init()
	return e
}
