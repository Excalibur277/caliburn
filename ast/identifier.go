package ast

type Identifier interface {
	Node
	IsIdentifier()
}

type IdentifierBase struct{ text string }

func NewIdentifer(text string) Identifier {
	return &IdentifierBase{text: text}
}

func (i *IdentifierBase) IsIdentifier() {}
func (i *IdentifierBase) String() string {
	return i.text
}
