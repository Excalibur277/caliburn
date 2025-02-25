package ast

import "github.com/antlr4-go/antlr/v4"

type BinaryOperation interface {
	IsBinaryOperation()
}

type UnaryOperation interface {
	IsUnaryOperation()
}

func NewBinaryOperation(token antlr.Token) BinaryOperation {
	return nil
}

func NewUnaryOperation(token antlr.Token) UnaryOperation {
	return nil
}
