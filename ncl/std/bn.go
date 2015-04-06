package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/halt"
)

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

func doMux(e *mux) {
	ncl.Step(e, func() {
		da, a := e.A.Select()
		ok, val := false, tri.NIL

		if da {
			if a == tri.TRUE {
				ok, val = e.T.Select()
			} else if a == tri.NIL {
				ok, val = e.N.Select()
			} else if a == tri.FALSE {
				ok, val = e.F.Select()
			}
		}
		e.B.Validate(ok, val)
	})
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

func doDemux(e *demux) {
	reset := func(e *demux) {
		e.T.Validate(false)
		e.N.Validate(false)
		e.F.Validate(false)
	}
	ncl.Step(e, func() {
		da, a := e.A.Select()
		if da {
			db, b := e.B.Select()
			if db {
				if a == tri.TRUE {
					e.T.Validate(db, b)
				} else if a == tri.NIL {
					e.N.Validate(db, b)
				} else if a == tri.FALSE {
					e.F.Validate(db, b)
				}
			} else {
				reset(e)
			}
		} else {
			reset(e)
		}
	})
}

func Mux() ncl.Element {
	e := &mux{A: NewIn(), B: NewOut(), T: NewIn(), N: NewIn(), F: NewIn()}
	go doMux(e)
	return e
}

func Demux() ncl.Element {
	e := &demux{A: NewIn(), B: NewIn(), T: NewOut(), N: NewOut(), F: NewOut()}
	go doDemux(e)
	return e
}
