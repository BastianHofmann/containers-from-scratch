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
		parent()
	case "child":
		child()
	default:
		fmt.Println("unkown argument")
	}
}

func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET,
	}


	if err := cmd.Run(); err != nil {
		fmt.Println("parent run error:", err)
	}

	syscall.Unmount("rootfs", syscall.MNT_FORCE)
}

func child() {
	if err := syscall.Mount("proc", "rootfs/proc", "proc", 0, ""); err != nil {
		fmt.Println("mount proc error:", err)
	}
	defer syscall.Unmount("/proc", 0)

	if err := syscall.Mount("rootfs", "rootfs", "rootfs", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		fmt.Println("mount rootfs error:", err)
	}

	if err := syscall.Chroot("rootfs"); err != nil {
		fmt.Println("chroot error:", err)
	}

	err := os.Chdir("/")
	if err != nil {
		fmt.Println("chdir error:", err)
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("child run error:", err)
	}
}
