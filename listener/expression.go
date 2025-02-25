package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// expressions: expression # ExpressionsInitial
func (l *CaliburnListener) ExitExpressionsInitial(c *parsing.ExpressionsInitialContext) {
	Push(l, []ast.Expression{Pop[ast.Expression](l)})
}

// expressions: expressions COMMA expression # ExpressionsAppend
func (l *CaliburnListener) ExitExpressionsAppend(c *parsing.ExpressionsAppendContext) {
	Push(l, append(Pop[[]ast.Expression](l), Pop[ast.Expression](l)))
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
// expression: expression OP_AND expression # BinaryExpression
// expression: expression OP_XOR expression # BinaryExpression
// expression: expression OP_OR expression # BinaryExpression
func (l *CaliburnListener) ExitBinaryExpression(c *parsing.BinaryExpressionContext) {
	Push(l, ast.NewBinaryExpression(Pop[ast.Expression](l), ast.NewBinaryOperation(c.GetOp()), Pop[ast.Expression](l)))
}

// expression: op = (OP_NOT | OP_ADD | OP_SUB) expression # UnaryExpression
func (l *CaliburnListener) ExitUnaryExpression(c *parsing.UnaryExpressionContext) {
	Push(l, ast.NewUnaryExpression(ast.NewUnaryOperation(c.GetOp()), Pop[ast.Expression](l)))
}

// expression: expression L_PAREN expressions_optional R_PAREN # CallExpression
func (l *CaliburnListener) ExitCallExpression(c *parsing.CallExpressionContext) {
	Push(l, ast.NewCallExpression(Pop[ast.Expression](l), Pop[[]ast.Expression](l)))
}

// expression: expression PERIOD identifier # AccessExpression
func (l *CaliburnListener) ExitAccessExpression(c *parsing.AccessExpressionContext) {
	Push(l, ast.NewAccessExpression(Pop[ast.Expression](l), Pop[ast.Identifier](l)))
}

// expression: expression L_S_BRACK expression R_S_BRACK # IndexExpression
func (l *CaliburnListener) ExitIndexExpression(c *parsing.IndexExpressionContext) {
	Push(l, ast.NewIndexExpression(Pop[ast.Expression](l), Pop[ast.Expression](l)))
}

// expression: expression L_S_BRACK expression COLON R_S_BRACK # SliceStartExpression
func (l *CaliburnListener) ExitSliceStartExpression(c *parsing.SliceStartExpressionContext) {
	Push(l, ast.NewSliceStartExpression(Pop[ast.Expression](l), Pop[ast.Expression](l)))
}

// expression: expression L_S_BRACK COLON expression R_S_BRACK # SliceEndExpression
func (l *CaliburnListener) ExitSliceEndExpression(c *parsing.SliceEndExpressionContext) {
	Push(l, ast.NewSliceEndExpression(Pop[ast.Expression](l), Pop[ast.Expression](l)))
}

// expression: expression L_S_BRACK expression COLON expression R_S_BRACK # SliceExpression
func (l *CaliburnListener) ExitSliceExpression(c *parsing.SliceExpressionContext) {
	Push(l, ast.NewSliceExpression(Pop[ast.Expression](l), Pop[ast.Expression](l), Pop[ast.Expression](l)))
}

// expression: identifier # IdentifierExpression
func (l *CaliburnListener) ExitIdentifierExpression(c *parsing.IdentifierExpressionContext) {
	Push(l, ast.NewIdentifierExpression(Pop[ast.Identifier](l)))
}

// expression: literal_atom # LiteralExpression
func (l *CaliburnListener) ExitLiteralExpression(c *parsing.LiteralExpressionContext) {
	Push(l, ast.NewLiteralExpression(Pop[ast.Literal](l)))
}

// expression: function_type L_PAREN parameters R_PAREN type_expression block # FunctionExpression
func (l *CaliburnListener) ExitFunctionExpression(c *parsing.FunctionExpressionContext) {
	Push(l, ast.NewFunctionExpression(Pop[ast.FunctionType](l), Pop[[]ast.Parameter](l), Pop[ast.TypeExpression](l), Pop[ast.Block](l)))
}

// expression: function_type L_PAREN parameters R_PAREN block # FunctionExpressionNoReturnType
func (l *CaliburnListener) ExitFunctionExpressionNoReturnType(c *parsing.FunctionExpressionNoReturnTypeContext) {
	Push(l, ast.NewFunctionExpression(Pop[ast.FunctionType](l), Pop[[]ast.Parameter](l), ast.NewEmptyType(), Pop[ast.Block](l)))
}

// expression: struct_type L_C_BRACK struct_values R_C_BRACK # StructExpression
func (l *CaliburnListener) ExitStructExpression(c *parsing.StructExpressionContext) {
	Push(l, ast.NewStructExpression(Pop[ast.StructType](l), Pop[ast.StructValues](l)))
}

// expression: type_expression L_S_BRACK collection_values R_S_BRACK # CollectionExpression
func (l *CaliburnListener) ExitCollectionExpression(c *parsing.CollectionExpressionContext) {
	Push(l, ast.NewCollectionExpression(Pop[ast.TypeExpression](l), Pop[[]ast.CollectionValue](l)))
}

// struct_values: named_struct_values COMMA? # StructValuesNamed
func (l *CaliburnListener) ExitStructValuesNamed(c *parsing.StructValuesNamedContext) {
	Push(l, ast.NewStructValuesNamed(Pop[[]ast.NamedStructValue](l)))
}

// struct_values: expressions_optional # StructValuesUnamed
func (l *CaliburnListener) ExitStructValuesUnamed(c *parsing.StructValuesUnamedContext) {
	Push(l, ast.NewStructValuesUnamed(Pop[[]ast.Expression](l)))
}

// named_struct_values: named_struct_value # NamedStructValuesInitial
func (l *CaliburnListener) ExitNamedStructValuesInitial(c *parsing.NamedStructValuesInitialContext) {
	Push(l, []ast.NamedStructValue{Pop[ast.NamedStructValue](l)})
}

// named_struct_values: named_struct_values COMMA named_struct_value # NamedStructValuesAppend
func (l *CaliburnListener) ExitNamedStructValuesAppend(c *parsing.NamedStructValuesAppendContext) {
	Push(l, append(Pop[[]ast.NamedStructValue](l), Pop[ast.NamedStructValue](l)))
}

// named_struct_value: identifier COLON expression # NamedStructValue
func (l *CaliburnListener) ExitNamedStructValue(c *parsing.NamedStructValueContext) {
	Push(l, ast.NewNamedStructValue(Pop[ast.Identifier](l), Pop[ast.Expression](l)))
}

// collection_values: # CollectionValuesInitial
func (l *CaliburnListener) ExitCollectionValuesInitial(c *parsing.CollectionValuesInitialContext) {
	Push(l, []ast.CollectionValue{})
}

// collection_values: collection_values COMMA collection_value # CollectionValuesAppend
func (l *CaliburnListener) ExitCollectionValuesAppend(c *parsing.CollectionValuesAppendContext) {
	Push(l, append(Pop[[]ast.CollectionValue](l), Pop[ast.CollectionValue](l)))
}

// collection_value: expression COMMA expression # CollectionValue
func (l *CaliburnListener) ExitCollectionValue(c *parsing.CollectionValueContext) {
	Push(l, ast.NewCollectionValue(Pop[ast.Expression](l), Pop[ast.Expression](l)))
}
