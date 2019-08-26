package core

import (
	"bufio"
	"fmt"
	"github.com/feifan00x/runner/info"
	"github.com/feifan00x/runner/utils"
	"os"
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
		reader := bufio.NewReader(os.Stdin)

		commandStr, _, _ := reader.ReadLine()
		execCommand(string(commandStr))
	}
}
