package listener

import (
	"caliburncode.com/m/parsing"
	"github.com/antlr4-go/antlr/v4"
)

var _ parsing.CaliburnParserListener = &CaliburnListener{}

type CaliburnListener struct {
	antlr.BaseParseTreeListener
	stack parseStack
}

func NewListener() *CaliburnListener {
	return &CaliburnListener{}
}

func (l *CaliburnListener) EnterModuleRule(c *parsing.ModuleRuleContext)                 {}
func (l *CaliburnListener) EnterDefinition(c *parsing.DefinitionContext)                 {}
func (l *CaliburnListener) EnterFunctionDefinition(c *parsing.FunctionDefinitionContext) {}
func (l *CaliburnListener) EnterParametersRule(c *parsing.ParametersRuleContext)         {}
func (l *CaliburnListener) EnterTypedParameter(c *parsing.TypedParameterContext)         {}
func (l *CaliburnListener) EnterUntypedParameter(c *parsing.UntypedParameterContext)     {}
func (l *CaliburnListener) EnterStructDestrutureParameter(c *parsing.StructDestrutureParameterContext) {
}
func (l *CaliburnListener) EnterTupleDestrutureParameter(c *parsing.TupleDestrutureParameterContext) {
}
func (l *CaliburnListener) EnterBlockRule(c *parsing.BlockRuleContext)                 {}
func (l *CaliburnListener) EnterStatement(c *parsing.StatementContext)                 {}
func (l *CaliburnListener) EnterJump_statement(c *parsing.Jump_statementContext)       {}
func (l *CaliburnListener) EnterReturnStatement(c *parsing.ReturnStatementContext)     {}
func (l *CaliburnListener) EnterBreakStatement(c *parsing.BreakStatementContext)       {}
func (l *CaliburnListener) EnterContinueStatement(c *parsing.ContinueStatementContext) {}
func (l *CaliburnListener) EnterControl_statement(c *parsing.Control_statementContext) {}
func (l *CaliburnListener) EnterIfStatement(c *parsing.IfStatementContext)             {}
func (l *CaliburnListener) EnterForStatement(c *parsing.ForStatementContext)           {}
func (l *CaliburnListener) EnterSwitchStatement(c *parsing.SwitchStatementContext)     {}
func (l *CaliburnListener) EnterCaseBlocks(c *parsing.CaseBlocksContext)               {}
func (l *CaliburnListener) EnterOptionCaseBlock(c *parsing.OptionCaseBlockContext)     {}
func (l *CaliburnListener) EnterDefaultCaseBlock(c *parsing.DefaultCaseBlockContext)   {}
func (l *CaliburnListener) EnterInline_statement(c *parsing.Inline_statementContext)   {}
func (l *CaliburnListener) EnterAssignStatement(c *parsing.AssignStatementContext)     {}
func (l *CaliburnListener) EnterAssignOperationStatement(c *parsing.AssignOperationStatementContext) {
}
func (l *CaliburnListener) EnterAssignExpressions(c *parsing.AssignExpressionsContext) {}
func (l *CaliburnListener) EnterExpressionAssignExpression(c *parsing.ExpressionAssignExpressionContext) {
}
func (l *CaliburnListener) EnterStructDestrutureAssignExpression(c *parsing.StructDestrutureAssignExpressionContext) {
}
func (l *CaliburnListener) EnterTupleDestrutureAssignExpression(c *parsing.TupleDestrutureAssignExpressionContext) {
}
func (l *CaliburnListener) EnterAssignDeclarations(c *parsing.AssignDeclarationsContext) {}
func (l *CaliburnListener) EnterExpressionAssignDeclaration(c *parsing.ExpressionAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterTypedAssignDeclarationDeclaration(c *parsing.TypedAssignDeclarationDeclarationContext) {
}
func (l *CaliburnListener) EnterUntypedAssignDeclarationDeclaration(c *parsing.UntypedAssignDeclarationDeclarationContext) {
}
func (l *CaliburnListener) EnterStructDestrutureAssignDeclaration(c *parsing.StructDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterTupleDestrutureAssignDeclaration(c *parsing.TupleDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterTypedAssignDeclarations(c *parsing.TypedAssignDeclarationsContext) {}
func (l *CaliburnListener) EnterTypedAssignDeclaration(c *parsing.TypedAssignDeclarationContext)   {}
func (l *CaliburnListener) EnterUntypedAssignDeclarations(c *parsing.UntypedAssignDeclarationsContext) {
}
func (l *CaliburnListener) EnterUntypedIdentifierAssignDeclaration(c *parsing.UntypedIdentifierAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterUntypedStructDestrutureAssignDeclaration(c *parsing.UntypedStructDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterUntypedTupleDestrutureAssignDeclaration(c *parsing.UntypedTupleDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) EnterExpressionStatement(c *parsing.ExpressionStatementContext)         {}
func (l *CaliburnListener) EnterExpressionsRule(c *parsing.ExpressionsRuleContext)                 {}
func (l *CaliburnListener) EnterBracketedExpression(c *parsing.BracketedExpressionContext)         {}
func (l *CaliburnListener) EnterLiteralExpression(c *parsing.LiteralExpressionContext)             {}
func (l *CaliburnListener) EnterSliceExpression(c *parsing.SliceExpressionContext)                 {}
func (l *CaliburnListener) EnterTupleExpression(c *parsing.TupleExpressionContext)                 {}
func (l *CaliburnListener) EnterBooleanBinaryExpression(c *parsing.BooleanBinaryExpressionContext) {}
func (l *CaliburnListener) EnterIndexExpression(c *parsing.IndexExpressionContext)                 {}
func (l *CaliburnListener) EnterUnaryExpression(c *parsing.UnaryExpressionContext)                 {}
func (l *CaliburnListener) EnterAccessExpression(c *parsing.AccessExpressionContext)               {}
func (l *CaliburnListener) EnterIdentifierExpression(c *parsing.IdentifierExpressionContext)       {}
func (l *CaliburnListener) EnterFunctionExpression(c *parsing.FunctionExpressionContext)           {}
func (l *CaliburnListener) EnterStructExpression(c *parsing.StructExpressionContext)               {}
func (l *CaliburnListener) EnterBinaryExpression(c *parsing.BinaryExpressionContext)               {}
func (l *CaliburnListener) EnterCastExpression(c *parsing.CastExpressionContext)                   {}
func (l *CaliburnListener) EnterCallExpression(c *parsing.CallExpressionContext)                   {}
func (l *CaliburnListener) EnterFunctionTypeFunc(c *parsing.FunctionTypeFuncContext)               {}
func (l *CaliburnListener) EnterFunctionTypeExpression(c *parsing.FunctionTypeExpressionContext)   {}
func (l *CaliburnListener) EnterStructTypeStruct(c *parsing.StructTypeStructContext)               {}
func (l *CaliburnListener) EnterStructTypeExpression(c *parsing.StructTypeExpressionContext)       {}
func (l *CaliburnListener) EnterTupleTypeTuple(c *parsing.TupleTypeTupleContext)                   {}
func (l *CaliburnListener) EnterTupleTypeExpression(c *parsing.TupleTypeExpressionContext)         {}
func (l *CaliburnListener) EnterAccessTypeExpression(c *parsing.AccessTypeExpressionContext)       {}
func (l *CaliburnListener) EnterIdentifierTypeExpression(c *parsing.IdentifierTypeExpressionContext) {
}
func (l *CaliburnListener) EnterIdentifiersRule(c *parsing.IdentifiersRuleContext)       {}
func (l *CaliburnListener) EnterUntypedLiteralAtom(c *parsing.UntypedLiteralAtomContext) {}
func (l *CaliburnListener) EnterTypedLiteralAtom(c *parsing.TypedLiteralAtomContext)     {}
func (l *CaliburnListener) ExitModuleRule(c *parsing.ModuleRuleContext)                  {}
func (l *CaliburnListener) ExitDefinition(c *parsing.DefinitionContext)                  {}
func (l *CaliburnListener) ExitFunctionDefinition(c *parsing.FunctionDefinitionContext)  {}
func (l *CaliburnListener) ExitParametersRule(c *parsing.ParametersRuleContext)          {}
func (l *CaliburnListener) ExitTypedParameter(c *parsing.TypedParameterContext)          {}
func (l *CaliburnListener) ExitUntypedParameter(c *parsing.UntypedParameterContext)      {}
func (l *CaliburnListener) ExitStructDestrutureParameter(c *parsing.StructDestrutureParameterContext) {
}
func (l *CaliburnListener) ExitTupleDestrutureParameter(c *parsing.TupleDestrutureParameterContext) {}
func (l *CaliburnListener) ExitBlockRule(c *parsing.BlockRuleContext)                               {}
func (l *CaliburnListener) ExitStatement(c *parsing.StatementContext)                               {}
func (l *CaliburnListener) ExitJump_statement(c *parsing.Jump_statementContext)                     {}
func (l *CaliburnListener) ExitReturnStatement(c *parsing.ReturnStatementContext)                   {}
func (l *CaliburnListener) ExitBreakStatement(c *parsing.BreakStatementContext)                     {}
func (l *CaliburnListener) ExitContinueStatement(c *parsing.ContinueStatementContext)               {}
func (l *CaliburnListener) ExitControl_statement(c *parsing.Control_statementContext)               {}
func (l *CaliburnListener) ExitIfStatement(c *parsing.IfStatementContext)                           {}
func (l *CaliburnListener) ExitForStatement(c *parsing.ForStatementContext)                         {}
func (l *CaliburnListener) ExitSwitchStatement(c *parsing.SwitchStatementContext)                   {}
func (l *CaliburnListener) ExitCaseBlocks(c *parsing.CaseBlocksContext)                             {}
func (l *CaliburnListener) ExitOptionCaseBlock(c *parsing.OptionCaseBlockContext)                   {}
func (l *CaliburnListener) ExitDefaultCaseBlock(c *parsing.DefaultCaseBlockContext)                 {}
func (l *CaliburnListener) ExitInline_statement(c *parsing.Inline_statementContext)                 {}
func (l *CaliburnListener) ExitAssignStatement(c *parsing.AssignStatementContext)                   {}
func (l *CaliburnListener) ExitAssignOperationStatement(c *parsing.AssignOperationStatementContext) {}
func (l *CaliburnListener) ExitAssignExpressions(c *parsing.AssignExpressionsContext)               {}
func (l *CaliburnListener) ExitExpressionAssignExpression(c *parsing.ExpressionAssignExpressionContext) {
}
func (l *CaliburnListener) ExitStructDestrutureAssignExpression(c *parsing.StructDestrutureAssignExpressionContext) {
}
func (l *CaliburnListener) ExitTupleDestrutureAssignExpression(c *parsing.TupleDestrutureAssignExpressionContext) {
}
func (l *CaliburnListener) ExitAssignDeclarations(c *parsing.AssignDeclarationsContext) {}
func (l *CaliburnListener) ExitExpressionAssignDeclaration(c *parsing.ExpressionAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitTypedAssignDeclarationDeclaration(c *parsing.TypedAssignDeclarationDeclarationContext) {
}
func (l *CaliburnListener) ExitUntypedAssignDeclarationDeclaration(c *parsing.UntypedAssignDeclarationDeclarationContext) {
}
func (l *CaliburnListener) ExitStructDestrutureAssignDeclaration(c *parsing.StructDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitTupleDestrutureAssignDeclaration(c *parsing.TupleDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitTypedAssignDeclarations(c *parsing.TypedAssignDeclarationsContext) {}
func (l *CaliburnListener) ExitTypedAssignDeclaration(c *parsing.TypedAssignDeclarationContext)   {}
func (l *CaliburnListener) ExitUntypedAssignDeclarations(c *parsing.UntypedAssignDeclarationsContext) {
}
func (l *CaliburnListener) ExitUntypedIdentifierAssignDeclaration(c *parsing.UntypedIdentifierAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitUntypedStructDestrutureAssignDeclaration(c *parsing.UntypedStructDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitUntypedTupleDestrutureAssignDeclaration(c *parsing.UntypedTupleDestrutureAssignDeclarationContext) {
}
func (l *CaliburnListener) ExitExpressionStatement(c *parsing.ExpressionStatementContext)           {}
func (l *CaliburnListener) ExitExpressionsRule(c *parsing.ExpressionsRuleContext)                   {}
func (l *CaliburnListener) ExitBracketedExpression(c *parsing.BracketedExpressionContext)           {}
func (l *CaliburnListener) ExitLiteralExpression(c *parsing.LiteralExpressionContext)               {}
func (l *CaliburnListener) ExitSliceExpression(c *parsing.SliceExpressionContext)                   {}
func (l *CaliburnListener) ExitTupleExpression(c *parsing.TupleExpressionContext)                   {}
func (l *CaliburnListener) ExitBooleanBinaryExpression(c *parsing.BooleanBinaryExpressionContext)   {}
func (l *CaliburnListener) ExitIndexExpression(c *parsing.IndexExpressionContext)                   {}
func (l *CaliburnListener) ExitUnaryExpression(c *parsing.UnaryExpressionContext)                   {}
func (l *CaliburnListener) ExitAccessExpression(c *parsing.AccessExpressionContext)                 {}
func (l *CaliburnListener) ExitIdentifierExpression(c *parsing.IdentifierExpressionContext)         {}
func (l *CaliburnListener) ExitFunctionExpression(c *parsing.FunctionExpressionContext)             {}
func (l *CaliburnListener) ExitStructExpression(c *parsing.StructExpressionContext)                 {}
func (l *CaliburnListener) ExitBinaryExpression(c *parsing.BinaryExpressionContext)                 {}
func (l *CaliburnListener) ExitCastExpression(c *parsing.CastExpressionContext)                     {}
func (l *CaliburnListener) ExitCallExpression(c *parsing.CallExpressionContext)                     {}
func (l *CaliburnListener) ExitFunctionTypeFunc(c *parsing.FunctionTypeFuncContext)                 {}
func (l *CaliburnListener) ExitFunctionTypeExpression(c *parsing.FunctionTypeExpressionContext)     {}
func (l *CaliburnListener) ExitStructTypeStruct(c *parsing.StructTypeStructContext)                 {}
func (l *CaliburnListener) ExitStructTypeExpression(c *parsing.StructTypeExpressionContext)         {}
func (l *CaliburnListener) ExitTupleTypeTuple(c *parsing.TupleTypeTupleContext)                     {}
func (l *CaliburnListener) ExitTupleTypeExpression(c *parsing.TupleTypeExpressionContext)           {}
func (l *CaliburnListener) ExitAccessTypeExpression(c *parsing.AccessTypeExpressionContext)         {}
func (l *CaliburnListener) ExitIdentifierTypeExpression(c *parsing.IdentifierTypeExpressionContext) {}
func (l *CaliburnListener) ExitIdentifiersRule(c *parsing.IdentifiersRuleContext)                   {}
func (l *CaliburnListener) ExitUntypedLiteralAtom(c *parsing.UntypedLiteralAtomContext)             {}
func (l *CaliburnListener) ExitTypedLiteralAtom(c *parsing.TypedLiteralAtomContext)                 {}
