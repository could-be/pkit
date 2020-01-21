package models

import (
	"github.com/could-be/tools/pkit/models/templatevar"
)

// 理想的情况下是一个目录树, 非叶子结点是目录, 叶子结点是文件名

const (
	goMod      = "goMod"
	configDev  = "configDev"
	makeFile   = "makefile"
	dockerfile = "dockerfile"
	// dockerCompose = "dockerComse"
	apiEndpoints = "apiEndpoints"
	apiService   = "apiService"
	apiTransport = "apiTransport"
	apiApi       = "apiApi"
	apiRun       = "apiRun"
	apiDoc       = "apiDoc"
	apiClient    = "apiClient"
	apiProto     = "apiProto"
	modelsConfig = "modelsConfig"
	dao          = "dao"
	service      = "service"
	main         = "main"
)

var ModuleFile = []string{
	goMod,
	configDev,
	makeFile,
	dockerfile,
	// dockerCompose,
	apiEndpoints,
	apiService,
	apiTransport,
	apiApi,
	apiRun,
	apiDoc,
	apiClient,
	apiProto,
	modelsConfig,
	dao,
	service,
	main,
}

var ProjectFiles = []string{
	goMod,
	configDev,
	makeFile,
	dockerfile,
	apiDoc,
	apiProto,
	modelsConfig,
	dao,
	service,
	main,
}

var KitFiles = []string{
	apiEndpoints,
	apiService,
	apiTransport,
	apiApi,
	apiRun,
	apiDoc,
	apiClient,
}

var TemplatesFile = map[string]string{
	goMod:      templatevar.GoModTemplate,
	configDev:  templatevar.ConfigDev,
	dockerfile: templatevar.DockerfileTemplate,
	makeFile:   templatevar.MakefileTemplate,
	// dockerCompose: templatevar.DockerComposeTemplate,
	apiEndpoints: templatevar.ApiEndpointsTemplate,
	apiService:   templatevar.ApiServiceTemplate,
	apiTransport: templatevar.ApiTransportTemplate,
	apiApi:       templatevar.ApiApiTemplate,
	apiRun:       templatevar.ApiRunTemplate,
	apiDoc:       templatevar.ApiDocTemplate,
	apiClient:    templatevar.ApiClientTemplate,
	apiProto:     templatevar.ApiProtoTemplate,
	modelsConfig: templatevar.ModelsConfigTemplate,
	dao:          templatevar.DaoTemplate,
	service:      templatevar.ServiceTemplate,
	main:         templatevar.MainTemplate,
}

func FullPath(projectName, module string) string {
	switch module {
	case goMod:
		return "go.mod"
	case configDev:
		return "config-dev.toml"
	case makeFile:
		return "Makefile"
	case dockerfile:
		return "Dockerfile"
	case apiEndpoints:
		return "api/endpoints.go"
	case apiService:
		return "api/service.go"
	case apiTransport:
		return "api/transport.go"
	case apiApi:
		return "api/api.go"
	case apiRun:
		return "api/run.go"
	case apiDoc:
		return "api/doc.go"
	case apiClient:
		return "api/client/client.go"
	case apiProto:
		return "api/" + projectName + ".proto"
	case modelsConfig:
		return "models/config.go"
	case dao:
		return "dao/dao.go"
	case service:
		return "service/service.go"
	case main:
		return "main.go"
	}

	return ""
}
