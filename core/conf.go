package core

import (
	"encoding/json"
	"fmt"
	"github.com/feifan00x/runner/utils"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"strconv"
)

type RunConfigs struct {
	Configs []RunConf
}

type RunConf struct {
	Name   string `json:"name"`
	Ver    string `json:"ver"`
	Cmd    string `json:"cmd"`
	Incl   string `json:"incl"`
	Status string `json:"status"`
	Pid    string `json:"pid"`
}

const RunConfPath string = "/usr/local/etc/runner/conf/"
const RunConfFile string = "conf.json"

// check
func checkFile() (bool, bool) {
	pathResult := true
	fileResult := true
	dir, err := os.Open(RunConfPath)
	if err != nil {
		pathResult = false
	} else {
		dir.Close()
	}
	file, err := os.OpenFile(RunConfPath+RunConfFile, os.O_RDONLY, 0)
	if err != nil {
		fileResult = false
	} else {
		defer file.Close()
	}
	if pathResult {
		utils.PrintlnColor(utils.DefaultSucColor, fmt.Sprint("path check: ", pathResult))
	} else {
		utils.PrintlnColor(utils.DefaultErrColor, fmt.Sprint("path check: ", pathResult))
	}
	if fileResult {
		utils.PrintlnColor(utils.DefaultSucColor, fmt.Sprint("file check: ", fileResult))
	} else {
		utils.PrintlnColor(utils.DefaultErrColor, fmt.Sprint("file check: ", fileResult))
	}
	return pathResult, fileResult
}

// init
func initFile(path bool, file bool) {
	if !path {
		fmt.Println(fmt.Sprint("create path: ", RunConfPath))
		mErr := os.MkdirAll(RunConfPath, os.ModePerm)
		if mErr != nil {
			utils.PrintlnColor(utils.DefaultErrColor, mErr.Error())
		}
	}
	if !file {
		fmt.Println(fmt.Sprint("create file: ", RunConfPath, RunConfFile))
		file, cErr := os.Create(RunConfPath + RunConfFile)
		if cErr != nil {
			utils.PrintlnColor(utils.DefaultErrColor, cErr.Error())
			return
		}
		_, wErr := file.WriteString("{}")
		if wErr != nil {
			utils.PrintlnColor(utils.DefaultErrColor, wErr.Error())
		}
	}
}

//load
func loadConf() {
	file, err := os.Open(RunConfPath + RunConfFile)
	if err != nil {
		utils.PrintlnColor(utils.DefaultErrColor, err.Error())
		return
	} else {
		defer file.Close()
	}
	var fileReader = make([]byte, 100)
	fileReader, err = ioutil.ReadFile(file.Name())
	if err != nil {
		utils.PrintlnColor(utils.DefaultErrColor, err.Error())
		return
	}
	var configs RunConfigs
	err = json.Unmarshal(fileReader, &configs)
	runtimeRunConfigs = &configs
}

//show table
func printTable() {
	var configs = *runtimeRunConfigs
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Name", "Remark", "Version", "Status", "PID"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	for k, v := range configs.Configs {
		table.Append([]string{strconv.Itoa(k), v.Name, v.Ver, v.Cmd, v.Status, v.Pid})
	}
	table.Render()
}
