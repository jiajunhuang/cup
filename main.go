package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/docker/docker/pkg/reexec"
)

func init() {
	reexec.Register("childProcess", childProcess)
	if reexec.Init() {
		os.Exit(0)
	}
}

func childProcess() {
	log.Printf("childProcess start...uid: %d, gid: %d\n", os.Getuid(), os.Getgid())
	if h, err := os.Hostname(); err == nil {
		log.Printf("child: hostname: %s\n", h)
	}
	if err := syscall.Sethostname([]byte("cup-host")); err != nil {
		log.Panicf("failed to set hostname: %s", err)
	}
	if h, err := os.Hostname(); err == nil {
		log.Printf("child: hostname: %s\n", h)
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Panicf("failed to run bash: %s", err)
	}
}

func main() {
	log.Println("main start...")
	cmd := reexec.Command("childProcess")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			syscall.SysProcIDMap{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			syscall.SysProcIDMap{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},
	}

	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to run command: %s", err)
	}
	if h, err := os.Hostname(); err == nil {
		log.Printf("main: hostname: %s\n", h)
	}
}
