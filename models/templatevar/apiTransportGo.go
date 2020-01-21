package templatevar

const (
	ApiTransportGoTemplate = `
package {{.ProjectName}}

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"

	"{{.Git}}/util/icontext/cachectx"
	"{{.Git}}/util/icontext/userctx"
	"{{.Git}}/util/instrumenting/id"
)


{{with .Interface}}
	{{$InterfaceName := .InterfaceName}}

	type grpcServer struct {
	{{- range .Apis}}
		{{FirstToLower .ApiName}}Handler   grpctransport.Handler
	{{- end }}
	}

	{{range .Apis -}}
	func (g *grpcServer) {{.ApiName}}(ctx context.Context, req {{.Params.TypeName}}) ({{.Results.TypeName}}, error) {
		_, ret, err := g.{{FirstToLower .ApiName}}Handler.ServeGRPC(ctx, req)
		if err != nil {
			return nil, err
		}

		return ret.({{.Results.TypeName}}), err
	}

	{{ end }}

	func MakeGRPCServer(s {{$InterfaceName}}, options ...grpctransport.ServerOption) {{$InterfaceName}} {

		endpoints := MakeEndpoints(s)

		return &grpcServer{
		{{- range .Apis }}
			{{FirstToLower .ApiName}}Handler: grpctransport.NewServer(
				endpoints.{{.ApiName}}Endpoint,
				Decode{{RemAsterisk .Params.TypeName}},
				Encode{{RemAsterisk .Results.TypeName}},
				ServerOptions("{{.ApiName}}", options...)...,
			),
		{{- end }}
		}
	}

	{{ range .Apis -}}
	func Decode{{RemAsterisk .Params.TypeName}}(_ context.Context, grpcReq interface{}) (interface{}, error) {
		req := grpcReq.({{.Params.TypeName}})
		return req, nil
	}
	{{ end }}

	{{ range .Apis -}}
	func Encode{{RemAsterisk .Results.TypeName}}(_ context.Context, response interface{}) (interface{}, error) {
		resp := response.({{.Results.TypeName}})
		return resp, nil
	}
	{{ end }}

	func ServerOptions(apiName string, options ...grpctransport.ServerOption) []grpctransport.ServerOption {
		serverOptions := []grpctransport.ServerOption{
			grpctransport.ServerBefore(
				metadataToContext,
				opentracing.GRPCToContext(instr.Tracer(), apiName, logger),
				id.GRPCToContext(),
				userctx.GRPCToContext(),
				cachectx.GRPCToContext(),
			),
		}
		serverOptions = append(serverOptions, options...)
		return serverOptions
	}

	func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
		for k, v := range md {
			if v != nil {
				// The key is added both in metadata format (k) which is all lower
				// and the http.CanonicalHeaderKey of the key so that it can be
				// accessed in either format
				ctx = context.WithValue(ctx, k, v[0])
				ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
			}
		}

		return ctx
	}
{{- end }}

`
)
