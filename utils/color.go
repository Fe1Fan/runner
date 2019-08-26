package utils

import "fmt"

type ColorConf struct {
	Conf int
	Bg   int
	Text int
}

var DefaultColor = ColorConf{0, 0, 0}
var DefaultSucColor = ColorConf{0, 0, 32}
var DefaultErrColor = ColorConf{0, 0, 31}

func colorStart() (string, int) {
	return "%c[%d;%d;%dm%s%c[0m", 0x1B
}

func colorEnd() int {
	return 0x1B
}

//printf
func PrintfColor(conf ColorConf, msg string) {
	start1, start2 := colorStart()
	fmt.Printf(start1, start2, conf.Conf, conf.Bg, conf.Text, msg, colorEnd())
}

//println
func PrintlnColor(conf ColorConf, msg string) {
	PrintfColor(conf, fmt.Sprint(msg, "\n"))
}

func PrintfColorMsg(message Message) {
	PrintfColor(message.Conf, message.Msg)
}

func PrintlnColorMsg(message Message) {
	PrintlnColor(message.Conf, message.Msg)
}
