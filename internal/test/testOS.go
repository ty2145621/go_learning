package main

import (
	"fmt"
	"runtime"
)

func testOS() {
	sysType := runtime.GOOS

	fmt.Printf("sysType:%s\n", sysType)
	if sysType == "linux" {
		// LINUX系统
		return
	}

	if sysType == "windows" {
		// windows系统
		return
	}
	return
}
