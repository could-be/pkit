package main

import (
	"flag"
	"log"

	"github.com/could-be/tools/pkit/internal/base"
	"github.com/could-be/tools/pkit/internal/kitcmd"
	"github.com/could-be/tools/pkit/monitor"
)

var (
	ShowFieldNameFlag = flag.Bool("ShowFieldName", false, "true show field name")
	pprof             = flag.Bool("pprof", false, "enable profiling")
)

func init() {
	base.Kit.Commands = []*base.Command{
		kitcmd.CmdInit,
		kitcmd.CmdUpdate,
	}
}

func main() {
	flag.Usage = base.Usage
	flag.Parse()
	log.SetFlags(log.Llongfile | log.LstdFlags)

	args := flag.Args()
	if len(args) < 1 {
		base.Usage()
	}

	if *pprof {
		monitor.StartCpuProf()
		defer monitor.StopCpuProf()
	}

	for _, cmd := range base.Kit.Commands {
		if cmd.Name() != args[0] {
			continue
		}
		cmd.Flag.Parse(args[1:])
		args = cmd.Flag.Args()
		cmd.Run(cmd, args)
		return
	}
}
