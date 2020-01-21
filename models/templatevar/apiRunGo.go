package templatevar

const (
	ApiRunGoTemplate = `
package {{.ProjectName}}

import (
    "net"

    "github.com/golang/glog"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "{{.Git}}/{{.ProjectName}}/models"
    "{{.Git}}/util/evolve"
    "{{.Git}}/util/igrpc"
    "{{.Git}}/util/instrumenting/dlog"
    "{{.Git}}/util/iregister"
)


func Run(cfg *models.Config, newAppFunc newFunc) {

    defer instr.Close()

    glog.CopyStandardLogTo("ERROR")
    defer glog.Flush()

    // listening
    listener, err := net.Listen("tcp", cfg.Addr)
    dlog.Fatal(err)

    grpcServer := grpc.NewServer()
    reflection.Register(grpcServer)

    // new app instance
    app := newAppFunc(cfg)
    dlog.Fatal(err)

    // 传递必要的信息到 grpc Context
    transportSvr := MakeGRPCServer(app)

    // register app to grpc server
    Register{{.Interface.InterfaceName}}Server(grpcServer, transportSvr)

    wg.Add(1)
    go igrpc.GraceHandler(wg, grpcServer, app)

    switch cfg.Etcd {
        case nil:
        glog.Info("start server using grpc")

        // use grpc directly
        err = grpcServer.Serve(listener)
        dlog.Fatal(err)
        default:
        glog.Info("start server using etcd")

        cfg.Etcd.Key = evolve.ServiceKey(ServiceName)
        cfg.Etcd.Value = cfg.Addr

        err = iregister.ServeWithEtcd(cfg.Etcd, grpcServer, listener, logger)
        dlog.Fatal(err)
    }

    wg.Wait()
}

`
)
