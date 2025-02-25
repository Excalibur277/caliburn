package ast

type Module interface {
	Node
	IsModule()
}

type ModuleBase struct {
	definitions []Definition
}

func (bm *ModuleBase) IsModule() {}

func NewModule(definitions []Definition) Module {
	return &ModuleBase{definitions: definitions}
}

func (mb *ModuleBase) String() string { return SliceToString(mb.definitions, "\n") }
