package main

import (
	"bufio"
	"fmt"
	"os"

	"caliburncode.com/m/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires file commandline argument")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	is := antlr.NewIoStream(bufio.NewReader(f))
	lexer := parsing.NewCaliburnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parsing.NewCaliburnParser(stream)

	listener := listener.NewListener()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Module())

}
