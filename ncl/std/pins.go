package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/assert"
	"ypk/halt"
)

type out struct {
	ncl.Out
	signal chan tri.Trit
	meta   chan tri.Trit
}

func (o *out) Validate(valid bool, value ...tri.Trit) {
	if valid {
		assert.For(len(value) == 1, 20)
		o.meta <- tri.TRUE
		o.signal <- value[0]
	} else {
		o.meta <- tri.FALSE
	}
}

type in struct {
	ncl.In
	signal chan tri.Trit
	meta   chan tri.Trit
}

func (i *in) Select() (valid bool, value tri.Trit) {
	meta := <-i.meta
	if meta == tri.TRUE {
		valid = true
		value = <-i.signal
	}
	return
}

type point struct {
	pins []ncl.Pin
}

func (p *point) Solder(pins ...ncl.Pin) {
	exists := func(pin ncl.Pin) bool {
		for _, x := range p.pins {
			if x == pin {
				return true
			}
		}
		return false
	}
	for _, pin := range pins {
		if !exists(pin) {
			p.pins = append(p.pins, pin)
		}
	}
}

func (p *point) sel() (meta tri.Trit, signal tri.Trit) {
	meta, signal = tri.FALSE, tri.NIL
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
			assert.For(meta == tri.FALSE, 100)
			meta := <-x.meta
			if meta == tri.TRUE {
				signal = <-x.signal
			}
		case *in:
		default:
			halt.As(100)
		}
	}
	return
}

func (p *point) set(meta tri.Trit, signal tri.Trit) {
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
		case *in:
			x.meta <- meta
			if meta == tri.TRUE {
				x.signal <- signal
			}
		default:
			halt.As(100)
		}
	}
}

func pt() (ret *point) {
	ret = &point{pins: make([]ncl.Pin, 0)}
	go func(p *point) {
		ncl.Step(p, func() {
			p.set(p.sel())
		})
	}(ret)
	return ret
}

func newOut() (ret *out) {
	ret = &out{}
	ret.signal = make(chan tri.Trit, 1)
	ret.meta = make(chan tri.Trit, 1)
	return
}

func newIn() (ret *in) {
	ret = &in{}
	ret.signal = make(chan tri.Trit)
	ret.meta = make(chan tri.Trit)
	return
}
