package ast

type Block interface {
	IsBlock()
}

type BlockBase struct {
	statements []Statement
}

func NewBlock(statements []Statement) *BlockBase {
	return &BlockBase{statements: statements}
}

func (b *BlockBase) IsBlock() {}
