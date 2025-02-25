package ast

import "fmt"

type StructValues interface {
	Node
	IsStructValues()
}

type StructValuesNamed struct {
	values []NamedStructValue
}

func NewStructValuesNamed(values []NamedStructValue) *StructValuesNamed {
	return &StructValuesNamed{values: values}
}

func (svn *StructValuesNamed) IsStructValues() {}
func (svn *StructValuesNamed) String() string  { return SliceToString(svn.values, " ") }

type StructValuesUnamed struct {
	values []Expression
}

func NewStructValuesUnamed(values []Expression) *StructValuesUnamed {
	return &StructValuesUnamed{values: values}
}

func (svu *StructValuesUnamed) IsStructValues() {}
func (svu *StructValuesUnamed) String() string  { return SliceToString(svu.values, " ") }

type NamedStructValue interface {
	Node
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
func (cv *NamedStructValueBase) String() string      { return fmt.Sprintf("%s : %s", cv.key, cv.value) }
