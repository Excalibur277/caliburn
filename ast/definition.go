package ast

import "fmt"

type Definition interface {
	Statement
	IsDefinition()
}

type DefinitionBase struct {
	StatementBase
}

func (db *DefinitionBase) IsDefinition() {}

type FunctionDefinition struct {
	DefinitionBase
	identifier      Identifier
	definedFunction Function
}

func NewFunctionDefinition(functionType FunctionType, identifier Identifier, parameters []Parameter, returnType Type, block Block) Definition {
	return &FunctionDefinition{
		identifier:      identifier,
		definedFunction: NewFunction(functionType, parameters, returnType, block),
	}
}

func (f *FunctionDefinition) String() string {
	return fmt.Sprintf("%s defined as %s", f.identifier, f.definedFunction)
}
