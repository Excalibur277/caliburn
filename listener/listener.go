package listener

import (
	"fmt"

	"caliburncode.com/m/ast"
	"caliburncode.com/m/listener/listenerstack"
	"caliburncode.com/m/parsing"
	"github.com/antlr4-go/antlr/v4"
)

var _ parsing.CaliburnParserListener = &CaliburnListener{}

type CaliburnListener struct {
	parsing.BaseCaliburnParserListener
	stack listenerstack.ListenerStack
	queue listenerstack.ListenerQueue
}

func Dequeue[T interface{}](l *CaliburnListener) T {
	if d, ok := l.queue.Dequeue().(T); ok {
		return d
	} else {
		panic(fmt.Sprintf("popping invalid type, got: %T, expected %T", d, new(T)))
	}
}

func Push[T interface{}](l *CaliburnListener, d T) {
	l.stack.Push(d)
}

func NewListener() *CaliburnListener {
	return &CaliburnListener{}
}

func (l *CaliburnListener) Pop(amount int) {
	if len(l.queue) > 0 {
		panic(fmt.Sprintf("queue has %d values remaining", len(l.queue)))
	}
	l.queue = l.stack.Pop(amount)
}

// module: definitions EOF # ModuleRule
func (l *CaliburnListener) ExitModuleRule(c *parsing.ModuleRuleContext) {
	l.Pop(1)
	Push(l, ast.NewModule(Dequeue[[]ast.Definition](l)))
}

func (l *CaliburnListener) EnterEveryRule(c antlr.ParserRuleContext) {
}

func (l *CaliburnListener) ExitEveryRule(c antlr.ParserRuleContext) {
}
