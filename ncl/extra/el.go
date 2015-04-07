package extra

import (
	"sim3/ncl"
	"sim3/ncl/tool"
)

func SM3() ncl.Element {
	t := &tool.Solder{}
	t.UserPin("A", tool.In)
	t.UserPin("B", tool.In)
	t.UserPin("C", tool.Out)
	t.UserPin("S", tool.Out)
	t.F("sm3.nl")
	return t.Root()
}

func init() {
	tool.Register("SM3", tool.Simple(SM3))
	tool.Register("SM3r", tool.Simple(SM3r))
	tool.Register("SM", tool.Simple(SM))
	tool.Register("SMr", tool.Simple(SMr))
}

func SM3r() ncl.Element {
	t := &tool.Solder{}
	t.UserPin("A", tool.In)
	t.UserPin("B", tool.In)
	t.UserPin("Cr", tool.Out)
	t.UserPin("Sr", tool.Out)
	t.F("sm3r.nl")
	return t.Root()
}

func SM() ncl.Element {
	t := &tool.Solder{}
	t.UserPin("C0", tool.In)
	t.UserPin("A", tool.In)
	t.UserPin("B", tool.In)
	t.UserPin("S", tool.Out)
	t.UserPin("C", tool.Out)
	t.F("sm.nl")
	return t.Root()
}

func SMr() ncl.Element {
	t := &tool.Solder{}
	t.UserPin("C0r", tool.In)
	t.UserPin("A", tool.In)
	t.UserPin("B", tool.In)
	t.UserPin("Sr", tool.Out)
	t.UserPin("Cr", tool.Out)
	t.F("smr.nl")
	return t.Root()
}
