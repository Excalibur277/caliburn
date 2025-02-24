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

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseCaliburnParserListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseCaliburnParserListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseCaliburnParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseCaliburnParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterTypedParameter is called when production TypedParameter is entered.
func (s *BaseCaliburnParserListener) EnterTypedParameter(ctx *TypedParameterContext) {}

// ExitTypedParameter is called when production TypedParameter is exited.
func (s *BaseCaliburnParserListener) ExitTypedParameter(ctx *TypedParameterContext) {}

// EnterUntypedParameter is called when production UntypedParameter is entered.
func (s *BaseCaliburnParserListener) EnterUntypedParameter(ctx *UntypedParameterContext) {}

// ExitUntypedParameter is called when production UntypedParameter is exited.
func (s *BaseCaliburnParserListener) ExitUntypedParameter(ctx *UntypedParameterContext) {}

// EnterStructDestrutureParameter is called when production StructDestrutureParameter is entered.
func (s *BaseCaliburnParserListener) EnterStructDestrutureParameter(ctx *StructDestrutureParameterContext) {
}

// ExitStructDestrutureParameter is called when production StructDestrutureParameter is exited.
func (s *BaseCaliburnParserListener) ExitStructDestrutureParameter(ctx *StructDestrutureParameterContext) {
}

// EnterTupleDestrutureParameter is called when production TupleDestrutureParameter is entered.
func (s *BaseCaliburnParserListener) EnterTupleDestrutureParameter(ctx *TupleDestrutureParameterContext) {
}

// ExitTupleDestrutureParameter is called when production TupleDestrutureParameter is exited.
func (s *BaseCaliburnParserListener) ExitTupleDestrutureParameter(ctx *TupleDestrutureParameterContext) {
}

// EnterBlock is called when production block is entered.
func (s *BaseCaliburnParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseCaliburnParserListener) ExitBlock(ctx *BlockContext) {}

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

// EnterCase_statements is called when production case_statements is entered.
func (s *BaseCaliburnParserListener) EnterCase_statements(ctx *Case_statementsContext) {}

// ExitCase_statements is called when production case_statements is exited.
func (s *BaseCaliburnParserListener) ExitCase_statements(ctx *Case_statementsContext) {}

// EnterOption_case_statement is called when production option_case_statement is entered.
func (s *BaseCaliburnParserListener) EnterOption_case_statement(ctx *Option_case_statementContext) {}

// ExitOption_case_statement is called when production option_case_statement is exited.
func (s *BaseCaliburnParserListener) ExitOption_case_statement(ctx *Option_case_statementContext) {}

// EnterDefault_case_statement is called when production default_case_statement is entered.
func (s *BaseCaliburnParserListener) EnterDefault_case_statement(ctx *Default_case_statementContext) {
}

// ExitDefault_case_statement is called when production default_case_statement is exited.
func (s *BaseCaliburnParserListener) ExitDefault_case_statement(ctx *Default_case_statementContext) {}

// EnterInline_statement is called when production inline_statement is entered.
func (s *BaseCaliburnParserListener) EnterInline_statement(ctx *Inline_statementContext) {}

// ExitInline_statement is called when production inline_statement is exited.
func (s *BaseCaliburnParserListener) ExitInline_statement(ctx *Inline_statementContext) {}

// EnterAssign_statement is called when production assign_statement is entered.
func (s *BaseCaliburnParserListener) EnterAssign_statement(ctx *Assign_statementContext) {}

// ExitAssign_statement is called when production assign_statement is exited.
func (s *BaseCaliburnParserListener) ExitAssign_statement(ctx *Assign_statementContext) {}

// EnterAssign_expressions is called when production assign_expressions is entered.
func (s *BaseCaliburnParserListener) EnterAssign_expressions(ctx *Assign_expressionsContext) {}

// ExitAssign_expressions is called when production assign_expressions is exited.
func (s *BaseCaliburnParserListener) ExitAssign_expressions(ctx *Assign_expressionsContext) {}

// EnterExpressionAssignExpression is called when production ExpressionAssignExpression is entered.
func (s *BaseCaliburnParserListener) EnterExpressionAssignExpression(ctx *ExpressionAssignExpressionContext) {
}

// ExitExpressionAssignExpression is called when production ExpressionAssignExpression is exited.
func (s *BaseCaliburnParserListener) ExitExpressionAssignExpression(ctx *ExpressionAssignExpressionContext) {
}

// EnterStructDestrutureAssignExpression is called when production StructDestrutureAssignExpression is entered.
func (s *BaseCaliburnParserListener) EnterStructDestrutureAssignExpression(ctx *StructDestrutureAssignExpressionContext) {
}

// ExitStructDestrutureAssignExpression is called when production StructDestrutureAssignExpression is exited.
func (s *BaseCaliburnParserListener) ExitStructDestrutureAssignExpression(ctx *StructDestrutureAssignExpressionContext) {
}

// EnterTupleDestrutureAssignExpression is called when production TupleDestrutureAssignExpression is entered.
func (s *BaseCaliburnParserListener) EnterTupleDestrutureAssignExpression(ctx *TupleDestrutureAssignExpressionContext) {
}

// ExitTupleDestrutureAssignExpression is called when production TupleDestrutureAssignExpression is exited.
func (s *BaseCaliburnParserListener) ExitTupleDestrutureAssignExpression(ctx *TupleDestrutureAssignExpressionContext) {
}

// EnterAssign_declarations is called when production assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterAssign_declarations(ctx *Assign_declarationsContext) {}

// ExitAssign_declarations is called when production assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitAssign_declarations(ctx *Assign_declarationsContext) {}

// EnterExpressionAssignDeclaration is called when production ExpressionAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) {
}

// ExitExpressionAssignDeclaration is called when production ExpressionAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) {
}

// EnterTypedAssignDeclaration is called when production TypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {
}

// ExitTypedAssignDeclaration is called when production TypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {}

// EnterUntypedAssignDeclaration is called when production UntypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedAssignDeclaration(ctx *UntypedAssignDeclarationContext) {
}

// ExitUntypedAssignDeclaration is called when production UntypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedAssignDeclaration(ctx *UntypedAssignDeclarationContext) {
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

// EnterTyped_assign_declarations is called when production typed_assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterTyped_assign_declarations(ctx *Typed_assign_declarationsContext) {
}

// ExitTyped_assign_declarations is called when production typed_assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitTyped_assign_declarations(ctx *Typed_assign_declarationsContext) {
}

// EnterTypedTypedAssignDeclaration is called when production TypedTypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedTypedAssignDeclaration(ctx *TypedTypedAssignDeclarationContext) {
}

// ExitTypedTypedAssignDeclaration is called when production TypedTypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedTypedAssignDeclaration(ctx *TypedTypedAssignDeclarationContext) {
}

// EnterUntyped_assign_declarations is called when production untyped_assign_declarations is entered.
func (s *BaseCaliburnParserListener) EnterUntyped_assign_declarations(ctx *Untyped_assign_declarationsContext) {
}

// ExitUntyped_assign_declarations is called when production untyped_assign_declarations is exited.
func (s *BaseCaliburnParserListener) ExitUntyped_assign_declarations(ctx *Untyped_assign_declarationsContext) {
}

// EnterUntypedAtomAssignDeclaration is called when production UntypedAtomAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedAtomAssignDeclaration(ctx *UntypedAtomAssignDeclarationContext) {
}

// ExitUntypedAtomAssignDeclaration is called when production UntypedAtomAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedAtomAssignDeclaration(ctx *UntypedAtomAssignDeclarationContext) {
}

// EnterUntypedStructDestrutureAssignDeclaration is called when production UntypedStructDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedStructDestrutureAssignDeclaration(ctx *UntypedStructDestrutureAssignDeclarationContext) {
}

// ExitUntypedStructDestrutureAssignDeclaration is called when production UntypedStructDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedStructDestrutureAssignDeclaration(ctx *UntypedStructDestrutureAssignDeclarationContext) {
}

// EnterUntypedTupleDestrutureAssignDeclaration is called when production UntypedTupleDestrutureAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedTupleDestrutureAssignDeclaration(ctx *UntypedTupleDestrutureAssignDeclarationContext) {
}

// ExitUntypedTupleDestrutureAssignDeclaration is called when production UntypedTupleDestrutureAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedTupleDestrutureAssignDeclaration(ctx *UntypedTupleDestrutureAssignDeclarationContext) {
}

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

// EnterFunction_type is called when production function_type is entered.
func (s *BaseCaliburnParserListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BaseCaliburnParserListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseCaliburnParserListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseCaliburnParserListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterTuple_type is called when production tuple_type is entered.
func (s *BaseCaliburnParserListener) EnterTuple_type(ctx *Tuple_typeContext) {}

// ExitTuple_type is called when production tuple_type is exited.
func (s *BaseCaliburnParserListener) ExitTuple_type(ctx *Tuple_typeContext) {}

// EnterType_expression is called when production type_expression is entered.
func (s *BaseCaliburnParserListener) EnterType_expression(ctx *Type_expressionContext) {}

// ExitType_expression is called when production type_expression is exited.
func (s *BaseCaliburnParserListener) ExitType_expression(ctx *Type_expressionContext) {}

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
