package std

import (
	"fmt"
	"sim3/api"
	"sim3/ncl"
	"sim3/tri"
	"ypk/assert"
)

type plus struct {
	ncl.Element
	O ncl.Out
}

func (t *plus) Pin(c ncl.PinCode) ncl.Pin {
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

var Static struct {
	Pos ncl.Element
}

func NewProbe(n string) (ret ncl.Element) {
	ret = &probe{I: newIn(), name: n}
	go func(p *probe) {
		ncl.Step(p, func() {
			meta, signal := p.I.Select()
			api.Log(&api.Item{Name: p.name, Type: "probe", Meta: meta, Signal: signal})
			fmt.Println(p.name, meta, signal)
		})
	}(ret.(*probe))
	return
}

func newPlus() (ret *plus) {
	ret = &plus{O: newOut()}
	go func(p *plus) {
		ncl.Step(p, func() {
			p.O.Validate(true, tri.TRUE)
		})
	}(ret)
	return
}

func init() {
	Static.Pos = newPlus()
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
	return ret
}
