package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// function_type: FUNC # FunctionTypeFunc
func (l *CaliburnListener) ExitFunctionTypeFunc(c *parsing.FunctionTypeFuncContext) {
	Push(l, ast.NewImplicitFunctionType())
}

// function_type: type_expression # FunctionTypeExpression
func (l *CaliburnListener) ExitFunctionTypeExpression(c *parsing.FunctionTypeExpressionContext) {
	Push(l, ast.NewFunctionTypeExpression(Pop[ast.TypeExpression](l)))
}

// struct_type: STRUCT # StructTypeStruct
func (l *CaliburnListener) ExitStructTypeStruct(c *parsing.StructTypeStructContext) {
	Push(l, ast.NewImplicitStructType())
}

// struct_type: TUPLE # StructTypeTuple
func (l *CaliburnListener) ExitStructTypeTuple(c *parsing.StructTypeTupleContext) {
	Push(l, ast.NewImplicitTupleType())
}

// struct_type: type_expression # StructTypeExpression
func (l *CaliburnListener) ExitStructTypeExpression(c *parsing.StructTypeExpressionContext) {
	Push(l, ast.NewStructTypeExpression(Pop[ast.TypeExpression](l)))
}

// type_expression: identifier # IdentifierTypeExpression
func (l *CaliburnListener) ExitIdentifierTypeExpression(c *parsing.IdentifierTypeExpressionContext) {
	Push(l, ast.NewIdentifierTypeExpression(Pop[ast.Identifier](l)))
}

// type_expression: type_expression PERIOD identifier # AccessTypeExpression
func (l *CaliburnListener) ExitAccessTypeExpression(c *parsing.AccessTypeExpressionContext) {
	Push(l, ast.NewAccessTypeExpression(Pop[ast.TypeExpression](l), Pop[ast.Identifier](l)))
}
