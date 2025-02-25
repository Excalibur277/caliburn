package main

import (
	"bufio"
	"fmt"
	"os"

	"caliburncode.com/m/listener"
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

	module := listener.ParseStream(is)

	fmt.Println(module)
}
