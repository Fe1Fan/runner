package core

import "github.com/feifan00x/runner/utils"

var runtimeRunConfigs *RunConfigs = nil

var runtimeMessage utils.Message

func updateRuntimeConfigs(config RunConf, index int) {
	runtimeRunConfigs.Configs[index] = config
	saveConfToFile()
}
