package ast

import "fmt"

type ForStatement interface {
	Node
	Statement
	IsForStatement()
}

type ForStatementBase struct {
	StatementBase
	priorStatements []InlineStatement
	expression      Expression
	block           Block
}

func NewForStatement(priorStatements []InlineStatement, expression Expression, block Block) *ForStatementBase {
	return &ForStatementBase{priorStatements: priorStatements, expression: expression, block: block}
}

func (fs *ForStatementBase) IsForStatement() {}
func (fs *ForStatementBase) String() string {
	return fmt.Sprintf("for %s %s %s", SliceToString(fs.priorStatements, " "), fs.expression, fs.block)
}

type ForStatementWithPost struct {
	ForStatementBase
	postStatements []InlineStatement
}

func NewForStatementWithPost(priorStatements []InlineStatement, expression Expression, postStatements []InlineStatement, block Block) *ForStatementWithPost {
	return &ForStatementWithPost{ForStatementBase: *NewForStatement(priorStatements, expression, block), postStatements: postStatements}
}

func (fs *ForStatementWithPost) String() string {
	return fmt.Sprintf("for %s %s -> %s %s", SliceToString(fs.priorStatements, " "), fs.expression, SliceToString(fs.postStatements, " "), fs.block)
}
