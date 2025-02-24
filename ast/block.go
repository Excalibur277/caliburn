package ast

type Block interface {
	IsBlock()
}

type BlockBase struct {
}

func (b *BlockBase) IsBlock() {}
