package models

type Project struct {
	LocalFlag bool   // 是否本地包
	UtilPath  string // util 包的本地目录

	Git            string // github.com/could-be
	ProjectName    string // project name
	DockerRegistry string // docker registry
}

// 生成对应的 template 静态源文件
type TemplateVars struct {
	TemplateName string // eg ApiRun --> (ApiRunTemplate 后面为 template 自动添加)
	TemplateSrc  string // `` template 内容, 注意 template 内容不能有 json 标签 和反引号冲突
}

type TemplateInfo struct {
	TemplateName string // eg: apiRun
	RelativePath string // eg: 相对路径 api/run.go
	TemplateSrc  string // `` template 内容, 注意 template 内容不能有 json 标签 和反引号冲突
	IsKit        bool   // api/目录文件, 由本工具维护, 非 api 目录的都为 project
}
