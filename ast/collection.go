package ast

import "fmt"

type CollectionValue interface {
	Node
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
func (cv *CollectionValueBase) String() string {
	return fmt.Sprintf("%s : %s", cv.keyExpression, cv.valueExpression)
}
