package util

import (
	"path/filepath"
	"strings"

	"github.com/could-be/tools/pkit/generator"
)

// {{workDir}}/project/template/api/run.go --> apiRun, api/run.go
// go.mod Makefile Dockerfile docker-compose
func PathInfo(prefix, abs string) (snake, relativePath string) {

	relativePath = strings.TrimPrefix(abs, prefix)
	// api/run.go.gotemplate ---> api/run.go
	relativePath = strings.TrimSuffix(relativePath, ".gotemplate")

	// docker-compose --> dockerCompose
	snake = RelativePath2Snake(relativePath)

	return
}

// 删除特殊字符, 并且特殊字符后大写 api/run.go --> apiRunGo
// 先删除.gotemplate
func RelativePath2Snake(str string) string {
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

// api/run.go --> run, .go
func GetPathName(path string) (name, ext string) {
	base := filepath.Base(path)
	ext = filepath.Ext(path)
	name = strings.TrimSuffix(base, ext)
	return
}

func ReplacePathType(path, newExt string) string {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext) + newExt
}
