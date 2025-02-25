package listener

import (
	"fmt"

	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func ParseStream(is antlr.CharStream) ast.Module {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	lexer := parsing.NewCaliburnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parsing.NewCaliburnParser(stream)

	listener := NewListener()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Module())

	listener.Pop(1)
	module := Dequeue[ast.Module](listener)
	return module
}
