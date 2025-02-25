package ast

type Literal interface {
	IsLiteral()
}

type LiteralBase struct{}

func (l *LiteralBase) IsLiteral() {}

type TypedLiteral interface {
	Literal
	IsTypedLiteral()
}

type TypedLiteralBase struct {
	LiteralBase
	token          LiteralToken
	typeExpression TypeExpression
}

func NewTypedLiteral(typeExpression TypeExpression, token LiteralToken) *TypedLiteralBase {
	return &TypedLiteralBase{typeExpression: typeExpression, token: token}
}

func (l *TypedLiteralBase) IsTypedLiteral() {}

type UntypedLiteral interface {
	Literal
	IsUntypedLiteral()
}

type UntypedLiteralBase struct {
	LiteralBase
	token LiteralToken
}

func NewUntypedLiteral(token LiteralToken) *UntypedLiteralBase {
	return &UntypedLiteralBase{token: token}
}

func (l *UntypedLiteralBase) IsUntypedLiteral() {}

type LiteralToken interface {
	IsLiteralToken()
}

type LiteralTokenBase struct {
	text string
}

func NewLiteralToken(text string) *LiteralTokenBase {
	return &LiteralTokenBase{text: text}
}

func (l *LiteralTokenBase) IsLiteIsLiteralTokenral() {}
