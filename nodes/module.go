package nodes

type Module interface {
	IsModule()
}

type BaseModule struct {
	Definitions []Definition
}

func (bm *BaseModule) IsModule() {}
