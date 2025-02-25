package ast

type StructValues interface {
	IsStructValues()
}

type StructValuesNamed struct {
	values []NamedStructValue
}

func NewStructValuesNamed(values []NamedStructValue) *StructValuesNamed {
	return &StructValuesNamed{values: values}
}

func (svn *StructValuesNamed) IsStructValues() {}

type StructValuesUnamed struct {
	values []Expression
}

func NewStructValuesUnamed(values []Expression) *StructValuesUnamed {
	return &StructValuesUnamed{values: values}
}

func (svu *StructValuesUnamed) IsStructValues() {}

type NamedStructValue interface {
	IsNamedStructValue()
}

type NamedStructValueBase struct {
	key   Identifier
	value Expression
}

func NewNamedStructValue(key Identifier, value Expression) *NamedStructValueBase {
	return &NamedStructValueBase{key: key, value: value}
}

func (cv *NamedStructValueBase) IsNamedStructValue() {}
