package ast

type Token interface {
	IsToken()
}

type TokenBase struct{}

func (l *TokenBase) IsToken() {}
