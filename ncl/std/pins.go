package std

import (
	"sim3/ncl"
	"sim3/tri"
	"ypk/assert"
	"ypk/halt"
)

type out struct {
	ncl.Out
	meta chan *tri.Trit
}

func (o *out) Validate(valid bool, value ...tri.Trit) {
	if valid {
		assert.For(len(value) == 1, 20)
		o.meta <- &value[0]
	} else {
		o.meta <- nil
	}
}

type in struct {
	ncl.In
	meta chan *tri.Trit
}

func (i *in) Select() (valid bool, value tri.Trit) {
	var tmp *tri.Trit
	select {
	case tmp = <-i.meta:
	}
	valid = tmp != nil
	if valid {
		value = *tmp
	}
	return
}

type point struct {
	pins []ncl.Pin
	name string
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

func (p *point) sel() (value *tri.Trit) {
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
			if value == nil {
				select {
				case value = <-x.meta:
				}
			} else {
				select {
				case _ = <-x.meta:
				}
			}
		case *in:
		default:
			halt.As(100)
		}
	}
	return
}

func (p *point) set(value *tri.Trit) {
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
		case *in:
			x.meta <- value
		default:
			halt.As(100)
		}
	}
}

func pt(n string) (ret *point) {
	ret = &point{pins: make([]ncl.Pin, 0), name: n}
	go func(p *point) {
		ncl.Step(p, func() {
			p.set(p.sel())
		})
	}(ret)
	return ret
}

func NewOut() (ret *out) {
	ret = &out{}
	ret.meta = make(chan *tri.Trit)
	return
}

func NewIn() (ret *in) {
	ret = &in{}
	ret.meta = make(chan *tri.Trit)
	return
}
