package std

import (
	"reflect"
	"sim3/api"
	"sim3/ncl"
	"sim3/tri"
	"ypk/assert"
	"ypk/halt"
)

type power struct {
	ncl.Element
	O     ncl.Out
	value tri.Trit
}

func (t *power) Pin(c ncl.PinCode) ncl.Pin {
	assert.For(c == ncl.O, 20)
	return t.O
}

type probe struct {
	ncl.Element
	I    ncl.In
	name string
}

func (p *probe) Pin(c ncl.PinCode) ncl.Pin {
	assert.For(c == ncl.I, 20)
	return p.I
}
func Probe(n string) (ret ncl.Element) {
	ret = &probe{I: NewIn(), name: n}
	go func(p *probe) {
		ncl.Step(p, func() {
			meta, signal := p.I.Select()
			api.Log(&api.Item{Name: p.name, Type: "probe", Meta: meta, Signal: signal})
			//fmt.Println(p.name, meta, signal)
		})
	}(ret.(*probe))
	return
}

func Source(trit tri.Trit) (ret *power) {
	ret = &power{O: NewOut(), value: trit}
	go func(p *power) {
		ncl.Step(p, func() {
			p.O.Validate(true, p.value)
		})
	}(ret)
	return
}

type gen struct {
	O   ncl.Out
	seq []tri.Trit
}

func (g *gen) Pin(c ncl.PinCode) ncl.Pin {
	assert.For(c == ncl.O, 20)
	return g.O
}

func Generator(s ...tri.Trit) (ret ncl.Element) {
	assert.For(len(s) > 0, 20)
	ret = &gen{O: NewOut(), seq: s}
	go func(g *gen) {
		i := 0
		valid := true
		ncl.Step(g, func() {
			g.O.Validate(valid, g.seq[i])
			if valid {
				i++
			}
			valid = !valid
			if i == len(g.seq) {
				i = 0
			}
		})
	}(ret.(*gen))
	return
}

var Static struct {
	Pos    ncl.Element
	Neg    ncl.Element
	Ground ncl.Element
}

func init() {
	Static.Pos = Source(tri.TRUE)
	Static.Neg = Source(tri.FALSE)
	Static.Ground = Source(tri.NIL)
}

type board struct {
	ncl.Compound
	points map[string]ncl.Point
	pins   map[ncl.PinCode]ncl.Pin
	_pins  map[ncl.PinCode]ncl.Pin
}

func (b *board) Pin(c ncl.PinCode) (ret ncl.Pin) {
	ret = b.pins[c]
	assert.For(ret != nil, 20)
	return
}

func (b *board) InnerPin(c ncl.PinCode) (ret ncl.Pin) {
	ret = b._pins[c]
	assert.For(ret != nil, 20)
	return
}

func (b *board) Point(x string) (ret ncl.Point) {
	ret = b.points[x]
	if ret == nil {
		ret = pt(x)
		b.points[x] = ret
	}
	return
}

func Board(pins map[ncl.PinCode]ncl.Pin) ncl.Compound {
	ret := &board{}
	ret.points = make(map[string]ncl.Point)
	ret.Point("T").Solder(Static.Pos.Pin(ncl.O))
	ret.Point("F").Solder(Static.Neg.Pin(ncl.O))
	ret.Point("N").Solder(Static.Ground.Pin(ncl.O))
	ret.pins = make(map[ncl.PinCode]ncl.Pin)
	ret._pins = make(map[ncl.PinCode]ncl.Pin)
	if pins != nil {
		for k, _p := range pins {
			b := Buffer().(*any2)
			switch p := _p.(type) {
			case ncl.In:
				b.I = p
				ret.pins[k] = p
				ret._pins[k] = b.O
			case ncl.Out:
				b.O = p
				ret.pins[k] = p
				ret._pins[k] = b.I
			default:
				halt.As(100, reflect.TypeOf(p))
			}
		}
	}
	return ret
}

type trig struct {
	D    ncl.Out
	S    ncl.In
	data tri.Trit
}

func (t *trig) Pin(c ncl.PinCode) ncl.Pin {
	switch c {
	case ncl.D:
		return t.D
	case ncl.S:
		return t.S
	}
	panic(0)
}

func Trigger() ncl.Element {
	t := &trig{D: NewOut(), S: NewIn(), data: tri.NIL}
	go func(t *trig) {
		ncl.Step(t, func() {
			ok, val := t.S.Select()
			if ok {
				t.data = val
			}
			t.D.Validate(true, t.data)
		})
	}(t)
	return t
}
