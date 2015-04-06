package extra

import (
	"sim3/ncl"
	"sim3/ncl/std"
	"sim3/ncl/tool"
)

func SM3() ncl.Element {
	t := &tool.Solder{}
	t.UserPin("A", std.NewIn())
	t.UserPin("B", std.NewIn())
	t.UserPin("C", std.NewOut())
	t.UserPin("S", std.NewOut())
	t.F("sm3.nl")
	return t.Root()
}

func init() {
	tool.Register("SM3", SM3)
}
