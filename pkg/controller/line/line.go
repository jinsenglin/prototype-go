package line

import (
	"time"

	"github.com/jinsenglin/prototype-go/pkg/model"
)

func Run() {
	go func() {
		line := model.NewLine()

		// DEMO CODE
		for {
			line.OpenChannel()
			line.Dump()
			time.Sleep(10e9)
			line.OpenChannel()
			line.Dump()
			time.Sleep(10e9)
			line.CloseChannel(1)
			line.Dump()
			time.Sleep(10e9)
			line.CloseChannel(0)
			line.Dump()
			time.Sleep(10e9)
			break
		}
		// END DEMO
	}()
}
