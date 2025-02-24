// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CaliburnParser.
type CaliburnParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CaliburnParser#ModuleRule.
	VisitModuleRule(ctx *ModuleRuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#FunctionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ParametersRule.
	VisitParametersRule(ctx *ParametersRuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedParameter.
	VisitTypedParameter(ctx *TypedParameterContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedParameter.
	VisitUntypedParameter(ctx *UntypedParameterContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructDestrutureParameter.
	VisitStructDestrutureParameter(ctx *StructDestrutureParameterContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleDestrutureParameter.
	VisitTupleDestrutureParameter(ctx *TupleDestrutureParameterContext) interface{}

	// Visit a parse tree produced by CaliburnParser#BlockRule.
	VisitBlockRule(ctx *BlockRuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#jump_statement.
	VisitJump_statement(ctx *Jump_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#control_statement.
	VisitControl_statement(ctx *Control_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#IfStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ForStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#SwitchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#CaseBlocks.
	VisitCaseBlocks(ctx *CaseBlocksContext) interface{}

	// Visit a parse tree produced by CaliburnParser#OptionCaseBlock.
	VisitOptionCaseBlock(ctx *OptionCaseBlockContext) interface{}

	// Visit a parse tree produced by CaliburnParser#DefaultCaseBlock.
	VisitDefaultCaseBlock(ctx *DefaultCaseBlockContext) interface{}

	// Visit a parse tree produced by CaliburnParser#inline_statement.
	VisitInline_statement(ctx *Inline_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AssignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AssignOperationStatement.
	VisitAssignOperationStatement(ctx *AssignOperationStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AssignExpressions.
	VisitAssignExpressions(ctx *AssignExpressionsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ExpressionAssignExpression.
	VisitExpressionAssignExpression(ctx *ExpressionAssignExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructDestrutureAssignExpression.
	VisitStructDestrutureAssignExpression(ctx *StructDestrutureAssignExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleDestrutureAssignExpression.
	VisitTupleDestrutureAssignExpression(ctx *TupleDestrutureAssignExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AssignDeclarations.
	VisitAssignDeclarations(ctx *AssignDeclarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ExpressionAssignDeclaration.
	VisitExpressionAssignDeclaration(ctx *ExpressionAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedAssignDeclarationDeclaration.
	VisitTypedAssignDeclarationDeclaration(ctx *TypedAssignDeclarationDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedAssignDeclarationDeclaration.
	VisitUntypedAssignDeclarationDeclaration(ctx *UntypedAssignDeclarationDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructDestrutureAssignDeclaration.
	VisitStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleDestrutureAssignDeclaration.
	VisitTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedAssignDeclarations.
	VisitTypedAssignDeclarations(ctx *TypedAssignDeclarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedAssignDeclaration.
	VisitTypedAssignDeclaration(ctx *TypedAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedAssignDeclarations.
	VisitUntypedAssignDeclarations(ctx *UntypedAssignDeclarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedIdentifierAssignDeclaration.
	VisitUntypedIdentifierAssignDeclaration(ctx *UntypedIdentifierAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedStructDestrutureAssignDeclaration.
	VisitUntypedStructDestrutureAssignDeclaration(ctx *UntypedStructDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedTupleDestrutureAssignDeclaration.
	VisitUntypedTupleDestrutureAssignDeclaration(ctx *UntypedTupleDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ExpressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#ExpressionsRule.
	VisitExpressionsRule(ctx *ExpressionsRuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#BracketedExpression.
	VisitBracketedExpression(ctx *BracketedExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#LiteralExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#SliceExpression.
	VisitSliceExpression(ctx *SliceExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleExpression.
	VisitTupleExpression(ctx *TupleExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#BooleanBinaryExpression.
	VisitBooleanBinaryExpression(ctx *BooleanBinaryExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#IndexExpression.
	VisitIndexExpression(ctx *IndexExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AccessExpression.
	VisitAccessExpression(ctx *AccessExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#IdentifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#FunctionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructExpression.
	VisitStructExpression(ctx *StructExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#CastExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#CallExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#FunctionTypeFunc.
	VisitFunctionTypeFunc(ctx *FunctionTypeFuncContext) interface{}

	// Visit a parse tree produced by CaliburnParser#FunctionTypeExpression.
	VisitFunctionTypeExpression(ctx *FunctionTypeExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructTypeStruct.
	VisitStructTypeStruct(ctx *StructTypeStructContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructTypeExpression.
	VisitStructTypeExpression(ctx *StructTypeExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleTypeTuple.
	VisitTupleTypeTuple(ctx *TupleTypeTupleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleTypeExpression.
	VisitTupleTypeExpression(ctx *TupleTypeExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#AccessTypeExpression.
	VisitAccessTypeExpression(ctx *AccessTypeExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#IdentifierTypeExpression.
	VisitIdentifierTypeExpression(ctx *IdentifierTypeExpressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#IdentifiersRule.
	VisitIdentifiersRule(ctx *IdentifiersRuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedLiteralAtom.
	VisitUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedLiteralAtom.
	VisitTypedLiteralAtom(ctx *TypedLiteralAtomContext) interface{}
}
