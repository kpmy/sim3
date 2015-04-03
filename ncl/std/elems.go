package std

import (
	"sim3/api"
	"sim3/ncl"
	"sim3/tri"
	"ypk/assert"
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
	ret = &probe{I: newIn(), name: n}
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
	ret = &power{O: newOut(), value: trit}
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
	ret = &gen{O: newOut(), seq: s}
	go func(g *gen) {
		i := 0
		valid := true
		ncl.Step(g, func() {
			g.O.Validate(valid, g.seq[i])
			valid = !valid
			i++
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
}

func (b *board) Point(x string, p ...ncl.Point) (ret ncl.Point) {
	ret = b.points[x]
	if ret == nil {
		ret = pt()
		b.points[x] = ret
	}
	return
}

func Board() ncl.Compound {
	ret := &board{}
	ret.points = make(map[string]ncl.Point)
	ret.Point("+").Solder(Static.Pos.Pin(ncl.O))
	ret.Point("-").Solder(Static.Neg.Pin(ncl.O))
	ret.Point("0").Solder(Static.Ground.Pin(ncl.O))
	return ret
}
