// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// CaliburnParserListener is a complete listener for a parse tree produced by CaliburnParser.
type CaliburnParserListener interface {
	antlr.ParseTreeListener

	// EnterModuleRule is called when entering the ModuleRule production.
	EnterModuleRule(c *ModuleRuleContext)

	// EnterDefinitionsAppend is called when entering the DefinitionsAppend production.
	EnterDefinitionsAppend(c *DefinitionsAppendContext)

	// EnterDefinitionsInitial is called when entering the DefinitionsInitial production.
	EnterDefinitionsInitial(c *DefinitionsInitialContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterFunctionDefinition is called when entering the FunctionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionDefinitionNoReturnType is called when entering the FunctionDefinitionNoReturnType production.
	EnterFunctionDefinitionNoReturnType(c *FunctionDefinitionNoReturnTypeContext)

	// EnterParametersEmpty is called when entering the ParametersEmpty production.
	EnterParametersEmpty(c *ParametersEmptyContext)

	// EnterParametersFilled is called when entering the ParametersFilled production.
	EnterParametersFilled(c *ParametersFilledContext)

	// EnterParametersListInitial is called when entering the ParametersListInitial production.
	EnterParametersListInitial(c *ParametersListInitialContext)

	// EnterParametersListAppend is called when entering the ParametersListAppend production.
	EnterParametersListAppend(c *ParametersListAppendContext)

	// EnterTypedParameter is called when entering the TypedParameter production.
	EnterTypedParameter(c *TypedParameterContext)

	// EnterUntypedParameter is called when entering the UntypedParameter production.
	EnterUntypedParameter(c *UntypedParameterContext)

	// EnterStructDestrutureParameter is called when entering the StructDestrutureParameter production.
	EnterStructDestrutureParameter(c *StructDestrutureParameterContext)

	// EnterTupleDestrutureParameter is called when entering the TupleDestrutureParameter production.
	EnterTupleDestrutureParameter(c *TupleDestrutureParameterContext)

	// EnterBlockRule is called when entering the BlockRule production.
	EnterBlockRule(c *BlockRuleContext)

	// EnterStatementsAppend is called when entering the StatementsAppend production.
	EnterStatementsAppend(c *StatementsAppendContext)

	// EnterStatementsInitial is called when entering the StatementsInitial production.
	EnterStatementsInitial(c *StatementsInitialContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterControl_statement is called when entering the control_statement production.
	EnterControl_statement(c *Control_statementContext)

	// EnterIfStatement is called when entering the IfStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfElseStatement is called when entering the IfElseStatement production.
	EnterIfElseStatement(c *IfElseStatementContext)

	// EnterElseStatement is called when entering the ElseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterElseIfStatement is called when entering the ElseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterForStatement is called when entering the ForStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForStatementWithAfter is called when entering the ForStatementWithAfter production.
	EnterForStatementWithAfter(c *ForStatementWithAfterContext)

	// EnterSwitchStatement is called when entering the SwitchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseBlocks is called when entering the CaseBlocks production.
	EnterCaseBlocks(c *CaseBlocksContext)

	// EnterOptionCaseBlock is called when entering the OptionCaseBlock production.
	EnterOptionCaseBlock(c *OptionCaseBlockContext)

	// EnterDefaultCaseBlock is called when entering the DefaultCaseBlock production.
	EnterDefaultCaseBlock(c *DefaultCaseBlockContext)

	// EnterInlineStatementsInitial is called when entering the InlineStatementsInitial production.
	EnterInlineStatementsInitial(c *InlineStatementsInitialContext)

	// EnterInlineStatementsAppend is called when entering the InlineStatementsAppend production.
	EnterInlineStatementsAppend(c *InlineStatementsAppendContext)

	// EnterInline_statement is called when entering the inline_statement production.
	EnterInline_statement(c *Inline_statementContext)

	// EnterAssignStatement is called when entering the AssignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterAssignOperationStatement is called when entering the AssignOperationStatement production.
	EnterAssignOperationStatement(c *AssignOperationStatementContext)

	// EnterAssignExpressions is called when entering the AssignExpressions production.
	EnterAssignExpressions(c *AssignExpressionsContext)

	// EnterExpressionAssignExpression is called when entering the ExpressionAssignExpression production.
	EnterExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// EnterStructDestrutureAssignExpression is called when entering the StructDestrutureAssignExpression production.
	EnterStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// EnterTupleDestrutureAssignExpression is called when entering the TupleDestrutureAssignExpression production.
	EnterTupleDestrutureAssignExpression(c *TupleDestrutureAssignExpressionContext)

	// EnterAssignDeclarations is called when entering the AssignDeclarations production.
	EnterAssignDeclarations(c *AssignDeclarationsContext)

	// EnterExpressionAssignDeclaration is called when entering the ExpressionAssignDeclaration production.
	EnterExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// EnterTypedAssignDeclarationDeclaration is called when entering the TypedAssignDeclarationDeclaration production.
	EnterTypedAssignDeclarationDeclaration(c *TypedAssignDeclarationDeclarationContext)

	// EnterUntypedAssignDeclarationDeclaration is called when entering the UntypedAssignDeclarationDeclaration production.
	EnterUntypedAssignDeclarationDeclaration(c *UntypedAssignDeclarationDeclarationContext)

	// EnterStructDestrutureAssignDeclaration is called when entering the StructDestrutureAssignDeclaration production.
	EnterStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// EnterTupleDestrutureAssignDeclaration is called when entering the TupleDestrutureAssignDeclaration production.
	EnterTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// EnterTypedAssignDeclarations is called when entering the TypedAssignDeclarations production.
	EnterTypedAssignDeclarations(c *TypedAssignDeclarationsContext)

	// EnterTypedAssignDeclaration is called when entering the TypedAssignDeclaration production.
	EnterTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// EnterUntypedAssignDeclarations is called when entering the UntypedAssignDeclarations production.
	EnterUntypedAssignDeclarations(c *UntypedAssignDeclarationsContext)

	// EnterUntypedIdentifierAssignDeclaration is called when entering the UntypedIdentifierAssignDeclaration production.
	EnterUntypedIdentifierAssignDeclaration(c *UntypedIdentifierAssignDeclarationContext)

	// EnterUntypedStructDestrutureAssignDeclaration is called when entering the UntypedStructDestrutureAssignDeclaration production.
	EnterUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// EnterUntypedTupleDestrutureAssignDeclaration is called when entering the UntypedTupleDestrutureAssignDeclaration production.
	EnterUntypedTupleDestrutureAssignDeclaration(c *UntypedTupleDestrutureAssignDeclarationContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterExpressionsRule is called when entering the ExpressionsRule production.
	EnterExpressionsRule(c *ExpressionsRuleContext)

	// EnterBracketedExpression is called when entering the BracketedExpression production.
	EnterBracketedExpression(c *BracketedExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterSliceExpression is called when entering the SliceExpression production.
	EnterSliceExpression(c *SliceExpressionContext)

	// EnterTupleExpression is called when entering the TupleExpression production.
	EnterTupleExpression(c *TupleExpressionContext)

	// EnterBooleanBinaryExpression is called when entering the BooleanBinaryExpression production.
	EnterBooleanBinaryExpression(c *BooleanBinaryExpressionContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterAccessExpression is called when entering the AccessExpression production.
	EnterAccessExpression(c *AccessExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterStructExpression is called when entering the StructExpression production.
	EnterStructExpression(c *StructExpressionContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterCastExpression is called when entering the CastExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterCallExpression is called when entering the CallExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterFunctionTypeFunc is called when entering the FunctionTypeFunc production.
	EnterFunctionTypeFunc(c *FunctionTypeFuncContext)

	// EnterFunctionTypeExpression is called when entering the FunctionTypeExpression production.
	EnterFunctionTypeExpression(c *FunctionTypeExpressionContext)

	// EnterStructTypeStruct is called when entering the StructTypeStruct production.
	EnterStructTypeStruct(c *StructTypeStructContext)

	// EnterStructTypeExpression is called when entering the StructTypeExpression production.
	EnterStructTypeExpression(c *StructTypeExpressionContext)

	// EnterTupleTypeTuple is called when entering the TupleTypeTuple production.
	EnterTupleTypeTuple(c *TupleTypeTupleContext)

	// EnterTupleTypeExpression is called when entering the TupleTypeExpression production.
	EnterTupleTypeExpression(c *TupleTypeExpressionContext)

	// EnterAccessTypeExpression is called when entering the AccessTypeExpression production.
	EnterAccessTypeExpression(c *AccessTypeExpressionContext)

	// EnterIdentifierTypeExpression is called when entering the IdentifierTypeExpression production.
	EnterIdentifierTypeExpression(c *IdentifierTypeExpressionContext)

	// EnterIdentifiersRule is called when entering the IdentifiersRule production.
	EnterIdentifiersRule(c *IdentifiersRuleContext)

	// EnterUntypedLiteral is called when entering the UntypedLiteral production.
	EnterUntypedLiteral(c *UntypedLiteralContext)

	// EnterTypedLiteral is called when entering the TypedLiteral production.
	EnterTypedLiteral(c *TypedLiteralContext)

	// ExitModuleRule is called when exiting the ModuleRule production.
	ExitModuleRule(c *ModuleRuleContext)

	// ExitDefinitionsAppend is called when exiting the DefinitionsAppend production.
	ExitDefinitionsAppend(c *DefinitionsAppendContext)

	// ExitDefinitionsInitial is called when exiting the DefinitionsInitial production.
	ExitDefinitionsInitial(c *DefinitionsInitialContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitFunctionDefinition is called when exiting the FunctionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionDefinitionNoReturnType is called when exiting the FunctionDefinitionNoReturnType production.
	ExitFunctionDefinitionNoReturnType(c *FunctionDefinitionNoReturnTypeContext)

	// ExitParametersEmpty is called when exiting the ParametersEmpty production.
	ExitParametersEmpty(c *ParametersEmptyContext)

	// ExitParametersFilled is called when exiting the ParametersFilled production.
	ExitParametersFilled(c *ParametersFilledContext)

	// ExitParametersListInitial is called when exiting the ParametersListInitial production.
	ExitParametersListInitial(c *ParametersListInitialContext)

	// ExitParametersListAppend is called when exiting the ParametersListAppend production.
	ExitParametersListAppend(c *ParametersListAppendContext)

	// ExitTypedParameter is called when exiting the TypedParameter production.
	ExitTypedParameter(c *TypedParameterContext)

	// ExitUntypedParameter is called when exiting the UntypedParameter production.
	ExitUntypedParameter(c *UntypedParameterContext)

	// ExitStructDestrutureParameter is called when exiting the StructDestrutureParameter production.
	ExitStructDestrutureParameter(c *StructDestrutureParameterContext)

	// ExitTupleDestrutureParameter is called when exiting the TupleDestrutureParameter production.
	ExitTupleDestrutureParameter(c *TupleDestrutureParameterContext)

	// ExitBlockRule is called when exiting the BlockRule production.
	ExitBlockRule(c *BlockRuleContext)

	// ExitStatementsAppend is called when exiting the StatementsAppend production.
	ExitStatementsAppend(c *StatementsAppendContext)

	// ExitStatementsInitial is called when exiting the StatementsInitial production.
	ExitStatementsInitial(c *StatementsInitialContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitControl_statement is called when exiting the control_statement production.
	ExitControl_statement(c *Control_statementContext)

	// ExitIfStatement is called when exiting the IfStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfElseStatement is called when exiting the IfElseStatement production.
	ExitIfElseStatement(c *IfElseStatementContext)

	// ExitElseStatement is called when exiting the ElseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitElseIfStatement is called when exiting the ElseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitForStatement is called when exiting the ForStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForStatementWithAfter is called when exiting the ForStatementWithAfter production.
	ExitForStatementWithAfter(c *ForStatementWithAfterContext)

	// ExitSwitchStatement is called when exiting the SwitchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseBlocks is called when exiting the CaseBlocks production.
	ExitCaseBlocks(c *CaseBlocksContext)

	// ExitOptionCaseBlock is called when exiting the OptionCaseBlock production.
	ExitOptionCaseBlock(c *OptionCaseBlockContext)

	// ExitDefaultCaseBlock is called when exiting the DefaultCaseBlock production.
	ExitDefaultCaseBlock(c *DefaultCaseBlockContext)

	// ExitInlineStatementsInitial is called when exiting the InlineStatementsInitial production.
	ExitInlineStatementsInitial(c *InlineStatementsInitialContext)

	// ExitInlineStatementsAppend is called when exiting the InlineStatementsAppend production.
	ExitInlineStatementsAppend(c *InlineStatementsAppendContext)

	// ExitInline_statement is called when exiting the inline_statement production.
	ExitInline_statement(c *Inline_statementContext)

	// ExitAssignStatement is called when exiting the AssignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitAssignOperationStatement is called when exiting the AssignOperationStatement production.
	ExitAssignOperationStatement(c *AssignOperationStatementContext)

	// ExitAssignExpressions is called when exiting the AssignExpressions production.
	ExitAssignExpressions(c *AssignExpressionsContext)

	// ExitExpressionAssignExpression is called when exiting the ExpressionAssignExpression production.
	ExitExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// ExitStructDestrutureAssignExpression is called when exiting the StructDestrutureAssignExpression production.
	ExitStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// ExitTupleDestrutureAssignExpression is called when exiting the TupleDestrutureAssignExpression production.
	ExitTupleDestrutureAssignExpression(c *TupleDestrutureAssignExpressionContext)

	// ExitAssignDeclarations is called when exiting the AssignDeclarations production.
	ExitAssignDeclarations(c *AssignDeclarationsContext)

	// ExitExpressionAssignDeclaration is called when exiting the ExpressionAssignDeclaration production.
	ExitExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// ExitTypedAssignDeclarationDeclaration is called when exiting the TypedAssignDeclarationDeclaration production.
	ExitTypedAssignDeclarationDeclaration(c *TypedAssignDeclarationDeclarationContext)

	// ExitUntypedAssignDeclarationDeclaration is called when exiting the UntypedAssignDeclarationDeclaration production.
	ExitUntypedAssignDeclarationDeclaration(c *UntypedAssignDeclarationDeclarationContext)

	// ExitStructDestrutureAssignDeclaration is called when exiting the StructDestrutureAssignDeclaration production.
	ExitStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// ExitTupleDestrutureAssignDeclaration is called when exiting the TupleDestrutureAssignDeclaration production.
	ExitTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// ExitTypedAssignDeclarations is called when exiting the TypedAssignDeclarations production.
	ExitTypedAssignDeclarations(c *TypedAssignDeclarationsContext)

	// ExitTypedAssignDeclaration is called when exiting the TypedAssignDeclaration production.
	ExitTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// ExitUntypedAssignDeclarations is called when exiting the UntypedAssignDeclarations production.
	ExitUntypedAssignDeclarations(c *UntypedAssignDeclarationsContext)

	// ExitUntypedIdentifierAssignDeclaration is called when exiting the UntypedIdentifierAssignDeclaration production.
	ExitUntypedIdentifierAssignDeclaration(c *UntypedIdentifierAssignDeclarationContext)

	// ExitUntypedStructDestrutureAssignDeclaration is called when exiting the UntypedStructDestrutureAssignDeclaration production.
	ExitUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// ExitUntypedTupleDestrutureAssignDeclaration is called when exiting the UntypedTupleDestrutureAssignDeclaration production.
	ExitUntypedTupleDestrutureAssignDeclaration(c *UntypedTupleDestrutureAssignDeclarationContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitExpressionsRule is called when exiting the ExpressionsRule production.
	ExitExpressionsRule(c *ExpressionsRuleContext)

	// ExitBracketedExpression is called when exiting the BracketedExpression production.
	ExitBracketedExpression(c *BracketedExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitSliceExpression is called when exiting the SliceExpression production.
	ExitSliceExpression(c *SliceExpressionContext)

	// ExitTupleExpression is called when exiting the TupleExpression production.
	ExitTupleExpression(c *TupleExpressionContext)

	// ExitBooleanBinaryExpression is called when exiting the BooleanBinaryExpression production.
	ExitBooleanBinaryExpression(c *BooleanBinaryExpressionContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitAccessExpression is called when exiting the AccessExpression production.
	ExitAccessExpression(c *AccessExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitStructExpression is called when exiting the StructExpression production.
	ExitStructExpression(c *StructExpressionContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitCastExpression is called when exiting the CastExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitCallExpression is called when exiting the CallExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitFunctionTypeFunc is called when exiting the FunctionTypeFunc production.
	ExitFunctionTypeFunc(c *FunctionTypeFuncContext)

	// ExitFunctionTypeExpression is called when exiting the FunctionTypeExpression production.
	ExitFunctionTypeExpression(c *FunctionTypeExpressionContext)

	// ExitStructTypeStruct is called when exiting the StructTypeStruct production.
	ExitStructTypeStruct(c *StructTypeStructContext)

	// ExitStructTypeExpression is called when exiting the StructTypeExpression production.
	ExitStructTypeExpression(c *StructTypeExpressionContext)

	// ExitTupleTypeTuple is called when exiting the TupleTypeTuple production.
	ExitTupleTypeTuple(c *TupleTypeTupleContext)

	// ExitTupleTypeExpression is called when exiting the TupleTypeExpression production.
	ExitTupleTypeExpression(c *TupleTypeExpressionContext)

	// ExitAccessTypeExpression is called when exiting the AccessTypeExpression production.
	ExitAccessTypeExpression(c *AccessTypeExpressionContext)

	// ExitIdentifierTypeExpression is called when exiting the IdentifierTypeExpression production.
	ExitIdentifierTypeExpression(c *IdentifierTypeExpressionContext)

	// ExitIdentifiersRule is called when exiting the IdentifiersRule production.
	ExitIdentifiersRule(c *IdentifiersRuleContext)

	// ExitUntypedLiteral is called when exiting the UntypedLiteral production.
	ExitUntypedLiteral(c *UntypedLiteralContext)

	// ExitTypedLiteral is called when exiting the TypedLiteral production.
	ExitTypedLiteral(c *TypedLiteralContext)
}
