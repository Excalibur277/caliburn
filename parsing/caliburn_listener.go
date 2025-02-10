// Code generated from Caliburn.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Caliburn
import "github.com/antlr4-go/antlr/v4"

// CaliburnListener is a complete listener for a parse tree produced by CaliburnParser.
type CaliburnListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperator_expression is called when entering the operator_expression production.
	EnterOperator_expression(c *Operator_expressionContext)

	// EnterExpression_atom is called when entering the expression_atom production.
	EnterExpression_atom(c *Expression_atomContext)

	// EnterValue_atom is called when entering the value_atom production.
	EnterValue_atom(c *Value_atomContext)

	// EnterType_atom is called when entering the type_atom production.
	EnterType_atom(c *Type_atomContext)

	// EnterUntyped_literal_atom is called when entering the untyped_literal_atom production.
	EnterUntyped_literal_atom(c *Untyped_literal_atomContext)

	// EnterTyped_literal_atom is called when entering the typed_literal_atom production.
	EnterTyped_literal_atom(c *Typed_literal_atomContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperator_expression is called when exiting the operator_expression production.
	ExitOperator_expression(c *Operator_expressionContext)

	// ExitExpression_atom is called when exiting the expression_atom production.
	ExitExpression_atom(c *Expression_atomContext)

	// ExitValue_atom is called when exiting the value_atom production.
	ExitValue_atom(c *Value_atomContext)

	// ExitType_atom is called when exiting the type_atom production.
	ExitType_atom(c *Type_atomContext)

	// ExitUntyped_literal_atom is called when exiting the untyped_literal_atom production.
	ExitUntyped_literal_atom(c *Untyped_literal_atomContext)

	// ExitTyped_literal_atom is called when exiting the typed_literal_atom production.
	ExitTyped_literal_atom(c *Typed_literal_atomContext)
}
