package std

import (
	"fmt"
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

func (t *power) Do() {
	t.O.Update(&t.value)
}

func (p *power) String() string {
	return fmt.Sprint(p.value)
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
	ret = &probe{name: n}
	ret.(*probe).I = NewIn(ret)
	return
}

func (p *probe) Do() {
	signal := p.I.Select()
	meta := signal != nil
	sig := tri.FALSE
	if meta {
		sig = *signal
	}
	api.Log(&api.Item{Name: p.name, Type: "probe", Meta: meta, Signal: sig})
}

func (p *probe) String() string {
	return fmt.Sprint(p.name)
}

func Source(trit tri.Trit) (ret *power) {
	ret = &power{value: trit}
	ret.O = NewOut(ret)
	ret.Do()
	return
}

type gen struct {
	O     ncl.Out
	seq   []tri.Trit
	i     int
	valid bool
}

func (g *gen) Pin(c ncl.PinCode) ncl.Pin {
	assert.For(c == ncl.O, 20)
	return g.O
}

func (g *gen) Do() {
	if g.valid {
		g.O.Update(&g.seq[g.i])
		g.i++
	} else {
		g.O.Update(nil)
	}
	g.valid = !g.valid
	if g.i == len(g.seq) {
		g.i = 0
	}
}

func Generator(s ...tri.Trit) (ret ncl.Element) {
	assert.For(len(s) > 0, 20)
	ret = &gen{seq: s}
	ret.(*gen).O = NewOut(ret)
	go func(g *gen) {
		ncl.Step(g, func() {
			g.Do()
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

func (b *board) Do() {}

func (b *board) String() string {
	return "board"
}

func (brd *board) Pins(pins ...map[ncl.PinCode]ncl.Pin) map[ncl.PinCode]ncl.Pin {
	if len(pins) > 0 {
		brd.pins = make(map[ncl.PinCode]ncl.Pin)
		brd._pins = make(map[ncl.PinCode]ncl.Pin)
		if pins[0] != nil {
			for k, _p := range pins[0] {
				b := Buffer().(*any2)
				switch p := _p.(type) {
				case *in:
					b.I = p
					brd.pins[k] = p
					p.owner = b
					brd._pins[k] = b.O
				case *out:
					b.O = p
					brd.pins[k] = p
					p.owner = b
					brd._pins[k] = b.I
				default:
					halt.As(100, reflect.TypeOf(p))
				}
			}
		}
	}
	return brd.pins
}

func Board() ncl.Compound {
	ret := &board{}
	ret.points = make(map[string]ncl.Point)
	ret.Point("T").Solder(Static.Pos.Pin(ncl.O))
	ret.Point("F").Solder(Static.Neg.Pin(ncl.O))
	ret.Point("N").Solder(Static.Ground.Pin(ncl.O))
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

func (t *trig) Value(v ...*tri.Trit) *tri.Trit {
	if len(v) > 0 {
		t.data = *v[0]
	}
	return &t.data
}

func (t *trig) Do() {
	val := t.S.Select()
	if val != nil {
		t.data = *val
	}
	t.D.Update(&t.data)
}
func (t *trig) String() string {
	return fmt.Sprint("T", ":", t.data)
}

func Trigger() ncl.Element {
	t := &trig{data: tri.NIL}
	t.D = NewOut(t)
	t.S = NewIn(t)
	return t
}
