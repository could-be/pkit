package generator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// 解析proto 文件, 返回文件名
func GenProto(protoFile string) {

	protoCmd := fmt.Sprintf("protoc %s --go_out=plugins=grpc:.", protoFile)
	cmd := exec.Command("/bin/sh", "-c", protoCmd)
	cmd.Stderr = os.Stderr
	if _, err := cmd.Output(); err != nil {
		log.Fatal(err)
	}
}
