package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func run() {
	fmt.Printf("running %v as PID %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Createflag
	//}

	must(os.Chdir("/"))
	must(cmd.Run())
}
func child() {
	fmt.Printf("running %v as PID %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/"))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
