package templatevar

const (
	ApiClientTemplate = `
{{$ProjectName := .ProjectName}}
package {{$ProjectName}}client

import (
    "context"
    "io"
    "time"

    "github.com/go-kit/kit/endpoint"
    "github.com/go-kit/kit/sd"
    "github.com/go-kit/kit/sd/etcdv3"
    "github.com/go-kit/kit/sd/lb"
    "github.com/go-kit/kit/tracing/opentracing"
    grpctransport "github.com/go-kit/kit/transport/grpc"
    "github.com/golang/glog"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"

    "{{.Git}}/{{$ProjectName}}/api"
    "{{.Git}}/util/evolve"
    "{{.Git}}/util/icontext/cachectx"
    "{{.Git}}/util/icontext/userctx"
    "{{.Git}}/util/igrpc"
    "{{.Git}}/util/instrumenting"
    "{{.Git}}/util/instrumenting/dlog"
    "{{.Git}}/util/instrumenting/id"
    "{{.Git}}/util/iregister"
)

{{with .Interface}}

    // 使用 grpc 直接连接
    func New(grpcAddr string, instr instrumenting.Instrumentation, options ...ClientOption) ({{$ProjectName}}.{{.InterfaceName}}, error) {
        conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
        dlog.FatalWithMessage("dial grpc failed", err)

        endpoints := &{{$ProjectName}}.Endpoints{
            {{- range .Apis }}
            {{.ApiName}}Endpoint: opentracing.TraceClient(instr.Tracer(), "{{.ApiName}}")(
                grpctransport.NewClient(
                    conn,
                    {{$ProjectName}}.ServiceName,
                    "{{.ApiName}}",
                    encodeGRPCRequest,
                    decodeGRPCResponse,
                    {{$ProjectName}}.{{StartExprToNotExpr .Results.TypeName}}{},
                    ClientOptions(instr, options...)...,
                ).Endpoint(),
            ),
            {{- end }}
        }

        return endpoints, nil
    }

    func NewWithDiscovery(discoveryConfig *iregister.EtcdConfig, instr instrumenting.Instrumentation) ({{$ProjectName}}.{{.InterfaceName}}, error) {

        // Build the client.
        cli, err := etcdv3.NewClient(context.TODO(), discoveryConfig.Addresses, etcdv3.ClientOptions{
            DialTimeout:   time.Duration(discoveryConfig.DialTimeout) * time.Second,
            DialKeepAlive: time.Duration(discoveryConfig.DialTimeout) * time.Second,
            Username:      discoveryConfig.Username,
            Password:      discoveryConfig.Password,
        })
        dlog.Fatal(err)

        instancer, err := etcdv3.NewInstancer(cli, evolve.ServiceKey({{$ProjectName}}.ServiceName), instr.Logger())
        dlog.Fatal(err)

        return &{{$ProjectName}}.Endpoints{
            {{- range .Apis }}
            {{.ApiName}}Endpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
                endpointer := sd.NewEndpointer(
                    instancer,
                    endpointFactory({{$ProjectName}}.ServiceName, "{{.ApiName}}", &{{$ProjectName}}.{{StartExprToNotExpr .Results.TypeName}}{}, instr),
                    instr.Logger())
                balancer := lb.NewRoundRobin(endpointer)

                retry := lb.Retry(3, 3*time.Second, balancer)
                if response, err = retry(ctx, request); err != nil {
                    if e, ok := err.(lb.RetryError); ok {
                        err = e.Final
                    }
                }
                return
            },
            {{- end }}
        }, nil
    }

    func encodeGRPCRequest(_ context.Context, request interface{}) (interface{}, error) {
        return request, nil
    }

    func decodeGRPCResponse(_ context.Context, reply interface{}) (interface{}, error) {
        return reply, nil
    }

    type clientConfig struct {
        headers []string
    }

    // ClientOption is a function that modifies the client config
    type ClientOption func(*clientConfig) error

    func ClientOptions(instr instrumenting.Instrumentation, options ...ClientOption) []grpctransport.ClientOption {
        var cc clientConfig

        for _, f := range options {
            err := f(&cc)
            if err != nil {
                // return nil, errors.Wrap(err, "cannot apply option")
                // TODO: 去掉这个逻辑
                glog.Errorf("cannot apply option", err)
            }
        }
        clientOptions := []grpctransport.ClientOption{
            grpctransport.ClientBefore(
                // TODO: 去掉这个逻辑
                contextValuesToGRPCMetadata(cc.headers),
                opentracing.ContextToGRPC(instr.Tracer(), instr.Logger()),
                id.ContextToGRPC(),
                userctx.ContextToGRPC(),
                cachectx.ContextToGRPC(),
            ),
        }

        return clientOptions
    }

    func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
        return func(ctx context.Context, md *metadata.MD) context.Context {
            var pairs []string
            for _, k := range keys {
            if v, ok := ctx.Value(k).(string); ok {
                pairs = append(pairs, k, v)
            }
        }

        if pairs != nil {
            *md = metadata.Join(*md, metadata.Pairs(pairs...))
        }

            return ctx
        }
    }


    func endpointFactory(serviceName, apiName string,
        reply interface{}, instr instrumenting.Instrumentation) sd.Factory {
        return func(instance string) (endpoint.Endpoint, io.Closer, error) {
            conn, closer, err := igrpc.Get(instance)
            if err != nil {
                return nil, nil, err
            }
            endpoint := grpctransport.NewClient(
                conn,
                serviceName,
                apiName,
                encodeGRPCRequest,
                decodeGRPCResponse,
                reply,
                grpctransport.ClientBefore(
                    opentracing.ContextToGRPC(instr.Tracer(), instr.Logger()),
                    id.ContextToGRPC(),
                    userctx.ContextToGRPC(),
                    cachectx.ContextToGRPC()),
            ).Endpoint()

            endpoint = opentracing.TraceClient(instr.Tracer(), apiName)(endpoint)

            return endpoint, closer, nil
        }
    }
{{- end }}
`
)
