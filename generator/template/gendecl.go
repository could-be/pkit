package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"text/template"
)

// declaration --> spec

// A GenDecl node (generic declaration node) represents an import,
// constant, type or variable declaration. A valid Lparen position
// (Lparen.IsValid()) indicates a parenthesized declaration.
//
// Relationship between Tok value and Specs element type:
//
//	token.IMPORT  *ImportSpec
//	token.CONST   *ValueSpec
//	token.TYPE    *TypeSpec
//	token.VAR     *ValueSpec

// Spec
// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.

var src = `
package goods

import "context"

// TODO: 工具需要生成这个代码
var ProjectName = "Goods"

type Goods interface {
	// 减库存
	StockReduction(context.Context, *StockReductionRequest) (resp *StockReductionResponse, err error)
	StockReduction2(context.Context, *StockReductionRequest2) (*StockReductionResponse2, error)
}

`

//
// 词法分析
// 语法分析

type CodeInfo struct {
	PackageName string
	ServiceName string
	Apis        []*Api
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

type Api struct {
	Name    string
	Params  *Field
	Results *Field
	//Params  []*Field
	//Results []*Field
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

//
// ast explorer
// https://astexplorer.net/

var (
	ShowFieldNameFlag = flag.Bool("ShowFieldName", false, "true show field name")
)

func main() {
	flag.Parse()

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", []byte(src), parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var codeInfo CodeInfo
	codeInfo.PackageName = node.Name.String()

	ast.Inspect(node, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.TypeSpec:
			name := node.Name.String()
			switch node := node.Type.(type) {
			case *ast.InterfaceType:
				codeInfo.ServiceName = name
				codeInfo.Apis = handleInterfaceType(node)
				return false
			}
		}
		return true
	})

	// fmt.Println(codeInfo)
	//generate("service.gotemplate", &codeInfo)
	//generate("endpoints.gotemplate", &codeInfo)
	//generate("transport.gotemplate", &codeInfo)
	//generate("run.gotemplate", &codeInfo)
	//generate("api.gotemplate", &codeInfo)
	//generate("doc.gotemplate", &codeInfo)
	generate("client.gotemplate", &codeInfo)
}

func FirstToLower(in string) string {
	if len(in) <= 1 {
		return strings.ToLower(in)
	}
	return strings.ToLower(in[:1]) + in[1:]
}

func StartExprToNotExpr(in string) string {
	if len(in) <= 1 {
		panic("invalid expression")
	}

	return in[1:]
}

var FuncMap = template.FuncMap{
	"FirstToLower":       FirstToLower,
	"ToLower":            strings.ToLower,
	"StartExprToNotExpr": StartExprToNotExpr,
}

func generate(tmpl string, codeInfo *CodeInfo) {

	t, err := template.New(tmpl).Funcs(FuncMap).ParseFiles(tmpl)
	if err != nil {
		panic(err)
	}

	var b []byte
	buffer := bytes.NewBuffer(b)
	if err = t.Execute(buffer, codeInfo); err != nil {
		panic(err)
	}

	//fmt.Println(buffer.String())
	source, err := format.Source(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(source))
}

func handleInterfaceType(node *ast.InterfaceType) []*Api {
	apis := make([]*Api, 0, len(node.Methods.List))
	for _, m := range node.Methods.List {
		name := m.Names[0].String()
		switch n := m.Type.(type) {
		case *ast.FuncType:
			apis = append(apis, handleInterfaceMethod(name, n))
		}
	}

	return apis
}

func handleInterfaceMethod(methodName string, node *ast.FuncType) *Api {
	//params := make([]*Field, 0, len(node.Params.List))
	//results := make([]*Field, 0, len(node.Results.List))
	var params, results *Field

	for _, param := range node.Params.List {
		var name string
		if param.Names != nil && len(param.Names) > 0 {
			name = param.Names[0].String()
		}
		//params = append(params,
		//	handleMethodField(name, PARAMS_FIELD_TYPE, param.Type))
		if params = handleMethodField(name, PARAMS_FIELD_TYPE, param.Type); params != nil {
			break
		}
	}

	for _, result := range node.Results.List {
		var name string
		if result.Names != nil && len(result.Names) > 0 {
			name = result.Names[0].String()
		}
		//results = append(results,
		//	handleMethodField(name, RESULT_FIELD_TYPE, result.Type))
		if results = handleMethodField(name, PARAMS_FIELD_TYPE, result.Type); results != nil {
			break
		}
	}
	if params == nil || results == nil {
		panic("results must be StarExpr ")
	}

	return &Api{
		Name:    methodName,
		Params:  params,
		Results: results,
	}
}

func handleMethodField(varName string, typ FieldType, expr ast.Expr) *Field {
	var typeName string
	switch e := expr.(type) {
	// 忽略这个
	//case *ast.SelectorExpr:
	//	//context.Context
	//	typeName = e.Sel.String()
	//	switch i := e.X.(type) {
	//	case *ast.Ident:
	//		typeName = i.String() + "." + typeName
	//		if varName == "" {
	//			varName = typ.String()
	//		}
	//	}

	case *ast.StarExpr:
		switch i := e.X.(type) {
		case *ast.Ident:
			typeName = "*" + i.String()
			if varName == "" {
				varName = typ.String()
			}
			field := &Field{
				// VarName:  varName,
				TypeName: typeName,
			}
			if *ShowFieldNameFlag {
				field.VarName = varName
			}
			return field
		}

		//err 固定的,忽略
		//case *ast.Ident:
		//	if typeName = e.String(); typeName == "error" && varName == "" {
		//		varName = "err"
		//	}
	}

	// interface 方案, 不要写参数变量
	//field := &Field{
	//	// VarName:  varName,
	//	TypeName: typeName,
	//}
	//if *ShowFieldNameFlag {
	//	field.VarName = varName
	//}
	//
	//return field

	return nil
}
