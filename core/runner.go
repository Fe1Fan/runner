package core

import (
	"fmt"
	"github.com/feifan00x/runner/info"
	"github.com/feifan00x/runner/utils"
	"regexp"
	"strconv"
)

const pattern string = `\d+`

var command string

func Runner() {
	fmt.Print(info.Banner)
	utils.PrintlnColor(utils.DefaultColor, info.Version)
	initFile(checkFile())
	loadConf()
	printTable()
	msg := runtimeMessage
	if msg.Show {
		utils.PrintlnColorMsg(msg)
	}
	fmt.Println("input s scan config or index number exec.")
	for {
		_, _ = fmt.Scanln(&command)
		if command == "s" || command == "S" {
			reload(Runner)
			return
		}
		rex, err := regexp.MatchString(pattern, command)
		if err != nil {
			runtimeMessage = utils.GenerateMessage(utils.DefaultErrColor, fmt.Sprint("regexp error", err.Error()))
			reload(Runner)
			return
		}
		if !rex {
			runtimeMessage = utils.GenerateMessage(utils.DefaultErrColor, "command error")
			reload(Runner)
			return
		}
		index, _ := strconv.Atoi(command)
		var configs = *runtimeRunConfigs
		if index > len(configs.Configs)-1 {
			runtimeMessage = utils.GenerateMessage(utils.DefaultErrColor, fmt.Sprint("not fund index: ", index, " err"))
			reload(Runner)
			return
		}
		var runConf = configs.Configs[index]
		fmt.Println(runConf)
	}
}

func reload(reload func()) {
	utils.ExecShell("clear")
	defer reload()
}
