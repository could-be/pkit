package main

import (
	"io/ioutil"

	"github.com/could-be/tools/pkit/internal/genstatic"
)

const (
	projectTemplateDir    = "project/template"
	staticTemplateVarsDir = "models/templatevar"
)

func main() {
	// 读取模板文件
	templateVarTmpl, err := ioutil.ReadFile("project/self/var.gotemplate")
	if err != nil {
		panic(err)
	}

	templateInfoTmpl, err := ioutil.ReadFile("project/self/info.gotemplate")
	if err != nil {
		panic(err)
	}

	genstatic.GenStaticFromProjectTemplate(projectTemplateDir, staticTemplateVarsDir, string(templateVarTmpl), string(templateInfoTmpl))
}
