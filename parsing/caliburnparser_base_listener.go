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

// EnterModule is called when production module is entered.
func (s *BaseCaliburnParserListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseCaliburnParserListener) ExitModule(ctx *ModuleContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseCaliburnParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseCaliburnParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterScoped_block is called when production scoped_block is entered.
func (s *BaseCaliburnParserListener) EnterScoped_block(ctx *Scoped_blockContext) {}

// ExitScoped_block is called when production scoped_block is exited.
func (s *BaseCaliburnParserListener) ExitScoped_block(ctx *Scoped_blockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCaliburnParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCaliburnParserListener) ExitStatement(ctx *StatementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseCaliburnParserListener) EnterJump_statement(ctx *Jump_statementContext) {}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseCaliburnParserListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseCaliburnParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseCaliburnParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BaseCaliburnParserListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BaseCaliburnParserListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BaseCaliburnParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BaseCaliburnParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

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

// EnterBracketedExpression is called when production BracketedExpression is entered.
func (s *BaseCaliburnParserListener) EnterBracketedExpression(ctx *BracketedExpressionContext) {}

// ExitBracketedExpression is called when production BracketedExpression is exited.
func (s *BaseCaliburnParserListener) ExitBracketedExpression(ctx *BracketedExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseCaliburnParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseCaliburnParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterSliceExpression is called when production SliceExpression is entered.
func (s *BaseCaliburnParserListener) EnterSliceExpression(ctx *SliceExpressionContext) {}

// ExitSliceExpression is called when production SliceExpression is exited.
func (s *BaseCaliburnParserListener) ExitSliceExpression(ctx *SliceExpressionContext) {}

// EnterTupleExpression is called when production TupleExpression is entered.
func (s *BaseCaliburnParserListener) EnterTupleExpression(ctx *TupleExpressionContext) {}

// ExitTupleExpression is called when production TupleExpression is exited.
func (s *BaseCaliburnParserListener) ExitTupleExpression(ctx *TupleExpressionContext) {}

// EnterBooleanBinaryExpression is called when production BooleanBinaryExpression is entered.
func (s *BaseCaliburnParserListener) EnterBooleanBinaryExpression(ctx *BooleanBinaryExpressionContext) {
}

// ExitBooleanBinaryExpression is called when production BooleanBinaryExpression is exited.
func (s *BaseCaliburnParserListener) ExitBooleanBinaryExpression(ctx *BooleanBinaryExpressionContext) {
}

// EnterIndexExpression is called when production IndexExpression is entered.
func (s *BaseCaliburnParserListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production IndexExpression is exited.
func (s *BaseCaliburnParserListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterUnaryExpression is called when production UnaryExpression is entered.
func (s *BaseCaliburnParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production UnaryExpression is exited.
func (s *BaseCaliburnParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterAccessExpression is called when production AccessExpression is entered.
func (s *BaseCaliburnParserListener) EnterAccessExpression(ctx *AccessExpressionContext) {}

// ExitAccessExpression is called when production AccessExpression is exited.
func (s *BaseCaliburnParserListener) ExitAccessExpression(ctx *AccessExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseCaliburnParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseCaliburnParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseCaliburnParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseCaliburnParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterStructExpression is called when production StructExpression is entered.
func (s *BaseCaliburnParserListener) EnterStructExpression(ctx *StructExpressionContext) {}

// ExitStructExpression is called when production StructExpression is exited.
func (s *BaseCaliburnParserListener) ExitStructExpression(ctx *StructExpressionContext) {}

// EnterBinaryExpression is called when production BinaryExpression is entered.
func (s *BaseCaliburnParserListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production BinaryExpression is exited.
func (s *BaseCaliburnParserListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterCastExpression is called when production CastExpression is entered.
func (s *BaseCaliburnParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production CastExpression is exited.
func (s *BaseCaliburnParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

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

// EnterType_expression is called when production type_expression is entered.
func (s *BaseCaliburnParserListener) EnterType_expression(ctx *Type_expressionContext) {}

// ExitType_expression is called when production type_expression is exited.
func (s *BaseCaliburnParserListener) ExitType_expression(ctx *Type_expressionContext) {}

// EnterComplex_type_expression is called when production complex_type_expression is entered.
func (s *BaseCaliburnParserListener) EnterComplex_type_expression(ctx *Complex_type_expressionContext) {
}

// ExitComplex_type_expression is called when production complex_type_expression is exited.
func (s *BaseCaliburnParserListener) ExitComplex_type_expression(ctx *Complex_type_expressionContext) {
}

// EnterIdentifiers is called when production identifiers is entered.
func (s *BaseCaliburnParserListener) EnterIdentifiers(ctx *IdentifiersContext) {}

// ExitIdentifiers is called when production identifiers is exited.
func (s *BaseCaliburnParserListener) ExitIdentifiers(ctx *IdentifiersContext) {}

// EnterUntypedLiteralAtom is called when production UntypedLiteralAtom is entered.
func (s *BaseCaliburnParserListener) EnterUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) {}

// ExitUntypedLiteralAtom is called when production UntypedLiteralAtom is exited.
func (s *BaseCaliburnParserListener) ExitUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) {}

// EnterTypedLiteralAtom is called when production TypedLiteralAtom is entered.
func (s *BaseCaliburnParserListener) EnterTypedLiteralAtom(ctx *TypedLiteralAtomContext) {}

// ExitTypedLiteralAtom is called when production TypedLiteralAtom is exited.
func (s *BaseCaliburnParserListener) ExitTypedLiteralAtom(ctx *TypedLiteralAtomContext) {}

// EnterUntyped_literal_atom is called when production untyped_literal_atom is entered.
func (s *BaseCaliburnParserListener) EnterUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// ExitUntyped_literal_atom is called when production untyped_literal_atom is exited.
func (s *BaseCaliburnParserListener) ExitUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// EnterTyped_literal_atom is called when production typed_literal_atom is entered.
func (s *BaseCaliburnParserListener) EnterTyped_literal_atom(ctx *Typed_literal_atomContext) {}

// ExitTyped_literal_atom is called when production typed_literal_atom is exited.
func (s *BaseCaliburnParserListener) ExitTyped_literal_atom(ctx *Typed_literal_atomContext) {}
