package listener

import (
	"caliburncode.com/m/parsing"
)

var _ parsing.CaliburnParserVisitor = &CaliburnVisitor{}

type CaliburnVisitor struct {
}

func NewVisitor() *CaliburnVisitor {
	return &CaliburnVisitor{}
}
