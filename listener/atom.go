package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// literal_atom: literal # UntypedLiteral
func (l *CaliburnListener) ExitUntypedLiteral(c *parsing.UntypedLiteralContext) {
	l.Pop(1)
	Push(l, ast.NewUntypedLiteral(Dequeue[ast.LiteralToken](l)))
}

// literal_atom: identifier literal # TypedLiteral
func (l *CaliburnListener) ExitTypedLiteral(c *parsing.TypedLiteralContext) {
	l.Pop(2)
	Push(l, ast.NewTypedLiteral(Dequeue[ast.TypeExpression](l), Dequeue[ast.LiteralToken](l)))
}

// literal: LiteralToken
func (l *CaliburnListener) ExitLiteral(c *parsing.LiteralContext) {
	Push(l, ast.NewLiteralToken(c.GetVal().GetText()))
}

// identifier: IdentifierToken
func (l *CaliburnListener) ExitIdentifier(c *parsing.IdentifierContext) {
	Push(l, ast.NewIdentifer(c.GetVal().GetText()))
}
