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
type ParameterBase struct {
}

func (f *ParameterBase) IsParameter() {}
