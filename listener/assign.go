package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// assign_statement: assign_declarations ASSIGN expressions Terminator # AssignStatement
func (l *CaliburnListener) ExitAssignStatement(c *parsing.AssignStatementContext) {
	l.Pop(2)
	Push(l, ast.NewAssignStatement(Dequeue[[]ast.AssignDeclaration](l), Dequeue[[]ast.Expression](l)))
}

// assign_statement: assign_expressions op = ( OP_ADD | OP_SUB | OP_MUL | OP_DIV | OP_MOD | OP_POW | OP_ROOT | OP_OR | OP_XOR | OP_AND | OP_LSHIFT | OP_RSHIFT) ASSIGN expressions Terminator # AssignOperationStatement
func (l *CaliburnListener) ExitAssignOperationStatement(c *parsing.AssignOperationStatementContext) {
	l.Pop(2)
	Push(l, ast.NewAssignOperationStatement(Dequeue[[]ast.AssignExpression](l), ast.NewBinaryOperation(c.GetOp()), Dequeue[[]ast.Expression](l)))
}

// assign_expressions: assign_expression # AssignExpressionsInitial
func (l *CaliburnListener) ExitAssignExpressionsInitial(c *parsing.AssignExpressionsInitialContext) {
	l.Pop(1)
	Push(l, []ast.AssignExpression{Dequeue[ast.AssignExpression](l)})
}

// assign_expressions: assign_expressions COMMA assign_expression # AssignExpressionsAppend
func (l *CaliburnListener) ExitAssignExpressionsAppend(c *parsing.AssignExpressionsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.AssignExpression](l), Dequeue[ast.AssignExpression](l)))
}

// aliasable_assign_expressions: aliasable_assign_expression # AliasableAssignExpressionsInitial
func (l *CaliburnListener) ExitAliasableAssignExpressionsInitial(c *parsing.AliasableAssignExpressionsInitialContext) {
	l.Pop(1)
	Push(l, []ast.AliasableAssignExpression{Dequeue[ast.AliasableAssignExpression](l)})
}

// aliasable_assign_expressions: aliasable_assign_expressions COMMA aliasable_assign_expression # AliasableAssignExpressionsAppend
func (l *CaliburnListener) ExitAliasableAssignExpressionsAppend(c *parsing.AliasableAssignExpressionsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.AliasableAssignExpression](l), Dequeue[ast.AliasableAssignExpression](l)))
}

// aliasable_assign_expression: assign_expression COLON identifier # AliasedAssignExpression
func (l *CaliburnListener) ExitAliasedAssignExpression(c *parsing.AliasedAssignExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewAliasedAssignExpression(Dequeue[ast.AssignExpression](l), Dequeue[ast.Identifier](l)))
}

// aliasable_assign_expression: assign_expression # UnaliasedAssignExpression
func (l *CaliburnListener) ExitUnaliasedAssignExpression(c *parsing.UnaliasedAssignExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewUnaliasedAssignExpression(Dequeue[ast.AssignExpression](l)))
}

// assign_expression: expression # ExpressionAssignExpression
func (l *CaliburnListener) ExitExpressionAssignExpression(c *parsing.ExpressionAssignExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewExpressionAssignExpression(Dequeue[ast.Expression](l)))
}

// assign_expression: L_C_BRACK aliasable_assign_expressions R_C_BRACK # StructDestrutureAssignExpression
func (l *CaliburnListener) ExitStructDestrutureAssignExpression(c *parsing.StructDestrutureAssignExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewStructDestrutureAssignExpression(Dequeue[[]ast.AliasableAssignExpression](l)))
}

// assign_declarations: assign_declaration # AssignDeclarationsInitial
func (l *CaliburnListener) ExitAssignDeclarationsInitial(c *parsing.AssignDeclarationsInitialContext) {
	l.Pop(1)
	Push(l, []ast.AssignDeclaration{Dequeue[ast.AssignDeclaration](l)})
}

// assign_declarations: assign_declarations COMMA assign_statement # AssignDeclarationsAppend
func (l *CaliburnListener) ExitAssignDeclarationsAppend(c *parsing.AssignDeclarationsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.AssignDeclaration](l), Dequeue[ast.AssignDeclaration](l)))
}

// aliasable_assign_declarations: aliasable_assign_declaration # AliasableAssignDeclarationsInitial
func (l *CaliburnListener) ExitAliasableAssignDeclarationsInitial(c *parsing.AliasableAssignDeclarationsInitialContext) {
	l.Pop(1)
	Push(l, []ast.AliasableAssignDeclaration{Dequeue[ast.AliasableAssignDeclaration](l)})
}

// aliasable_assign_declarations: aliasable_assign_declarations COMMA aliasable_assign_declaration # AliasableAssignDeclarationsAppend
func (l *CaliburnListener) ExitAliasableAssignDeclarationsAppend(c *parsing.AliasableAssignDeclarationsAppendContext) {
	l.Pop(1)
	Push(l, []ast.AliasableAssignDeclaration{Dequeue[ast.AliasableAssignDeclaration](l)})
}

// aliasable_assign_declaration: assign_declaration COLON identifier # AliasedAssignDeclaration
func (l *CaliburnListener) ExitAliasedAssignDeclaration(c *parsing.AliasedAssignDeclarationContext) {
	l.Pop(2)
	Push(l, ast.NewAliasedAssignDeclaration(Dequeue[ast.AssignDeclaration](l), Dequeue[ast.Identifier](l)))
}

// aliasable_assign_declaration: assign_declaration # UnaliasedAssignDeclaration
func (l *CaliburnListener) ExitUnaliasedAssignDeclaration(c *parsing.UnaliasedAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewUnaliasedAssignDeclaration(Dequeue[ast.AssignDeclaration](l)))
}

// assign_declaration: expression # ExpressionAssignDeclaration
func (l *CaliburnListener) ExitExpressionAssignDeclaration(c *parsing.ExpressionAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewExpressionAssignDeclaration(Dequeue[ast.Expression](l)))
}

// assign_declaration: typed_assign_declaration # TypedAssignDeclarationDeclaration
func (l *CaliburnListener) ExitTypedAssignDeclarationDeclaration(c *parsing.TypedAssignDeclarationDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewTypedAssignDeclarationDeclaration(Dequeue[ast.TypedAssignDeclaration](l)))
}

// assign_declaration: VAR untyped_assign_declaration # UntypedAssignDeclarationDeclaration
func (l *CaliburnListener) ExitUntypedAssignDeclarationDeclaration(c *parsing.UntypedAssignDeclarationDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewUntypedAssignDeclarationDeclaration(Dequeue[ast.UntypedAssignDeclaration](l)))
}

// assign_declaration: L_C_BRACK aliasable_assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
func (l *CaliburnListener) ExitStructDestrutureAssignDeclaration(c *parsing.StructDestrutureAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewStructDestrutureAssignDeclaration(Dequeue[[]ast.AliasableAssignDeclaration](l)))
}

// typed_assign_declaration: type_expression untyped_assign_declaration # TypedAssignDeclaration
func (l *CaliburnListener) ExitTypedAssignDeclaration(c *parsing.TypedAssignDeclarationContext) {
	l.Pop(2)
	Push(l, ast.NewTypedAssignDeclaration(Dequeue[ast.TypeExpression](l), Dequeue[ast.UntypedAssignDeclaration](l)))
}

// aliasable_untyped_assign_declarations: aliasable_untyped_assign_declaration # AliasableUntypedAssignDeclarationsInitial
func (l *CaliburnListener) ExitAliasableUntypedAssignDeclarationsInitial(c *parsing.AliasableUntypedAssignDeclarationsInitialContext) {
	l.Pop(1)
	Push(l, []ast.AliasableUntypedAssignDeclaration{Dequeue[ast.AliasableUntypedAssignDeclaration](l)})
}

// aliasable_untyped_assign_declarations: aliasable_untyped_assign_declarations COMMA aliasable_untyped_assign_declaration # AliasableUntypedAssignDeclarationsAppend
func (l *CaliburnListener) ExitAliasableUntypedAssignDeclarationsAppend(c *parsing.AliasableUntypedAssignDeclarationsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.AliasableUntypedAssignDeclaration](l), Dequeue[ast.AliasableUntypedAssignDeclaration](l)))
}

// aliasable_untyped_assign_declaration: untyped_assign_declaration COLON identifier # AliasedUntypedAssignDeclaration
func (l *CaliburnListener) ExitAliasedUntypedAssignDeclaration(c *parsing.AliasedUntypedAssignDeclarationContext) {
	l.Pop(2)
	Push(l, ast.NewAliasedUntypedAssignDeclaration(Dequeue[ast.UntypedAssignDeclaration](l), Dequeue[ast.Identifier](l)))
}

// aliasable_untyped_assign_declaration: untyped_assign_declaration # UnaliasedUntypedAssignDeclaration
func (l *CaliburnListener) ExitUnaliasedUntypedAssignDeclaration(c *parsing.UnaliasedUntypedAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewUnaliasedUntypedAssignDeclaration(Dequeue[ast.UntypedAssignDeclaration](l)))
}

// untyped_assign_declaration: var = identifier # UntypedIdentifierAssignDeclaration
func (l *CaliburnListener) ExitUntypedIdentifierAssignDeclaration(c *parsing.UntypedIdentifierAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewUntypedIdentifierAssignDeclaration(Dequeue[ast.Identifier](l)))
}

// untyped_assign_declaration: L_C_BRACK aliasable_untyped_assign_declarations R_C_BRACK # UntypedStructDestrutureAssignDeclaration
func (l *CaliburnListener) ExitUntypedStructDestrutureAssignDeclaration(c *parsing.UntypedStructDestrutureAssignDeclarationContext) {
	l.Pop(1)
	Push(l, ast.NewUntypedStructDestrutureAssignDeclaration(Dequeue[[]ast.AliasableUntypedAssignDeclaration](l)))
}
