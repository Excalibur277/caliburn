package ast

type SwitchStatement struct {
	StatementBase
	priorStatements []InlineStatement
	expression      Expression
	caseBlocks      CaseBlocks
}

func NewSwitchStatement(priorStatements []InlineStatement, expression Expression, caseBlocks CaseBlocks) *SwitchStatement {
	return &SwitchStatement{priorStatements: priorStatements, expression: expression, caseBlocks: caseBlocks}
}

type CaseBlocks interface {
	IsCaseBlocks()
}

type CaseBlocksBase struct {
	optionCaseBlocks []OptionCaseBlock
}

func NewCaseBlocks(optionCaseBlocks []OptionCaseBlock) *CaseBlocksBase {
	return &CaseBlocksBase{optionCaseBlocks: optionCaseBlocks}
}

func (cs *CaseBlocksBase) IsCaseBlocks() {}

type CaseBlocksDefault struct {
	CaseBlocksBase
	defaultCaseBlock DefaultCaseBlock
}

func NewCaseBlocksDefault(optionCaseBlocks []OptionCaseBlock, defaultCaseBlock DefaultCaseBlock) *CaseBlocksDefault {
	return &CaseBlocksDefault{CaseBlocksBase: *NewCaseBlocks(optionCaseBlocks), defaultCaseBlock: defaultCaseBlock}
}

type DefaultCaseBlock interface {
	IsDefaultCaseBlock()
}

type DefaultCaseBlockBase struct {
	block Block
}

func NewDefaultCaseBlock(block Block) *DefaultCaseBlockBase {
	return &DefaultCaseBlockBase{block: block}
}

func (dcs *DefaultCaseBlockBase) IsDefaultCaseBlock() {}

type OptionCaseBlock interface {
	IsOptionCaseBlock()
}

type OptionCaseBlockBase struct {
	expression Expression
	block      Block
}

func NewOptionCaseBlock(expression Expression, block Block) *OptionCaseBlockBase {
	return &OptionCaseBlockBase{expression: expression, block: block}
}

func (dcs *OptionCaseBlockBase) IsOptionCaseBlock() {}
