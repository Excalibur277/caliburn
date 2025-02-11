// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// CaliburnParserListener is a complete listener for a parse tree produced by CaliburnParser.
type CaliburnParserListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterScoped_block is called when entering the scoped_block production.
	EnterScoped_block(c *Scoped_blockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterControl_statement is called when entering the control_statement production.
	EnterControl_statement(c *Control_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterSwitch_statement is called when entering the switch_statement production.
	EnterSwitch_statement(c *Switch_statementContext)

	// EnterInline_statement is called when entering the inline_statement production.
	EnterInline_statement(c *Inline_statementContext)

	// EnterFunction_declaration_statement is called when entering the function_declaration_statement production.
	EnterFunction_declaration_statement(c *Function_declaration_statementContext)

	// EnterAssign_statement is called when entering the assign_statement production.
	EnterAssign_statement(c *Assign_statementContext)

	// EnterAssign_declarations is called when entering the assign_declarations production.
	EnterAssign_declarations(c *Assign_declarationsContext)

	// EnterDeclaredAssignDeclaration is called when entering the DeclaredAssignDeclaration production.
	EnterDeclaredAssignDeclaration(c *DeclaredAssignDeclarationContext)

	// EnterUndeclaredAssignDeclaration is called when entering the UndeclaredAssignDeclaration production.
	EnterUndeclaredAssignDeclaration(c *UndeclaredAssignDeclarationContext)

	// EnterStructDestrutureAssignDeclaration is called when entering the StructDestrutureAssignDeclaration production.
	EnterStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// EnterTupleDestrutureAssignDeclaration is called when entering the TupleDestrutureAssignDeclaration production.
	EnterTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// EnterDeclared_assign_declarations is called when entering the declared_assign_declarations production.
	EnterDeclared_assign_declarations(c *Declared_assign_declarationsContext)

	// EnterTypedDeclaredAssignDeclaration is called when entering the TypedDeclaredAssignDeclaration production.
	EnterTypedDeclaredAssignDeclaration(c *TypedDeclaredAssignDeclarationContext)

	// EnterUntypedDeclaredAssignDeclaration is called when entering the UntypedDeclaredAssignDeclaration production.
	EnterUntypedDeclaredAssignDeclaration(c *UntypedDeclaredAssignDeclarationContext)

	// EnterUndeclared_assign_declarations is called when entering the undeclared_assign_declarations production.
	EnterUndeclared_assign_declarations(c *Undeclared_assign_declarationsContext)

	// EnterUndeclaredAtomAssignDeclaration is called when entering the UndeclaredAtomAssignDeclaration production.
	EnterUndeclaredAtomAssignDeclaration(c *UndeclaredAtomAssignDeclarationContext)

	// EnterUndeclaredStructDestrutureAssignDeclaration is called when entering the UndeclaredStructDestrutureAssignDeclaration production.
	EnterUndeclaredStructDestrutureAssignDeclaration(c *UndeclaredStructDestrutureAssignDeclarationContext)

	// EnterUndeclaredTupleDestrutureAssignDeclaration is called when entering the UndeclaredTupleDestrutureAssignDeclaration production.
	EnterUndeclaredTupleDestrutureAssignDeclaration(c *UndeclaredTupleDestrutureAssignDeclarationContext)

	// EnterOperatorAssignment is called when entering the OperatorAssignment production.
	EnterOperatorAssignment(c *OperatorAssignmentContext)

	// EnterExpression_statement is called when entering the expression_statement production.
	EnterExpression_statement(c *Expression_statementContext)

	// EnterExpressions is called when entering the expressions production.
	EnterExpressions(c *ExpressionsContext)

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

	// EnterFunction_expression is called when entering the function_expression production.
	EnterFunction_expression(c *Function_expressionContext)

	// EnterStruct_expression is called when entering the struct_expression production.
	EnterStruct_expression(c *Struct_expressionContext)

	// EnterTuple_expression is called when entering the tuple_expression production.
	EnterTuple_expression(c *Tuple_expressionContext)

	// EnterType_expression is called when entering the type_expression production.
	EnterType_expression(c *Type_expressionContext)

	// EnterComplex_type_expression is called when entering the complex_type_expression production.
	EnterComplex_type_expression(c *Complex_type_expressionContext)

	// EnterIdentifiers is called when entering the identifiers production.
	EnterIdentifiers(c *IdentifiersContext)

	// EnterUntypedLiteralAtom is called when entering the UntypedLiteralAtom production.
	EnterUntypedLiteralAtom(c *UntypedLiteralAtomContext)

	// EnterTypedLiteralAtom is called when entering the TypedLiteralAtom production.
	EnterTypedLiteralAtom(c *TypedLiteralAtomContext)

	// EnterUntyped_literal_atom is called when entering the untyped_literal_atom production.
	EnterUntyped_literal_atom(c *Untyped_literal_atomContext)

	// EnterTyped_literal_atom is called when entering the typed_literal_atom production.
	EnterTyped_literal_atom(c *Typed_literal_atomContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitScoped_block is called when exiting the scoped_block production.
	ExitScoped_block(c *Scoped_blockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitControl_statement is called when exiting the control_statement production.
	ExitControl_statement(c *Control_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitSwitch_statement is called when exiting the switch_statement production.
	ExitSwitch_statement(c *Switch_statementContext)

	// ExitInline_statement is called when exiting the inline_statement production.
	ExitInline_statement(c *Inline_statementContext)

	// ExitFunction_declaration_statement is called when exiting the function_declaration_statement production.
	ExitFunction_declaration_statement(c *Function_declaration_statementContext)

	// ExitAssign_statement is called when exiting the assign_statement production.
	ExitAssign_statement(c *Assign_statementContext)

	// ExitAssign_declarations is called when exiting the assign_declarations production.
	ExitAssign_declarations(c *Assign_declarationsContext)

	// ExitDeclaredAssignDeclaration is called when exiting the DeclaredAssignDeclaration production.
	ExitDeclaredAssignDeclaration(c *DeclaredAssignDeclarationContext)

	// ExitUndeclaredAssignDeclaration is called when exiting the UndeclaredAssignDeclaration production.
	ExitUndeclaredAssignDeclaration(c *UndeclaredAssignDeclarationContext)

	// ExitStructDestrutureAssignDeclaration is called when exiting the StructDestrutureAssignDeclaration production.
	ExitStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// ExitTupleDestrutureAssignDeclaration is called when exiting the TupleDestrutureAssignDeclaration production.
	ExitTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// ExitDeclared_assign_declarations is called when exiting the declared_assign_declarations production.
	ExitDeclared_assign_declarations(c *Declared_assign_declarationsContext)

	// ExitTypedDeclaredAssignDeclaration is called when exiting the TypedDeclaredAssignDeclaration production.
	ExitTypedDeclaredAssignDeclaration(c *TypedDeclaredAssignDeclarationContext)

	// ExitUntypedDeclaredAssignDeclaration is called when exiting the UntypedDeclaredAssignDeclaration production.
	ExitUntypedDeclaredAssignDeclaration(c *UntypedDeclaredAssignDeclarationContext)

	// ExitUndeclared_assign_declarations is called when exiting the undeclared_assign_declarations production.
	ExitUndeclared_assign_declarations(c *Undeclared_assign_declarationsContext)

	// ExitUndeclaredAtomAssignDeclaration is called when exiting the UndeclaredAtomAssignDeclaration production.
	ExitUndeclaredAtomAssignDeclaration(c *UndeclaredAtomAssignDeclarationContext)

	// ExitUndeclaredStructDestrutureAssignDeclaration is called when exiting the UndeclaredStructDestrutureAssignDeclaration production.
	ExitUndeclaredStructDestrutureAssignDeclaration(c *UndeclaredStructDestrutureAssignDeclarationContext)

	// ExitUndeclaredTupleDestrutureAssignDeclaration is called when exiting the UndeclaredTupleDestrutureAssignDeclaration production.
	ExitUndeclaredTupleDestrutureAssignDeclaration(c *UndeclaredTupleDestrutureAssignDeclarationContext)

	// ExitOperatorAssignment is called when exiting the OperatorAssignment production.
	ExitOperatorAssignment(c *OperatorAssignmentContext)

	// ExitExpression_statement is called when exiting the expression_statement production.
	ExitExpression_statement(c *Expression_statementContext)

	// ExitExpressions is called when exiting the expressions production.
	ExitExpressions(c *ExpressionsContext)

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

	// ExitFunction_expression is called when exiting the function_expression production.
	ExitFunction_expression(c *Function_expressionContext)

	// ExitStruct_expression is called when exiting the struct_expression production.
	ExitStruct_expression(c *Struct_expressionContext)

	// ExitTuple_expression is called when exiting the tuple_expression production.
	ExitTuple_expression(c *Tuple_expressionContext)

	// ExitType_expression is called when exiting the type_expression production.
	ExitType_expression(c *Type_expressionContext)

	// ExitComplex_type_expression is called when exiting the complex_type_expression production.
	ExitComplex_type_expression(c *Complex_type_expressionContext)

	// ExitIdentifiers is called when exiting the identifiers production.
	ExitIdentifiers(c *IdentifiersContext)

	// ExitUntypedLiteralAtom is called when exiting the UntypedLiteralAtom production.
	ExitUntypedLiteralAtom(c *UntypedLiteralAtomContext)

	// ExitTypedLiteralAtom is called when exiting the TypedLiteralAtom production.
	ExitTypedLiteralAtom(c *TypedLiteralAtomContext)

	// ExitUntyped_literal_atom is called when exiting the untyped_literal_atom production.
	ExitUntyped_literal_atom(c *Untyped_literal_atomContext)

	// ExitTyped_literal_atom is called when exiting the typed_literal_atom production.
	ExitTyped_literal_atom(c *Typed_literal_atomContext)
}
