package ast

type Identifier interface {
	IsIdentifier()
}

type IdentifierBase struct{}

func (i *IdentifierBase) IsIdentifier() {}
