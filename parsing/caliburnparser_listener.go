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

	// EnterBlockRule is called when entering the BlockRule production.
	EnterBlockRule(c *BlockRuleContext)

	// EnterStatementsAppend is called when entering the StatementsAppend production.
	EnterStatementsAppend(c *StatementsAppendContext)

	// EnterStatementsInitial is called when entering the StatementsInitial production.
	EnterStatementsInitial(c *StatementsInitialContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterInlineStatementsInitial is called when entering the InlineStatementsInitial production.
	EnterInlineStatementsInitial(c *InlineStatementsInitialContext)

	// EnterInlineStatementsAppend is called when entering the InlineStatementsAppend production.
	EnterInlineStatementsAppend(c *InlineStatementsAppendContext)

	// EnterInline_statement is called when entering the inline_statement production.
	EnterInline_statement(c *Inline_statementContext)

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

	// EnterCaseBlocksDefault is called when entering the CaseBlocksDefault production.
	EnterCaseBlocksDefault(c *CaseBlocksDefaultContext)

	// EnterOptionCaseBlocksAppend is called when entering the OptionCaseBlocksAppend production.
	EnterOptionCaseBlocksAppend(c *OptionCaseBlocksAppendContext)

	// EnterOptionCaseBlocksInitial is called when entering the OptionCaseBlocksInitial production.
	EnterOptionCaseBlocksInitial(c *OptionCaseBlocksInitialContext)

	// EnterOptionCaseBlock is called when entering the OptionCaseBlock production.
	EnterOptionCaseBlock(c *OptionCaseBlockContext)

	// EnterDefaultCaseBlock is called when entering the DefaultCaseBlock production.
	EnterDefaultCaseBlock(c *DefaultCaseBlockContext)

	// EnterAssignStatement is called when entering the AssignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterAssignOperationStatement is called when entering the AssignOperationStatement production.
	EnterAssignOperationStatement(c *AssignOperationStatementContext)

	// EnterAssignExpressionsAppend is called when entering the AssignExpressionsAppend production.
	EnterAssignExpressionsAppend(c *AssignExpressionsAppendContext)

	// EnterAssignExpressionsInitial is called when entering the AssignExpressionsInitial production.
	EnterAssignExpressionsInitial(c *AssignExpressionsInitialContext)

	// EnterAliasableAssignExpressionsInitial is called when entering the AliasableAssignExpressionsInitial production.
	EnterAliasableAssignExpressionsInitial(c *AliasableAssignExpressionsInitialContext)

	// EnterAliasableAssignExpressionsAppend is called when entering the AliasableAssignExpressionsAppend production.
	EnterAliasableAssignExpressionsAppend(c *AliasableAssignExpressionsAppendContext)

	// EnterAliasedAssignExpression is called when entering the AliasedAssignExpression production.
	EnterAliasedAssignExpression(c *AliasedAssignExpressionContext)

	// EnterUnaliasedAssignExpression is called when entering the UnaliasedAssignExpression production.
	EnterUnaliasedAssignExpression(c *UnaliasedAssignExpressionContext)

	// EnterExpressionAssignExpression is called when entering the ExpressionAssignExpression production.
	EnterExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// EnterStructDestrutureAssignExpression is called when entering the StructDestrutureAssignExpression production.
	EnterStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// EnterAssignDeclarationsAppend is called when entering the AssignDeclarationsAppend production.
	EnterAssignDeclarationsAppend(c *AssignDeclarationsAppendContext)

	// EnterAssignDeclarationsInitial is called when entering the AssignDeclarationsInitial production.
	EnterAssignDeclarationsInitial(c *AssignDeclarationsInitialContext)

	// EnterAliasableAssignDeclarationsAppend is called when entering the AliasableAssignDeclarationsAppend production.
	EnterAliasableAssignDeclarationsAppend(c *AliasableAssignDeclarationsAppendContext)

	// EnterAliasableAssignDeclarationsInitial is called when entering the AliasableAssignDeclarationsInitial production.
	EnterAliasableAssignDeclarationsInitial(c *AliasableAssignDeclarationsInitialContext)

	// EnterAliasedAssignDeclaration is called when entering the AliasedAssignDeclaration production.
	EnterAliasedAssignDeclaration(c *AliasedAssignDeclarationContext)

	// EnterUnaliasedAssignDeclaration is called when entering the UnaliasedAssignDeclaration production.
	EnterUnaliasedAssignDeclaration(c *UnaliasedAssignDeclarationContext)

	// EnterExpressionAssignDeclaration is called when entering the ExpressionAssignDeclaration production.
	EnterExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// EnterTypedAssignDeclarationDeclaration is called when entering the TypedAssignDeclarationDeclaration production.
	EnterTypedAssignDeclarationDeclaration(c *TypedAssignDeclarationDeclarationContext)

	// EnterUntypedAssignDeclarationDeclaration is called when entering the UntypedAssignDeclarationDeclaration production.
	EnterUntypedAssignDeclarationDeclaration(c *UntypedAssignDeclarationDeclarationContext)

	// EnterStructDestrutureAssignDeclaration is called when entering the StructDestrutureAssignDeclaration production.
	EnterStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// EnterTypedAssignDeclaration is called when entering the TypedAssignDeclaration production.
	EnterTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// EnterAliasableUntypedAssignDeclarationsInitial is called when entering the AliasableUntypedAssignDeclarationsInitial production.
	EnterAliasableUntypedAssignDeclarationsInitial(c *AliasableUntypedAssignDeclarationsInitialContext)

	// EnterAliasableUntypedAssignDeclarationsAppend is called when entering the AliasableUntypedAssignDeclarationsAppend production.
	EnterAliasableUntypedAssignDeclarationsAppend(c *AliasableUntypedAssignDeclarationsAppendContext)

	// EnterAliasedUntypedAssignDeclaration is called when entering the AliasedUntypedAssignDeclaration production.
	EnterAliasedUntypedAssignDeclaration(c *AliasedUntypedAssignDeclarationContext)

	// EnterUnaliasedUntypedAssignDeclaration is called when entering the UnaliasedUntypedAssignDeclaration production.
	EnterUnaliasedUntypedAssignDeclaration(c *UnaliasedUntypedAssignDeclarationContext)

	// EnterUntypedIdentifierAssignDeclaration is called when entering the UntypedIdentifierAssignDeclaration production.
	EnterUntypedIdentifierAssignDeclaration(c *UntypedIdentifierAssignDeclarationContext)

	// EnterUntypedStructDestrutureAssignDeclaration is called when entering the UntypedStructDestrutureAssignDeclaration production.
	EnterUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// EnterExpressionStatement is called when entering the ExpressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterExpressionsInitial is called when entering the ExpressionsInitial production.
	EnterExpressionsInitial(c *ExpressionsInitialContext)

	// EnterExpressionsAppend is called when entering the ExpressionsAppend production.
	EnterExpressionsAppend(c *ExpressionsAppendContext)

	// EnterExpressionsOptional is called when entering the ExpressionsOptional production.
	EnterExpressionsOptional(c *ExpressionsOptionalContext)

	// EnterExpressionsOptionalNone is called when entering the ExpressionsOptionalNone production.
	EnterExpressionsOptionalNone(c *ExpressionsOptionalNoneContext)

	// EnterBracketedExpression is called when entering the BracketedExpression production.
	EnterBracketedExpression(c *BracketedExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterSliceExpression is called when entering the SliceExpression production.
	EnterSliceExpression(c *SliceExpressionContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterAccessExpression is called when entering the AccessExpression production.
	EnterAccessExpression(c *AccessExpressionContext)

	// EnterSliceEndExpression is called when entering the SliceEndExpression production.
	EnterSliceEndExpression(c *SliceEndExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterStructExpression is called when entering the StructExpression production.
	EnterStructExpression(c *StructExpressionContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterSliceStartExpression is called when entering the SliceStartExpression production.
	EnterSliceStartExpression(c *SliceStartExpressionContext)

	// EnterFunctionExpressionNoReturnType is called when entering the FunctionExpressionNoReturnType production.
	EnterFunctionExpressionNoReturnType(c *FunctionExpressionNoReturnTypeContext)

	// EnterCollectionExpression is called when entering the CollectionExpression production.
	EnterCollectionExpression(c *CollectionExpressionContext)

	// EnterCallExpression is called when entering the CallExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterStructValuesNamed is called when entering the StructValuesNamed production.
	EnterStructValuesNamed(c *StructValuesNamedContext)

	// EnterStructValuesUnamed is called when entering the StructValuesUnamed production.
	EnterStructValuesUnamed(c *StructValuesUnamedContext)

	// EnterNamedStructValuesAppend is called when entering the NamedStructValuesAppend production.
	EnterNamedStructValuesAppend(c *NamedStructValuesAppendContext)

	// EnterNamedStructValuesInitial is called when entering the NamedStructValuesInitial production.
	EnterNamedStructValuesInitial(c *NamedStructValuesInitialContext)

	// EnterNamedStructValue is called when entering the NamedStructValue production.
	EnterNamedStructValue(c *NamedStructValueContext)

	// EnterCollectionValuesInitial is called when entering the CollectionValuesInitial production.
	EnterCollectionValuesInitial(c *CollectionValuesInitialContext)

	// EnterCollectionValuesAppend is called when entering the CollectionValuesAppend production.
	EnterCollectionValuesAppend(c *CollectionValuesAppendContext)

	// EnterCollectionValue is called when entering the CollectionValue production.
	EnterCollectionValue(c *CollectionValueContext)

	// EnterFunctionTypeFunc is called when entering the FunctionTypeFunc production.
	EnterFunctionTypeFunc(c *FunctionTypeFuncContext)

	// EnterFunctionTypeExpression is called when entering the FunctionTypeExpression production.
	EnterFunctionTypeExpression(c *FunctionTypeExpressionContext)

	// EnterStructTypeStruct is called when entering the StructTypeStruct production.
	EnterStructTypeStruct(c *StructTypeStructContext)

	// EnterStructTypeTuple is called when entering the StructTypeTuple production.
	EnterStructTypeTuple(c *StructTypeTupleContext)

	// EnterStructTypeExpression is called when entering the StructTypeExpression production.
	EnterStructTypeExpression(c *StructTypeExpressionContext)

	// EnterAccessTypeExpression is called when entering the AccessTypeExpression production.
	EnterAccessTypeExpression(c *AccessTypeExpressionContext)

	// EnterIdentifierTypeExpression is called when entering the IdentifierTypeExpression production.
	EnterIdentifierTypeExpression(c *IdentifierTypeExpressionContext)

	// EnterUntypedLiteral is called when entering the UntypedLiteral production.
	EnterUntypedLiteral(c *UntypedLiteralContext)

	// EnterTypedLiteral is called when entering the TypedLiteral production.
	EnterTypedLiteral(c *TypedLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

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

	// ExitBlockRule is called when exiting the BlockRule production.
	ExitBlockRule(c *BlockRuleContext)

	// ExitStatementsAppend is called when exiting the StatementsAppend production.
	ExitStatementsAppend(c *StatementsAppendContext)

	// ExitStatementsInitial is called when exiting the StatementsInitial production.
	ExitStatementsInitial(c *StatementsInitialContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitInlineStatementsInitial is called when exiting the InlineStatementsInitial production.
	ExitInlineStatementsInitial(c *InlineStatementsInitialContext)

	// ExitInlineStatementsAppend is called when exiting the InlineStatementsAppend production.
	ExitInlineStatementsAppend(c *InlineStatementsAppendContext)

	// ExitInline_statement is called when exiting the inline_statement production.
	ExitInline_statement(c *Inline_statementContext)

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

	// ExitCaseBlocksDefault is called when exiting the CaseBlocksDefault production.
	ExitCaseBlocksDefault(c *CaseBlocksDefaultContext)

	// ExitOptionCaseBlocksAppend is called when exiting the OptionCaseBlocksAppend production.
	ExitOptionCaseBlocksAppend(c *OptionCaseBlocksAppendContext)

	// ExitOptionCaseBlocksInitial is called when exiting the OptionCaseBlocksInitial production.
	ExitOptionCaseBlocksInitial(c *OptionCaseBlocksInitialContext)

	// ExitOptionCaseBlock is called when exiting the OptionCaseBlock production.
	ExitOptionCaseBlock(c *OptionCaseBlockContext)

	// ExitDefaultCaseBlock is called when exiting the DefaultCaseBlock production.
	ExitDefaultCaseBlock(c *DefaultCaseBlockContext)

	// ExitAssignStatement is called when exiting the AssignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitAssignOperationStatement is called when exiting the AssignOperationStatement production.
	ExitAssignOperationStatement(c *AssignOperationStatementContext)

	// ExitAssignExpressionsAppend is called when exiting the AssignExpressionsAppend production.
	ExitAssignExpressionsAppend(c *AssignExpressionsAppendContext)

	// ExitAssignExpressionsInitial is called when exiting the AssignExpressionsInitial production.
	ExitAssignExpressionsInitial(c *AssignExpressionsInitialContext)

	// ExitAliasableAssignExpressionsInitial is called when exiting the AliasableAssignExpressionsInitial production.
	ExitAliasableAssignExpressionsInitial(c *AliasableAssignExpressionsInitialContext)

	// ExitAliasableAssignExpressionsAppend is called when exiting the AliasableAssignExpressionsAppend production.
	ExitAliasableAssignExpressionsAppend(c *AliasableAssignExpressionsAppendContext)

	// ExitAliasedAssignExpression is called when exiting the AliasedAssignExpression production.
	ExitAliasedAssignExpression(c *AliasedAssignExpressionContext)

	// ExitUnaliasedAssignExpression is called when exiting the UnaliasedAssignExpression production.
	ExitUnaliasedAssignExpression(c *UnaliasedAssignExpressionContext)

	// ExitExpressionAssignExpression is called when exiting the ExpressionAssignExpression production.
	ExitExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// ExitStructDestrutureAssignExpression is called when exiting the StructDestrutureAssignExpression production.
	ExitStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// ExitAssignDeclarationsAppend is called when exiting the AssignDeclarationsAppend production.
	ExitAssignDeclarationsAppend(c *AssignDeclarationsAppendContext)

	// ExitAssignDeclarationsInitial is called when exiting the AssignDeclarationsInitial production.
	ExitAssignDeclarationsInitial(c *AssignDeclarationsInitialContext)

	// ExitAliasableAssignDeclarationsAppend is called when exiting the AliasableAssignDeclarationsAppend production.
	ExitAliasableAssignDeclarationsAppend(c *AliasableAssignDeclarationsAppendContext)

	// ExitAliasableAssignDeclarationsInitial is called when exiting the AliasableAssignDeclarationsInitial production.
	ExitAliasableAssignDeclarationsInitial(c *AliasableAssignDeclarationsInitialContext)

	// ExitAliasedAssignDeclaration is called when exiting the AliasedAssignDeclaration production.
	ExitAliasedAssignDeclaration(c *AliasedAssignDeclarationContext)

	// ExitUnaliasedAssignDeclaration is called when exiting the UnaliasedAssignDeclaration production.
	ExitUnaliasedAssignDeclaration(c *UnaliasedAssignDeclarationContext)

	// ExitExpressionAssignDeclaration is called when exiting the ExpressionAssignDeclaration production.
	ExitExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// ExitTypedAssignDeclarationDeclaration is called when exiting the TypedAssignDeclarationDeclaration production.
	ExitTypedAssignDeclarationDeclaration(c *TypedAssignDeclarationDeclarationContext)

	// ExitUntypedAssignDeclarationDeclaration is called when exiting the UntypedAssignDeclarationDeclaration production.
	ExitUntypedAssignDeclarationDeclaration(c *UntypedAssignDeclarationDeclarationContext)

	// ExitStructDestrutureAssignDeclaration is called when exiting the StructDestrutureAssignDeclaration production.
	ExitStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// ExitTypedAssignDeclaration is called when exiting the TypedAssignDeclaration production.
	ExitTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// ExitAliasableUntypedAssignDeclarationsInitial is called when exiting the AliasableUntypedAssignDeclarationsInitial production.
	ExitAliasableUntypedAssignDeclarationsInitial(c *AliasableUntypedAssignDeclarationsInitialContext)

	// ExitAliasableUntypedAssignDeclarationsAppend is called when exiting the AliasableUntypedAssignDeclarationsAppend production.
	ExitAliasableUntypedAssignDeclarationsAppend(c *AliasableUntypedAssignDeclarationsAppendContext)

	// ExitAliasedUntypedAssignDeclaration is called when exiting the AliasedUntypedAssignDeclaration production.
	ExitAliasedUntypedAssignDeclaration(c *AliasedUntypedAssignDeclarationContext)

	// ExitUnaliasedUntypedAssignDeclaration is called when exiting the UnaliasedUntypedAssignDeclaration production.
	ExitUnaliasedUntypedAssignDeclaration(c *UnaliasedUntypedAssignDeclarationContext)

	// ExitUntypedIdentifierAssignDeclaration is called when exiting the UntypedIdentifierAssignDeclaration production.
	ExitUntypedIdentifierAssignDeclaration(c *UntypedIdentifierAssignDeclarationContext)

	// ExitUntypedStructDestrutureAssignDeclaration is called when exiting the UntypedStructDestrutureAssignDeclaration production.
	ExitUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// ExitExpressionStatement is called when exiting the ExpressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitExpressionsInitial is called when exiting the ExpressionsInitial production.
	ExitExpressionsInitial(c *ExpressionsInitialContext)

	// ExitExpressionsAppend is called when exiting the ExpressionsAppend production.
	ExitExpressionsAppend(c *ExpressionsAppendContext)

	// ExitExpressionsOptional is called when exiting the ExpressionsOptional production.
	ExitExpressionsOptional(c *ExpressionsOptionalContext)

	// ExitExpressionsOptionalNone is called when exiting the ExpressionsOptionalNone production.
	ExitExpressionsOptionalNone(c *ExpressionsOptionalNoneContext)

	// ExitBracketedExpression is called when exiting the BracketedExpression production.
	ExitBracketedExpression(c *BracketedExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitSliceExpression is called when exiting the SliceExpression production.
	ExitSliceExpression(c *SliceExpressionContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitAccessExpression is called when exiting the AccessExpression production.
	ExitAccessExpression(c *AccessExpressionContext)

	// ExitSliceEndExpression is called when exiting the SliceEndExpression production.
	ExitSliceEndExpression(c *SliceEndExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitStructExpression is called when exiting the StructExpression production.
	ExitStructExpression(c *StructExpressionContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitSliceStartExpression is called when exiting the SliceStartExpression production.
	ExitSliceStartExpression(c *SliceStartExpressionContext)

	// ExitFunctionExpressionNoReturnType is called when exiting the FunctionExpressionNoReturnType production.
	ExitFunctionExpressionNoReturnType(c *FunctionExpressionNoReturnTypeContext)

	// ExitCollectionExpression is called when exiting the CollectionExpression production.
	ExitCollectionExpression(c *CollectionExpressionContext)

	// ExitCallExpression is called when exiting the CallExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitStructValuesNamed is called when exiting the StructValuesNamed production.
	ExitStructValuesNamed(c *StructValuesNamedContext)

	// ExitStructValuesUnamed is called when exiting the StructValuesUnamed production.
	ExitStructValuesUnamed(c *StructValuesUnamedContext)

	// ExitNamedStructValuesAppend is called when exiting the NamedStructValuesAppend production.
	ExitNamedStructValuesAppend(c *NamedStructValuesAppendContext)

	// ExitNamedStructValuesInitial is called when exiting the NamedStructValuesInitial production.
	ExitNamedStructValuesInitial(c *NamedStructValuesInitialContext)

	// ExitNamedStructValue is called when exiting the NamedStructValue production.
	ExitNamedStructValue(c *NamedStructValueContext)

	// ExitCollectionValuesInitial is called when exiting the CollectionValuesInitial production.
	ExitCollectionValuesInitial(c *CollectionValuesInitialContext)

	// ExitCollectionValuesAppend is called when exiting the CollectionValuesAppend production.
	ExitCollectionValuesAppend(c *CollectionValuesAppendContext)

	// ExitCollectionValue is called when exiting the CollectionValue production.
	ExitCollectionValue(c *CollectionValueContext)

	// ExitFunctionTypeFunc is called when exiting the FunctionTypeFunc production.
	ExitFunctionTypeFunc(c *FunctionTypeFuncContext)

	// ExitFunctionTypeExpression is called when exiting the FunctionTypeExpression production.
	ExitFunctionTypeExpression(c *FunctionTypeExpressionContext)

	// ExitStructTypeStruct is called when exiting the StructTypeStruct production.
	ExitStructTypeStruct(c *StructTypeStructContext)

	// ExitStructTypeTuple is called when exiting the StructTypeTuple production.
	ExitStructTypeTuple(c *StructTypeTupleContext)

	// ExitStructTypeExpression is called when exiting the StructTypeExpression production.
	ExitStructTypeExpression(c *StructTypeExpressionContext)

	// ExitAccessTypeExpression is called when exiting the AccessTypeExpression production.
	ExitAccessTypeExpression(c *AccessTypeExpressionContext)

	// ExitIdentifierTypeExpression is called when exiting the IdentifierTypeExpression production.
	ExitIdentifierTypeExpression(c *IdentifierTypeExpressionContext)

	// ExitUntypedLiteral is called when exiting the UntypedLiteral production.
	ExitUntypedLiteral(c *UntypedLiteralContext)

	// ExitTypedLiteral is called when exiting the TypedLiteral production.
	ExitTypedLiteral(c *TypedLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
