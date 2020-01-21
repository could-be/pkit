package templatevar

import (
	"github.com/could-be/tools/pkit/models"
)

// 自动生成: 依据project/template 目录结构自动生成的本程序需要的静态文件
var Templates = []models.TemplateInfo{
	{
		TemplateName: "Dockerfile",
		RelativePath: "Dockerfile",
		TemplateSrc:  DockerfileTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "Makefile",
		RelativePath: "Makefile",
		TemplateSrc:  MakefileTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "apiApiGo",
		RelativePath: "api/api.go",
		TemplateSrc:  ApiApiGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "apiClientClientGo",
		RelativePath: "api/client/client.go",
		TemplateSrc:  ApiClientClientGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "apiEndpointsGo",
		RelativePath: "api/endpoints.go",
		TemplateSrc:  ApiEndpointsGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "apiExampleProto",
		RelativePath: "api/example.proto",
		TemplateSrc:  ApiExampleProtoTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "apiRunGo",
		RelativePath: "api/run.go",
		TemplateSrc:  ApiRunGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "apiServiceGo",
		RelativePath: "api/service.go",
		TemplateSrc:  ApiServiceGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "apiTransportGo",
		RelativePath: "api/transport.go",
		TemplateSrc:  ApiTransportGoTemplate,
		IsKit:        true,
	},
	{
		TemplateName: "configDevYaml",
		RelativePath: "config-dev.yaml",
		TemplateSrc:  ConfigDevYamlTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "daoDaoGo",
		RelativePath: "dao/dao.go",
		TemplateSrc:  DaoDaoGoTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "dockerComposeYaml",
		RelativePath: "docker-compose.yaml",
		TemplateSrc:  DockerComposeYamlTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "goMod",
		RelativePath: "go.mod",
		TemplateSrc:  GoModTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "mainGo",
		RelativePath: "main.go",
		TemplateSrc:  MainGoTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "modelsConfigGo",
		RelativePath: "models/config.go",
		TemplateSrc:  ModelsConfigGoTemplate,
		IsKit:        false,
	},
	{
		TemplateName: "serviceServiceGo",
		RelativePath: "service/service.go",
		TemplateSrc:  ServiceServiceGoTemplate,
		IsKit:        false,
	},
}
