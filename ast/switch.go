package ast

import (
	"fmt"
)

type SwitchStatement struct {
	StatementBase
	priorStatements []InlineStatement
	expression      Expression
	caseBlocks      CaseBlocks
}

func NewSwitchStatement(priorStatements []InlineStatement, expression Expression, caseBlocks CaseBlocks) *SwitchStatement {
	return &SwitchStatement{priorStatements: priorStatements, expression: expression, caseBlocks: caseBlocks}
}

func (ss *SwitchStatement) String() string {
	return fmt.Sprintf("switch %s %s %s", ss.priorStatements, ss.expression, ss.caseBlocks)
}

type CaseBlocks interface {
	Node
	IsCaseBlocks()
}

type CaseBlocksBase struct {
	optionCaseBlocks []OptionCaseBlock
}

func NewCaseBlocks(optionCaseBlocks []OptionCaseBlock) *CaseBlocksBase {
	return &CaseBlocksBase{optionCaseBlocks: optionCaseBlocks}
}

func (cs *CaseBlocksBase) IsCaseBlocks()  {}
func (cs *CaseBlocksBase) String() string { return SliceToString(cs.optionCaseBlocks, " ") }

type CaseBlocksDefault struct {
	CaseBlocksBase
	defaultCaseBlock DefaultCaseBlock
}

func NewCaseBlocksDefault(optionCaseBlocks []OptionCaseBlock, defaultCaseBlock DefaultCaseBlock) *CaseBlocksDefault {
	return &CaseBlocksDefault{CaseBlocksBase: *NewCaseBlocks(optionCaseBlocks), defaultCaseBlock: defaultCaseBlock}
}

func (cs *CaseBlocksDefault) String() string {
	return fmt.Sprintf("%s %s", cs.CaseBlocksBase, cs.defaultCaseBlock)
}

type DefaultCaseBlock interface {
	Node
	IsDefaultCaseBlock()
}

type DefaultCaseBlockBase struct {
	block Block
}

func NewDefaultCaseBlock(block Block) *DefaultCaseBlockBase {
	return &DefaultCaseBlockBase{block: block}
}

func (dcs *DefaultCaseBlockBase) IsDefaultCaseBlock() {}
func (dcs *DefaultCaseBlockBase) String() string      { return fmt.Sprintf("default %s", dcs.block) }

type OptionCaseBlock interface {
	Node
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
func (dcs *OptionCaseBlockBase) String() string {
	return fmt.Sprintf("case %s %s", dcs.expression, dcs.block)
}
