package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// expressions: expression # ExpressionsInitial
func (l *CaliburnListener) ExitExpressionsInitial(c *parsing.ExpressionsInitialContext) {
	l.Pop(1)
	Push(l, []ast.Expression{Dequeue[ast.Expression](l)})
}

// expressions: expressions COMMA expression # ExpressionsAppend
func (l *CaliburnListener) ExitExpressionsAppend(c *parsing.ExpressionsAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.Expression](l), Dequeue[ast.Expression](l)))
}

// expressions_optional: # ExpressionsOptionalNone
func (l *CaliburnListener) ExitExpressionsOptionalNone(c *parsing.ExpressionsOptionalNoneContext) {
	Push(l, []ast.Expression{})
}

// expression: L_PAREN expression R_PAREN # BracketedExpression
func (l *CaliburnListener) ExitBracketedExpression(c *parsing.BracketedExpressionContext) {
	// Nothing
}

// expression: expression op = (OP_POW | OP_ROOT) expression # BinaryExpression
// expression: expression op = (OP_MUL | OP_DIV | OP_MOD) expression # BinaryExpression
// expression: expression op = (OP_ADD | OP_SUB) expression # BinaryExpression
// expression: expression op = (OP_LSHIFT | OP_RSHIFT) expression # BinaryExpression
// expression: expression op = (OP_EQU | OP_NEQ | OP_LTE | OP_GTE | OP_LST | OP_GRT) expression # BinaryExpression
// expression: expression op = OP_AND expression # BinaryExpression
// expression: expression op = OP_XOR expression # BinaryExpression
// expression: expression op = OP_OR expression # BinaryExpression
func (l *CaliburnListener) ExitBinaryExpression(c *parsing.BinaryExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewBinaryExpression(Dequeue[ast.Expression](l), ast.NewBinaryOperation(c.GetOp()), Dequeue[ast.Expression](l)))
}

// expression: op = (OP_NOT | OP_ADD | OP_SUB) expression # UnaryExpression
func (l *CaliburnListener) ExitUnaryExpression(c *parsing.UnaryExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewUnaryExpression(ast.NewUnaryOperation(c.GetOp()), Dequeue[ast.Expression](l)))
}

// expression: expression L_PAREN expressions_optional R_PAREN # CallExpression
func (l *CaliburnListener) ExitCallExpression(c *parsing.CallExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewCallExpression(Dequeue[ast.Expression](l), Dequeue[[]ast.Expression](l)))
}

// expression: expression PERIOD identifier # AccessExpression
func (l *CaliburnListener) ExitAccessExpression(c *parsing.AccessExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewAccessExpression(Dequeue[ast.Expression](l), Dequeue[ast.Identifier](l)))
}

// expression: expression L_S_BRACK expression R_S_BRACK # IndexExpression
func (l *CaliburnListener) ExitIndexExpression(c *parsing.IndexExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewIndexExpression(Dequeue[ast.Expression](l), Dequeue[ast.Expression](l)))
}

// expression: expression L_S_BRACK expression COLON R_S_BRACK # SliceStartExpression
func (l *CaliburnListener) ExitSliceStartExpression(c *parsing.SliceStartExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewSliceStartExpression(Dequeue[ast.Expression](l), Dequeue[ast.Expression](l)))
}

// expression: expression L_S_BRACK COLON expression R_S_BRACK # SliceEndExpression
func (l *CaliburnListener) ExitSliceEndExpression(c *parsing.SliceEndExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewSliceEndExpression(Dequeue[ast.Expression](l), Dequeue[ast.Expression](l)))
}

// expression: expression L_S_BRACK expression COLON expression R_S_BRACK # SliceExpression
func (l *CaliburnListener) ExitSliceExpression(c *parsing.SliceExpressionContext) {
	l.Pop(3)
	Push(l, ast.NewSliceExpression(Dequeue[ast.Expression](l), Dequeue[ast.Expression](l), Dequeue[ast.Expression](l)))
}

// expression: identifier # IdentifierExpression
func (l *CaliburnListener) ExitIdentifierExpression(c *parsing.IdentifierExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewIdentifierExpression(Dequeue[ast.Identifier](l)))
}

// expression: literal_atom # LiteralExpression
func (l *CaliburnListener) ExitLiteralExpression(c *parsing.LiteralExpressionContext) {
	l.Pop(1)
	Push(l, ast.NewLiteralExpression(Dequeue[ast.Literal](l)))
}

// expression: function_type L_PAREN parameters R_PAREN type_expression block # FunctionExpression
func (l *CaliburnListener) ExitFunctionExpression(c *parsing.FunctionExpressionContext) {
	l.Pop(4)
	Push(l, ast.NewFunctionExpression(Dequeue[ast.FunctionType](l), Dequeue[[]ast.Parameter](l), Dequeue[ast.TypeExpression](l), Dequeue[ast.Block](l)))
}

// expression: function_type L_PAREN parameters R_PAREN block # FunctionExpressionNoReturnType
func (l *CaliburnListener) ExitFunctionExpressionNoReturnType(c *parsing.FunctionExpressionNoReturnTypeContext) {
	l.Pop(3)
	Push(l, ast.NewFunctionExpression(Dequeue[ast.FunctionType](l), Dequeue[[]ast.Parameter](l), ast.NewEmptyType(), Dequeue[ast.Block](l)))
}

// expression: struct_type L_C_BRACK struct_values R_C_BRACK # StructExpression
func (l *CaliburnListener) ExitStructExpression(c *parsing.StructExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewStructExpression(Dequeue[ast.StructType](l), Dequeue[ast.StructValues](l)))
}

// expression: type_expression L_S_BRACK collection_values R_S_BRACK # CollectionExpression
func (l *CaliburnListener) ExitCollectionExpression(c *parsing.CollectionExpressionContext) {
	l.Pop(2)
	Push(l, ast.NewCollectionExpression(Dequeue[ast.TypeExpression](l), Dequeue[[]ast.CollectionValue](l)))
}

// struct_values: named_struct_values COMMA? # StructValuesNamed
func (l *CaliburnListener) ExitStructValuesNamed(c *parsing.StructValuesNamedContext) {
	l.Pop(1)
	Push(l, ast.NewStructValuesNamed(Dequeue[[]ast.NamedStructValue](l)))
}

// struct_values: expressions_optional # StructValuesUnamed
func (l *CaliburnListener) ExitStructValuesUnamed(c *parsing.StructValuesUnamedContext) {
	l.Pop(1)
	Push(l, ast.NewStructValuesUnamed(Dequeue[[]ast.Expression](l)))
}

// named_struct_values: named_struct_value # NamedStructValuesInitial
func (l *CaliburnListener) ExitNamedStructValuesInitial(c *parsing.NamedStructValuesInitialContext) {
	l.Pop(1)
	Push(l, []ast.NamedStructValue{Dequeue[ast.NamedStructValue](l)})
}

// named_struct_values: named_struct_values COMMA named_struct_value # NamedStructValuesAppend
func (l *CaliburnListener) ExitNamedStructValuesAppend(c *parsing.NamedStructValuesAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.NamedStructValue](l), Dequeue[ast.NamedStructValue](l)))
}

// named_struct_value: identifier COLON expression # NamedStructValue
func (l *CaliburnListener) ExitNamedStructValue(c *parsing.NamedStructValueContext) {
	l.Pop(2)
	Push(l, ast.NewNamedStructValue(Dequeue[ast.Identifier](l), Dequeue[ast.Expression](l)))
}

// collection_values: # CollectionValuesInitial
func (l *CaliburnListener) ExitCollectionValuesInitial(c *parsing.CollectionValuesInitialContext) {
	Push(l, []ast.CollectionValue{})
}

// collection_values: collection_values COMMA collection_value # CollectionValuesAppend
func (l *CaliburnListener) ExitCollectionValuesAppend(c *parsing.CollectionValuesAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.CollectionValue](l), Dequeue[ast.CollectionValue](l)))
}

// collection_value: expression COMMA expression # CollectionValue
func (l *CaliburnListener) ExitCollectionValue(c *parsing.CollectionValueContext) {
	l.Pop(2)
	Push(l, ast.NewCollectionValue(Dequeue[ast.Expression](l), Dequeue[ast.Expression](l)))
}
