package main

import (
	"fmt"
	"github.com/feifan00x/runner/config"
	"github.com/feifan00x/runner/core"
	"github.com/feifan00x/runner/plugins"
)

const banner string = `__________                                  
\______   \__ __  ____   ____   ___________ 
 |       _/  |  \/    \ /    \_/ __ \_  __ \
 |    |   \  |  /   |  \   |  \  ___/|  | \/
 |____|_  /____/|___|  /___|  /\___  >__|   
        \/           \/     \/     \/ ` + "\n"

var command string

func main() {
	fmt.Println(banner)
	plugins.InitPlugins(func() {
		fmt.Println("Welcome Runner V 0.0.1, Plugin Init Success")
	})
	go core.Task()
	plugins.FormatPlugins()
	config.ShowPluginsTable()
	config.GetMessage()
	fmt.Println("input s or number")
	_, _ = fmt.Scanln(&command)
	core.Runner(command, main)
}
