package ast

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

type TypeDefinition struct {
	DefinitionBase
}

type ClassDefinition struct {
	DefinitionBase
}

type UsingDefinition struct {
	DefinitionBase
}
