package ast

import "fmt"

type Literal interface {
	Node
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
	typeExpression TypeExpression
	token          LiteralToken
}

func NewTypedLiteral(typeExpression TypeExpression, token LiteralToken) *TypedLiteralBase {
	return &TypedLiteralBase{typeExpression: typeExpression, token: token}
}

func (l *TypedLiteralBase) String() string {
	return fmt.Sprintf("%s %s", l.typeExpression, l.token)
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
func (l *UntypedLiteralBase) String() string {
	return l.token.String()
}

type LiteralToken interface {
	Node
	IsLiteralToken()
}

type LiteralTokenBase struct {
	text string
}

func NewLiteralToken(text string) *LiteralTokenBase {
	return &LiteralTokenBase{text: text}
}

func (l *LiteralTokenBase) IsLiteralToken() {}
func (l *LiteralTokenBase) String() string {
	return l.text
}
