package ast

import "fmt"

type Function interface {
	Node
	IsFunction()
}

type FunctionBase struct {
	functionType FunctionType
	parameters   []Parameter
	returnType   Type
	block        Block
}

func NewFunction(functionType FunctionType, parameters []Parameter, returnType Type, block Block) *FunctionBase {
	return &FunctionBase{
		functionType: functionType,
		parameters:   parameters,
		returnType:   returnType,
		block:        block,
	}
}
func (f *FunctionBase) IsFunction() {}
func (f *FunctionBase) String() string {
	return fmt.Sprintf("%s (%s) %s %s", f.functionType, SliceToString(f.parameters, ", "), f.returnType, f.block)
}

type Parameter interface {
	Node
	IsParameter()
}
type TypedParameter struct {
	typedAssignDeclaration TypedAssignDeclaration
}

func NewTypedParameter(typedAssignDeclaration TypedAssignDeclaration) *TypedParameter {
	return &TypedParameter{typedAssignDeclaration: typedAssignDeclaration}
}

func (p *TypedParameter) IsParameter() {}
func (p *TypedParameter) String() string {
	return p.typedAssignDeclaration.String()
}

type UntypedParameter struct {
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewUntypedParameter(untypedAssignDeclaration UntypedAssignDeclaration) *UntypedParameter {
	return &UntypedParameter{untypedAssignDeclaration: untypedAssignDeclaration}
}

func (p *UntypedParameter) IsParameter() {}
func (p *UntypedParameter) String() string {
	return p.untypedAssignDeclaration.String()
}

type StructDestrutureParameter struct {
	parameters []Parameter
}

func NewStructDestructureParameter(parameters []Parameter) *StructDestrutureParameter {
	return &StructDestrutureParameter{parameters: parameters}
}

func (p *StructDestrutureParameter) IsParameter() {}
func (p *StructDestrutureParameter) String() string {
	return fmt.Sprintf("{ %s }", SliceToString(p.parameters, ", "))
}
