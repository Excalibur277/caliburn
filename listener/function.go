package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// parameters: # ParametersEmpty
func (l *CaliburnListener) ExitParametersEmpty(c *parsing.ParametersEmptyContext) {
	Push(l, []ast.Parameter{})
}

// parameters: parameters_list # ParametersFilled
func (l *CaliburnListener) ExitParametersList(c *parsing.ParametersEmptyContext) {
	// Nothing
}

// parameters_list: parameter # ParametersListInitial
func (l *CaliburnListener) ExitParametersListInitial(c *parsing.ParametersListInitialContext) {
	Push(l, []ast.Parameter{Pop[ast.Parameter](l)})
}

// parameters_list: parameters_list COMMA parameter # ParametersListAppend
func (l *CaliburnListener) ExitParametersListAppend(c *parsing.ParametersListAppendContext) {
	Push(l, append(Pop[[]ast.Parameter](l), Pop[ast.Parameter](l)))
}

// parameter: typed_assign_declaration # TypedParameter
func (l *CaliburnListener) ExitTypedParameter(c *parsing.TypedParameterContext) {
	Push(l, ast.NewTypedParameter(Pop[ast.TypedAssignDeclaration](l)))
}

// parameter: untyped_assign_declaration # UntypedParameter
func (l *CaliburnListener) ExitUntypedParameter(c *parsing.UntypedParameterContext) {
	Push(l, ast.NewUntypedParameter(Pop[ast.UntypedAssignDeclaration](l)))
}

// parameter: L_C_BRACK parameters R_C_BRACK # StructDestrutureParameter
func (l *CaliburnListener) ExitStructDestrutureParameter(c *parsing.StructDestrutureParameterContext) {
	Push(l, ast.NewStructDestructureParameter(Pop[[]ast.Parameter](l)))
}
