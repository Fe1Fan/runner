package core

import (
	"bufio"
	"fmt"
	"github.com/feifan00x/runner/info"
	"github.com/feifan00x/runner/utils"
	"os"
)

func init() {
	initFile(checkFile())
	loadConf()
}

func Runner() {
	fmt.Print(info.Banner)
	utils.PrintlnColor(utils.DefaultColor, info.Version)
	printTable()
	msg := GetRuntimeMessage()
	if msg.Show {
		utils.PrintlnColorMsg(msg)
		msg.Show = false
		UpdateRuntimeMessage(msg)
	}
	fmt.Println("input s scan config or index number exec.")
	for {
		reader := bufio.NewReader(os.Stdin)

		commandStr, _, _ := reader.ReadLine()
		ExecCommand(string(commandStr), Runner)
	}
}
