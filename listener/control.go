package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

func (l *CaliburnListener) ExitIfStatement(c *parsing.IfStatementContext) {
	Push(l, ast.NewIfStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.Block](l)))
}

func (l *CaliburnListener) ExitIfElseStatement(c *parsing.IfElseStatementContext) {
	Push(l, ast.NewIfElseStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.Block](l), Pop[ast.ElseStatement](l)))
}
