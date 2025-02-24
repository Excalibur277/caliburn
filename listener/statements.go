package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

func (l *CaliburnListener) ExitStatementsInitial(c *parsing.StatementsInitialContext) {
	Push(l, []ast.Statement{})
}

func (l *CaliburnListener) ExitStatementsAppend(c *parsing.StatementsAppendContext) {
	Push(l, append(Pop[[]ast.Statement](l), Pop[ast.Statement](l)))
}

func (l *CaliburnListener) ExitInlineStatementsInitial(c *parsing.InlineStatementsInitialContext) {
	Push(l, []ast.InlineStatement{})
}

func (l *CaliburnListener) ExitInlineStatementsAppend(c *parsing.InlineStatementsAppendContext) {
	Push(l, append(Pop[[]ast.InlineStatement](l), Pop[ast.InlineStatement](l)))
}

func (l *CaliburnListener) ExitReturnStatement(c *parsing.ReturnStatementContext) {
	Push(l, ast.NewReturnStatement(Pop[ast.Expression](l)))
}

func (l *CaliburnListener) ExitBreakStatement(c *parsing.BreakStatementContext) {
	Push(l, ast.NewBreakStatement())
}

func (l *CaliburnListener) ExitContinueStatement(c *parsing.ContinueStatementContext) {
	Push(l, ast.NewContinueStatement())
}
