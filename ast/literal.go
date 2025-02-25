package ast

type Literal interface {
	IsLiteral()
}

type LiteralBase struct{}

func (l *LiteralBase) IsLiteral() {}

type TypedLiteral interface {
	IsTypedLiteral()
}

type TypedLiteralBase struct{}

func (l *TypedLiteralBase) ITypedLiteral() {}

type UntypedLiteral interface {
	IsUntypedLiteral()
}

type UntypedLiteralBase struct{}

func (l *UntypedLiteralBase) IsUntypedLiteral() {}
