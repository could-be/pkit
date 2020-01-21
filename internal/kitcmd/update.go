package kitcmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/could-be/tools/pkit/generator"
	"github.com/could-be/tools/pkit/internal/base"
	"github.com/could-be/tools/pkit/models"
	"github.com/could-be/tools/pkit/models/templatevar"
	"github.com/could-be/tools/pkit/util"
)

var CmdUpdate = &base.Command{
	UsageLine: "kit update [path/project.proto]",
	Short:     "update [project.proto]",
}

func init() {
	CmdUpdate.Run = runUpdate
}

// 无参数,自动当前目录递归查找所需proto 文件
// kit update project.proto
// 一个偷懒的坐反, 直接执行 go generate ./...
// TODO: 不指定参数只能查找
// 递归当前目录向下查找
func runUpdate(cmd *base.Command, args []string) {

	switch len(args) {
	default:
		// 指定 proto 文件, 且不在 projectName目录下, 切换到ProjectName目录
		// 检查 proto 文件
		if !strings.HasSuffix(args[0], ".proto") {
			log.Fatal("unknown proto type")
		}

		projectDir := fmt.Sprintf("%s/..", filepath.Dir(args[0]))
		if err := os.Chdir(projectDir); err != nil {
			log.Fatal(err)
		}

	case 0:
		// 没指定 proto 文件, 智能推测, 这种情况只能是在当前项目里,
		// projectName/
		// projectName/api

		// 项目根目录
		if checkProjectDir(".") {
			break
		}

		if checkProjectDir("..") {
			if err := os.Chdir(".."); err != nil {
				panic(err)
			}
			break
		}
	}

	update()
}

// 根据 go.mod 推算
func GetProject() (projectName string) {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && info.Name() == "go.mod" {
			f, err := os.Open(info.Name())
			if err != nil {
				return err
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), "module ") {
					list := strings.Split(scanner.Text(), "/")
					projectName = list[len(list)-1]
					break
				}
			}

			if err := scanner.Err(); err != nil {
				return err
			}
			return errors.New("")
		}
		return nil
	})

	return
}

// 检测是否在项目根目录
func checkProjectDir(path string) bool {
	goMod := fmt.Sprintf("%s/go.mod", path)
	// 第一级检测 是否在项目根目录
	// 检测 go.mod 文件
	if fi, err := os.Stat(goMod); err == nil && !fi.IsDir() {
		return true
	}
	return false
}

// 这里工作目录已经切换到了当前项目的 api 目录
// project/api/
// pbgo := protoDir + "/" + strings.TrimSuffix(protoBase, ".proto") + ".pb.go"
func update() {
	var codeInfo *models.CodeInfo

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(info.Name(), ".proto") {

			generator.GenProto(path)
			pbGo := util.ReplacePathType(path, ".pb.go")
			codeInfo = generator.ParseGoFile(pbGo)
			return errors.New("")
		}
		return nil
	})

	for _, tmplInfo := range templatevar.Templates {
		if tmplInfo.IsKit {
			generator.Generate(tmplInfo.RelativePath, tmplInfo.TemplateSrc, codeInfo)
		}
	}
}
