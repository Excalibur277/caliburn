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

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterTypedParameter is called when entering the TypedParameter production.
	EnterTypedParameter(c *TypedParameterContext)

	// EnterUntypedParameter is called when entering the UntypedParameter production.
	EnterUntypedParameter(c *UntypedParameterContext)

	// EnterStructDestrutureParameter is called when entering the StructDestrutureParameter production.
	EnterStructDestrutureParameter(c *StructDestrutureParameterContext)

	// EnterTupleDestrutureParameter is called when entering the TupleDestrutureParameter production.
	EnterTupleDestrutureParameter(c *TupleDestrutureParameterContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

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

	// EnterCase_statements is called when entering the case_statements production.
	EnterCase_statements(c *Case_statementsContext)

	// EnterOption_case_statement is called when entering the option_case_statement production.
	EnterOption_case_statement(c *Option_case_statementContext)

	// EnterDefault_case_statement is called when entering the default_case_statement production.
	EnterDefault_case_statement(c *Default_case_statementContext)

	// EnterInline_statement is called when entering the inline_statement production.
	EnterInline_statement(c *Inline_statementContext)

	// EnterAssign_statement is called when entering the assign_statement production.
	EnterAssign_statement(c *Assign_statementContext)

	// EnterAssign_expressions is called when entering the assign_expressions production.
	EnterAssign_expressions(c *Assign_expressionsContext)

	// EnterExpressionAssignExpression is called when entering the ExpressionAssignExpression production.
	EnterExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// EnterStructDestrutureAssignExpression is called when entering the StructDestrutureAssignExpression production.
	EnterStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// EnterTupleDestrutureAssignExpression is called when entering the TupleDestrutureAssignExpression production.
	EnterTupleDestrutureAssignExpression(c *TupleDestrutureAssignExpressionContext)

	// EnterAssign_declarations is called when entering the assign_declarations production.
	EnterAssign_declarations(c *Assign_declarationsContext)

	// EnterExpressionAssignDeclaration is called when entering the ExpressionAssignDeclaration production.
	EnterExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// EnterTypedAssignDeclaration is called when entering the TypedAssignDeclaration production.
	EnterTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// EnterUntypedAssignDeclaration is called when entering the UntypedAssignDeclaration production.
	EnterUntypedAssignDeclaration(c *UntypedAssignDeclarationContext)

	// EnterStructDestrutureAssignDeclaration is called when entering the StructDestrutureAssignDeclaration production.
	EnterStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// EnterTupleDestrutureAssignDeclaration is called when entering the TupleDestrutureAssignDeclaration production.
	EnterTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// EnterTyped_assign_declarations is called when entering the typed_assign_declarations production.
	EnterTyped_assign_declarations(c *Typed_assign_declarationsContext)

	// EnterTypedTypedAssignDeclaration is called when entering the TypedTypedAssignDeclaration production.
	EnterTypedTypedAssignDeclaration(c *TypedTypedAssignDeclarationContext)

	// EnterUntyped_assign_declarations is called when entering the untyped_assign_declarations production.
	EnterUntyped_assign_declarations(c *Untyped_assign_declarationsContext)

	// EnterUntypedAtomAssignDeclaration is called when entering the UntypedAtomAssignDeclaration production.
	EnterUntypedAtomAssignDeclaration(c *UntypedAtomAssignDeclarationContext)

	// EnterUntypedStructDestrutureAssignDeclaration is called when entering the UntypedStructDestrutureAssignDeclaration production.
	EnterUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// EnterUntypedTupleDestrutureAssignDeclaration is called when entering the UntypedTupleDestrutureAssignDeclaration production.
	EnterUntypedTupleDestrutureAssignDeclaration(c *UntypedTupleDestrutureAssignDeclarationContext)

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

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterTuple_type is called when entering the tuple_type production.
	EnterTuple_type(c *Tuple_typeContext)

	// EnterType_expression is called when entering the type_expression production.
	EnterType_expression(c *Type_expressionContext)

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

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitTypedParameter is called when exiting the TypedParameter production.
	ExitTypedParameter(c *TypedParameterContext)

	// ExitUntypedParameter is called when exiting the UntypedParameter production.
	ExitUntypedParameter(c *UntypedParameterContext)

	// ExitStructDestrutureParameter is called when exiting the StructDestrutureParameter production.
	ExitStructDestrutureParameter(c *StructDestrutureParameterContext)

	// ExitTupleDestrutureParameter is called when exiting the TupleDestrutureParameter production.
	ExitTupleDestrutureParameter(c *TupleDestrutureParameterContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

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

	// ExitCase_statements is called when exiting the case_statements production.
	ExitCase_statements(c *Case_statementsContext)

	// ExitOption_case_statement is called when exiting the option_case_statement production.
	ExitOption_case_statement(c *Option_case_statementContext)

	// ExitDefault_case_statement is called when exiting the default_case_statement production.
	ExitDefault_case_statement(c *Default_case_statementContext)

	// ExitInline_statement is called when exiting the inline_statement production.
	ExitInline_statement(c *Inline_statementContext)

	// ExitAssign_statement is called when exiting the assign_statement production.
	ExitAssign_statement(c *Assign_statementContext)

	// ExitAssign_expressions is called when exiting the assign_expressions production.
	ExitAssign_expressions(c *Assign_expressionsContext)

	// ExitExpressionAssignExpression is called when exiting the ExpressionAssignExpression production.
	ExitExpressionAssignExpression(c *ExpressionAssignExpressionContext)

	// ExitStructDestrutureAssignExpression is called when exiting the StructDestrutureAssignExpression production.
	ExitStructDestrutureAssignExpression(c *StructDestrutureAssignExpressionContext)

	// ExitTupleDestrutureAssignExpression is called when exiting the TupleDestrutureAssignExpression production.
	ExitTupleDestrutureAssignExpression(c *TupleDestrutureAssignExpressionContext)

	// ExitAssign_declarations is called when exiting the assign_declarations production.
	ExitAssign_declarations(c *Assign_declarationsContext)

	// ExitExpressionAssignDeclaration is called when exiting the ExpressionAssignDeclaration production.
	ExitExpressionAssignDeclaration(c *ExpressionAssignDeclarationContext)

	// ExitTypedAssignDeclaration is called when exiting the TypedAssignDeclaration production.
	ExitTypedAssignDeclaration(c *TypedAssignDeclarationContext)

	// ExitUntypedAssignDeclaration is called when exiting the UntypedAssignDeclaration production.
	ExitUntypedAssignDeclaration(c *UntypedAssignDeclarationContext)

	// ExitStructDestrutureAssignDeclaration is called when exiting the StructDestrutureAssignDeclaration production.
	ExitStructDestrutureAssignDeclaration(c *StructDestrutureAssignDeclarationContext)

	// ExitTupleDestrutureAssignDeclaration is called when exiting the TupleDestrutureAssignDeclaration production.
	ExitTupleDestrutureAssignDeclaration(c *TupleDestrutureAssignDeclarationContext)

	// ExitTyped_assign_declarations is called when exiting the typed_assign_declarations production.
	ExitTyped_assign_declarations(c *Typed_assign_declarationsContext)

	// ExitTypedTypedAssignDeclaration is called when exiting the TypedTypedAssignDeclaration production.
	ExitTypedTypedAssignDeclaration(c *TypedTypedAssignDeclarationContext)

	// ExitUntyped_assign_declarations is called when exiting the untyped_assign_declarations production.
	ExitUntyped_assign_declarations(c *Untyped_assign_declarationsContext)

	// ExitUntypedAtomAssignDeclaration is called when exiting the UntypedAtomAssignDeclaration production.
	ExitUntypedAtomAssignDeclaration(c *UntypedAtomAssignDeclarationContext)

	// ExitUntypedStructDestrutureAssignDeclaration is called when exiting the UntypedStructDestrutureAssignDeclaration production.
	ExitUntypedStructDestrutureAssignDeclaration(c *UntypedStructDestrutureAssignDeclarationContext)

	// ExitUntypedTupleDestrutureAssignDeclaration is called when exiting the UntypedTupleDestrutureAssignDeclaration production.
	ExitUntypedTupleDestrutureAssignDeclaration(c *UntypedTupleDestrutureAssignDeclarationContext)

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

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitTuple_type is called when exiting the tuple_type production.
	ExitTuple_type(c *Tuple_typeContext)

	// ExitType_expression is called when exiting the type_expression production.
	ExitType_expression(c *Type_expressionContext)

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
