package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// literal_atom: literal # UntypedLiteral
func (l *CaliburnListener) ExitUntypedLiteral(c *parsing.UntypedLiteralContext) {
	Push(l, ast.NewUntypedLiteral(Pop[ast.Literal](l)))
}

// literal_atom: identifier literal # TypedLiteral
func (l *CaliburnListener) ExitTypedLiteral(c *parsing.TypedLiteralContext) {
	Push(l, ast.NewTypedLiteral(Pop[ast.TypeExpression](l), Pop[ast.Literal](l)))
}

// literal: LiteralToken
func (l *CaliburnListener) ExitLiteral(c *parsing.LiteralContext) {
	Push(l, ast.NewLiteral(c.GetVal().GetText()))
}

// identifier: IdentifierToken
func (l *CaliburnListener) ExitIdentifier(c *parsing.IdentifierContext) {
	Push(l, ast.NewIdentifer(c.GetVal().GetText()))
}
