package tri

var TRUE Trit = Trit{n: false, t: true}
var FALSE Trit = Trit{n: false, t: false}
var NIL Trit = Trit{n: true, t: false}

type Trit struct {
	n bool
	t bool
}

func (t Trit) String() string {
	if t.n {
		return "%nil"
	} else if t.t {
		return "%true"
	} else {
		return "%false"
	}
}
