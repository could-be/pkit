package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

//第一个字符转换我小写
func FirstToLower(in string) string {
	if len(in) <= 1 {
		return strings.ToLower(in)
	}
	return strings.ToLower(in[:1]) + in[1:]
}

//指针类型转换为非指针类型
func StartExprToNotExpr(in string) string {
	if len(in) <= 1 {
		panic("invalid expression")
	}

	return in[1:]
}

// 去重的星号
func RemAsterisk(in string) string {
	if len(in) <= 1 {
		panic("invalid expression")
	}
	if strings.HasPrefix(in, "*") {
		return strings.TrimPrefix(in, "*")
	}
	return in
}

func FirstToUpper(in string) string {
	if len(in) <= 0 {
		panic("invalid expression")
	}
	return strings.ToUpper(in[:1]) + in[1:]
}

var FuncMap = template.FuncMap{
	"FirstToLower":       FirstToLower,
	"FirstToUpper":       FirstToUpper,
	"ToLower":            strings.ToLower,
	"StartExprToNotExpr": StartExprToNotExpr,
	"RemAsterisk":        RemAsterisk,
}

// 全路径,例如 api/client/client.go
func Generate(outFile, tmpl string, data interface{}) {

	var err error
	t := template.Must(template.New(outFile).Funcs(FuncMap).Parse(tmpl))

	var b []byte
	buffer := bytes.NewBuffer(b)

	if err = t.Execute(buffer, data); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(buffer.Bytes()))
	// go文件 格式化
	if ext := filepath.Ext(outFile); ext == ".go" {
		b, err = format.Source(buffer.Bytes())
		if err != nil {
			log.Fatal(fmt.Sprintf("%s:%s", outFile, err.Error()))
		}
	} else {
		b = buffer.Bytes()
	}

	// 写入文件
	// check dir  防止误删文件夹, 需要检查文件目录是否存在
	if err := ioutil.WriteFile(outFile, b, 0666); err != nil {
		log.Fatal(err)
	}
}
