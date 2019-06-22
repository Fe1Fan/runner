package core

import (
	"fmt"
	"github.com/feifan00x/runner/config"
	"github.com/feifan00x/runner/plugins"
	"io/ioutil"
	"os/exec"
	"strconv"
)

func Runner(command string, resultFunc func()) {
	defer resultFunc()
	defer clean()
	var pluginsInfo = *config.PluginsAddress
	if command == "S" || command == "s" {
		config.ErrorMessage = ""
		config.InfoMessage = "scanner local plugins success"
		plugins.ScannerPlugins()
	} else {
		index, err := strconv.Atoi(command)
		if err != nil {
			config.InfoMessage = ""
			config.ErrorMessage = "place input number or 's' "
			return
		}
		if index > len(pluginsInfo.Plugin)-1 {
			config.InfoMessage = ""
			config.ErrorMessage = fmt.Sprint("not found ", index)
			return
		}
		plugin := pluginsInfo.Plugin[index]
		fmt.Println(plugin)
		config.ErrorMessage = ""
		config.InfoMessage = execShell(plugin.Uri)
		plugins.UpdatePluginsState(index, "start", config.InfoMessage)
	}
}

func clean() {
	fmt.Println("\033[H\033[2J")
}

func execShell(command string) string {
	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return err.Error()
	}

	outBytes, _ := ioutil.ReadAll(stdout)
	_ = stdout.Close()

	if string(outBytes) == "" {
		fmt.Println("return null")
		return "return null"
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return err.Error()
	}
	fmt.Println(string(outBytes))
	return string(outBytes)
}
