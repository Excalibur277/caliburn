package ast

type Type interface {
	IsType()
}

type TypeBase struct{}

func (t *TypeBase) IsType() {}

type FunctionType interface {
	Type
	IsFunctionType()
}

type FunctionTypeBase struct {
	TypeBase
}

func (ft *FunctionTypeBase) IsFunctionType() {}

type StructType interface {
	Type
	IsStructType()
}

type StructTypeBase struct {
	TypeBase
}

func (st *StructTypeBase) IsStructType() {}

type TupleType interface {
	Type
	IsTupleType()
}

type TupleTypeBase struct {
	TypeBase
}

func (tt *TupleTypeBase) IsTupleType() {}

type TypeExpression interface {
	Type
	IsTypeExpression()
}

type TypeExpressionBase struct {
	TypeBase
}

func (te *TypeExpressionBase) IsTypeExpression() {}

type EmptyType interface {
	Type
	FunctionType
	StructType
	TupleType
	TypeExpression
	IsEmptyType()
}

type EmptyTypeBase struct {
	TypeBase
}

func (et *EmptyTypeBase) IsFunctionType()   {}
func (et *EmptyTypeBase) IsStructType()     {}
func (et *EmptyTypeBase) IsTupleType()      {}
func (et *EmptyTypeBase) IsTypeExpression() {}
func (et *EmptyTypeBase) IsEmptyType()      {}

func NewEmptyType() EmptyType {
	return &EmptyTypeBase{}
}
