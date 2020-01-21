package templatevar

const (
	ApiApiTemplate = `
package {{.ProjectName}}

import (
    "sync"

    "{{.Git}}/util/instrumenting"
    "github.com/go-kit/kit/log"
)

var (
    logger = log.NewNopLogger()
    // TODO: labels
    instr = instrumenting.InitGlobalInstrumentation(ServiceName, []string{})
    wg  = &sync.WaitGroup{}
)

// Config contains the required fields for running a server
type Config struct {
    DebugAddr string 
    Addr      string 
}

type newFunc func(cfg interface{}) ServiceCloser

// 返回基础设施层,供上层调用
func Instrumentation() instrumenting.Instrumentation {
    return instr
}
`
)
