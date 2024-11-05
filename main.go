package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

var logger = log.Default()

func main() {
	var reLaunch = 0
	defer logger.Printf("Total relaunch`s %d\n", reLaunch)

	logger.Print("reLauncher start")

	args := os.Args[1:]
	if len(args) == 0 {
		logger.Fatalf("The first argument must specify the executable file.")
	}

	for {
		reLaunch++
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			logger.Printf("%d. Command(%s) start with error=%s\n", reLaunch, args[0], err)
			time.Sleep(time.Second * 5)
			continue
		}
		pid := cmd.Process.Pid
		logger.Printf("%d. Process(%d) started\n", reLaunch, pid)

		err = cmd.Wait()
		if err != nil {
			logger.Printf("%d. Process(%d) exited with err=%s\n", reLaunch, pid, err)
		} else {
			logger.Printf("%d. Process(%d) exited\n", reLaunch, pid)
		}

		time.Sleep(time.Second)
	}
}
