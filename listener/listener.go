package listener

import (
	"caliburncode.com/m/parsing"
)

var _ parsing.CaliburnParserListener = &CaliburnListener{}

type CaliburnListener struct {
}

func NewListener() *CaliburnListener {
	return &CaliburnListener{}
}
