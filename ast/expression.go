package ast

import "fmt"

type Expression interface {
	Node
	IsExpression()
}

type ExpressionBase struct {
}

func (db *ExpressionBase) IsExpression() {}

type BinaryExpression struct {
	ExpressionBase
	lhsExpression Expression
	operation     BinaryOperation
	rhsExpression Expression
}

func NewBinaryExpression(lhsExpression Expression, operation BinaryOperation, rhsExpression Expression) *BinaryExpression {
	return &BinaryExpression{lhsExpression: lhsExpression, operation: operation, rhsExpression: rhsExpression}
}

func (e *BinaryExpression) String() string {
	return fmt.Sprintf("%s %s %s", e.lhsExpression, e.operation, e.rhsExpression)
}

type UnaryExpression struct {
	ExpressionBase
	operation  UnaryOperation
	expression Expression
}

func NewUnaryExpression(operation UnaryOperation, expression Expression) *UnaryExpression {
	return &UnaryExpression{operation: operation, expression: expression}
}

func (e *UnaryExpression) String() string {
	return fmt.Sprintf("%s %s", e.operation, e.expression)
}

type CallExpression struct {
	ExpressionBase
	callableExpression Expression
	arguments          []Expression
}

func NewCallExpression(callableExpression Expression, arguments []Expression) *CallExpression {
	return &CallExpression{callableExpression: callableExpression, arguments: arguments}
}

func (e *CallExpression) String() string {
	return fmt.Sprintf("%s(%s)", e.callableExpression, SliceToString(e.arguments, ", "))
}

type AccessExpression struct {
	ExpressionBase
	expression Expression
	identifier Identifier
}

func NewAccessExpression(expression Expression, identifier Identifier) *AccessExpression {
	return &AccessExpression{expression: expression, identifier: identifier}
}

func (e *AccessExpression) String() string {
	return fmt.Sprintf("%s.%s", e.expression, e.identifier)
}

type IndexExpression struct {
	ExpressionBase
	valueExpression Expression
	indexExpression Expression
}

func NewIndexExpression(valueExpression Expression, indexExpression Expression) *IndexExpression {
	return &IndexExpression{valueExpression: valueExpression, indexExpression: indexExpression}
}

func (e *IndexExpression) String() string {
	return fmt.Sprintf("%s[%s]", e.valueExpression, e.indexExpression)
}

type SliceStartExpression struct {
	ExpressionBase
	valueExpression      Expression
	startIndexExpression Expression
}

func NewSliceStartExpression(valueExpression Expression, startIndexExpression Expression) *SliceStartExpression {
	return &SliceStartExpression{valueExpression: valueExpression, startIndexExpression: startIndexExpression}
}

func (e *SliceStartExpression) String() string {
	return fmt.Sprintf("%s[%s:]", e.valueExpression, e.startIndexExpression)
}

type SliceEndExpression struct {
	ExpressionBase
	valueExpression    Expression
	endIndexExpression Expression
}

func NewSliceEndExpression(valueExpression Expression, endIndexExpression Expression) *SliceEndExpression {
	return &SliceEndExpression{valueExpression: valueExpression, endIndexExpression: endIndexExpression}
}

func (e *SliceEndExpression) String() string {
	return fmt.Sprintf("%s[:%s]", e.valueExpression, e.endIndexExpression)
}

type SliceExpression struct {
	ExpressionBase
	valueExpression      Expression
	startIndexExpression Expression
	endIndexExpression   Expression
}

func NewSliceExpression(valueExpression Expression, startIndexExpression, endIndexExpression Expression) *SliceExpression {
	return &SliceExpression{valueExpression: valueExpression, startIndexExpression: startIndexExpression, endIndexExpression: endIndexExpression}
}

func (e *SliceExpression) String() string {
	return fmt.Sprintf("%s[%s:%s]", e.valueExpression, e.startIndexExpression, e.endIndexExpression)
}

type IdentifierExpression struct {
	ExpressionBase
	identifier Identifier
}

func NewIdentifierExpression(identifier Identifier) *IdentifierExpression {
	return &IdentifierExpression{identifier: identifier}
}

func (e *IdentifierExpression) String() string {
	return e.identifier.String()
}

type LiteralExpression struct {
	ExpressionBase
	literal Literal
}

func NewLiteralExpression(literal Literal) *LiteralExpression {
	return &LiteralExpression{literal: literal}
}

func (e *LiteralExpression) String() string {
	return e.literal.String()
}

type FunctionExpression struct {
	ExpressionBase
	function Function
}

func NewFunctionExpression(functionType FunctionType, parameters []Parameter, returnType Type, block Block) *FunctionExpression {
	return &FunctionExpression{
		function: NewFunction(functionType, parameters, returnType, block)}
}

func (e *FunctionExpression) String() string {
	return e.function.String()
}

type StructExpression struct {
	ExpressionBase
	structType   StructType
	structValues StructValues
}

func NewStructExpression(structType StructType, structValues StructValues) *StructExpression {
	return &StructExpression{
		structType:   structType,
		structValues: structValues,
	}
}

func (e *StructExpression) String() string {
	return fmt.Sprintf("%s{%s}", e.structType.String(), e.structValues.String())
}

type CollectionExpression struct {
	ExpressionBase
	typeExpression   TypeExpression
	collectionValues []CollectionValue
}

func NewCollectionExpression(typeExpression TypeExpression, collectionValues []CollectionValue) *CollectionExpression {
	return &CollectionExpression{
		typeExpression:   typeExpression,
		collectionValues: collectionValues,
	}
}

func (e *CollectionExpression) String() string {
	return fmt.Sprintf("%s[%s]", e.typeExpression.String(), SliceToString(e.collectionValues, ", "))
}
