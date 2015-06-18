package std

import (
	"fmt"
	"github.com/kpmy/trigo"
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"sim3/ncl"
	"sync"
)

type out struct {
	ncl.Out
	val   *tri.Trit
	owner ncl.Element
}

func (o *out) String() string {
	return fmt.Sprint(o.owner, ":", o.val, ".", "out")
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

func (i *in) String() string {
	return fmt.Sprint(i.owner, ":", i.val, ".", "in")
}

func (i *in) Select() *tri.Trit {
	return i.val
}

type point struct {
	pins []ncl.Pin
	name string
}

func (p *point) dump(t ...string) (ret string) {
	ok := false
	for i := 0; i < len(t) && !ok; i++ {
		ok = t[i] == p.name
	}
	if ok {
		ret = p.name
		for _, x := range p.pins {
			ret = fmt.Sprint(ret, " ", x)
		}
		fmt.Println(ret)
	}
	return
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
		} else {
			halt.As(100, pin)
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

func (p *point) run() {
	wg := &sync.WaitGroup{}
	for _, _x := range p.pins {
		switch x := _x.(type) {
		case *out:
			x.owner.Do()
		case *in:
			wg.Add(1)
			x.meta <- wg
		default:
			halt.As(100)
		}
	}
	wg.Wait()
}

func pt(n string) (ret *point) {
	ret = &point{pins: make([]ncl.Pin, 0), name: n}
	go func(p *point) {
		ncl.Step(p, func() {
			p.set(p.sel())
			//p.dump("a", "b", "c", "s")
			p.run()
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
		ncl.Step(i, func() {
			wg := <-i.meta
			i.owner.Do()
			wg.Done()
		})
	}(ret)
	return
}
