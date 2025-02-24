package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

func (l *CaliburnListener) ExitParametersEmpty(c *parsing.ParametersEmptyContext) {
	Push(l, []ast.Parameter{})
}

func (l *CaliburnListener) ExitParametersList(c *parsing.ParametersEmptyContext) {
	// Nothing
}

func (l *CaliburnListener) ExitParametersListInitial(c *parsing.ParametersListInitialContext) {
	Push(l, []ast.Parameter{Pop[ast.Parameter](l)})
}

func (l *CaliburnListener) ExitParametersListAppend(c *parsing.ParametersListAppendContext) {
	Push(l, append(Pop[[]ast.Parameter](l), Pop[ast.Parameter](l)))
}

func (l *CaliburnListener) ExitTypedParameter(c *parsing.TypedParameterContext) {
	Push(l, ast.NewTypedParameter(Pop[ast.TypedAssignDeclaration](l)))
}

func (l *CaliburnListener) ExitUntypedParameter(c *parsing.UntypedParameterContext) {
	Push(l, ast.NewUntypedParameter(Pop[ast.UntypedAssignDeclaration](l)))
}

func (l *CaliburnListener) ExitStructDestrutureParameter(c *parsing.StructDestrutureParameterContext) {
	Push(l, ast.NewStructDestructureParameter(Pop[[]ast.Parameter](l)))
}

func (l *CaliburnListener) ExitTupleDestrutureParameter(c *parsing.TupleDestrutureParameterContext) {
	Push(l, ast.NewTupleDestructureParameter(Pop[[]ast.Parameter](l)))
}
