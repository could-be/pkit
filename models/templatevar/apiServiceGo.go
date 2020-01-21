package templatevar

const (
	ApiServiceGoTemplate = `
package {{.ProjectName}}

import (
    "context"
    "{{.Git}}/util/igrpc"
)

{{with .Interface}}
var ServiceName = _{{.InterfaceName}}_serviceDesc.ServiceName

type ServiceCloser interface {
    {{.InterfaceName}}
    igrpc.Closer
}

type {{.InterfaceName}} interface {
    {{- range .Apis }}
        {{.ApiName}}(context.Context, {{.Params.TypeName}}) ({{.Results.TypeName}}, error)
    {{- end }}
}
{{- end }}



`
)
