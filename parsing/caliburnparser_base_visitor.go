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

func (v *BaseCaliburnParserVisitor) VisitTupleDestrutureParameter(ctx *TupleDestrutureParameterContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitOptionCaseBlock(ctx *OptionCaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefaultCaseBlock(ctx *DefaultCaseBlockContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignOperationStatement(ctx *AssignOperationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignExpressions(ctx *AssignExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionAssignExpression(ctx *ExpressionAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructDestrutureAssignExpression(ctx *StructDestrutureAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTupleDestrutureAssignExpression(ctx *TupleDestrutureAssignExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssignDeclarations(ctx *AssignDeclarationsContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedAssignDeclarations(ctx *TypedAssignDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedAssignDeclarations(ctx *UntypedAssignDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedIdentifierAssignDeclaration(ctx *UntypedIdentifierAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedStructDestrutureAssignDeclaration(ctx *UntypedStructDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedTupleDestrutureAssignDeclaration(ctx *UntypedTupleDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressionsRule(ctx *ExpressionsRuleContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBooleanBinaryExpression(ctx *BooleanBinaryExpressionContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitCastExpression(ctx *CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitStructTypeExpression(ctx *StructTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTupleTypeTuple(ctx *TupleTypeTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTupleTypeExpression(ctx *TupleTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAccessTypeExpression(ctx *AccessTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifierTypeExpression(ctx *IdentifierTypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifiersRule(ctx *IdentifiersRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedLiteral(ctx *UntypedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedLiteral(ctx *TypedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
