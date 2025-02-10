// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// BaseCaliburnParserListener is a complete listener for a parse tree produced by CaliburnParser.
type BaseCaliburnParserListener struct{}

var _ CaliburnParserListener = &BaseCaliburnParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCaliburnParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCaliburnParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCaliburnParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCaliburnParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScoped_block is called when production scoped_block is entered.
func (s *BaseCaliburnParserListener) EnterScoped_block(ctx *Scoped_blockContext) {}

// ExitScoped_block is called when production scoped_block is exited.
func (s *BaseCaliburnParserListener) ExitScoped_block(ctx *Scoped_blockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCaliburnParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCaliburnParserListener) ExitStatement(ctx *StatementContext) {}

// EnterControl_statement is called when production control_statement is entered.
func (s *BaseCaliburnParserListener) EnterControl_statement(ctx *Control_statementContext) {}

// ExitControl_statement is called when production control_statement is exited.
func (s *BaseCaliburnParserListener) ExitControl_statement(ctx *Control_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseCaliburnParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseCaliburnParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseCaliburnParserListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseCaliburnParserListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterSwitch_statement is called when production switch_statement is entered.
func (s *BaseCaliburnParserListener) EnterSwitch_statement(ctx *Switch_statementContext) {}

// ExitSwitch_statement is called when production switch_statement is exited.
func (s *BaseCaliburnParserListener) ExitSwitch_statement(ctx *Switch_statementContext) {}

// EnterInline_statement is called when production inline_statement is entered.
func (s *BaseCaliburnParserListener) EnterInline_statement(ctx *Inline_statementContext) {}

// ExitInline_statement is called when production inline_statement is exited.
func (s *BaseCaliburnParserListener) ExitInline_statement(ctx *Inline_statementContext) {}

// EnterFunction_declaration_statement is called when production function_declaration_statement is entered.
func (s *BaseCaliburnParserListener) EnterFunction_declaration_statement(ctx *Function_declaration_statementContext) {
}

// ExitFunction_declaration_statement is called when production function_declaration_statement is exited.
func (s *BaseCaliburnParserListener) ExitFunction_declaration_statement(ctx *Function_declaration_statementContext) {
}

// EnterAssign_statement is called when production assign_statement is entered.
func (s *BaseCaliburnParserListener) EnterAssign_statement(ctx *Assign_statementContext) {}

// ExitAssign_statement is called when production assign_statement is exited.
func (s *BaseCaliburnParserListener) ExitAssign_statement(ctx *Assign_statementContext) {}

// EnterAssign_declarations is called when production assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterAssign_declarations(ctx *Assign_declarationsContext) {}

// ExitAssign_declarations is called when production assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitAssign_declarations(ctx *Assign_declarationsContext) {}

// EnterDeclaredAssignDeclaration is called when production DeclaredAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterDeclaredAssignDeclaration(ctx *DeclaredAssignDeclarationContext) {
}

// ExitDeclaredAssignDeclaration is called when production DeclaredAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitDeclaredAssignDeclaration(ctx *DeclaredAssignDeclarationContext) {
}

// EnterUndeclaredAssignDeclaration is called when production UndeclaredAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUndeclaredAssignDeclaration(ctx *UndeclaredAssignDeclarationContext) {
}

// ExitUndeclaredAssignDeclaration is called when production UndeclaredAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUndeclaredAssignDeclaration(ctx *UndeclaredAssignDeclarationContext) {
}

// EnterStructDestrutureAssignDeclaration is called when production StructDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) {
}

// ExitStructDestrutureAssignDeclaration is called when production StructDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) {
}

// EnterTupleDestrutureAssignDeclaration is called when production TupleDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) {
}

// ExitTupleDestrutureAssignDeclaration is called when production TupleDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) {
}

// EnterDeclared_assign_declarations is called when production declared_assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterDeclared_assign_declarations(ctx *Declared_assign_declarationsContext) {
}

// ExitDeclared_assign_declarations is called when production declared_assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitDeclared_assign_declarations(ctx *Declared_assign_declarationsContext) {
}

// EnterTypedDeclaredAssignDeclaration is called when production TypedDeclaredAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedDeclaredAssignDeclaration(ctx *TypedDeclaredAssignDeclarationContext) {
}

// ExitTypedDeclaredAssignDeclaration is called when production TypedDeclaredAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedDeclaredAssignDeclaration(ctx *TypedDeclaredAssignDeclarationContext) {
}

// EnterUntypedDeclaredAssignDeclaration is called when production UntypedDeclaredAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedDeclaredAssignDeclaration(ctx *UntypedDeclaredAssignDeclarationContext) {
}

// ExitUntypedDeclaredAssignDeclaration is called when production UntypedDeclaredAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedDeclaredAssignDeclaration(ctx *UntypedDeclaredAssignDeclarationContext) {
}

// EnterUndeclared_assign_declarations is called when production undeclared_assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterUndeclared_assign_declarations(ctx *Undeclared_assign_declarationsContext) {
}

// ExitUndeclared_assign_declarations is called when production undeclared_assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitUndeclared_assign_declarations(ctx *Undeclared_assign_declarationsContext) {
}

// EnterUndeclaredAtomAssignDeclaration is called when production UndeclaredAtomAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUndeclaredAtomAssignDeclaration(ctx *UndeclaredAtomAssignDeclarationContext) {
}

// ExitUndeclaredAtomAssignDeclaration is called when production UndeclaredAtomAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUndeclaredAtomAssignDeclaration(ctx *UndeclaredAtomAssignDeclarationContext) {
}

// EnterUndeclaredStructDestrutureAssignDeclaration is called when production UndeclaredStructDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUndeclaredStructDestrutureAssignDeclaration(ctx *UndeclaredStructDestrutureAssignDeclarationContext) {
}

// ExitUndeclaredStructDestrutureAssignDeclaration is called when production UndeclaredStructDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUndeclaredStructDestrutureAssignDeclaration(ctx *UndeclaredStructDestrutureAssignDeclarationContext) {
}

// EnterUndeclaredTupleDestrutureAssignDeclaration is called when production UndeclaredTupleDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUndeclaredTupleDestrutureAssignDeclaration(ctx *UndeclaredTupleDestrutureAssignDeclarationContext) {
}

// ExitUndeclaredTupleDestrutureAssignDeclaration is called when production UndeclaredTupleDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUndeclaredTupleDestrutureAssignDeclaration(ctx *UndeclaredTupleDestrutureAssignDeclarationContext) {
}

// EnterOperatorAssignment is called when production OperatorAssignment is entered.
func (s *BaseCaliburnParserListener) EnterOperatorAssignment(ctx *OperatorAssignmentContext) {}

// ExitOperatorAssignment is called when production OperatorAssignment is exited.
func (s *BaseCaliburnParserListener) ExitOperatorAssignment(ctx *OperatorAssignmentContext) {}

// EnterExpression_statement is called when production expression_statement is entered.
func (s *BaseCaliburnParserListener) EnterExpression_statement(ctx *Expression_statementContext) {}

// ExitExpression_statement is called when production expression_statement is exited.
func (s *BaseCaliburnParserListener) ExitExpression_statement(ctx *Expression_statementContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseCaliburnParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseCaliburnParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterExponentExpression is called when production ExponentExpression is entered.
func (s *BaseCaliburnParserListener) EnterExponentExpression(ctx *ExponentExpressionContext) {}

// ExitExponentExpression is called when production ExponentExpression is exited.
func (s *BaseCaliburnParserListener) ExitExponentExpression(ctx *ExponentExpressionContext) {}

// EnterAdditionExpression is called when production AdditionExpression is entered.
func (s *BaseCaliburnParserListener) EnterAdditionExpression(ctx *AdditionExpressionContext) {}

// ExitAdditionExpression is called when production AdditionExpression is exited.
func (s *BaseCaliburnParserListener) ExitAdditionExpression(ctx *AdditionExpressionContext) {}

// EnterBracketedExpression is called when production BracketedExpression is entered.
func (s *BaseCaliburnParserListener) EnterBracketedExpression(ctx *BracketedExpressionContext) {}

// ExitBracketedExpression is called when production BracketedExpression is exited.
func (s *BaseCaliburnParserListener) ExitBracketedExpression(ctx *BracketedExpressionContext) {}

// EnterSliceExpression is called when production SliceExpression is entered.
func (s *BaseCaliburnParserListener) EnterSliceExpression(ctx *SliceExpressionContext) {}

// ExitSliceExpression is called when production SliceExpression is exited.
func (s *BaseCaliburnParserListener) ExitSliceExpression(ctx *SliceExpressionContext) {}

// EnterTupleExpression is called when production TupleExpression is entered.
func (s *BaseCaliburnParserListener) EnterTupleExpression(ctx *TupleExpressionContext) {}

// ExitTupleExpression is called when production TupleExpression is exited.
func (s *BaseCaliburnParserListener) ExitTupleExpression(ctx *TupleExpressionContext) {}

// EnterIndexExpression is called when production IndexExpression is entered.
func (s *BaseCaliburnParserListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production IndexExpression is exited.
func (s *BaseCaliburnParserListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterUnaryExpression is called when production UnaryExpression is entered.
func (s *BaseCaliburnParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production UnaryExpression is exited.
func (s *BaseCaliburnParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterAtomExpression is called when production AtomExpression is entered.
func (s *BaseCaliburnParserListener) EnterAtomExpression(ctx *AtomExpressionContext) {}

// ExitAtomExpression is called when production AtomExpression is exited.
func (s *BaseCaliburnParserListener) ExitAtomExpression(ctx *AtomExpressionContext) {}

// EnterShiftExpression is called when production ShiftExpression is entered.
func (s *BaseCaliburnParserListener) EnterShiftExpression(ctx *ShiftExpressionContext) {}

// ExitShiftExpression is called when production ShiftExpression is exited.
func (s *BaseCaliburnParserListener) ExitShiftExpression(ctx *ShiftExpressionContext) {}

// EnterOrExpression is called when production OrExpression is entered.
func (s *BaseCaliburnParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production OrExpression is exited.
func (s *BaseCaliburnParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAccessExpression is called when production AccessExpression is entered.
func (s *BaseCaliburnParserListener) EnterAccessExpression(ctx *AccessExpressionContext) {}

// ExitAccessExpression is called when production AccessExpression is exited.
func (s *BaseCaliburnParserListener) ExitAccessExpression(ctx *AccessExpressionContext) {}

// EnterComparisonExpression is called when production ComparisonExpression is entered.
func (s *BaseCaliburnParserListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production ComparisonExpression is exited.
func (s *BaseCaliburnParserListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterXOrExpression is called when production XOrExpression is entered.
func (s *BaseCaliburnParserListener) EnterXOrExpression(ctx *XOrExpressionContext) {}

// ExitXOrExpression is called when production XOrExpression is exited.
func (s *BaseCaliburnParserListener) ExitXOrExpression(ctx *XOrExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseCaliburnParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseCaliburnParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterStructExpression is called when production StructExpression is entered.
func (s *BaseCaliburnParserListener) EnterStructExpression(ctx *StructExpressionContext) {}

// ExitStructExpression is called when production StructExpression is exited.
func (s *BaseCaliburnParserListener) ExitStructExpression(ctx *StructExpressionContext) {}

// EnterAndExpression is called when production AndExpression is entered.
func (s *BaseCaliburnParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production AndExpression is exited.
func (s *BaseCaliburnParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterCastExpression is called when production CastExpression is entered.
func (s *BaseCaliburnParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production CastExpression is exited.
func (s *BaseCaliburnParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseCaliburnParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseCaliburnParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterCallExpression is called when production CallExpression is entered.
func (s *BaseCaliburnParserListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production CallExpression is exited.
func (s *BaseCaliburnParserListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterFunction_expression is called when production function_expression is entered.
func (s *BaseCaliburnParserListener) EnterFunction_expression(ctx *Function_expressionContext) {}

// ExitFunction_expression is called when production function_expression is exited.
func (s *BaseCaliburnParserListener) ExitFunction_expression(ctx *Function_expressionContext) {}

// EnterStruct_expression is called when production struct_expression is entered.
func (s *BaseCaliburnParserListener) EnterStruct_expression(ctx *Struct_expressionContext) {}

// ExitStruct_expression is called when production struct_expression is exited.
func (s *BaseCaliburnParserListener) ExitStruct_expression(ctx *Struct_expressionContext) {}

// EnterTuple_expression is called when production tuple_expression is entered.
func (s *BaseCaliburnParserListener) EnterTuple_expression(ctx *Tuple_expressionContext) {}

// ExitTuple_expression is called when production tuple_expression is exited.
func (s *BaseCaliburnParserListener) ExitTuple_expression(ctx *Tuple_expressionContext) {}

// EnterExpression_atom is called when production expression_atom is entered.
func (s *BaseCaliburnParserListener) EnterExpression_atom(ctx *Expression_atomContext) {}

// ExitExpression_atom is called when production expression_atom is exited.
func (s *BaseCaliburnParserListener) ExitExpression_atom(ctx *Expression_atomContext) {}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BaseCaliburnParserListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BaseCaliburnParserListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterLiteral_atom is called when production literal_atom is entered.
func (s *BaseCaliburnParserListener) EnterLiteral_atom(ctx *Literal_atomContext) {}

// ExitLiteral_atom is called when production literal_atom is exited.
func (s *BaseCaliburnParserListener) ExitLiteral_atom(ctx *Literal_atomContext) {}

// EnterUntyped_literal_atom is called when production untyped_literal_atom is entered.
func (s *BaseCaliburnParserListener) EnterUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// ExitUntyped_literal_atom is called when production untyped_literal_atom is exited.
func (s *BaseCaliburnParserListener) ExitUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// EnterTyped_literal_atom is called when production typed_literal_atom is entered.
func (s *BaseCaliburnParserListener) EnterTyped_literal_atom(ctx *Typed_literal_atomContext) {}

// ExitTyped_literal_atom is called when production typed_literal_atom is exited.
func (s *BaseCaliburnParserListener) ExitTyped_literal_atom(ctx *Typed_literal_atomContext) {}
