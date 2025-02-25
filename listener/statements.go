package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// statements: # StatementsInitial
func (l *CaliburnListener) ExitStatementsInitial(c *parsing.StatementsInitialContext) {
	Push(l, []ast.Statement{})
}

// statements: statements statement # StatementsAppend
func (l *CaliburnListener) ExitStatementsAppend(c *parsing.StatementsAppendContext) {
	Push(l, append(Pop[[]ast.Statement](l), Pop[ast.Statement](l)))
}

// inline_statement: # InlineStatementsInitial
func (l *CaliburnListener) ExitInlineStatementsInitial(c *parsing.InlineStatementsInitialContext) {
	Push(l, []ast.InlineStatement{})
}

// inline_statement: inline_statements inline_statement # InlineStatementsAppend
func (l *CaliburnListener) ExitInlineStatementsAppend(c *parsing.InlineStatementsAppendContext) {
	Push(l, append(Pop[[]ast.InlineStatement](l), Pop[ast.InlineStatement](l)))
}

// return_statement: RETURN expression Terminator # ReturnStatement
func (l *CaliburnListener) ExitReturnStatement(c *parsing.ReturnStatementContext) {
	Push(l, ast.NewReturnStatement(Pop[ast.Expression](l)))
}

// return_statement: BREAK Terminator # BreakStatement
func (l *CaliburnListener) ExitBreakStatement(c *parsing.BreakStatementContext) {
	Push(l, ast.NewBreakStatement())
}

// return_statement: CONTINUE Terminator # ContinueStatement
func (l *CaliburnListener) ExitContinueStatement(c *parsing.ContinueStatementContext) {
	Push(l, ast.NewContinueStatement())
}
