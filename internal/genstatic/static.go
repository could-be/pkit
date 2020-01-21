package genstatic

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"

	"github.com/could-be/tools/pkit/generator"
	"github.com/could-be/tools/pkit/models"
)

// 删除特殊字符, 并且特殊字符后大写 api/run.go --> apiRunGo
// 先删除.gotemplate
func trimSpecialCharacters(str string) string {
	for _, sep := range []string{".", "-", "/"} {
		list := strings.Split(str, sep)
		for i := range list {
			if i == 0 {
				str = list[0]
				continue
			}
			// 特殊字符后面第一个大写
			str += generator.FirstToUpper(list[i])
		}
	}

	return str
}

// {{workDir}}/project/template/api/run.go --> apiRun, api/run.go
// go.mod Makefile Dockerfile docker-compose
func pathInfo(prefix, abs string) (snake, relativePath string) {

	relativePath = strings.TrimPrefix(abs, prefix)
	// api/run.go.gotemplate ---> api/run.go
	relativePath = strings.TrimSuffix(relativePath, ".gotemplate")

	// docker-compose --> dockerCompose
	snake = trimSpecialCharacters(relativePath)

	return
}

// 是否是 kit 相关内容, 即 api/目录下加内容
func isKit(path string) bool {
	dir := filepath.Dir(path)
	if strings.HasSuffix(dir, "api/client") {
		return true
	}
	if strings.HasSuffix(dir, "api") &&
		!strings.Contains(path, ".proto") {
		return true
	}
	return false
}

func check(info os.FileInfo, path string) bool {
	if info.IsDir() {
		return false
	}
	// 仅解析*.gotemplate 文件
	if !strings.HasSuffix(path, ".gotemplate") {
		return false
	}

	// 忽略_开头的 gotemplate 文件
	base := filepath.Base(path)
	if strings.HasPrefix(base, "_") {
		return false
	}

	return true
}

// GenSelf("project/self/var.gotemplate", "models/templatevar")
// 依据 project/template 里面的文件,生成对应的 Go 模板变量, 文件~静态变量
// 对应生成目录都为 models/templatevar
func GenStaticFromProjectTemplate(projectTemplateDir, templateVarDir, templateVarTmpl, templateInfoTmpl string) {
	var templateVars []models.TemplateVars
	var templateInfoTmpls []models.TemplateInfo

	// path: project/template/
	prefix := projectTemplateDir + "/"

	err := filepath.Walk(projectTemplateDir, func(path string, info os.FileInfo, err error) error {
		// 读取文件, 生成对应静态的 Go 文件 const ApiApiTemplate = ``
		if !check(info, path) {
			return nil
		}

		// 读取 template
		b, err := ioutil.ReadFile(path)
		if err != nil {
			glog.Error(err)
			return err
		}

		name, relativePath := pathInfo(prefix, path)

		templateVars = append(templateVars, models.TemplateVars{
			TemplateName: name, // eg: ApiRun
			TemplateSrc:  string(b),
		})

		templateInfoTmpls = append(templateInfoTmpls, models.TemplateInfo{
			TemplateName: name,         // eg: APIRun
			RelativePath: relativePath, // eg: api/run.go
			IsKit:        isKit(relativePath),
		})
		return nil
	})

	if err != nil {
		panic(err)
	}

	// const ApiRunTemplate = `` static var
	for _, v := range templateVars {
		file := fmt.Sprintf("%s/%s.go", templateVarDir, generator.FirstToLower(v.TemplateName))
		generator.Generate(file, templateVarTmpl, &v)
	}

	// 生成对应的目录数详情 Templates = []models.TemplateInfo{}
	generator.Generate(fmt.Sprintf("%s/%s", templateVarDir, "templates.go"), templateInfoTmpl, &templateInfoTmpls)
}
