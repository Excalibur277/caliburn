package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

func (l *CaliburnListener) ExitBlock(c *parsing.BlockContext) {
	Push(l, ast.NewBlock(Pop[[]ast.Statement](l)))
}
