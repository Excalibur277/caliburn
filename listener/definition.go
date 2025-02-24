package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

func (l *CaliburnListener) ExitDefinitionsInitial(c *parsing.DefinitionsInitialContext) {
	Push(l, []ast.Definition{})
}

func (l *CaliburnListener) ExitDefinitionsAppend(c *parsing.DefinitionsAppendContext) {
	Push(l, append(Pop[[]ast.Definition](l), Pop[ast.Definition](l)))
}

func (l *CaliburnListener) ExitDefinition(c *parsing.DefinitionContext) {
	// Nothing
}

func (l *CaliburnListener) ExitFunctionDefinition(c *parsing.FunctionDefinitionContext) {
	Push(l, ast.NewFunctionDefinition(
		Pop[ast.FunctionType](l),
		Pop[ast.Identifier](l),
		Pop[[]ast.Parameter](l),
		Pop[ast.TypeExpression](l),
		Pop[ast.Block](l),
	))
}

func (l *CaliburnListener) ExitFunctionDefinitionNoReturnType(c *parsing.FunctionDefinitionNoReturnTypeContext) {
	Push(l, ast.NewFunctionDefinition(
		Pop[ast.FunctionType](l),
		Pop[ast.Identifier](l),
		Pop[[]ast.Parameter](l),
		ast.NewEmptyType(),
		Pop[ast.Block](l),
	))
}
