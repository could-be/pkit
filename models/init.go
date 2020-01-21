package models

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

var (
	gitAddress        = "github.com"
	defaultUser       = "could-be"
	localFlag         = false
	utilPath          = "../util"
	showFieldNameFlag = false
	dockerRegistry    = ""
)

// 删除前导/后导 0
func TrimSpace(str string) string {
	str = strings.TrimSuffix(str, " ")
	return strings.TrimPrefix(str, " ")
}

// 读取环境变量 GIT 设置 git 仓库地址
func init() {
	if git := os.Getenv("GIT"); git != "" {
		if _, err := url.Parse(git); err != nil {
			log.Print("invalid git")
			return
		}
		gitAddress = TrimSpace(git)
	}

	if user := os.Getenv("GIT_USER"); user != "" {
		defaultUser = TrimSpace(user)
	}

	if local := os.Getenv("GIT_LOCAL_FLAG"); local != "" {
		localFlag = true
	}

	if utilPathStr := os.Getenv("GIT_UTIL_PATH"); utilPathStr != "" {
		utilPath = TrimSpace(utilPathStr)
	}

	if showFieldName := os.Getenv("Show_Field_Name"); showFieldName == "true" {
		showFieldNameFlag = true
	}
}

func Git() string {
	return gitAddress + "/" + defaultUser
}

// git remote add origin git@github.com:could-be/activity.git
func Repositories(projectName string) string {
	return fmt.Sprintf("git@%s:%s/%s.git", gitAddress, defaultUser, projectName)
}

func Local() bool {
	return localFlag
}

// 项目的根目录
func UtilPath() string {
	return utilPath
}

func ShowFieldName() bool {
	return showFieldNameFlag
}

func DockerRegistry() string {
	if repo := os.Getenv("DOCKER_REPO"); repo != "" {
		dockerRegistry = repo
	}

	return dockerRegistry
}
