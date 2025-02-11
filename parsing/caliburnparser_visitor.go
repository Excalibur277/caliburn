// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CaliburnParser.
type CaliburnParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CaliburnParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by CaliburnParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#scoped_block.
	VisitScoped_block(ctx *Scoped_blockContext) interface{}

	// Visit a parse tree produced by CaliburnParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#jump_statement.
	VisitJump_statement(ctx *Jump_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#control_statement.
	VisitControl_statement(ctx *Control_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#switch_statement.
	VisitSwitch_statement(ctx *Switch_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#inline_statement.
	VisitInline_statement(ctx *Inline_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#function_declaration_statement.
	VisitFunction_declaration_statement(ctx *Function_declaration_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#assign_statement.
	VisitAssign_statement(ctx *Assign_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#assign_declarations.
	VisitAssign_declarations(ctx *Assign_declarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#DeclaredAssignDeclaration.
	VisitDeclaredAssignDeclaration(ctx *DeclaredAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UndeclaredAssignDeclaration.
	VisitUndeclaredAssignDeclaration(ctx *UndeclaredAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#StructDestrutureAssignDeclaration.
	VisitStructDestrutureAssignDeclaration(ctx *StructDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TupleDestrutureAssignDeclaration.
	VisitTupleDestrutureAssignDeclaration(ctx *TupleDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#declared_assign_declarations.
	VisitDeclared_assign_declarations(ctx *Declared_assign_declarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedDeclaredAssignDeclaration.
	VisitTypedDeclaredAssignDeclaration(ctx *TypedDeclaredAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedDeclaredAssignDeclaration.
	VisitUntypedDeclaredAssignDeclaration(ctx *UntypedDeclaredAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#undeclared_assign_declarations.
	VisitUndeclared_assign_declarations(ctx *Undeclared_assign_declarationsContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UndeclaredAtomAssignDeclaration.
	VisitUndeclaredAtomAssignDeclaration(ctx *UndeclaredAtomAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UndeclaredStructDestrutureAssignDeclaration.
	VisitUndeclaredStructDestrutureAssignDeclaration(ctx *UndeclaredStructDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UndeclaredTupleDestrutureAssignDeclaration.
	VisitUndeclaredTupleDestrutureAssignDeclaration(ctx *UndeclaredTupleDestrutureAssignDeclarationContext) interface{}

	// Visit a parse tree produced by CaliburnParser#OperatorAssignment.
	VisitOperatorAssignment(ctx *OperatorAssignmentContext) interface{}

	// Visit a parse tree produced by CaliburnParser#expression_statement.
	VisitExpression_statement(ctx *Expression_statementContext) interface{}

	// Visit a parse tree produced by CaliburnParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) interface{}

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

	// Visit a parse tree produced by CaliburnParser#function_expression.
	VisitFunction_expression(ctx *Function_expressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#struct_expression.
	VisitStruct_expression(ctx *Struct_expressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#tuple_expression.
	VisitTuple_expression(ctx *Tuple_expressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#type_expression.
	VisitType_expression(ctx *Type_expressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#complex_type_expression.
	VisitComplex_type_expression(ctx *Complex_type_expressionContext) interface{}

	// Visit a parse tree produced by CaliburnParser#identifiers.
	VisitIdentifiers(ctx *IdentifiersContext) interface{}

	// Visit a parse tree produced by CaliburnParser#UntypedLiteralAtom.
	VisitUntypedLiteralAtom(ctx *UntypedLiteralAtomContext) interface{}

	// Visit a parse tree produced by CaliburnParser#TypedLiteralAtom.
	VisitTypedLiteralAtom(ctx *TypedLiteralAtomContext) interface{}

	// Visit a parse tree produced by CaliburnParser#untyped_literal_atom.
	VisitUntyped_literal_atom(ctx *Untyped_literal_atomContext) interface{}

	// Visit a parse tree produced by CaliburnParser#typed_literal_atom.
	VisitTyped_literal_atom(ctx *Typed_literal_atomContext) interface{}
}
