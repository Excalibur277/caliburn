package ast

type Literal interface {
	IsLiteral()
}

type LiteralBase struct{}

func (l *LiteralBase) IsLiteral() {}
