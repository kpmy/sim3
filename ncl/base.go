package ncl

import (
	"sim3/tri"
)

type PinCode int

const (
	I PinCode = iota
	O
)

type Pin interface {
}

type In interface {
	Pin
	Select() (bool, tri.Trit)
}

type Out interface {
	Pin
	Validate(bool, ...tri.Trit)
}

type Point interface {
	Solder(...Pin)
}

type Element interface {
	Pin(PinCode) Pin
}

type Compound interface {
	Element
	Point(string, ...Point) Point
}
