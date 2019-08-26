package utils

type Message struct {
	Conf ColorConf
	Msg  string
	Show bool
}

func GenerateMessage(conf ColorConf, msg string) Message {
	return Message{Conf: conf, Msg: msg, Show: true}
}
