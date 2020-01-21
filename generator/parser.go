package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/could-be/tools/pkit/models"
)

func ParseGoFile(path string) *models.CodeInfo {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	codeInfo := &models.CodeInfo{
		ProjectName: strings.ToLower(node.Name.String()),
		Git:         models.Git(),
	}

	ast.Inspect(node, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.TypeSpec:
			name := node.Name.String()
			switch node := node.Type.(type) {
			case *ast.InterfaceType:
				if strings.HasSuffix(name, "Server") {
					codeInfo.Interface = handleInterfaceType(strings.TrimSuffix(name, "Server"), node)
				}
				return false
			}
		}
		return true
	})

	return codeInfo
}

func handleInterfaceType(name string, node *ast.InterfaceType) *models.InterfaceType {
	apis := make([]*models.Api, 0, len(node.Methods.List))
	for _, m := range node.Methods.List {
		name := m.Names[0].String()
		switch n := m.Type.(type) {
		case *ast.FuncType:
			apis = append(apis, handleInterfaceMethod(name, n))
		}
	}

	return &models.InterfaceType{
		InterfaceName: name,
		Apis:          apis,
	}
}

func handleInterfaceMethod(methodName string, node *ast.FuncType) *models.Api {
	var params, results *models.Field

	for _, param := range node.Params.List {
		var name string
		if param.Names != nil && len(param.Names) > 0 {
			name = param.Names[0].String()
		}
		if params = handleMethodField(name, models.PARAMS_FIELD_TYPE, param.Type); params != nil {
			break
		}
	}

	for _, result := range node.Results.List {
		var name string
		if result.Names != nil && len(result.Names) > 0 {
			name = result.Names[0].String()
		}
		if results = handleMethodField(name, models.PARAMS_FIELD_TYPE, result.Type); results != nil {
			break
		}
	}
	if params == nil || results == nil {
		panic("params or results must be StarExpr ")
	}

	return &models.Api{
		ApiName: methodName,
		Params:  params,
		Results: results,
	}
}

// 根据需求,只关心需要的字段
func handleMethodField(varName string, typ models.FieldType, expr ast.Expr) *models.Field {
	var typeName string
	switch e := expr.(type) {

	case *ast.StarExpr:
		switch i := e.X.(type) {
		case *ast.Ident:
			typeName = "*" + i.String()
			if varName == "" {
				varName = typ.String()
			}
			field := &models.Field{
				TypeName: typeName,
			}
			if models.ShowFieldName() {
				field.VarName = varName
			}
			return field
		}

	}

	return nil
}
