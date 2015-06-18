package std

import (
	"github.com/kpmy/trigo"
	"github.com/kpmy/ypk/halt"
	"sim3/ncl"
)

type sw struct {
	I ncl.In
	O ncl.Out
	A ncl.In
}

func (e *sw) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.A:
		return e.A
	case ncl.I:
		return e.I
	case ncl.O:
		return e.O
	default:
		halt.As(100)
	}
	panic(0)
}

func (e *sw) init() {
	e.A = NewIn(e)
	e.I = NewIn(e)
	e.O = NewOut(e)
}

func (e *sw) Do() {
	_a := e.A.Select()
	var val *tri.Trit
	if _a != nil {
		a := *_a
		if a == tri.TRUE {
			val = e.I.Select()
		}
	}
	e.O.Update(val)

}

func Sw() ncl.Element {
	e := &sw{}
	e.init()
	return e
}

type mux struct {
	T, N, F ncl.In
	A       ncl.In
	B       ncl.Out
}

func (e *mux) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.A:
		return e.A
	case ncl.T:
		return e.T
	case ncl.N:
		return e.N
	case ncl.F:
		return e.F
	case ncl.B:
		return e.B
	default:
		halt.As(100)
	}
	panic(0)
}

func (e *mux) init() {
	e.A = NewIn(e)
	e.B = NewOut(e)
	e.T = NewIn(e)
	e.N = NewIn(e)
	e.F = NewIn(e)
}

func (e *mux) Do() {
	_a := e.A.Select()
	var val *tri.Trit
	if _a != nil {
		a := *_a
		if a == tri.TRUE {
			val = e.T.Select()
		} else if a == tri.NIL {
			val = e.N.Select()
		} else if a == tri.FALSE {
			val = e.F.Select()
		}
	}
	e.B.Update(val)

}

type demux struct {
	T, N, F ncl.Out
	A       ncl.In
	B       ncl.In
}

func (e *demux) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.A:
		return e.A
	case ncl.T:
		return e.T
	case ncl.N:
		return e.N
	case ncl.F:
		return e.F
	case ncl.B:
		return e.B
	default:
		halt.As(100)
	}
	panic(0)
}

func (e *demux) init() {
	e.A = NewIn(e)
	e.B = NewIn(e)
	e.T = NewOut(e)
	e.N = NewOut(e)
	e.F = NewOut(e)
}

func (e *demux) Do() {
	reset := func(e *demux) {
		e.T.Update(nil)
		e.N.Update(nil)
		e.F.Update(nil)
	}
	_a := e.A.Select()
	if _a != nil {
		b := e.B.Select()
		if b != nil {
			a := *_a
			if a == tri.TRUE {
				e.T.Update(b)
			} else if a == tri.NIL {
				e.N.Update(b)
			} else if a == tri.FALSE {
				e.F.Update(b)
			}
		} else {
			reset(e)
		}
	} else {
		reset(e)
	}
}

func Mux() ncl.Element {
	e := &mux{}
	e.init()
	return e
}

func Demux() ncl.Element {
	e := &demux{}
	e.init()
	return e
}
