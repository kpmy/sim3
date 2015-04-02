package ncl

import (
	"time"
)

func Step(obj interface{}, step func()) {
	do := func() {
		for {
			step()
			time.Sleep(time.Duration(time.Millisecond * 200))
		}
	}
	do()
}
