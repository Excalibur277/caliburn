package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// block: L_C_BRACK statement* R_C_BRACK # BlockRule
func (l *CaliburnListener) ExitBlock(c *parsing.BlockContext) {
	Push(l, ast.NewBlock(Pop[[]ast.Statement](l)))
}
