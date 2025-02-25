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
	Push(l, append(Pop[[]ast.Definition](l), Pop[ast.Definition](l)))
}

// function_definition: function_type identifier L_PAREN parameters R_PAREN type_expression block # FunctionDefinition
func (l *CaliburnListener) ExitFunctionDefinition(c *parsing.FunctionDefinitionContext) {
	Push(l, ast.NewFunctionDefinition(
		Pop[ast.FunctionType](l),
		Pop[ast.Identifier](l),
		Pop[[]ast.Parameter](l),
		Pop[ast.TypeExpression](l),
		Pop[ast.Block](l),
	))
}

// function_definition: function_type identifier L_PAREN parameters R_PAREN block                 # FunctionDefinitionNoReturnType
func (l *CaliburnListener) ExitFunctionDefinitionNoReturnType(c *parsing.FunctionDefinitionNoReturnTypeContext) {
	Push(l, ast.NewFunctionDefinition(
		Pop[ast.FunctionType](l),
		Pop[ast.Identifier](l),
		Pop[[]ast.Parameter](l),
		ast.NewEmptyType(),
		Pop[ast.Block](l),
	))
}
