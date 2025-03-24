package ast

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type BinaryOperation interface {
	Node
	IsBinaryOperation()
}

type BinaryOperationBase struct{}

func (bo BinaryOperationBase) IsBinaryOperation() {}

type BinaryOperationADD struct{ BinaryOperationBase }

func (bo BinaryOperationADD) String() string { return "+" }

type BinaryOperationSUB struct{ BinaryOperationBase }

func (bo BinaryOperationSUB) String() string { return "-" }

type BinaryOperationMUL struct{ BinaryOperationBase }

func (bo BinaryOperationMUL) String() string { return "*" }

type BinaryOperationDIV struct{ BinaryOperationBase }

func (bo BinaryOperationDIV) String() string { return "/" }

type BinaryOperationMOD struct{ BinaryOperationBase }

func (bo BinaryOperationMOD) String() string { return "%" }

type BinaryOperationPOW struct{ BinaryOperationBase }

func (bo BinaryOperationPOW) String() string { return "^" }

type BinaryOperationROOT struct{ BinaryOperationBase }

func (bo BinaryOperationROOT) String() string { return "~" }

type BinaryOperationOR struct{ BinaryOperationBase }

func (bo BinaryOperationOR) String() string { return "|" }

type BinaryOperationXOR struct{ BinaryOperationBase }

func (bo BinaryOperationXOR) String() string { return "&!|" }

type BinaryOperationAND struct{ BinaryOperationBase }

func (bo BinaryOperationAND) String() string { return "&" }

type BinaryOperationLSHIFT struct{ BinaryOperationBase }

func (bo BinaryOperationLSHIFT) String() string { return "<<" }

type BinaryOperationRSHIFT struct{ BinaryOperationBase }

func (bo BinaryOperationRSHIFT) String() string { return ">>" }

type BinaryOperationEQU struct{ BinaryOperationBase }

func (bo BinaryOperationEQU) String() string { return "==" }

type BinaryOperationNEQ struct{ BinaryOperationBase }

func (bo BinaryOperationNEQ) String() string { return "!=" }

type BinaryOperationGTE struct{ BinaryOperationBase }

func (bo BinaryOperationGTE) String() string { return ">=" }

type BinaryOperationGRT struct{ BinaryOperationBase }

func (bo BinaryOperationGRT) String() string { return ">" }

type BinaryOperationLTE struct{ BinaryOperationBase }

func (bo BinaryOperationLTE) String() string { return "<=" }

type BinaryOperationLST struct{ BinaryOperationBase }

func (bo BinaryOperationLST) String() string { return "<" }

func NewBinaryOperation(token antlr.Token) BinaryOperation {
	switch t := token.GetText(); t {
	case "+":
		return BinaryOperationADD{}
	case "-":
		return BinaryOperationSUB{}
	case "*":
		return BinaryOperationMUL{}
	case "/":
		return BinaryOperationDIV{}
	case "%":
		return BinaryOperationMOD{}
	case "^":
		return BinaryOperationPOW{}
	case "~":
		return BinaryOperationROOT{}
	case "|":
		return BinaryOperationOR{}
	case "|!&": // TODO Change this
		return BinaryOperationXOR{}
	case "&":
		return BinaryOperationAND{}
	case "<<":
		return BinaryOperationLSHIFT{}
	case ">>":
		return BinaryOperationRSHIFT{}
	case "==":
		return BinaryOperationEQU{}
	case "!=":
		return BinaryOperationNEQ{}
	case ">=":
		return BinaryOperationGTE{}
	case ">":
		return BinaryOperationGRT{}
	case "<=":
		return BinaryOperationLTE{}
	case "<":
		return BinaryOperationLST{}
	default:
		panic(fmt.Sprintf("Could not find operator: %s", t))
	}
}

type UnaryOperation interface {
	Node
	IsUnaryOperation()
}

type UnaryOperationBase struct{}

func (uo UnaryOperationBase) IsUnaryOperation() {}

type UnaryOperationPOS struct{ UnaryOperationBase }

func (uo UnaryOperationPOS) String() string { return "+" }

type UnaryOperationNEG struct{ UnaryOperationBase }

func (uo UnaryOperationNEG) String() string { return "-" }

type UnaryOperationNOT struct{ UnaryOperationBase }

func (uo UnaryOperationNOT) String() string { return "!" }

func NewUnaryOperation(token antlr.Token) UnaryOperation {
	switch t := token.GetText(); t {
	case "+":
		return UnaryOperationPOS{}
	case "-":
		return UnaryOperationNEG{}
	case "!":
		return UnaryOperationNOT{}
	default:
		panic(fmt.Sprintf("Could not find operator: %s", t))
	}
}
