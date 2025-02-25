package ast

type Type interface {
	IsType()
	IsImplicit() bool
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

type ImplicitFunctionType struct {
	FunctionTypeBase
}

func NewImplicitFunctionType() *ImplicitFunctionType {
	return &ImplicitFunctionType{}
}

func (ft *ImplicitFunctionType) IsImplicit() bool { return true }

type FunctionTypeExpression struct {
	FunctionTypeBase
	typeExpression TypeExpression
}

func NewFunctionTypeExpression(typeExpression TypeExpression) *FunctionTypeExpression {
	return &FunctionTypeExpression{typeExpression: typeExpression}
}

func (fte *FunctionTypeExpression) IsImplicit() bool { return false }

type StructType interface {
	Type
	IsStructType()
	IsImplicitNamed() bool
}
type StructTypeBase struct {
	TypeBase
}

func (st *StructTypeBase) IsStructType() {}

type ImplicitStructType struct {
	StructTypeBase
}

func NewImplicitStructType() *ImplicitStructType {
	return &ImplicitStructType{}
}

func (st *ImplicitStructType) IsImplicit() bool      { return true }
func (st *ImplicitStructType) IsImplicitNamed() bool { return true }

type ImplicitTupleType struct {
	StructTypeBase
}

func NewImplicitTupleType() *ImplicitTupleType {
	return &ImplicitTupleType{}
}

func (tt *ImplicitTupleType) IsImplicit() bool      { return true }
func (tt *ImplicitTupleType) IsImplicitNamed() bool { return false }

type StructTypeExpression struct {
	StructTypeBase
	typeExpression TypeExpression
}

func NewStructTypeExpression(typeExpression TypeExpression) *StructTypeExpression {
	return &StructTypeExpression{typeExpression: typeExpression}
}

func (ste *StructTypeExpression) IsImplicit() bool      { return false }
func (ste *StructTypeExpression) IsImplicitNamed() bool { return false }

type TypeExpression interface {
	Type
	IsTypeExpression()
}

type TypeExpressionBase struct {
	TypeBase
}

func (te *TypeExpressionBase) IsTypeExpression() {}
func (te *TypeExpressionBase) IsImplicit() bool  { return false }

type IdentifierTypeExpression struct {
	TypeExpressionBase
	identifier Identifier
}

func NewIdentifierTypeExpression(identifier Identifier) *IdentifierTypeExpression {
	return &IdentifierTypeExpression{identifier: identifier}
}

type AccessTypeExpression struct {
	TypeExpressionBase
	typeExpression   TypeExpression
	accessIdentifier Identifier
}

func NewAccessTypeExpression(typeExpression TypeExpression, accessIdentifier Identifier) *AccessTypeExpression {
	return &AccessTypeExpression{typeExpression: typeExpression, accessIdentifier: accessIdentifier}
}

type EmptyType interface {
	Type
	FunctionType
	StructType
	TypeExpression
	IsEmptyType()
}

type EmptyTypeBase struct {
	TypeBase
}

func (et *EmptyTypeBase) IsFunctionType()       {}
func (et *EmptyTypeBase) IsStructType()         {}
func (et *EmptyTypeBase) IsTupleType()          {}
func (et *EmptyTypeBase) IsTypeExpression()     {}
func (et *EmptyTypeBase) IsEmptyType()          {}
func (et *EmptyTypeBase) IsImplicit() bool      { return false }
func (et *EmptyTypeBase) IsImplicitNamed() bool { return false }

func NewEmptyType() EmptyType {
	return &EmptyTypeBase{}
}
