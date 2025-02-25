package ast

type Token interface {
	Node
	IsToken()
}

type TokenBase struct{}

func (l *TokenBase) IsToken() {}
