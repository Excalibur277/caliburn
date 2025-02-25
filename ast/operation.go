package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type BinaryOperation interface {
	Node
	IsBinaryOperation()
}

type UnaryOperation interface {
	Node
	IsUnaryOperation()
}

func NewBinaryOperation(token antlr.Token) BinaryOperation {
	fmt.Println(token.GetText())
	return nil // TODO
}

func NewUnaryOperation(token antlr.Token) UnaryOperation {
	fmt.Println(token.GetText())
	return nil // TODO
}
