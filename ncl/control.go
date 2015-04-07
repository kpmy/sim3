package ncl

import "time"

var slow time.Duration = time.Duration(time.Millisecond * 1)

func Step(obj interface{}, step func()) {
	do := func() {
		for {
			step()
			time.Sleep(slow)
		}
	}
	do()
}
