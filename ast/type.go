package ast

import "fmt"

type Type interface {
	Node
	IsType()
	IsImplicit() bool
}

type TypeBase struct{}

func (t *TypeBase) IsType() {}

type FunctionType interface {
	Node
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

func (ft *ImplicitFunctionType) String() string { return "func" }

type FunctionTypeExpression struct {
	FunctionTypeBase
	typeExpression TypeExpression
}

func NewFunctionTypeExpression(typeExpression TypeExpression) *FunctionTypeExpression {
	return &FunctionTypeExpression{typeExpression: typeExpression}
}

func (fte *FunctionTypeExpression) IsImplicit() bool { return false }
func (fte *FunctionTypeExpression) String() string   { return fte.typeExpression.String() }

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
func (st *ImplicitStructType) String() string        { return "struct" }

type ImplicitTupleType struct {
	StructTypeBase
}

func NewImplicitTupleType() *ImplicitTupleType {
	return &ImplicitTupleType{}
}

func (tt *ImplicitTupleType) IsImplicit() bool      { return true }
func (tt *ImplicitTupleType) IsImplicitNamed() bool { return false }
func (tt *ImplicitTupleType) String() string        { return "tuple" }

type StructTypeExpression struct {
	StructTypeBase
	typeExpression TypeExpression
}

func NewStructTypeExpression(typeExpression TypeExpression) *StructTypeExpression {
	return &StructTypeExpression{typeExpression: typeExpression}
}

func (ste *StructTypeExpression) IsImplicit() bool      { return false }
func (ste *StructTypeExpression) IsImplicitNamed() bool { return false }
func (ste *StructTypeExpression) String() string        { return ste.typeExpression.String() }

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

func (ite *IdentifierTypeExpression) String() string { return ite.identifier.String() }

type AccessTypeExpression struct {
	TypeExpressionBase
	typeExpression   TypeExpression
	accessIdentifier Identifier
}

func NewAccessTypeExpression(typeExpression TypeExpression, accessIdentifier Identifier) *AccessTypeExpression {
	return &AccessTypeExpression{typeExpression: typeExpression, accessIdentifier: accessIdentifier}
}

func (ate *AccessTypeExpression) String() string {
	return fmt.Sprintf("%s %s", ate.typeExpression, ate.accessIdentifier)
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

func (et *EmptyTypeBase) String() string { return "" }
