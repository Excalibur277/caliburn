package ast

type Expression interface {
	IsExpression()
}

type ExpressionBase struct {
}

func (db *ExpressionBase) IsExpression() {}
