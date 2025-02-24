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

// EnterModuleRule is called when production ModuleRule is entered.
func (s *BaseCaliburnParserListener) EnterModuleRule(ctx *ModuleRuleContext) {}

// ExitModuleRule is called when production ModuleRule is exited.
func (s *BaseCaliburnParserListener) ExitModuleRule(ctx *ModuleRuleContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseCaliburnParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseCaliburnParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterFunctionDefinition is called when production FunctionDefinition is entered.
func (s *BaseCaliburnParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production FunctionDefinition is exited.
func (s *BaseCaliburnParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParametersRule is called when production ParametersRule is entered.
func (s *BaseCaliburnParserListener) EnterParametersRule(ctx *ParametersRuleContext) {}

// ExitParametersRule is called when production ParametersRule is exited.
func (s *BaseCaliburnParserListener) ExitParametersRule(ctx *ParametersRuleContext) {}

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

// EnterBlockRule is called when production BlockRule is entered.
func (s *BaseCaliburnParserListener) EnterBlockRule(ctx *BlockRuleContext) {}

// ExitBlockRule is called when production BlockRule is exited.
func (s *BaseCaliburnParserListener) ExitBlockRule(ctx *BlockRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCaliburnParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCaliburnParserListener) ExitStatement(ctx *StatementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseCaliburnParserListener) EnterJump_statement(ctx *Jump_statementContext) {}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseCaliburnParserListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseCaliburnParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseCaliburnParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterBreakStatement is called when production BreakStatement is entered.
func (s *BaseCaliburnParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production BreakStatement is exited.
func (s *BaseCaliburnParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production ContinueStatement is entered.
func (s *BaseCaliburnParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production ContinueStatement is exited.
func (s *BaseCaliburnParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterControl_statement is called when production control_statement is entered.
func (s *BaseCaliburnParserListener) EnterControl_statement(ctx *Control_statementContext) {}

// ExitControl_statement is called when production control_statement is exited.
func (s *BaseCaliburnParserListener) ExitControl_statement(ctx *Control_statementContext) {}

// EnterIfStatement is called when production IfStatement is entered.
func (s *BaseCaliburnParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production IfStatement is exited.
func (s *BaseCaliburnParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseCaliburnParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseCaliburnParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterSwitchStatement is called when production SwitchStatement is entered.
func (s *BaseCaliburnParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production SwitchStatement is exited.
func (s *BaseCaliburnParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlocks is called when production CaseBlocks is entered.
func (s *BaseCaliburnParserListener) EnterCaseBlocks(ctx *CaseBlocksContext) {}

// ExitCaseBlocks is called when production CaseBlocks is exited.
func (s *BaseCaliburnParserListener) ExitCaseBlocks(ctx *CaseBlocksContext) {}

// EnterOptionCaseBlock is called when production OptionCaseBlock is entered.
func (s *BaseCaliburnParserListener) EnterOptionCaseBlock(ctx *OptionCaseBlockContext) {}

// ExitOptionCaseBlock is called when production OptionCaseBlock is exited.
func (s *BaseCaliburnParserListener) ExitOptionCaseBlock(ctx *OptionCaseBlockContext) {}

// EnterDefaultCaseBlock is called when production DefaultCaseBlock is entered.
func (s *BaseCaliburnParserListener) EnterDefaultCaseBlock(ctx *DefaultCaseBlockContext) {}

// ExitDefaultCaseBlock is called when production DefaultCaseBlock is exited.
func (s *BaseCaliburnParserListener) ExitDefaultCaseBlock(ctx *DefaultCaseBlockContext) {}

// EnterInline_statement is called when production inline_statement is entered.
func (s *BaseCaliburnParserListener) EnterInline_statement(ctx *Inline_statementContext) {}

// ExitInline_statement is called when production inline_statement is exited.
func (s *BaseCaliburnParserListener) ExitInline_statement(ctx *Inline_statementContext) {}

// EnterAssignStatement is called when production AssignStatement is entered.
func (s *BaseCaliburnParserListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production AssignStatement is exited.
func (s *BaseCaliburnParserListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterAssignOperationStatement is called when production AssignOperationStatement is entered.
func (s *BaseCaliburnParserListener) EnterAssignOperationStatement(ctx *AssignOperationStatementContext) {
}

// ExitAssignOperationStatement is called when production AssignOperationStatement is exited.
func (s *BaseCaliburnParserListener) ExitAssignOperationStatement(ctx *AssignOperationStatementContext) {
}

// EnterAssignExpressions is called when production AssignExpressions is entered.
func (s *BaseCaliburnParserListener) EnterAssignExpressions(ctx *AssignExpressionsContext) {}

// ExitAssignExpressions is called when production AssignExpressions is exited.
func (s *BaseCaliburnParserListener) ExitAssignExpressions(ctx *AssignExpressionsContext) {}

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

// EnterAssignDeclarations is called when production AssignDeclarations is entered.
func (s *BaseCaliburnParserListener) EnterAssignDeclarations(ctx *AssignDeclarationsContext) {}

// ExitAssignDeclarations is called when production AssignDeclarations is exited.
func (s *BaseCaliburnParserListener) ExitAssignDeclarations(ctx *AssignDeclarationsContext) {}

// EnterExpressionAssignDeclaration is called when production ExpressionAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) {
}

// ExitExpressionAssignDeclaration is called when production ExpressionAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) {
}

// EnterTypedAssignDeclarationDeclaration is called when production TypedAssignDeclarationDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedAssignDeclarationDeclaration(ctx *TypedAssignDeclarationDeclarationContext) {
}

// ExitTypedAssignDeclarationDeclaration is called when production TypedAssignDeclarationDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedAssignDeclarationDeclaration(ctx *TypedAssignDeclarationDeclarationContext) {
}

// EnterUntypedAssignDeclarationDeclaration is called when production UntypedAssignDeclarationDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedAssignDeclarationDeclaration(ctx *UntypedAssignDeclarationDeclarationContext) {
}

// ExitUntypedAssignDeclarationDeclaration is called when production UntypedAssignDeclarationDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedAssignDeclarationDeclaration(ctx *UntypedAssignDeclarationDeclarationContext) {
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

// EnterTypedAssignDeclarations is called when production TypedAssignDeclarations is entered.
func (s *BaseCaliburnParserListener) EnterTypedAssignDeclarations(ctx *TypedAssignDeclarationsContext) {
}

// ExitTypedAssignDeclarations is called when production TypedAssignDeclarations is exited.
func (s *BaseCaliburnParserListener) ExitTypedAssignDeclarations(ctx *TypedAssignDeclarationsContext) {
}

// EnterTypedAssignDeclaration is called when production TypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {
}

// ExitTypedAssignDeclaration is called when production TypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {}

// EnterUntypedAssignDeclarations is called when production UntypedAssignDeclarations is entered.
func (s *BaseCaliburnParserListener) EnterUntypedAssignDeclarations(ctx *UntypedAssignDeclarationsContext) {
}

// ExitUntypedAssignDeclarations is called when production UntypedAssignDeclarations is exited.
func (s *BaseCaliburnParserListener) ExitUntypedAssignDeclarations(ctx *UntypedAssignDeclarationsContext) {
}

// EnterUntypedIdentifierAssignDeclaration is called when production UntypedIdentifierAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUntypedIdentifierAssignDeclaration(ctx *UntypedIdentifierAssignDeclarationContext) {
}

// ExitUntypedIdentifierAssignDeclaration is called when production UntypedIdentifierAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUntypedIdentifierAssignDeclaration(ctx *UntypedIdentifierAssignDeclarationContext) {
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

// EnterExpressionStatement is called when production ExpressionStatement is entered.
func (s *BaseCaliburnParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production ExpressionStatement is exited.
func (s *BaseCaliburnParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpressionsRule is called when production ExpressionsRule is entered.
func (s *BaseCaliburnParserListener) EnterExpressionsRule(ctx *ExpressionsRuleContext) {}

// ExitExpressionsRule is called when production ExpressionsRule is exited.
func (s *BaseCaliburnParserListener) ExitExpressionsRule(ctx *ExpressionsRuleContext) {}

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

// EnterFunctionTypeFunc is called when production FunctionTypeFunc is entered.
func (s *BaseCaliburnParserListener) EnterFunctionTypeFunc(ctx *FunctionTypeFuncContext) {}

// ExitFunctionTypeFunc is called when production FunctionTypeFunc is exited.
func (s *BaseCaliburnParserListener) ExitFunctionTypeFunc(ctx *FunctionTypeFuncContext) {}

// EnterFunctionTypeExpression is called when production FunctionTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterFunctionTypeExpression(ctx *FunctionTypeExpressionContext) {
}

// ExitFunctionTypeExpression is called when production FunctionTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitFunctionTypeExpression(ctx *FunctionTypeExpressionContext) {}

// EnterStructTypeStruct is called when production StructTypeStruct is entered.
func (s *BaseCaliburnParserListener) EnterStructTypeStruct(ctx *StructTypeStructContext) {}

// ExitStructTypeStruct is called when production StructTypeStruct is exited.
func (s *BaseCaliburnParserListener) ExitStructTypeStruct(ctx *StructTypeStructContext) {}

// EnterStructTypeExpression is called when production StructTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterStructTypeExpression(ctx *StructTypeExpressionContext) {}

// ExitStructTypeExpression is called when production StructTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitStructTypeExpression(ctx *StructTypeExpressionContext) {}

// EnterTupleTypeTuple is called when production TupleTypeTuple is entered.
func (s *BaseCaliburnParserListener) EnterTupleTypeTuple(ctx *TupleTypeTupleContext) {}

// ExitTupleTypeTuple is called when production TupleTypeTuple is exited.
func (s *BaseCaliburnParserListener) ExitTupleTypeTuple(ctx *TupleTypeTupleContext) {}

// EnterTupleTypeExpression is called when production TupleTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterTupleTypeExpression(ctx *TupleTypeExpressionContext) {}

// ExitTupleTypeExpression is called when production TupleTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitTupleTypeExpression(ctx *TupleTypeExpressionContext) {}

// EnterAccessTypeExpression is called when production AccessTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterAccessTypeExpression(ctx *AccessTypeExpressionContext) {}

// ExitAccessTypeExpression is called when production AccessTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitAccessTypeExpression(ctx *AccessTypeExpressionContext) {}

// EnterIdentifierTypeExpression is called when production IdentifierTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterIdentifierTypeExpression(ctx *IdentifierTypeExpressionContext) {
}

// ExitIdentifierTypeExpression is called when production IdentifierTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitIdentifierTypeExpression(ctx *IdentifierTypeExpressionContext) {
}

// EnterIdentifiersRule is called when production IdentifiersRule is entered.
func (s *BaseCaliburnParserListener) EnterIdentifiersRule(ctx *IdentifiersRuleContext) {}

// ExitIdentifiersRule is called when production IdentifiersRule is exited.
func (s *BaseCaliburnParserListener) ExitIdentifiersRule(ctx *IdentifiersRuleContext) {}

// EnterUntypedLiteralAtom is called when production UntypedLiteralAtom is entered.
func (s *BaseCaliburnParserListener) EnterUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) {}

// ExitUntypedLiteralAtom is called when production UntypedLiteralAtom is exited.
func (s *BaseCaliburnParserListener) ExitUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) {}

// EnterTypedLiteralAtom is called when production TypedLiteralAtom is entered.
func (s *BaseCaliburnParserListener) EnterTypedLiteralAtom(ctx *TypedLiteralAtomContext) {}

// ExitTypedLiteralAtom is called when production TypedLiteralAtom is exited.
func (s *BaseCaliburnParserListener) ExitTypedLiteralAtom(ctx *TypedLiteralAtomContext) {}
