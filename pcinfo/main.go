package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type GoInfoObject struct {
	GoOS     string
	Kernel   string
	Core     string
	Platform string
	OS       string
	Hostname string
	CPUs     int
}

func main() {
	cmd := exec.Command("cmd", "lsb_release -a")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	outStr := out.String()
	outStrArr := strings.Split(outStr, ":")
	fmt.Println(outStrArr[len(outStrArr)-1])
}
