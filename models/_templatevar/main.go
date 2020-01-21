package templatevar

const (
	MainTemplate = `
package main

import (
	"flag"

	"{{.Git}}/{{.ProjectName}}/api"
	"{{.Git}}/{{.ProjectName}}/models"
	"{{.Git}}/{{.ProjectName}}/service"
	"{{.Git}}/util/iconfig"
	"{{.Git}}/util/instrumenting/dlog"
)

func main() {
	flag.Parse()

	err := iconfig.LoadAndWatch(models.DefaultConfig)
	dlog.Fatal(err)

	{{.ProjectName}}.Run(models.DefaultConfig, service.New)
}
`
)
