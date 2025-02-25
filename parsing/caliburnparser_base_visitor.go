// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

type BaseCaliburnParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCaliburnParserVisitor) VisitModuleRule(ctx *ModuleRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefinitionsAppend(ctx *DefinitionsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefinitionsInitial(ctx *DefinitionsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionDefinitionNoReturnType(ctx *FunctionDefinitionNoReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitParametersEmpty(ctx *ParametersEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitParametersFilled(ctx *ParametersFilledContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitParametersListInitial(ctx *ParametersListInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitParametersListAppend(ctx *ParametersListAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedParameter(ctx *TypedParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedParameter(ctx *UntypedParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructDestrutureParameter(ctx *StructDestrutureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBlockRule(ctx *BlockRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStatementsAppend(ctx *StatementsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStatementsInitial(ctx *StatementsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitInlineStatementsInitial(ctx *InlineStatementsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitInlineStatementsAppend(ctx *InlineStatementsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitInline_statement(ctx *Inline_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitJump_statement(ctx *Jump_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitControl_statement(ctx *Control_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitElseIfStatement(ctx *ElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitForStatementWithAfter(ctx *ForStatementWithAfterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCaseBlocks(ctx *CaseBlocksContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCaseBlocksDefault(ctx *CaseBlocksDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitOptionCaseBlocksAppend(ctx *OptionCaseBlocksAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitOptionCaseBlocksInitial(ctx *OptionCaseBlocksInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitOptionCaseBlock(ctx *OptionCaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefaultCaseBlock(ctx *DefaultCaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignOperationStatement(ctx *AssignOperationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignExpressionsAppend(ctx *AssignExpressionsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignExpressionsInitial(ctx *AssignExpressionsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableAssignExpressionsInitial(ctx *AliasableAssignExpressionsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableAssignExpressionsAppend(ctx *AliasableAssignExpressionsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasedAssignExpression(ctx *AliasedAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUnaliasedAssignExpression(ctx *UnaliasedAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionAssignExpression(ctx *ExpressionAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructDestrutureAssignExpression(ctx *StructDestrutureAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignDeclarationsAppend(ctx *AssignDeclarationsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignDeclarationsInitial(ctx *AssignDeclarationsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableAssignDeclarationsAppend(ctx *AliasableAssignDeclarationsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableAssignDeclarationsInitial(ctx *AliasableAssignDeclarationsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasedAssignDeclaration(ctx *AliasedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUnaliasedAssignDeclaration(ctx *UnaliasedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedAssignDeclarationDeclaration(ctx *TypedAssignDeclarationDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedAssignDeclarationDeclaration(ctx *UntypedAssignDeclarationDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableUntypedAssignDeclarationsInitial(ctx *AliasableUntypedAssignDeclarationsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasableUntypedAssignDeclarationsAppend(ctx *AliasableUntypedAssignDeclarationsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAliasedUntypedAssignDeclaration(ctx *AliasedUntypedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUnaliasedUntypedAssignDeclaration(ctx *UnaliasedUntypedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedIdentifierAssignDeclaration(ctx *UntypedIdentifierAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedStructDestrutureAssignDeclaration(ctx *UntypedStructDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionsInitial(ctx *ExpressionsInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionsAppend(ctx *ExpressionsAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionsOptional(ctx *ExpressionsOptionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionsOptionalNone(ctx *ExpressionsOptionalNoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBracketedExpression(ctx *BracketedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitSliceExpression(ctx *SliceExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIndexExpression(ctx *IndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAccessExpression(ctx *AccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitSliceEndExpression(ctx *SliceEndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructExpression(ctx *StructExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitSliceStartExpression(ctx *SliceStartExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionExpressionNoReturnType(ctx *FunctionExpressionNoReturnTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCollectionExpression(ctx *CollectionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructValuesNamed(ctx *StructValuesNamedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructValuesUnamed(ctx *StructValuesUnamedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitNamedStructValuesAppend(ctx *NamedStructValuesAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitNamedStructValuesInitial(ctx *NamedStructValuesInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitNamedStructValue(ctx *NamedStructValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCollectionValuesInitial(ctx *CollectionValuesInitialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCollectionValuesAppend(ctx *CollectionValuesAppendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCollectionValue(ctx *CollectionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionTypeFunc(ctx *FunctionTypeFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunctionTypeExpression(ctx *FunctionTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructTypeStruct(ctx *StructTypeStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructTypeTuple(ctx *StructTypeTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructTypeExpression(ctx *StructTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAccessTypeExpression(ctx *AccessTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifierTypeExpression(ctx *IdentifierTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedLiteral(ctx *UntypedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedLiteral(ctx *TypedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
