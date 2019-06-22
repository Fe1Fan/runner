package plugins

import (
	"encoding/json"
	"fmt"
	"github.com/feifan00x/runner/config"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//init
func InitPlugins(resultFunc func()) {
	dir, err := os.Open(config.PluginsPath)
	if err != nil {
		cErr := os.MkdirAll(config.PluginsPath, os.ModePerm)
		if cErr != nil {
			fmt.Println(cErr)
		}
		return
	}
	jsonFile := config.PluginsPath + config.PluginsConfig
	file, err := os.OpenFile(jsonFile, os.O_RDONLY, 0)
	if err != nil {
		_, cErr := os.Create(jsonFile)
		if cErr != nil {
			fmt.Println(cErr)
		}
		return
	}
	dir.Close()
	file.Close()
	defer resultFunc()
}

// format json to type
func FormatPlugins() {
	configPath := config.PluginsPath + config.PluginsConfig
	file, err := os.OpenFile(configPath, os.O_RDONLY, 0)
	if err != nil {
		InitPlugins(FormatPlugins)
		return
	}
	var fileByte = make([]byte, 100)
	fileByte, err = ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(fileByte) <= 0 {
		fmt.Println("NULL!")
		return
	}
	var plugins config.Plugins
	err = json.Unmarshal(fileByte, &plugins)
	if err != nil {
		fmt.Println(err)
		return
	}
	config.PluginsAddress = &plugins

}

func ScannerPlugins() {
	err := filepath.Walk(config.PluginsPath, func(path string, info os.FileInfo, err error) error {
		addPlugins(path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func addPlugins(path string) {
	paths := strings.Split(path, "/")
	files := strings.Split(paths[len(paths)-1], ".")
	if len(files) < 2 || files[1] != "sh" {
		return
	} else {
		_, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		example := []string{"name", "remark", "version", "envs"}
		var plugin = config.Plugin{}
		for i := 0; i < len(example); i++ {
			var command = "cat " + path + " |grep " + example[i]
			cmd := exec.Command("/bin/bash", "-c", command)

			stdout, _ := cmd.StdoutPipe()

			if err := cmd.Start(); err != nil {
				fmt.Println("Execute failed when Start:" + err.Error())
				return
			}

			outBytes, _ := ioutil.ReadAll(stdout)
			_ = stdout.Close()

			if string(outBytes) == "" {
				return
			}

			if err := cmd.Wait(); err != nil {
				fmt.Println("Execute failed when Wait:" + err.Error())
				return
			}
			key := strings.Split(string(outBytes), "=")[0]
			value := strings.Split(strings.Split(string(outBytes), "=")[1], "\n")[0]

			switch key {
			case "name":
				plugin.Name = strings.Replace(value, `"`, "", 2)
				break
			case "remark":
				plugin.Remark = strings.Replace(value, `"`, "", 2)
				break
			case "version":
				plugin.Version = strings.Replace(value, `"`, "", 2)
				break
			case "envs":
				plugin.Envs = strings.Split(strings.Replace(strings.Replace(strings.Replace(value, "(", "", 1), ")", "", 1), `"`, "", 1000), ",")
				break
			}
		}
		plugin.Status = "stop"
		plugin.Pid = "nil"
		plugin.Uri = path
		var plugins = *config.PluginsAddress
		add := true
		for index := 0; index < len(plugins.Plugin); index++ {
			if plugins.Plugin[index].Uri == plugin.Uri {
				add = false
			}
		}
		if add {
			newPlugin := append(plugins.Plugin, plugin)
			newPlugins := config.Plugins{}
			newPlugins.Plugin = newPlugin
			data, err := json.MarshalIndent(newPlugins, "", "  ")
			if err != nil {
				fmt.Println("ERROR:", err)
			}
			_ = ioutil.WriteFile(config.PluginsPath+config.PluginsConfig, data, 0)
		}
	}
}

//update state
func UpdatePluginsState(index int, state, pid string) {
	var plugins = *config.PluginsAddress
	plugins.Plugin[index].Status = state
	plugins.Plugin[index].Pid = pid
	data, err := json.MarshalIndent(plugins, "", "  ")
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	_ = ioutil.WriteFile(config.PluginsPath+config.PluginsConfig, data, 0)
}
