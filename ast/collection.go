package ast

type CollectionValue interface {
	IsCollectionValue()
}

type CollectionValueBase struct {
	keyExpression   Expression
	valueExpression Expression
}

func NewCollectionValue(keyExpression, valueExpression Expression) *CollectionValueBase {
	return &CollectionValueBase{keyExpression: keyExpression, valueExpression: valueExpression}
}

func (cv *CollectionValueBase) IsCollectionValue() {}
