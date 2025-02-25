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

// EnterDefinitionsAppend is called when production DefinitionsAppend is entered.
func (s *BaseCaliburnParserListener) EnterDefinitionsAppend(ctx *DefinitionsAppendContext) {}

// ExitDefinitionsAppend is called when production DefinitionsAppend is exited.
func (s *BaseCaliburnParserListener) ExitDefinitionsAppend(ctx *DefinitionsAppendContext) {}

// EnterDefinitionsInitial is called when production DefinitionsInitial is entered.
func (s *BaseCaliburnParserListener) EnterDefinitionsInitial(ctx *DefinitionsInitialContext) {}

// ExitDefinitionsInitial is called when production DefinitionsInitial is exited.
func (s *BaseCaliburnParserListener) ExitDefinitionsInitial(ctx *DefinitionsInitialContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseCaliburnParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseCaliburnParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterFunctionDefinition is called when production FunctionDefinition is entered.
func (s *BaseCaliburnParserListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production FunctionDefinition is exited.
func (s *BaseCaliburnParserListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionDefinitionNoReturnType is called when production FunctionDefinitionNoReturnType is entered.
func (s *BaseCaliburnParserListener) EnterFunctionDefinitionNoReturnType(ctx *FunctionDefinitionNoReturnTypeContext) {
}

// ExitFunctionDefinitionNoReturnType is called when production FunctionDefinitionNoReturnType is exited.
func (s *BaseCaliburnParserListener) ExitFunctionDefinitionNoReturnType(ctx *FunctionDefinitionNoReturnTypeContext) {
}

// EnterParametersEmpty is called when production ParametersEmpty is entered.
func (s *BaseCaliburnParserListener) EnterParametersEmpty(ctx *ParametersEmptyContext) {}

// ExitParametersEmpty is called when production ParametersEmpty is exited.
func (s *BaseCaliburnParserListener) ExitParametersEmpty(ctx *ParametersEmptyContext) {}

// EnterParametersFilled is called when production ParametersFilled is entered.
func (s *BaseCaliburnParserListener) EnterParametersFilled(ctx *ParametersFilledContext) {}

// ExitParametersFilled is called when production ParametersFilled is exited.
func (s *BaseCaliburnParserListener) ExitParametersFilled(ctx *ParametersFilledContext) {}

// EnterParametersListInitial is called when production ParametersListInitial is entered.
func (s *BaseCaliburnParserListener) EnterParametersListInitial(ctx *ParametersListInitialContext) {}

// ExitParametersListInitial is called when production ParametersListInitial is exited.
func (s *BaseCaliburnParserListener) ExitParametersListInitial(ctx *ParametersListInitialContext) {}

// EnterParametersListAppend is called when production ParametersListAppend is entered.
func (s *BaseCaliburnParserListener) EnterParametersListAppend(ctx *ParametersListAppendContext) {}

// ExitParametersListAppend is called when production ParametersListAppend is exited.
func (s *BaseCaliburnParserListener) ExitParametersListAppend(ctx *ParametersListAppendContext) {}

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

// EnterBlockRule is called when production BlockRule is entered.
func (s *BaseCaliburnParserListener) EnterBlockRule(ctx *BlockRuleContext) {}

// ExitBlockRule is called when production BlockRule is exited.
func (s *BaseCaliburnParserListener) ExitBlockRule(ctx *BlockRuleContext) {}

// EnterStatementsAppend is called when production StatementsAppend is entered.
func (s *BaseCaliburnParserListener) EnterStatementsAppend(ctx *StatementsAppendContext) {}

// ExitStatementsAppend is called when production StatementsAppend is exited.
func (s *BaseCaliburnParserListener) ExitStatementsAppend(ctx *StatementsAppendContext) {}

// EnterStatementsInitial is called when production StatementsInitial is entered.
func (s *BaseCaliburnParserListener) EnterStatementsInitial(ctx *StatementsInitialContext) {}

// ExitStatementsInitial is called when production StatementsInitial is exited.
func (s *BaseCaliburnParserListener) ExitStatementsInitial(ctx *StatementsInitialContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCaliburnParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCaliburnParserListener) ExitStatement(ctx *StatementContext) {}

// EnterInlineStatementsInitial is called when production InlineStatementsInitial is entered.
func (s *BaseCaliburnParserListener) EnterInlineStatementsInitial(ctx *InlineStatementsInitialContext) {
}

// ExitInlineStatementsInitial is called when production InlineStatementsInitial is exited.
func (s *BaseCaliburnParserListener) ExitInlineStatementsInitial(ctx *InlineStatementsInitialContext) {
}

// EnterInlineStatementsAppend is called when production InlineStatementsAppend is entered.
func (s *BaseCaliburnParserListener) EnterInlineStatementsAppend(ctx *InlineStatementsAppendContext) {
}

// ExitInlineStatementsAppend is called when production InlineStatementsAppend is exited.
func (s *BaseCaliburnParserListener) ExitInlineStatementsAppend(ctx *InlineStatementsAppendContext) {}

// EnterInline_statement is called when production inline_statement is entered.
func (s *BaseCaliburnParserListener) EnterInline_statement(ctx *Inline_statementContext) {}

// ExitInline_statement is called when production inline_statement is exited.
func (s *BaseCaliburnParserListener) ExitInline_statement(ctx *Inline_statementContext) {}

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

// EnterIfElseStatement is called when production IfElseStatement is entered.
func (s *BaseCaliburnParserListener) EnterIfElseStatement(ctx *IfElseStatementContext) {}

// ExitIfElseStatement is called when production IfElseStatement is exited.
func (s *BaseCaliburnParserListener) ExitIfElseStatement(ctx *IfElseStatementContext) {}

// EnterElseStatement is called when production ElseStatement is entered.
func (s *BaseCaliburnParserListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production ElseStatement is exited.
func (s *BaseCaliburnParserListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterElseIfStatement is called when production ElseIfStatement is entered.
func (s *BaseCaliburnParserListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production ElseIfStatement is exited.
func (s *BaseCaliburnParserListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseCaliburnParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseCaliburnParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForStatementWithAfter is called when production ForStatementWithAfter is entered.
func (s *BaseCaliburnParserListener) EnterForStatementWithAfter(ctx *ForStatementWithAfterContext) {}

// ExitForStatementWithAfter is called when production ForStatementWithAfter is exited.
func (s *BaseCaliburnParserListener) ExitForStatementWithAfter(ctx *ForStatementWithAfterContext) {}

// EnterSwitchStatement is called when production SwitchStatement is entered.
func (s *BaseCaliburnParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production SwitchStatement is exited.
func (s *BaseCaliburnParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlocks is called when production CaseBlocks is entered.
func (s *BaseCaliburnParserListener) EnterCaseBlocks(ctx *CaseBlocksContext) {}

// ExitCaseBlocks is called when production CaseBlocks is exited.
func (s *BaseCaliburnParserListener) ExitCaseBlocks(ctx *CaseBlocksContext) {}

// EnterCaseBlocksDefault is called when production CaseBlocksDefault is entered.
func (s *BaseCaliburnParserListener) EnterCaseBlocksDefault(ctx *CaseBlocksDefaultContext) {}

// ExitCaseBlocksDefault is called when production CaseBlocksDefault is exited.
func (s *BaseCaliburnParserListener) ExitCaseBlocksDefault(ctx *CaseBlocksDefaultContext) {}

// EnterOptionCaseBlocksAppend is called when production OptionCaseBlocksAppend is entered.
func (s *BaseCaliburnParserListener) EnterOptionCaseBlocksAppend(ctx *OptionCaseBlocksAppendContext) {
}

// ExitOptionCaseBlocksAppend is called when production OptionCaseBlocksAppend is exited.
func (s *BaseCaliburnParserListener) ExitOptionCaseBlocksAppend(ctx *OptionCaseBlocksAppendContext) {}

// EnterOptionCaseBlocksInitial is called when production OptionCaseBlocksInitial is entered.
func (s *BaseCaliburnParserListener) EnterOptionCaseBlocksInitial(ctx *OptionCaseBlocksInitialContext) {
}

// ExitOptionCaseBlocksInitial is called when production OptionCaseBlocksInitial is exited.
func (s *BaseCaliburnParserListener) ExitOptionCaseBlocksInitial(ctx *OptionCaseBlocksInitialContext) {
}

// EnterOptionCaseBlock is called when production OptionCaseBlock is entered.
func (s *BaseCaliburnParserListener) EnterOptionCaseBlock(ctx *OptionCaseBlockContext) {}

// ExitOptionCaseBlock is called when production OptionCaseBlock is exited.
func (s *BaseCaliburnParserListener) ExitOptionCaseBlock(ctx *OptionCaseBlockContext) {}

// EnterDefaultCaseBlock is called when production DefaultCaseBlock is entered.
func (s *BaseCaliburnParserListener) EnterDefaultCaseBlock(ctx *DefaultCaseBlockContext) {}

// ExitDefaultCaseBlock is called when production DefaultCaseBlock is exited.
func (s *BaseCaliburnParserListener) ExitDefaultCaseBlock(ctx *DefaultCaseBlockContext) {}

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

// EnterAssignExpressionsAppend is called when production AssignExpressionsAppend is entered.
func (s *BaseCaliburnParserListener) EnterAssignExpressionsAppend(ctx *AssignExpressionsAppendContext) {
}

// ExitAssignExpressionsAppend is called when production AssignExpressionsAppend is exited.
func (s *BaseCaliburnParserListener) ExitAssignExpressionsAppend(ctx *AssignExpressionsAppendContext) {
}

// EnterAssignExpressionsInitial is called when production AssignExpressionsInitial is entered.
func (s *BaseCaliburnParserListener) EnterAssignExpressionsInitial(ctx *AssignExpressionsInitialContext) {
}

// ExitAssignExpressionsInitial is called when production AssignExpressionsInitial is exited.
func (s *BaseCaliburnParserListener) ExitAssignExpressionsInitial(ctx *AssignExpressionsInitialContext) {
}

// EnterAliasableAssignExpressionsInitial is called when production AliasableAssignExpressionsInitial is entered.
func (s *BaseCaliburnParserListener) EnterAliasableAssignExpressionsInitial(ctx *AliasableAssignExpressionsInitialContext) {
}

// ExitAliasableAssignExpressionsInitial is called when production AliasableAssignExpressionsInitial is exited.
func (s *BaseCaliburnParserListener) ExitAliasableAssignExpressionsInitial(ctx *AliasableAssignExpressionsInitialContext) {
}

// EnterAliasableAssignExpressionsAppend is called when production AliasableAssignExpressionsAppend is entered.
func (s *BaseCaliburnParserListener) EnterAliasableAssignExpressionsAppend(ctx *AliasableAssignExpressionsAppendContext) {
}

// ExitAliasableAssignExpressionsAppend is called when production AliasableAssignExpressionsAppend is exited.
func (s *BaseCaliburnParserListener) ExitAliasableAssignExpressionsAppend(ctx *AliasableAssignExpressionsAppendContext) {
}

// EnterAliasedAssignExpression is called when production AliasedAssignExpression is entered.
func (s *BaseCaliburnParserListener) EnterAliasedAssignExpression(ctx *AliasedAssignExpressionContext) {
}

// ExitAliasedAssignExpression is called when production AliasedAssignExpression is exited.
func (s *BaseCaliburnParserListener) ExitAliasedAssignExpression(ctx *AliasedAssignExpressionContext) {
}

// EnterUnaliasedAssignExpression is called when production UnaliasedAssignExpression is entered.
func (s *BaseCaliburnParserListener) EnterUnaliasedAssignExpression(ctx *UnaliasedAssignExpressionContext) {
}

// ExitUnaliasedAssignExpression is called when production UnaliasedAssignExpression is exited.
func (s *BaseCaliburnParserListener) ExitUnaliasedAssignExpression(ctx *UnaliasedAssignExpressionContext) {
}

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

// EnterAssignDeclarationsAppend is called when production AssignDeclarationsAppend is entered.
func (s *BaseCaliburnParserListener) EnterAssignDeclarationsAppend(ctx *AssignDeclarationsAppendContext) {
}

// ExitAssignDeclarationsAppend is called when production AssignDeclarationsAppend is exited.
func (s *BaseCaliburnParserListener) ExitAssignDeclarationsAppend(ctx *AssignDeclarationsAppendContext) {
}

// EnterAssignDeclarationsInitial is called when production AssignDeclarationsInitial is entered.
func (s *BaseCaliburnParserListener) EnterAssignDeclarationsInitial(ctx *AssignDeclarationsInitialContext) {
}

// ExitAssignDeclarationsInitial is called when production AssignDeclarationsInitial is exited.
func (s *BaseCaliburnParserListener) ExitAssignDeclarationsInitial(ctx *AssignDeclarationsInitialContext) {
}

// EnterAliasableAssignDeclarationsAppend is called when production AliasableAssignDeclarationsAppend is entered.
func (s *BaseCaliburnParserListener) EnterAliasableAssignDeclarationsAppend(ctx *AliasableAssignDeclarationsAppendContext) {
}

// ExitAliasableAssignDeclarationsAppend is called when production AliasableAssignDeclarationsAppend is exited.
func (s *BaseCaliburnParserListener) ExitAliasableAssignDeclarationsAppend(ctx *AliasableAssignDeclarationsAppendContext) {
}

// EnterAliasableAssignDeclarationsInitial is called when production AliasableAssignDeclarationsInitial is entered.
func (s *BaseCaliburnParserListener) EnterAliasableAssignDeclarationsInitial(ctx *AliasableAssignDeclarationsInitialContext) {
}

// ExitAliasableAssignDeclarationsInitial is called when production AliasableAssignDeclarationsInitial is exited.
func (s *BaseCaliburnParserListener) ExitAliasableAssignDeclarationsInitial(ctx *AliasableAssignDeclarationsInitialContext) {
}

// EnterAliasedAssignDeclaration is called when production AliasedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterAliasedAssignDeclaration(ctx *AliasedAssignDeclarationContext) {
}

// ExitAliasedAssignDeclaration is called when production AliasedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitAliasedAssignDeclaration(ctx *AliasedAssignDeclarationContext) {
}

// EnterUnaliasedAssignDeclaration is called when production UnaliasedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUnaliasedAssignDeclaration(ctx *UnaliasedAssignDeclarationContext) {
}

// ExitUnaliasedAssignDeclaration is called when production UnaliasedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUnaliasedAssignDeclaration(ctx *UnaliasedAssignDeclarationContext) {
}

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

// EnterTypedAssignDeclaration is called when production TypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {
}

// ExitTypedAssignDeclaration is called when production TypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) {}

// EnterAliasableUntypedAssignDeclarationsInitial is called when production AliasableUntypedAssignDeclarationsInitial is entered.
func (s *BaseCaliburnParserListener) EnterAliasableUntypedAssignDeclarationsInitial(ctx *AliasableUntypedAssignDeclarationsInitialContext) {
}

// ExitAliasableUntypedAssignDeclarationsInitial is called when production AliasableUntypedAssignDeclarationsInitial is exited.
func (s *BaseCaliburnParserListener) ExitAliasableUntypedAssignDeclarationsInitial(ctx *AliasableUntypedAssignDeclarationsInitialContext) {
}

// EnterAliasableUntypedAssignDeclarationsAppend is called when production AliasableUntypedAssignDeclarationsAppend is entered.
func (s *BaseCaliburnParserListener) EnterAliasableUntypedAssignDeclarationsAppend(ctx *AliasableUntypedAssignDeclarationsAppendContext) {
}

// ExitAliasableUntypedAssignDeclarationsAppend is called when production AliasableUntypedAssignDeclarationsAppend is exited.
func (s *BaseCaliburnParserListener) ExitAliasableUntypedAssignDeclarationsAppend(ctx *AliasableUntypedAssignDeclarationsAppendContext) {
}

// EnterAliasedUntypedAssignDeclaration is called when production AliasedUntypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterAliasedUntypedAssignDeclaration(ctx *AliasedUntypedAssignDeclarationContext) {
}

// ExitAliasedUntypedAssignDeclaration is called when production AliasedUntypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitAliasedUntypedAssignDeclaration(ctx *AliasedUntypedAssignDeclarationContext) {
}

// EnterUnaliasedUntypedAssignDeclaration is called when production UnaliasedUntypedAssignDeclaration is entered.
func (s *BaseCaliburnParserListener) EnterUnaliasedUntypedAssignDeclaration(ctx *UnaliasedUntypedAssignDeclarationContext) {
}

// ExitUnaliasedUntypedAssignDeclaration is called when production UnaliasedUntypedAssignDeclaration is exited.
func (s *BaseCaliburnParserListener) ExitUnaliasedUntypedAssignDeclaration(ctx *UnaliasedUntypedAssignDeclarationContext) {
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

// EnterExpressionStatement is called when production ExpressionStatement is entered.
func (s *BaseCaliburnParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production ExpressionStatement is exited.
func (s *BaseCaliburnParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpressionsInitial is called when production ExpressionsInitial is entered.
func (s *BaseCaliburnParserListener) EnterExpressionsInitial(ctx *ExpressionsInitialContext) {}

// ExitExpressionsInitial is called when production ExpressionsInitial is exited.
func (s *BaseCaliburnParserListener) ExitExpressionsInitial(ctx *ExpressionsInitialContext) {}

// EnterExpressionsAppend is called when production ExpressionsAppend is entered.
func (s *BaseCaliburnParserListener) EnterExpressionsAppend(ctx *ExpressionsAppendContext) {}

// ExitExpressionsAppend is called when production ExpressionsAppend is exited.
func (s *BaseCaliburnParserListener) ExitExpressionsAppend(ctx *ExpressionsAppendContext) {}

// EnterExpressionsOptional is called when production ExpressionsOptional is entered.
func (s *BaseCaliburnParserListener) EnterExpressionsOptional(ctx *ExpressionsOptionalContext) {}

// ExitExpressionsOptional is called when production ExpressionsOptional is exited.
func (s *BaseCaliburnParserListener) ExitExpressionsOptional(ctx *ExpressionsOptionalContext) {}

// EnterExpressionsOptionalNone is called when production ExpressionsOptionalNone is entered.
func (s *BaseCaliburnParserListener) EnterExpressionsOptionalNone(ctx *ExpressionsOptionalNoneContext) {
}

// ExitExpressionsOptionalNone is called when production ExpressionsOptionalNone is exited.
func (s *BaseCaliburnParserListener) ExitExpressionsOptionalNone(ctx *ExpressionsOptionalNoneContext) {
}

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

// EnterSliceEndExpression is called when production SliceEndExpression is entered.
func (s *BaseCaliburnParserListener) EnterSliceEndExpression(ctx *SliceEndExpressionContext) {}

// ExitSliceEndExpression is called when production SliceEndExpression is exited.
func (s *BaseCaliburnParserListener) ExitSliceEndExpression(ctx *SliceEndExpressionContext) {}

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

// EnterSliceStartExpression is called when production SliceStartExpression is entered.
func (s *BaseCaliburnParserListener) EnterSliceStartExpression(ctx *SliceStartExpressionContext) {}

// ExitSliceStartExpression is called when production SliceStartExpression is exited.
func (s *BaseCaliburnParserListener) ExitSliceStartExpression(ctx *SliceStartExpressionContext) {}

// EnterFunctionExpressionNoReturnType is called when production FunctionExpressionNoReturnType is entered.
func (s *BaseCaliburnParserListener) EnterFunctionExpressionNoReturnType(ctx *FunctionExpressionNoReturnTypeContext) {
}

// ExitFunctionExpressionNoReturnType is called when production FunctionExpressionNoReturnType is exited.
func (s *BaseCaliburnParserListener) ExitFunctionExpressionNoReturnType(ctx *FunctionExpressionNoReturnTypeContext) {
}

// EnterCollectionExpression is called when production CollectionExpression is entered.
func (s *BaseCaliburnParserListener) EnterCollectionExpression(ctx *CollectionExpressionContext) {}

// ExitCollectionExpression is called when production CollectionExpression is exited.
func (s *BaseCaliburnParserListener) ExitCollectionExpression(ctx *CollectionExpressionContext) {}

// EnterCallExpression is called when production CallExpression is entered.
func (s *BaseCaliburnParserListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production CallExpression is exited.
func (s *BaseCaliburnParserListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterStructValuesNamed is called when production StructValuesNamed is entered.
func (s *BaseCaliburnParserListener) EnterStructValuesNamed(ctx *StructValuesNamedContext) {}

// ExitStructValuesNamed is called when production StructValuesNamed is exited.
func (s *BaseCaliburnParserListener) ExitStructValuesNamed(ctx *StructValuesNamedContext) {}

// EnterStructValuesUnamed is called when production StructValuesUnamed is entered.
func (s *BaseCaliburnParserListener) EnterStructValuesUnamed(ctx *StructValuesUnamedContext) {}

// ExitStructValuesUnamed is called when production StructValuesUnamed is exited.
func (s *BaseCaliburnParserListener) ExitStructValuesUnamed(ctx *StructValuesUnamedContext) {}

// EnterNamedStructValuesAppend is called when production NamedStructValuesAppend is entered.
func (s *BaseCaliburnParserListener) EnterNamedStructValuesAppend(ctx *NamedStructValuesAppendContext) {
}

// ExitNamedStructValuesAppend is called when production NamedStructValuesAppend is exited.
func (s *BaseCaliburnParserListener) ExitNamedStructValuesAppend(ctx *NamedStructValuesAppendContext) {
}

// EnterNamedStructValuesInitial is called when production NamedStructValuesInitial is entered.
func (s *BaseCaliburnParserListener) EnterNamedStructValuesInitial(ctx *NamedStructValuesInitialContext) {
}

// ExitNamedStructValuesInitial is called when production NamedStructValuesInitial is exited.
func (s *BaseCaliburnParserListener) ExitNamedStructValuesInitial(ctx *NamedStructValuesInitialContext) {
}

// EnterNamedStructValue is called when production NamedStructValue is entered.
func (s *BaseCaliburnParserListener) EnterNamedStructValue(ctx *NamedStructValueContext) {}

// ExitNamedStructValue is called when production NamedStructValue is exited.
func (s *BaseCaliburnParserListener) ExitNamedStructValue(ctx *NamedStructValueContext) {}

// EnterCollectionValuesInitial is called when production CollectionValuesInitial is entered.
func (s *BaseCaliburnParserListener) EnterCollectionValuesInitial(ctx *CollectionValuesInitialContext) {
}

// ExitCollectionValuesInitial is called when production CollectionValuesInitial is exited.
func (s *BaseCaliburnParserListener) ExitCollectionValuesInitial(ctx *CollectionValuesInitialContext) {
}

// EnterCollectionValuesAppend is called when production CollectionValuesAppend is entered.
func (s *BaseCaliburnParserListener) EnterCollectionValuesAppend(ctx *CollectionValuesAppendContext) {
}

// ExitCollectionValuesAppend is called when production CollectionValuesAppend is exited.
func (s *BaseCaliburnParserListener) ExitCollectionValuesAppend(ctx *CollectionValuesAppendContext) {}

// EnterCollectionValue is called when production CollectionValue is entered.
func (s *BaseCaliburnParserListener) EnterCollectionValue(ctx *CollectionValueContext) {}

// ExitCollectionValue is called when production CollectionValue is exited.
func (s *BaseCaliburnParserListener) ExitCollectionValue(ctx *CollectionValueContext) {}

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

// EnterStructTypeTuple is called when production StructTypeTuple is entered.
func (s *BaseCaliburnParserListener) EnterStructTypeTuple(ctx *StructTypeTupleContext) {}

// ExitStructTypeTuple is called when production StructTypeTuple is exited.
func (s *BaseCaliburnParserListener) ExitStructTypeTuple(ctx *StructTypeTupleContext) {}

// EnterStructTypeExpression is called when production StructTypeExpression is entered.
func (s *BaseCaliburnParserListener) EnterStructTypeExpression(ctx *StructTypeExpressionContext) {}

// ExitStructTypeExpression is called when production StructTypeExpression is exited.
func (s *BaseCaliburnParserListener) ExitStructTypeExpression(ctx *StructTypeExpressionContext) {}

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

// EnterUntypedLiteral is called when production UntypedLiteral is entered.
func (s *BaseCaliburnParserListener) EnterUntypedLiteral(ctx *UntypedLiteralContext) {}

// ExitUntypedLiteral is called when production UntypedLiteral is exited.
func (s *BaseCaliburnParserListener) ExitUntypedLiteral(ctx *UntypedLiteralContext) {}

// EnterTypedLiteral is called when production TypedLiteral is entered.
func (s *BaseCaliburnParserListener) EnterTypedLiteral(ctx *TypedLiteralContext) {}

// ExitTypedLiteral is called when production TypedLiteral is exited.
func (s *BaseCaliburnParserListener) ExitTypedLiteral(ctx *TypedLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseCaliburnParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseCaliburnParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseCaliburnParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseCaliburnParserListener) ExitIdentifier(ctx *IdentifierContext) {}
