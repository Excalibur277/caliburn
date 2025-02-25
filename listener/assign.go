package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// assign_statement: assign_declarations ASSIGN expressions Terminator # AssignStatement
func (l *CaliburnListener) ExitAssignStatement(c *parsing.AssignStatementContext) {
	Push(l, ast.NewAssignStatement(Pop[[]ast.AssignDeclaration](l), Pop[[]ast.Expression](l)))
}

// assign_statement: assign_expressions op = ( OP_ADD | OP_SUB | OP_MUL | OP_DIV | OP_MOD | OP_POW | OP_ROOT | OP_OR | OP_XOR | OP_AND | OP_LSHIFT | OP_RSHIFT) ASSIGN expressions Terminator # AssignOperationStatement
func (l *CaliburnListener) ExitAssignOperationStatement(c *parsing.AssignOperationStatementContext) {
	Push(l, ast.NewAssignOperationStatement(Pop[[]ast.AssignExpression](l), ast.NewBinaryOperation(c.GetOp()), Pop[[]ast.Expression](l)))
}

// assign_expressions: assign_expression # AssignExpressionsInitial
func (l *CaliburnListener) ExitAssignExpressionsInitial(c *parsing.AssignExpressionsInitialContext) {
	Push(l, []ast.AssignExpression{Pop[ast.AssignExpression](l)})
}

// assign_expressions: assign_expressions COMMA assign_expression # AssignExpressionsAppend
func (l *CaliburnListener) ExitAssignExpressionsAppend(c *parsing.AssignExpressionsAppendContext) {
	Push(l, append(Pop[[]ast.AssignExpression](l), Pop[ast.AssignExpression](l)))
}

// aliasable_assign_expressions: aliasable_assign_expression # AliasableAssignExpressionsInitial
func (l *CaliburnListener) ExitAliasableAssignExpressionsInitial(c *parsing.AliasableAssignExpressionsInitialContext) {
	Push(l, []ast.AliasableAssignExpression{Pop[ast.AliasableAssignExpression](l)})
}

// aliasable_assign_expressions: aliasable_assign_expressions COMMA aliasable_assign_expression # AliasableAssignExpressionsAppend
func (l *CaliburnListener) ExitAliasableAssignExpressionsAppend(c *parsing.AliasableAssignExpressionsAppendContext) {
	Push(l, append(Pop[[]ast.AliasableAssignExpression](l), Pop[ast.AliasableAssignExpression](l)))
}

// aliasable_assign_expression: assign_expression COLON identifier # AliasedAssignExpression
func (l *CaliburnListener) ExitAliasedAssignExpression(c *parsing.AliasedAssignExpressionContext) {
	Push(l, ast.NewAliasedAssignExpression(Pop[ast.AssignExpression](l), Pop[ast.Identifier](l)))
}

// aliasable_assign_expression: assign_expression # UnaliasedAssignExpression
func (l *CaliburnListener) ExitUnaliasedAssignExpression(c *parsing.UnaliasedAssignExpressionContext) {
	Push(l, ast.NewUnaliasedAssignExpression(Pop[ast.AssignExpression](l)))
}

// assign_expression: expression # ExpressionAssignExpression
func (l *CaliburnListener) ExitExpressionAssignExpression(c *parsing.ExpressionAssignExpressionContext) {
	Push(l, ast.NewExpressionAssignExpression(Pop[ast.Expression](l)))
}

// assign_expression: L_C_BRACK aliasable_assign_expressions R_C_BRACK # StructDestrutureAssignExpression
func (l *CaliburnListener) ExitStructDestrutureAssignExpression(c *parsing.StructDestrutureAssignExpressionContext) {
	Push(l, ast.NewStructDestrutureAssignExpression(Pop[[]ast.AliasableAssignExpression](l)))
}

// assign_declarations: assign_declaration # AssignDeclarationsInitial
func (l *CaliburnListener) ExitAssignDeclarationsInitial(c *parsing.AssignDeclarationsInitialContext) {
	Push(l, []ast.AssignDeclaration{Pop[ast.AssignDeclaration](l)})
}

// assign_declarations: assign_declarations COMMA assign_statement # AssignDeclarationsAppend
func (l *CaliburnListener) ExitAssignDeclarationsAppend(c *parsing.AssignDeclarationsAppendContext) {
	Push(l, append(Pop[[]ast.AssignDeclaration](l), Pop[ast.AssignDeclaration](l)))
}

// aliasable_assign_declarations: aliasable_assign_declaration # AliasableAssignDeclarationsInitial
func (l *CaliburnListener) ExitAliasableAssignDeclarationsInitial(c *parsing.AliasableAssignDeclarationsInitialContext) {
	Push(l, []ast.AliasableAssignDeclaration{Pop[ast.AliasableAssignDeclaration](l)})
}

// aliasable_assign_declarations: aliasable_assign_declarations COMMA aliasable_assign_declaration # AliasableAssignDeclarationsAppend
func (l *CaliburnListener) ExitAliasableAssignDeclarationsAppend(c *parsing.AliasableAssignDeclarationsAppendContext) {
	Push(l, []ast.AliasableAssignDeclaration{Pop[ast.AliasableAssignDeclaration](l)})
}

// aliasable_assign_declaration: assign_declaration COLON identifier # AliasedAssignDeclaration
func (l *CaliburnListener) ExitAliasedAssignDeclaration(c *parsing.AliasedAssignDeclarationContext) {
	Push(l, ast.NewAliasedAssignDeclaration(Pop[ast.AssignDeclaration](l), Pop[ast.Identifier](l)))
}

// aliasable_assign_declaration: assign_declaration # UnaliasedAssignDeclaration
func (l *CaliburnListener) ExitUnaliasedAssignDeclaration(c *parsing.UnaliasedAssignDeclarationContext) {
	Push(l, ast.NewUnaliasedAssignDeclaration(Pop[ast.AssignDeclaration](l)))
}

// assign_declaration: expression # ExpressionAssignDeclaration
func (l *CaliburnListener) ExitExpressionAssignDeclaration(c *parsing.ExpressionAssignDeclarationContext) {
	Push(l, ast.NewExpressionAssignDeclaration(Pop[ast.Expression](l)))
}

// assign_declaration: typed_assign_declaration # TypedAssignDeclarationDeclaration
func (l *CaliburnListener) ExitTypedAssignDeclarationDeclaration(c *parsing.TypedAssignDeclarationDeclarationContext) {
	Push(l, ast.NewTypedAssignDeclarationDeclaration(Pop[ast.TypedAssignDeclaration](l)))
}

// assign_declaration: VAR untyped_assign_declaration # UntypedAssignDeclarationDeclaration
func (l *CaliburnListener) ExitUntypedAssignDeclarationDeclaration(c *parsing.UntypedAssignDeclarationDeclarationContext) {
	Push(l, ast.NewUntypedAssignDeclarationDeclaration(Pop[ast.UntypedAssignDeclaration](l)))
}

// assign_declaration: L_C_BRACK aliasable_assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
func (l *CaliburnListener) ExitStructDestrutureAssignDeclaration(c *parsing.StructDestrutureAssignDeclarationContext) {
	Push(l, ast.NewStructDestrutureAssignDeclaration(Pop[[]ast.AliasableAssignDeclaration](l)))
}

// typed_assign_declaration: type_expression untyped_assign_declaration # TypedAssignDeclaration
func (l *CaliburnListener) ExitTypedAssignDeclaration(c *parsing.TypedAssignDeclarationContext) {
	Push(l, ast.NewTypedAssignDeclaration(Pop[ast.TypeExpression](l), Pop[ast.UntypedAssignDeclaration](l)))
}

// aliasable_untyped_assign_declarations: aliasable_untyped_assign_declaration # AliasableUntypedAssignDeclarationsInitial
func (l *CaliburnListener) ExitAliasableUntypedAssignDeclarationsInitial(c *parsing.AliasableUntypedAssignDeclarationsInitialContext) {
	Push(l, []ast.AliasableUntypedAssignDeclaration{Pop[ast.AliasableUntypedAssignDeclaration](l)})
}

// aliasable_untyped_assign_declarations: aliasable_untyped_assign_declarations COMMA aliasable_untyped_assign_declaration # AliasableUntypedAssignDeclarationsAppend
func (l *CaliburnListener) ExitAliasableUntypedAssignDeclarationsAppend(c *parsing.AliasableUntypedAssignDeclarationsAppendContext) {
	Push(l, append(Pop[[]ast.AliasableUntypedAssignDeclaration](l), Pop[ast.AliasableUntypedAssignDeclaration](l)))
}

// aliasable_untyped_assign_declaration: untyped_assign_declaration COLON identifier # AliasedUntypedAssignDeclaration
func (l *CaliburnListener) ExitAliasedUntypedAssignDeclaration(c *parsing.AliasedUntypedAssignDeclarationContext) {
	Push(l, ast.NewAliasedUntypedAssignDeclaration(Pop[ast.UntypedAssignDeclaration](l), Pop[ast.Identifier](l)))
}

// aliasable_untyped_assign_declaration: untyped_assign_declaration # UnaliasedUntypedAssignDeclaration
func (l *CaliburnListener) ExitUnaliasedUntypedAssignDeclaration(c *parsing.UnaliasedUntypedAssignDeclarationContext) {
	Push(l, ast.NewUnaliasedUntypedAssignDeclaration(Pop[ast.UntypedAssignDeclaration](l)))
}

// untyped_assign_declaration: var = identifier # UntypedIdentifierAssignDeclaration
func (l *CaliburnListener) ExitUntypedIdentifierAssignDeclaration(c *parsing.UntypedIdentifierAssignDeclarationContext) {
	Push(l, ast.NewUntypedIdentifierAssignDeclaration(Pop[ast.Identifier](l)))
}

// untyped_assign_declaration: L_C_BRACK aliasable_untyped_assign_declarations R_C_BRACK # UntypedStructDestrutureAssignDeclaration
func (l *CaliburnListener) ExitUntypedStructDestrutureAssignDeclaration(c *parsing.UntypedStructDestrutureAssignDeclarationContext) {
	Push(l, ast.NewUntypedStructDestrutureAssignDeclaration(Pop[[]ast.AliasableUntypedAssignDeclaration](l)))
}
