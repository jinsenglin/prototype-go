package line

import (
	"github.com/jinsenglin/prototype-go/pkg/model"
)

func Run() (line *model.Line) {
	line = model.NewLine()

	go line.Listen()

	return
}
