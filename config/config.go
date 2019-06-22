package config

import (
	"fmt"
	"os"
	"strconv"
)

import (
	"github.com/olekukonko/tablewriter"
)

// message config
type messageColor struct {
	Conf int
	Bg   int
	Text int
}

var ErrorMessageColor = messageColor{0, 0, 31}
var InfoMessageColor = messageColor{0, 0, 32}

var ErrorMessage = ""

var InfoMessage = ""

func GetMessage() {
	if InfoMessage != "" {
		fmt.Printf("\n %c[%d;%d;%dm%s%c[0m\n\n", 0x1B, InfoMessageColor.Conf, InfoMessageColor.Bg,
			InfoMessageColor.Text, fmt.Sprint("InfoMessage:", InfoMessage), 0x1B)
	}
	if ErrorMessage != "" {
		fmt.Printf("\n %c[%d;%d;%dm%s%c[0m\n\n", 0x1B, ErrorMessageColor.Conf, ErrorMessageColor.Bg,
			ErrorMessageColor.Text, fmt.Sprint("ErrorMessage:", ErrorMessage), 0x1B)
	}
}

//plugins config
const PluginsPath string = "/usr/local/etc/runner/plugins"
const PluginsConfig string = "/plugins.json"

type Plugins struct {
	Plugin  []Plugin
	Include []string
}

type Plugin struct {
	Name    string
	Remark  string
	Uri     string
	Version string
	Status  string
	Pid     string
	Envs    []string
}

var PluginsAddress *Plugins = nil

//show plugins table
func ShowPluginsTable() {
	fmt.Println()
	fmt.Println("Plugins: input 'S' to scanner new add plugins")
	fmt.Println()
	var pluginsInfo = *PluginsAddress
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Name", "Remark", "Version", "Status", "PID"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	for k, v := range pluginsInfo.Plugin {
		table.Append([]string{strconv.Itoa(k), v.Name, v.Remark, v.Version, v.Status, v.Pid})
	}
	table.Render()
}
