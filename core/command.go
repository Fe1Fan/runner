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
	pattern      string
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
		{name: "q", pattern: "^q$", remark: "exit", execInt: commandExit, execIntCode: 0},
		{name: "s", pattern: "^s$", remark: "scan", execFunc: commandReload, execFuncName: Runner},
		{name: "h", pattern: "h", remark: "help", execNil: commandHelp, execFunc: commandReload, execFuncName: Runner},
		{name: "run", pattern: "^run (\\d+|\\S+)", remark: "run index or name", execStr: commandRun},
		{name: "stop", pattern: "^stop (\\d+|\\S+)", remark: "run index or name", execStr: commandStop},
	}
	for _, obj := range commands {
		b, _ := regexp.MatchString(obj.pattern, name)
		if b {
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
				if obj.execStrVal == "" || obj.execStrVal == " " {
					obj.execStrVal = name
				}
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
func commandRun(val string) {
	var vals = strings.Fields(val)
	if len(vals) > 1 {
		val = vals[1]
	}
	var configs = *runtimeRunConfigs
	var conf RunConf
	index, err := strconv.Atoi(val)
	if err != nil {
		//find by name
		fmt.Println(fmt.Sprint("find by name", val))
		for i, obj := range configs.Configs {
			if obj.Name == val {
				conf = configs.Configs[i]
				index = i + 1
			}
		}
	} else {
		//find by index
		fmt.Println(fmt.Sprint("find by index", index))
		if index > len(configs.Configs) {
			utils.PrintlnColor(utils.DefaultErrColor, "index not fund")
		}
		conf = configs.Configs[index-1]
	}

	if conf == (RunConf{}) {
		runtimeMessage = utils.GenerateMessage(utils.DefaultErrColor, fmt.Sprint("Not fund ", val))
		commandReload(Runner)
		return
	}

	if conf.Cmd != "" && conf.Cmd != " " {
		result := utils.ExecShell(fmt.Sprint(conf.Cmd, "&& echo $$"))
		// fuck \n
		resultAndPid := strings.Split(result, `
`)
		conf.Result = resultAndPid[0]
		conf.LRT = time.Now().Format("2006-01-02 15:04:05")
		conf.Pid = resultAndPid[1]
	}
	updateRuntimeConfigs(conf, index-1)
	commandReload(Runner)
}

func commandStop(val string) {

}
