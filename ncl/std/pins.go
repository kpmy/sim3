package std

import (
	"fmt"
	"sim3/ncl"
	"sim3/tri"
	"sync"
	"ypk/assert"
	"ypk/halt"
)

type out struct {
	ncl.Out
	val   *tri.Trit
	owner ncl.Element
}

func (o *out) Update(value *tri.Trit) {
	o.val = value
}

type in struct {
	ncl.In
	val   *tri.Trit
	meta  chan *sync.WaitGroup
	owner ncl.Element
}

func (i *in) Select() *tri.Trit {
	return i.val
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
				value = x.val
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
			x.val = value
		default:
			halt.As(100)
		}
	}
}

func (p *point) run(wg *sync.WaitGroup) {
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
		case *in:
			fmt.Println("run")
			wg.Add(1)
			x.meta <- wg
		default:
			halt.As(100)
		}
	}
}

func pt(n string) (ret *point) {
	ret = &point{pins: make([]ncl.Pin, 0), name: n}
	go func(p *point) {
		ncl.Step(p, func() {
			fmt.Println("point", p.name)
			wg := &sync.WaitGroup{}
			p.set(p.sel())
			p.run(wg)
			wg.Wait()
		})
	}(ret)
	return ret
}

func NewOut(o ncl.Element) (ret *out) {
	ret = &out{owner: o}
	return
}

func NewIn(o ncl.Element) (ret *in) {
	assert.For(o != nil, 20)
	ret = &in{owner: o}
	ret.meta = make(chan *sync.WaitGroup)
	go func(i *in) {
		wg := <-i.meta
		i.owner.Do()
		wg.Done()
		fmt.Println("done")
	}(ret)
	return
}
