package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// parameters: # ParametersEmpty
func (l *CaliburnListener) ExitParametersEmpty(c *parsing.ParametersEmptyContext) {
	Push(l, []ast.Parameter{})
}

// parameters_list: parameter # ParametersListInitial
func (l *CaliburnListener) ExitParametersListInitial(c *parsing.ParametersListInitialContext) {
	l.Pop(1)
	Push(l, []ast.Parameter{Dequeue[ast.Parameter](l)})
}

// parameters_list: parameters_list COMMA parameter # ParametersListAppend
func (l *CaliburnListener) ExitParametersListAppend(c *parsing.ParametersListAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.Parameter](l), Dequeue[ast.Parameter](l)))
}

// parameter: typed_assign_declaration # TypedParameter
func (l *CaliburnListener) ExitTypedParameter(c *parsing.TypedParameterContext) {
	l.Pop(1)
	Push(l, ast.NewTypedParameter(Dequeue[ast.TypedAssignDeclaration](l)))
}

// parameter: untyped_assign_declaration # UntypedParameter
func (l *CaliburnListener) ExitUntypedParameter(c *parsing.UntypedParameterContext) {
	l.Pop(1)
	Push(l, ast.NewUntypedParameter(Dequeue[ast.UntypedAssignDeclaration](l)))
}

// parameter: L_C_BRACK parameters R_C_BRACK # StructDestrutureParameter
func (l *CaliburnListener) ExitStructDestrutureParameter(c *parsing.StructDestrutureParameterContext) {
	l.Pop(1)
	Push(l, ast.NewStructDestructureParameter(Dequeue[[]ast.Parameter](l)))
}
