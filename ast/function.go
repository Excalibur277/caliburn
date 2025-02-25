package ast

type Function interface {
	IsFunction()
}

type FunctionBase struct {
	functionType FunctionType
	parameters   []Parameter
	returnType   Type
	block        Block
}

func (f *FunctionBase) IsFunction() {}

func NewFunction(functionType FunctionType, parameters []Parameter, returnType Type, block Block) Function {
	return &FunctionBase{
		functionType: functionType,
		parameters:   parameters,
		returnType:   returnType,
		block:        block,
	}
}

type Parameter interface {
	IsParameter()
}
type TypedParameter struct {
	typedAssignDeclaration TypedAssignDeclaration
}

func NewTypedParameter(typedAssignDeclaration TypedAssignDeclaration) *TypedParameter {
	return &TypedParameter{typedAssignDeclaration: typedAssignDeclaration}
}

func (p *TypedParameter) IsParameter() {}

type UntypedParameter struct {
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewUntypedParameter(untypedAssignDeclaration UntypedAssignDeclaration) *UntypedParameter {
	return &UntypedParameter{untypedAssignDeclaration: untypedAssignDeclaration}
}

func (p *UntypedParameter) IsParameter() {}

type StructDestrutureParameter struct {
	parameters []Parameter
}

func NewStructDestructureParameter(parameters []Parameter) *StructDestrutureParameter {
	return &StructDestrutureParameter{parameters: parameters}
}

func (p *StructDestrutureParameter) IsParameter() {}
