package std

import (
	"fmt"
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
	I ncl.In
}

func (p *probe) Pin(c ncl.PinCode) ncl.Pin {
	assert.For(c == ncl.I, 20)
	return p.I
}

var Static struct {
	Pos ncl.Element
}

func NewProbe() (ret ncl.Element) {
	ret = &probe{I: newIn()}
	go func(p *probe) {
		fmt.Println("probe")
		for {
			meta, signal := p.I.Select()
			fmt.Println("probe", meta, signal)
		}
	}(ret.(*probe))
	return
}

func newPlus() (ret *plus) {
	ret = &plus{O: newOut()}
	go func(p *plus) {
		fmt.Println("plus")
		for {
			p.O.Validate(true, tri.TRUE)
			fmt.Print("+")
		}
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
