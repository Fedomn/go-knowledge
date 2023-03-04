package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestKillProcessGroup(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "watch date > date.txt")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	start := time.Now()
	time.AfterFunc(10*time.Second, func() {
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}
