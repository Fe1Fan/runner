package core

import (
	"fmt"
	"github.com/feifan00x/runner/utils"
	"os"
	"regexp"
	"strconv"
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
func commandExecIndexShell(code int) {
	fmt.Print("input")
	fmt.Println(code)
}
