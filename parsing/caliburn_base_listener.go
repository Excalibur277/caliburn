// Code generated from Caliburn.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Caliburn
import "github.com/antlr4-go/antlr/v4"

// BaseCaliburnListener is a complete listener for a parse tree produced by CaliburnParser.
type BaseCaliburnListener struct{}

var _ CaliburnListener = &BaseCaliburnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCaliburnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCaliburnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCaliburnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCaliburnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCaliburnListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCaliburnListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOperator_expression is called when production operator_expression is entered.
func (s *BaseCaliburnListener) EnterOperator_expression(ctx *Operator_expressionContext) {}

// ExitOperator_expression is called when production operator_expression is exited.
func (s *BaseCaliburnListener) ExitOperator_expression(ctx *Operator_expressionContext) {}

// EnterExpression_atom is called when production expression_atom is entered.
func (s *BaseCaliburnListener) EnterExpression_atom(ctx *Expression_atomContext) {}

// ExitExpression_atom is called when production expression_atom is exited.
func (s *BaseCaliburnListener) ExitExpression_atom(ctx *Expression_atomContext) {}

// EnterValue_atom is called when production value_atom is entered.
func (s *BaseCaliburnListener) EnterValue_atom(ctx *Value_atomContext) {}

// ExitValue_atom is called when production value_atom is exited.
func (s *BaseCaliburnListener) ExitValue_atom(ctx *Value_atomContext) {}

// EnterType_atom is called when production type_atom is entered.
func (s *BaseCaliburnListener) EnterType_atom(ctx *Type_atomContext) {}

// ExitType_atom is called when production type_atom is exited.
func (s *BaseCaliburnListener) ExitType_atom(ctx *Type_atomContext) {}

// EnterUntyped_literal_atom is called when production untyped_literal_atom is entered.
func (s *BaseCaliburnListener) EnterUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// ExitUntyped_literal_atom is called when production untyped_literal_atom is exited.
func (s *BaseCaliburnListener) ExitUntyped_literal_atom(ctx *Untyped_literal_atomContext) {}

// EnterTyped_literal_atom is called when production typed_literal_atom is entered.
func (s *BaseCaliburnListener) EnterTyped_literal_atom(ctx *Typed_literal_atomContext) {}

// ExitTyped_literal_atom is called when production typed_literal_atom is exited.
func (s *BaseCaliburnListener) ExitTyped_literal_atom(ctx *Typed_literal_atomContext) {}
