package ast

type Identifier interface {
	IsIdentifier()
}

type IdentifierBase struct{ text string }

func NewIdentifer(text string) Identifier {
	return &IdentifierBase{text: text}
}

func (i *IdentifierBase) IsIdentifier() {}
