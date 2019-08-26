package core

import (
	"fmt"
	"github.com/feifan00x/runner/info"
	"github.com/feifan00x/runner/utils"
)

var commandStr string

func Runner() {
	fmt.Print(info.Banner)
	utils.PrintlnColor(utils.DefaultColor, info.Version)
	initFile(checkFile())
	loadConf()
	printTable()
	msg := runtimeMessage
	if msg.Show {
		utils.PrintlnColorMsg(msg)
		runtimeMessage.Show = false
	}
	fmt.Println("input s scan config or index number exec.")
	for {
		_, _ = fmt.Scanln(&commandStr)
		execCommand(commandStr)
	}
}
