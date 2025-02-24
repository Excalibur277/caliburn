package ast

type Module interface {
	IsModule()
}

type ModuleBase struct {
	definitions []Definition
}

func (bm *ModuleBase) IsModule() {}

func NewModule(definitions []Definition) Module {
	return &ModuleBase{definitions: definitions}
}
