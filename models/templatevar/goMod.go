package templatevar

const (
	GoModTemplate = `
module {{.Git}}/{{.ProjectName}}

go 1.13

require (
    {{.Git}}/util v0.0.0-00010101000000-000000000000
    github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
    github.com/golang/protobuf v1.3.2
    github.com/go-kit/kit v0.9.0
    github.com/jinzhu/gorm v1.9.11
    github.com/pelletier/go-toml v1.6.0 // indirect
    google.golang.org/grpc v1.25.1
    gopkg.in/redis.v5 v5.2.9
)

{{if (eq .LocalFlag true) -}}
    replace {{.Git}}/util => {{.UtilPath}}
{{- end}}

`
)
