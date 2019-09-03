package main

import (
	"fmt"
	"github.com/feifan00x/runner/core"
	"github.com/feifan00x/runner/utils"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		core.ExecCommand(fmt.Sprint("run ", os.Args[1]), func() {
			//get run result
			msg := core.GetRuntimeMessage()
			if msg.Show {
				utils.PrintlnColorMsg(msg)
				msg.Show = false
				core.UpdateRuntimeMessage(msg)
			} else {
				utils.PrintlnColor(utils.DefaultColor, "runner success")
			}
		})
	} else {
		core.Runner()
	}
}
