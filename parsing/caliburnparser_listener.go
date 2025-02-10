// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
import "github.com/antlr4-go/antlr/v4"

// CaliburnParserListener is a complete listener for a parse tree produced by CaliburnParser.
type CaliburnParserListener interface {
	antlr.ParseTreeListener

	// EnterScoped_block is called when entering the scoped_block production.
	EnterScoped_block(c *Scoped_blockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

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

	// EnterExponentExpression is called when entering the ExponentExpression production.
	EnterExponentExpression(c *ExponentExpressionContext)

	// EnterAdditionExpression is called when entering the AdditionExpression production.
	EnterAdditionExpression(c *AdditionExpressionContext)

	// EnterBracketedExpression is called when entering the BracketedExpression production.
	EnterBracketedExpression(c *BracketedExpressionContext)

	// EnterSliceExpression is called when entering the SliceExpression production.
	EnterSliceExpression(c *SliceExpressionContext)

	// EnterTupleExpression is called when entering the TupleExpression production.
	EnterTupleExpression(c *TupleExpressionContext)

	// EnterIndexExpression is called when entering the IndexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterAtomExpression is called when entering the AtomExpression production.
	EnterAtomExpression(c *AtomExpressionContext)

	// EnterShiftExpression is called when entering the ShiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterOrExpression is called when entering the OrExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterAccessExpression is called when entering the AccessExpression production.
	EnterAccessExpression(c *AccessExpressionContext)

	// EnterComparisonExpression is called when entering the ComparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterXOrExpression is called when entering the XOrExpression production.
	EnterXOrExpression(c *XOrExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterStructExpression is called when entering the StructExpression production.
	EnterStructExpression(c *StructExpressionContext)

	// EnterAndExpression is called when entering the AndExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterCastExpression is called when entering the CastExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterCallExpression is called when entering the CallExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterFunction_expression is called when entering the function_expression production.
	EnterFunction_expression(c *Function_expressionContext)

	// EnterStruct_expression is called when entering the struct_expression production.
	EnterStruct_expression(c *Struct_expressionContext)

	// EnterTuple_expression is called when entering the tuple_expression production.
	EnterTuple_expression(c *Tuple_expressionContext)

	// EnterExpression_atom is called when entering the expression_atom production.
	EnterExpression_atom(c *Expression_atomContext)

	// EnterIdentifiers is called when entering the identifiers production.
	EnterIdentifiers(c *IdentifiersContext)

	// EnterLiteral_atom is called when entering the literal_atom production.
	EnterLiteral_atom(c *Literal_atomContext)

	// EnterUntyped_literal_atom is called when entering the untyped_literal_atom production.
	EnterUntyped_literal_atom(c *Untyped_literal_atomContext)

	// EnterTyped_literal_atom is called when entering the typed_literal_atom production.
	EnterTyped_literal_atom(c *Typed_literal_atomContext)

	// ExitScoped_block is called when exiting the scoped_block production.
	ExitScoped_block(c *Scoped_blockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

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

	// ExitExponentExpression is called when exiting the ExponentExpression production.
	ExitExponentExpression(c *ExponentExpressionContext)

	// ExitAdditionExpression is called when exiting the AdditionExpression production.
	ExitAdditionExpression(c *AdditionExpressionContext)

	// ExitBracketedExpression is called when exiting the BracketedExpression production.
	ExitBracketedExpression(c *BracketedExpressionContext)

	// ExitSliceExpression is called when exiting the SliceExpression production.
	ExitSliceExpression(c *SliceExpressionContext)

	// ExitTupleExpression is called when exiting the TupleExpression production.
	ExitTupleExpression(c *TupleExpressionContext)

	// ExitIndexExpression is called when exiting the IndexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitAtomExpression is called when exiting the AtomExpression production.
	ExitAtomExpression(c *AtomExpressionContext)

	// ExitShiftExpression is called when exiting the ShiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitOrExpression is called when exiting the OrExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitAccessExpression is called when exiting the AccessExpression production.
	ExitAccessExpression(c *AccessExpressionContext)

	// ExitComparisonExpression is called when exiting the ComparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitXOrExpression is called when exiting the XOrExpression production.
	ExitXOrExpression(c *XOrExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitStructExpression is called when exiting the StructExpression production.
	ExitStructExpression(c *StructExpressionContext)

	// ExitAndExpression is called when exiting the AndExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitCastExpression is called when exiting the CastExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitCallExpression is called when exiting the CallExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitFunction_expression is called when exiting the function_expression production.
	ExitFunction_expression(c *Function_expressionContext)

	// ExitStruct_expression is called when exiting the struct_expression production.
	ExitStruct_expression(c *Struct_expressionContext)

	// ExitTuple_expression is called when exiting the tuple_expression production.
	ExitTuple_expression(c *Tuple_expressionContext)

	// ExitExpression_atom is called when exiting the expression_atom production.
	ExitExpression_atom(c *Expression_atomContext)

	// ExitIdentifiers is called when exiting the identifiers production.
	ExitIdentifiers(c *IdentifiersContext)

	// ExitLiteral_atom is called when exiting the literal_atom production.
	ExitLiteral_atom(c *Literal_atomContext)

	// ExitUntyped_literal_atom is called when exiting the untyped_literal_atom production.
	ExitUntyped_literal_atom(c *Untyped_literal_atomContext)

	// ExitTyped_literal_atom is called when exiting the typed_literal_atom production.
	ExitTyped_literal_atom(c *Typed_literal_atomContext)
}
