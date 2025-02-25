package listener

import (
	"caliburncode.com/m/ast"
	"caliburncode.com/m/parsing"
)

// if_statement: IF inline_statements expression block # IfStatement
func (l *CaliburnListener) ExitIfStatement(c *parsing.IfStatementContext) {
	Push(l, ast.NewIfStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.Block](l)))
}

// if_statement: IF inline_statements expression block else_statement # IfElseStatement
func (l *CaliburnListener) ExitIfElseStatement(c *parsing.IfElseStatementContext) {
	Push(l, ast.NewIfElseStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.Block](l), Pop[ast.ElseStatement](l)))
}

// else_statement: ELSE block # ElseStatement
func (l *CaliburnListener) ExitElseStatement(c *parsing.ElseStatementContext) {
	Push(l, ast.NewElseStatement(Pop[[]ast.Block](l)))
}

// else_statement: ELSE if_statement # ElseIfStatement
func (l *CaliburnListener) ExitElseIfStatement(c *parsing.ElseIfStatementContext) {
	Push(l, ast.NewElseIfStatement(Pop[[]ast.IfStatement](l)))
}

// for_statement: FOR inline_statements expression block # ForStatement
func (l *CaliburnListener) ExitForStatement(c *parsing.ForStatementContext) {
	Push(l, ast.NewForStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.Block](l)))
}

// for_statement: FOR inline_statements expression ARROW inline_statements block # ForStatementWithAfter
func (l *CaliburnListener) ExitForStatementWithAfter(c *parsing.ForStatementWithAfterContext) {
	Push(l, ast.NewForStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[[]ast.InlineStatement](l), Pop[ast.Block](l)))
}

// switch_statement: SWITCH inline_statements expression L_C_BRACK case_blocks R_C_BRACK # SwitchStatement
func (l *CaliburnListener) ExitSwitchStatement(c *parsing.SwitchStatementContext) {
	Push(l, ast.NewSwitchStatement(Pop[[]ast.InlineStatement](l), Pop[ast.Expression](l), Pop[ast.CaseBlocks](l)))
}

// case_blocks: option_case_blocks # CaseBlocks
func (l *CaliburnListener) ExitCaseBlocks(c *parsing.CaseBlocksContext) {
	Push(l, ast.NewCaseBlocks(Pop[[]ast.OptionCaseBlock](l)))
}

// case_blocks: option_case_blocks default_case_block # CaseBlocksDefault
func (l *CaliburnListener) ExitCaseBlocksDefault(c *parsing.CaseBlocksContext) {
	Push(l, ast.NewCaseBlocksDefault(Pop[[]ast.OptionCaseBlock](l), Pop[ast.DefaultCaseBlock](l)))
}

// option_case_blocks: # OptionCaseBlocksInitial
func (l *CaliburnListener) ExitOptionCaseBlocksInitial(c *parsing.OptionCaseBlocksInitialContext) {
	Push(l, []ast.OptionCaseBlock{})
}

// option_case_blocks: option_case_blocks option_case_block # OptionCaseBlocksAppend
func (l *CaliburnListener) ExitOptionCaseBlocksAppend(c *parsing.OptionCaseBlocksAppendContext) {
	Push(l, append(Pop[[]ast.OptionCaseBlock](l), Pop[ast.OptionCaseBlock](l)))
}

// option_case_block: CASE expression block # OptionCaseBlock
func (l *CaliburnListener) ExitOptionCaseBlock(c *parsing.OptionCaseBlockContext) {
	Push(l, ast.NewOptionCaseBlock(Pop[ast.Expression](l), Pop[ast.Block](l)))
}

// default_case_block: DEFAULT block # DefaultCaseBlock
func (l *CaliburnListener) ExitDefaultCaseBlock(c *parsing.DefaultCaseBlockContext) {
	Push(l, ast.NewDefaultCaseBlock(Pop[ast.Block](l)))
}
