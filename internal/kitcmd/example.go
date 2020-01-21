package kitcmd

import (
	"github.com/could-be/tools/pkit/internal/base"
)

var CmdExample = &base.Command{
	UsageLine:   "kit example",
	CustomFlags: true,
	Short:       "generate an example",
}

func init() {
	CmdExample.Run = runExample
}

// TODO
func runExample(cmd *base.Command, args []string) {
}
