package tool

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sim3/ncl"
	"sim3/ncl/std"
	"ypk/assert"
)

type Import func() ncl.Element

var imps map[string]Import = make(map[string]Import)

func init() {
	imps["NOT"] = std.Not
	imps["PROBE"] = func() ncl.Element {
		return std.Probe("hi!")
	}
	imps["SUM3"] = std.Sum3
}

type Solder struct {
	imp  map[string]Import
	ent  map[string]ncl.Element
	root ncl.Compound
}

type Pin map[string]string

type PinList []Pin

type NetList struct {
	Import   []string
	Entities map[string]string
	Netlist  map[string]PinList
}

func parsePin(p string) ncl.PinCode {
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
	default:
		panic(0)
	}
}

func (s *Solder) handle(n *NetList) {
	for _, i := range n.Import {
		s.imp[i] = imps[i]
		fmt.Println(i)
	}
	for k, v := range n.Entities {
		assert.For(s.ent[k] == nil, 27, k, v)
		assert.For(s.imp[v] != nil, 28, v)
		s.ent[k] = s.imp[v]()
		fmt.Println(k, v)
	}
	for k, v := range n.Netlist {
		p := s.root.Point(k)
		for _, i := range v {
			for _e, io := range i {
				e := s.ent[_e]
				assert.For(e != nil, 29, _e)
				pin := e.Pin(parsePin(io))
				assert.For(pin != nil, 30, e, io)
				p.Solder(pin)
				fmt.Println(_e, io)
			}
		}
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
	s.imp = make(map[string]Import)
	s.ent = make(map[string]ncl.Element)
	s.parse(`import: [NOT, SUM3, PROBE]`)
}

func (s *Solder) Y(y string) {
	if s.root == nil {
		s.init()
	}
	s.parse(y)
}

func (s *Solder) F(fn string) {
	if s.root == nil {
		s.init()
	}
	path, _ := os.Getwd()
	input, err := os.Open(filepath.Join(path, "netlist", fn))
	assert.For(err == nil, 40, path, fn)
	data, _ := ioutil.ReadAll(input)
	nl := &NetList{}
	err = yaml.Unmarshal(data, nl)
	assert.For(err == nil, 41, err)
	s.handle(nl)
}