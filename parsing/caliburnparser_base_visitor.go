// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

type BaseCaliburnParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCaliburnParserVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitScoped_block(ctx *Scoped_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitJump_statement(ctx *Jump_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitControl_statement(ctx *Control_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitSwitch_statement(ctx *Switch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitInline_statement(ctx *Inline_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitFunction_declaration_statement(ctx *Function_declaration_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssign_statement(ctx *Assign_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitAssign_declarations(ctx *Assign_declarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDeclaredAssignDeclaration(ctx *DeclaredAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUndeclaredAssignDeclaration(ctx *UndeclaredAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitDeclared_assign_declarations(ctx *Declared_assign_declarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedDeclaredAssignDeclaration(ctx *TypedDeclaredAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedDeclaredAssignDeclaration(ctx *UntypedDeclaredAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUndeclared_assign_declarations(ctx *Undeclared_assign_declarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUndeclaredAtomAssignDeclaration(ctx *UndeclaredAtomAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUndeclaredStructDestrutureAssignDeclaration(ctx *UndeclaredStructDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUndeclaredTupleDestrutureAssignDeclaration(ctx *UndeclaredTupleDestrutureAssignDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitOperatorAssignment(ctx *OperatorAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpression_statement(ctx *Expression_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
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

func (v *BaseCaliburnParserVisitor) VisitFunction_expression(ctx *Function_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitStruct_expression(ctx *Struct_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTuple_expression(ctx *Tuple_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitType_expression(ctx *Type_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitComplex_type_expression(ctx *Complex_type_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitIdentifiers(ctx *IdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTypedLiteralAtom(ctx *TypedLiteralAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitUntyped_literal_atom(ctx *Untyped_literal_atomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCaliburnParserVisitor) VisitTyped_literal_atom(ctx *Typed_literal_atomContext) interface{} {
	return v.VisitChildren(ctx)
}
