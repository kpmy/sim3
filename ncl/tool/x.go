package tool

import (
	"github.com/kpmy/ypk/assert"
	"gopkg.in/yaml.v2"
	"hash/fnv"
	"reflect"
	"sim3/ncl"
	"sim3/ncl/std"
	"sim3/tri"
)

type Import func(...interface{}) ncl.Element

type DataSource interface {
	Get(name string) (data []byte, err error)
}

var imps map[string]Import = make(map[string]Import)

func Simple(f func() ncl.Element) Import {
	return func(...interface{}) ncl.Element {
		return f()
	}
}

type PinClosure func(ncl.Element) ncl.Pin

func In(e ncl.Element) ncl.Pin {
	return std.NewIn(e)
}

func Out(e ncl.Element) ncl.Pin {
	return std.NewOut(e)
}

func init() {
	imps["NOT"] = Simple(std.Not)
	imps["PROBE"] = func(opts ...interface{}) ncl.Element {
		assert.For(len(opts) != 0, 20)
		return std.Probe(opts[0].(string))
	}
	imps["SUM3"] = Simple(std.Sum3)
	imps["SUM3r"] = Simple(std.Sum3r)
	imps["DMX"] = Simple(std.Demux)
	imps["MX"] = Simple(std.Mux)
	imps["NAND"] = Simple(std.AndNot)
	imps["NOR"] = Simple(std.OrNot)
	imps["CAR3s"] = Simple(std.Car3s)
	imps["CAR3m"] = Simple(std.Car3m)
	imps["CAR3sr"] = Simple(std.Car3sr)
	imps["CMP"] = Simple(std.Cmp)
	imps["CL"] = Simple(std.CycleLeft)
	imps["CR"] = Simple(std.CycleRight)
	imps["T"] = Simple(std.Trigger)
	imps["SW"] = Simple(std.Sw)
}

func Register(name string, i Import) {
	imps[name] = i
}

type Solder struct {
	imp  map[string]Import
	ent  map[string]ncl.Element
	root ncl.Compound
	pins map[ncl.PinCode]PinClosure
	Data DataSource
}

type Pin map[string]string

type PinList []Pin

type NetList struct {
	Import   []string
	Entities map[string]string
	Netlist  map[string]PinList
	Init     map[string]string
}

func value(v string) tri.Trit {
	switch v {
	case "T":
		return tri.TRUE
	case "N":
		return tri.NIL
	case "F":
		return tri.FALSE
	default:
		panic(0)
	}
}

func encodePin(p string) ncl.PinCode {
	switch p {
	case "A":
		return ncl.A
	case "B":
		return ncl.B
	case "N":
		return ncl.N
	case "T":
		return ncl.T
	case "F":
		return ncl.F
	case "I":
		return ncl.I
	case "O":
		return ncl.O
	case "S":
		return ncl.S
	case "D":
		return ncl.D
	case "C":
		return ncl.C
	default:
		h := fnv.New32()
		h.Write([]byte(p))
		return ncl.PinCode(h.Sum32())
	}
}

func (s *Solder) handle(n *NetList) {
	for _, i := range n.Import {
		s.imp[i] = imps[i]
	}
	for k, v := range n.Entities {
		assert.For(s.ent[k] == nil, 27, k, v)
		assert.For(s.imp[v] != nil, 28, v)
		s.ent[k] = s.imp[v](k)
	}
	for k, v := range n.Netlist {
		p := s.root.Point(k)
		for _, i := range v {
			for _e, io := range i {
				e := s.ent[_e]
				assert.For(e != nil, 29, _e)
				var pin ncl.Pin
				if _e == "$" {
					pin = s.root.InnerPin(encodePin(io))
				} else {
					pin = e.Pin(encodePin(io))
				}
				assert.For(pin != nil, 30, e, io)
				p.Solder(pin)
			}
		}
	}
	for k, v := range n.Init {
		t, ok := s.ent[k].(ncl.Trigger)
		assert.For(ok, 20, reflect.TypeOf(s.ent[k]))
		val := value(v)
		t.Value(&val)
	}
}

func (s *Solder) parse(data string) {
	n := &NetList{}
	err := yaml.Unmarshal([]byte(data), n)
	assert.For(err == nil, 39, data, err)
	s.handle(n)
}

func (s *Solder) init() {
	s.root = std.Board()
	pins := make(map[ncl.PinCode]ncl.Pin)
	for k, v := range s.pins {
		pins[k] = v(s.root)
	}
	s.root.Pins(pins)
	s.imp = make(map[string]Import)
	s.ent = make(map[string]ncl.Element)
	s.parse(`import: [PROBE]`)
	s.ent["$"] = s.root
}

func (s *Solder) UserPin(name string, p PinClosure) {
	assert.For(name != "", 20)
	assert.For(p != nil, 21)
	assert.For(s.root == nil, 22)
	if s.pins == nil {
		s.pins = make(map[ncl.PinCode]PinClosure)
	}
	s.pins[encodePin(name)] = p
}

func (s *Solder) Y(y string) {
	if s.root == nil {
		s.init()
	}
	s.parse(y)
}

var Src DataSource

func (s *Solder) F(fn string) {
	if s.root == nil {
		s.init()
	}
	if s.Data == nil {
		s.Data = Src
	}
	assert.For(s.Data != nil, 20)
	data, _ := s.Data.Get(fn)
	nl := &NetList{}
	err := yaml.Unmarshal(data, nl)
	assert.For(err == nil, 41, err)
	s.handle(nl)
}

func (s *Solder) Root() ncl.Element {
	assert.For(s.root != nil, 20)
	return s.root
}
