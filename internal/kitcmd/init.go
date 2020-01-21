package kitcmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/could-be/tools/pkit/generator"
	"github.com/could-be/tools/pkit/internal/base"
	"github.com/could-be/tools/pkit/models"
	"github.com/could-be/tools/pkit/models/templatevar"
	"github.com/could-be/tools/pkit/util"
)

var CmdInit = &base.Command{
	UsageLine: "kit init [projectName]",
	Short:     "init a service with go-kit supported",
}

func init() {
	CmdInit.Run = runInit
}

// kit init projectName
func runInit(cmd *base.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("too few arguments ", len(args))
	}

	projectPath := args[0]
	// 判断同名文件或文件夹是否已经存在
	if _, err := os.Stat(projectPath); err == nil {
		log.Fatalf("project %s already exists", projectPath)
	}

	// 创建工程目录结构 并且切换到工程目录 project
	if err := CreateProject(projectPath); err != nil {
		log.Fatal(err)
	}

	projectName := strings.ToLower(filepath.Base(projectPath))

	// 已经切换到了指定 projectPath 目录, 生成默认工程文件, eg: go mod等
	if err := NewProjectTemplateFiles(projectName); err != nil {
		log.Fatal(err)
	}

	// 初始化仓库
	if err := GitInit(projectName); err != nil {
		log.Fatal(err)
	}
}

// 创建目录结构
func CreateProject(projectPath string) error {

	// 创建工程目录
	if err := os.MkdirAll(projectPath, 0777); err != nil {
		return err
	}

	// 切换工作目录
	if err := os.Chdir(projectPath); err != nil {
		return err
	}

	for _, tmplInfo := range templatevar.Templates {
		if dir := filepath.Dir(tmplInfo.RelativePath); dir != "." {
			if err := os.MkdirAll(dir, 0777); err != nil {
				return fmt.Errorf("create directory %v: %v", dir, err)
			}
		}
	}
	return nil
}

// 创建工程模板文件, 各种源码文件
func NewProjectTemplateFiles(projectName string) error {
	project := &models.Project{
		LocalFlag:      models.Local(),
		UtilPath:       models.UtilPath(),
		Git:            models.Git(),
		ProjectName:    projectName,
		DockerRegistry: models.DockerRegistry(),
	}

	for _, tmplInfo := range templatevar.Templates {
		if !tmplInfo.IsKit {
			DynamicUpdateFileName(projectName, tmplInfo)
			generator.Generate(tmplInfo.RelativePath, tmplInfo.TemplateSrc, project)
		}
	}
	return nil
}

// config.yaml
func DynamicUpdateFileName(projectName string, tmplInfo *models.TemplateInfo) {
	_, ext := util.GetPathName(tmplInfo.RelativePath)

	// 修改 example.proto 文件名为当前项目名
	if filepath.Ext(tmplInfo.RelativePath) == ".proto" {
		tmplInfo.RelativePath = filepath.Dir(tmplInfo.RelativePath) + "/" + projectName + ext
		tmplInfo.TemplateName = util.RelativePath2Snake(tmplInfo.RelativePath)
		return
	}

	base := filepath.Base(tmplInfo.RelativePath)
	if base == "config.toml" || base == "config.json" || base == "config.yaml" {
		tmplInfo.RelativePath = filepath.Dir(tmplInfo.RelativePath) + "/" + projectName + ext
		tmplInfo.TemplateName = util.RelativePath2Snake(tmplInfo.RelativePath)
	}
}

// ▶ git init
// ▶ git remote add origin git@github.com:could-be/activity.git
func GitInit(projectName string) error {
	_, err := exec.Command("/bin/sh", "-c", "git init").Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("/bin/sh", "-c", "git remote add origin "+models.Repositories(projectName)).Output()
	if err != nil {
		return err
	}

	return nil
}
