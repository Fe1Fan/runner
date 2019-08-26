package utils

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func ExecShell(command string) string {
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
