package models

import "encoding/json"

type CodeInfo struct {
	Git         string
	ProjectName string // 包名
	Interface   *InterfaceType
}

func (c *CodeInfo) String() string {
	if c == nil {
		return ""
	}

	byt, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(byt)
}

type InterfaceType struct {
	InterfaceName string
	Apis          []*Api
}

func (i *InterfaceType) String() string {
	if i == nil {
		return ""
	}

	byt, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(byt)
}

type Api struct {
	ApiName string
	Params  *Field
	Results *Field
}

func (m *Api) String() string {
	if m == nil {
		return ""
	}

	byt, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(byt)
}

type Field struct {
	VarName  string
	TypeName string
}

func (f *Field) String() string {
	if f == nil {
		return ""
	}

	byt, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(byt)
}

type FieldType int

const (
	INVALID_FIELD_TYPE FieldType = iota
	PARAMS_FIELD_TYPE
	RESULT_FIELD_TYPE
)

func (f *FieldType) String() string {
	if f == nil {
		return ""
	}
	switch *f {
	case PARAMS_FIELD_TYPE:
		return "req"
	case RESULT_FIELD_TYPE:
		return "resp"
	}
	return ""
}
