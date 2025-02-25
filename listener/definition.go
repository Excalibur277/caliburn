package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// definitions: # DefinitionsInitial
func (l *CaliburnListener) ExitDefinitionsInitial(c *parsing.DefinitionsInitialContext) {
	Push(l, []ast.Definition{})
}

// definitions: definitions definition # DefinitionsAppend
func (l *CaliburnListener) ExitDefinitionsAppend(c *parsing.DefinitionsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.Definition](l), Dequeue[ast.Definition](l)))
}

// function_definition: function_type identifier L_PAREN parameters R_PAREN type_expression block # FunctionDefinition
func (l *CaliburnListener) ExitFunctionDefinition(c *parsing.FunctionDefinitionContext) {
	l.Pop(5)
	Push(l, ast.NewFunctionDefinition(
		Dequeue[ast.FunctionType](l),
		Dequeue[ast.Identifier](l),
		Dequeue[[]ast.Parameter](l),
		Dequeue[ast.TypeExpression](l),
		Dequeue[ast.Block](l),
	))
}

// function_definition: function_type identifier L_PAREN parameters R_PAREN block                 # FunctionDefinitionNoReturnType
func (l *CaliburnListener) ExitFunctionDefinitionNoReturnType(c *parsing.FunctionDefinitionNoReturnTypeContext) {
	l.Pop(4)
	Push(l, ast.NewFunctionDefinition(
		Dequeue[ast.FunctionType](l),
		Dequeue[ast.Identifier](l),
		Dequeue[[]ast.Parameter](l),
		ast.NewEmptyType(),
		Dequeue[ast.Block](l),
	))
}
