package ast

import "fmt"

type Block interface {
	Node
	IsBlock()
}

type BlockBase struct {
	statements []Statement
}

func NewBlock(statements []Statement) *BlockBase {
	return &BlockBase{statements: statements}
}

func (b *BlockBase) IsBlock() {}
func (b *BlockBase) String() string {
	return fmt.Sprintf("{\n%s }", SliceToString(b.statements, "\n"))
}
