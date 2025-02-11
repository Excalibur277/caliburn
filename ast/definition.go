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
	DefinedFunction Function
}

type TypeDefinition struct {
	DefinitionBase
	DefinedType Type
}

type ClassDefinition struct {
	DefinitionBase
	DefinedClass Class
}

type UsingDefinition struct {
	DefinitionBase
	DefinedUsing Using
}
