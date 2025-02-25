package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// if_statement: IF inline_statements expression block # IfStatement
func (l *CaliburnListener) ExitIfStatement(c *parsing.IfStatementContext) {
	l.Pop(3)
	Push(l, ast.NewIfStatement(Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Expression](l), Dequeue[ast.Block](l)))
}

// if_statement: IF inline_statements expression block else_statement # IfElseStatement
func (l *CaliburnListener) ExitIfElseStatement(c *parsing.IfElseStatementContext) {
	l.Pop(4)
	Push(l, ast.NewIfElseStatement(Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Expression](l), Dequeue[ast.Block](l), Dequeue[ast.ElseStatement](l)))
}

// else_statement: ELSE block # ElseStatement
func (l *CaliburnListener) ExitElseStatement(c *parsing.ElseStatementContext) {
	l.Pop(1)
	Push(l, ast.NewElseStatement(Dequeue[ast.Block](l)))
}

// else_statement: ELSE if_statement # ElseIfStatement
func (l *CaliburnListener) ExitElseIfStatement(c *parsing.ElseIfStatementContext) {
	l.Pop(1)
	Push(l, ast.NewElseIfStatement(Dequeue[ast.IfStatement](l)))
}

// for_statement: FOR inline_statements expression block # ForStatement
func (l *CaliburnListener) ExitForStatement(c *parsing.ForStatementContext) {
	l.Pop(3)
	Push(l, ast.NewForStatement(Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Expression](l), Dequeue[ast.Block](l)))
}

// for_statement: FOR inline_statements expression ARROW inline_statements block # ForStatementWithAfter
func (l *CaliburnListener) ExitForStatementWithAfter(c *parsing.ForStatementWithAfterContext) {
	l.Pop(4)
	Push(l, ast.NewForStatementWithPost(Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Expression](l), Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Block](l)))
}

// switch_statement: SWITCH inline_statements expression L_C_BRACK case_blocks R_C_BRACK # SwitchStatement
func (l *CaliburnListener) ExitSwitchStatement(c *parsing.SwitchStatementContext) {
	l.Pop(3)
	Push(l, ast.NewSwitchStatement(Dequeue[[]ast.InlineStatement](l), Dequeue[ast.Expression](l), Dequeue[ast.CaseBlocks](l)))
}

// case_blocks: option_case_blocks # CaseBlocks
func (l *CaliburnListener) ExitCaseBlocks(c *parsing.CaseBlocksContext) {
	l.Pop(1)
	Push(l, ast.NewCaseBlocks(Dequeue[[]ast.OptionCaseBlock](l)))
}

// case_blocks: option_case_blocks default_case_block # CaseBlocksDefault
func (l *CaliburnListener) ExitCaseBlocksDefault(c *parsing.CaseBlocksDefaultContext) {
	l.Pop(2)
	Push(l, ast.NewCaseBlocksDefault(Dequeue[[]ast.OptionCaseBlock](l), Dequeue[ast.DefaultCaseBlock](l)))
}

// option_case_blocks: # OptionCaseBlocksInitial
func (l *CaliburnListener) ExitOptionCaseBlocksInitial(c *parsing.OptionCaseBlocksInitialContext) {
	Push(l, []ast.OptionCaseBlock{})
}

// option_case_blocks: option_case_blocks option_case_block # OptionCaseBlocksAppend
func (l *CaliburnListener) ExitOptionCaseBlocksAppend(c *parsing.OptionCaseBlocksAppendContext) {
	l.Pop(2)
	Push(l, append(Dequeue[[]ast.OptionCaseBlock](l), Dequeue[ast.OptionCaseBlock](l)))
}

// option_case_block: CASE expression block # OptionCaseBlock
func (l *CaliburnListener) ExitOptionCaseBlock(c *parsing.OptionCaseBlockContext) {
	l.Pop(2)
	Push(l, ast.NewOptionCaseBlock(Dequeue[ast.Expression](l), Dequeue[ast.Block](l)))
}

// default_case_block: DEFAULT block # DefaultCaseBlock
func (l *CaliburnListener) ExitDefaultCaseBlock(c *parsing.DefaultCaseBlockContext) {
	l.Pop(1)
	Push(l, ast.NewDefaultCaseBlock(Dequeue[ast.Block](l)))
}
