package core

import (
	"fmt"
	"github.com/feifan00x/runner/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type command struct {
	name         string
	remark       string
	execNil      func()
	execInt      func(code int)
	execIntCode  int
	execStr      func(val string)
	execStrVal   string
	execFunc     func(func())
	execFuncName func()
}

var commands []command

func execCommand(name string) {
	commands = []command{
		{name: "q", remark: "exit", execInt: commandExit, execIntCode: 0},
		{name: "s", remark: "scan", execFunc: commandReload, execFuncName: Runner},
		{name: `\d+`, remark: "run index", execInt: commandExecIndexShell},
		{name: `h`, remark: "help", execNil: commandHelp, execFunc: commandReload, execFuncName: Runner},
	}
	for _, obj := range commands {
		b, _ := regexp.MatchString(obj.name, name)
		if obj.name == name || b {
			if obj.execNil != nil {
				obj.execNil()
			}
			if obj.execInt != nil {
				if obj.execIntCode == 0 {
					obj.execIntCode, _ = strconv.Atoi(name)
				}
				obj.execInt(obj.execIntCode)
			}
			if obj.execStr != nil {
				obj.execStr(obj.execStrVal)
			}
			if obj.execFunc != nil {
				obj.execFunc(obj.execFuncName)
			}
			return
		}
	}
	runtimeMessage = utils.GenerateMessage(utils.DefaultErrColor, "command error")
	commandReload(Runner)
}

//help
func commandHelp() {
	var commandHelp string
	for index, obj := range commands {
		commandHelp = fmt.Sprint(commandHelp, index, ",", obj.name, ",", obj.remark, "\n")
	}
	runtimeMessage = utils.GenerateMessage(utils.DefaultColor, commandHelp)
	utils.ExecShell("clear")
}

//reload
func commandReload(reload func()) {
	utils.ExecShell("clear")
	defer reload()
}

//exit
func commandExit(code int) {
	os.Exit(code)
}

//exec shell
func commandExecIndexShell(index int) {
	var configs = *runtimeRunConfigs
	if index > len(configs.Configs) {
		utils.PrintlnColor(utils.DefaultErrColor, "index not fund")
	}
	conf := configs.Configs[index-1]
	if conf.Cmd != "" && conf.Cmd != " " {
		result := utils.ExecShell(fmt.Sprint(conf.Cmd, "&& echo $$"))
		resultAndPid := strings.Split(result, `
`)
		fmt.Println(len(resultAndPid))
		fmt.Println(resultAndPid)
		conf.Result = strings.Split(result, `
`)[0]
		conf.LRT = time.Now().Format("2006-01-02 15:04:05")
		conf.Pid = strings.Split(result, `
`)[1]
	}
	updateRuntimeConfigs(conf, index-1)
	commandReload(Runner)
}
