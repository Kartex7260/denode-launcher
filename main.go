package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var logger = log.Default()

func main() {
	var reLaunch = 0
	defer logger.Printf("Total relaunch`s %d\n", reLaunch)

	exeFileName := GetExeFileName()

	for {
		reLaunch++
		cmd := exec.Command(exeFileName, os.Args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			logger.Printf("%d. Start node with error=%s\n", reLaunch, err)
			break
		}
		pid := cmd.Process.Pid
		logger.Printf("%d. Node(pid: %d) started\n", reLaunch, pid)

		err = cmd.Wait()
		if err != nil {
			logger.Printf("%d. Node(pid: %d) exited with err=%s\n", reLaunch, pid, err)
		} else {
			logger.Printf("%d. Node(pid: %d) exited\n", reLaunch, pid)
		}

		time.Sleep(time.Second)
	}
}

func GetExeFileName() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	var postfix = ""
	if goos == "windows" {
		postfix = ".exe"
	}

	return "./denode-" + goos + "-" + goarch + postfix
}
