package ast

type Expression interface {
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

type UnaryExpression struct {
	ExpressionBase
	operation  UnaryOperation
	expression Expression
}

func NewUnaryExpression(operation UnaryOperation, expression Expression) *UnaryExpression {
	return &UnaryExpression{operation: operation, expression: expression}
}

type CallExpression struct {
	ExpressionBase
	callableExpression Expression
	arguments          []Expression
}

func NewCallExpression(callableExpression Expression, arguments []Expression) *CallExpression {
	return &CallExpression{callableExpression: callableExpression, arguments: arguments}
}

type AccessExpression struct {
	ExpressionBase
	expression Expression
	identifier Identifier
}

func NewAccessExpression(expression Expression, identifier Identifier) *AccessExpression {
	return &AccessExpression{expression: expression, identifier: identifier}
}

type IndexExpression struct {
	ExpressionBase
	valueExpression Expression
	indexExpression Expression
}

func NewIndexExpression(valueExpression Expression, indexExpression Expression) *IndexExpression {
	return &IndexExpression{valueExpression: valueExpression, indexExpression: indexExpression}
}

type SliceStartExpression struct {
	ExpressionBase
	valueExpression      Expression
	startIndexExpression Expression
}

func NewSliceStartExpression(valueExpression Expression, startIndexExpression Expression) *SliceStartExpression {
	return &SliceStartExpression{valueExpression: valueExpression, startIndexExpression: startIndexExpression}
}

type SliceEndExpression struct {
	ExpressionBase
	valueExpression      Expression
	startIndexExpression Expression
}

func NewSliceEndExpression(valueExpression Expression, startIndexExpression Expression) *SliceEndExpression {
	return &SliceEndExpression{valueExpression: valueExpression, startIndexExpression: startIndexExpression}
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

type IdentifierExpression struct {
	ExpressionBase
	identifier Identifier
}

func NewIdentifierExpression(identifier Identifier) *IdentifierExpression {
	return &IdentifierExpression{identifier: identifier}
}

type LiteralExpression struct {
	ExpressionBase
	literal Literal
}

func NewLiteralExpression(literal Literal) *LiteralExpression {
	return &LiteralExpression{literal: literal}
}

type FunctionExpression struct {
	ExpressionBase
	function Function
}

func NewFunctionExpression(functionType FunctionType, parameters []Parameter, returnType Type, block Block) *FunctionExpression {
	return &FunctionExpression{
		function: NewFunction(functionType, parameters, returnType, block)}
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
