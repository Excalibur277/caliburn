package ast

type SwitchStatement struct{ StatementBase }

type CaseBlocks interface {
	IsCaseBlocks()
}

type CaseBlocksBase struct{}

func (cs *CaseBlocksBase) IsCaseBlocks() {}

type DefaultCaseBlock interface {
	IsDefaultCaseBlock()
}

type DefaultCaseBlockBase struct{}

func (dcs *DefaultCaseBlockBase) IsDefaultCaseBlock() {}

type OptionCaseBlock interface {
	IsOptionCaseBlock()
}

type OptionCaseBlockBase struct{}

func (dcs *OptionCaseBlockBase) IsOptionCaseBlock() {}
